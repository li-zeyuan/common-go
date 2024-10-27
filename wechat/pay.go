package wechat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/li-zeyuan/common-go/model"
	"github.com/li-zeyuan/common-go/mylogger"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"go.uber.org/zap"
)

const (
	TradeStateSucc   = "SUCCESS"
	TradeStateNotPay = "NOTPAY"
)

func (w *Wechat) Prepay(ctx context.Context, prepayRequest jsapi.PrepayRequest) (string, *model.PrepaySign, error) {
	if w.Conf.Pay.Enable == false || w.payClient == nil {
		return "", nil, errors.New("disable wechat pay")
	}

	svc := jsapi.JsapiApiService{Client: w.payClient}
	prepayRequest.Appid = core.String(w.Conf.AppId)
	prepayRequest.Mchid = core.String(w.Conf.Pay.MchID)
	prepayRequest.NotifyUrl = core.String(w.Conf.Pay.NotifyUrl)
	resp, result, err := svc.Prepay(ctx, prepayRequest)
	if err != nil {
		mylogger.Error(ctx, "call prepay fail", zap.Any("prepay_req", prepayRequest), zap.Error(err))
		return "", nil, err
	}
	defer result.Response.Body.Close()

	if result.Response.StatusCode != http.StatusOK {
		mylogger.Error(ctx, "call prepay not ok", zap.Any("response", result.Response))
		return "", nil, errors.New("call prepay not ok")
	}

	nonce, err := utils.GenerateNonce()
	if err != nil {
		mylogger.Error(ctx, "generate nonce fail", zap.Error(err))
		return "", nil, err
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	pkg := fmt.Sprintf("prepay_id=%s", *resp.PrepayId)
	paySign, err := utils.SignSHA256WithRSA(fmt.Sprintf("%s\n%s\n%s\n%s\n", w.Conf.AppId, timestamp, nonce, pkg), w.mchPrivateKey)
	if err != nil {
		mylogger.Error(ctx, "wc pay sign fail", zap.Error(err))
		return "", nil, err
	}

	prePaySign := new(model.PrepaySign)
	prePaySign.TimeStamp = timestamp
	prePaySign.NonceStr = nonce
	prePaySign.Package = pkg
	prePaySign.SignType = "RSA"
	prePaySign.PaySign = paySign

	return *resp.PrepayId, prePaySign, nil
}

func (w *Wechat) ParseNotifyReqParams(ctx context.Context, request *http.Request) (*payments.Transaction, error) {
	handler, err := notify.NewRSANotifyHandler(w.Conf.Pay.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(core.NewCertificateMap(wx.downloader.GetAll(ctx))))
	if err != nil {
		mylogger.Error(ctx, "new notify handler fail", zap.Error(err))
		return nil, err
	}

	transaction := new(payments.Transaction)
	notifyReq, err := handler.ParseNotifyRequest(ctx, request, transaction)
	if err != nil {
		mylogger.Error(ctx, "new notify handler fail", zap.Error(err))
		return nil, err
	}

	mylogger.Info(ctx, "notify params", zap.Any("transaction", transaction), zap.Any("notify_request", notifyReq))

	if notifyReq.EventType != "TRANSACTION.SUCCESS" {
		mylogger.Error(ctx, "not success event type", zap.String("event_type", notifyReq.EventType))
		return nil, errInvalidEventType
	}

	return transaction, nil
}

func (w *Wechat) QueryOrderByOutTradeNo(ctx context.Context, orderID int64) (*payments.Transaction, error) {
	if w.Conf.Pay.Enable == false || w.payClient == nil {
		return nil, errors.New("disable wechat pay")
	}

	reqParams := jsapi.QueryOrderByOutTradeNoRequest{}
	reqParams.OutTradeNo = core.String(strconv.FormatInt(orderID, 10))
	reqParams.Mchid = core.String(w.Conf.Pay.MchID)
	svc := jsapi.JsapiApiService{Client: w.payClient}
	resp, result, err := svc.QueryOrderByOutTradeNo(ctx, reqParams)
	if err != nil {
		mylogger.Error(ctx, "call order by out trade no fail", zap.Any("params", reqParams), zap.Error(err))
		return nil, err
	}
	defer result.Response.Body.Close()

	if result.Response.StatusCode != http.StatusOK {
		mylogger.Error(ctx, "call order by out trade no not ok", zap.Any("response", result.Response))
		return nil, errors.New("call order by out trade no not ok")
	}

	return resp, nil
}
