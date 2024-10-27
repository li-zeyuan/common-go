package wechat

import (
	"bytes"
	"context"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/li-zeyuan/common-go/httptransfer"
	"github.com/li-zeyuan/common-go/model"
	"github.com/li-zeyuan/common-go/mylogger"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"go.uber.org/zap"
)

var (
	wx                  *Wechat
	errInvalidEventType = errors.New("not success event type")
	errNilWechat        = errors.New("nil wechat")
)

type Wechat struct {
	baseUrl       string
	Conf          *Config
	mchPrivateKey *rsa.PrivateKey
	payClient     *core.Client
	downloader    *downloader.CertificateDownloader
}

func NewWechat(ctx context.Context, conf *Config) (*Wechat, error) {
	if wx != nil {
		return wx, nil
	}

	wx = new(Wechat)
	wx.Conf = conf
	wx.baseUrl = "https://api.weixin.qq.com"

	var err error
	if conf.Pay.Enable {
		wx.mchPrivateKey, err = utils.LoadPrivateKeyWithPath(conf.Pay.PrivateKeyPath)
		if err != nil {
			mylogger.Error(ctx, "load private key fail", zap.String("private_key_path", conf.Pay.PrivateKeyPath), zap.Error(err))
			return nil, err
		}

		wx.payClient, err = core.NewClient(ctx, option.WithWechatPayAutoAuthCipher(conf.Pay.MchID, conf.Pay.MchCertificateSerialNumber, wx.mchPrivateKey, conf.Pay.MchAPIv3Key))
		if err != nil {
			mylogger.Error(ctx, "new wechat pay client fail", zap.Error(err))
			return nil, err
		}

		wx.downloader, err = downloader.NewCertificateDownloader(ctx, conf.Pay.MchID, wx.mchPrivateKey, conf.Pay.MchCertificateSerialNumber, conf.Pay.MchAPIv3Key)
		if err != nil {
			mylogger.Error(ctx, "new certificate downloader fail", zap.Error(err))
			return nil, err
		}
	}

	return wx, nil
}

func (w *Wechat) QueryWxSession(ctx context.Context, code string) (*model.WXSessionRet, error) {
	baseWXSessionUrl := w.baseUrl + "/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	url := fmt.Sprintf(baseWXSessionUrl, w.Conf.AppId, w.Conf.Secret, code)
	resp, err := http.Get(url)
	if err != nil {
		mylogger.Error(ctx, "get weChat session error: ", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	wxResp := model.WXSessionRet{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	if wxResp.ErrCode != 0 {
		mylogger.Error(ctx, "weChat session code error: ", zap.String("err_msg", wxResp.ErrMsg))
		return nil, httptransfer.ErrorCode{Code: wxResp.ErrCode, Msg: wxResp.ErrMsg}
	}

	return &wxResp, nil
}

func (w *Wechat) GetAccessToken(ctx context.Context) (string, error) {
	httpResp, err := http.Get(fmt.Sprintf(w.baseUrl+"/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", w.Conf.AppId, w.Conf.Secret))
	if err != nil {
		mylogger.Error(ctx, "get token error: ", zap.Error(err))
		return "", err
	}
	defer httpResp.Body.Close()

	resp := new(model.GetAccessTokenResp)
	decoder := json.NewDecoder(httpResp.Body)
	if err = decoder.Decode(&resp); err != nil {
		return "", err
	}

	if resp.ErrCode != 0 {
		mylogger.Error(ctx, "get access token error: ", zap.String("err_msg", resp.ErrMsg))
		return "", httptransfer.ErrorCode{Code: resp.ErrCode, Msg: resp.ErrMsg}
	}

	return resp.AccessToken, nil
}

func (w *Wechat) GetUserPhone(ctx context.Context, code string) (string, string, error) {
	token, err := w.GetAccessToken(ctx)
	if err != nil {
		return "", "", err
	}

	req := new(model.WeChatLoginReq)
	req.Code = code
	reqB, err := json.Marshal(req)
	if err != nil {
		mylogger.Error(ctx, "marshal req fail", zap.Error(err))
		return "", "", err
	}

	httpResp, err := http.Post(fmt.Sprintf(w.baseUrl+"/wxa/business/getuserphonenumber?access_token=%s", token), "application/json", bytes.NewReader(reqB))
	if err != nil {
		mylogger.Error(ctx, "get user phone error: ", zap.Error(err))
		return "", "", err
	}
	defer httpResp.Body.Close()

	resp := new(model.WeChatGetUserPhoneResp)
	decoder := json.NewDecoder(httpResp.Body)
	if err = decoder.Decode(&resp); err != nil {
		mylogger.Error(ctx, "decode resp fail", zap.Error(err))
		return "", "", err
	}

	if resp.ErrCode != 0 {
		mylogger.Error(ctx, "no zero error code: ", zap.String("err_msg", resp.ErrMsg))
		return "", "", httptransfer.ErrorCode{Code: resp.ErrCode, Msg: resp.ErrMsg}
	}

	return resp.PhoneInfo.PurePhoneNumber, resp.PhoneInfo.CountryCode, nil
}
