package txcloud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/li-zeyuan/common-go/mylogger"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	v20180321 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
	"github.com/tencentyun/cos-go-sdk-v5"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"go.uber.org/zap"
)

const (
	accelerateFormat = "https://%s-%s.cos.accelerate.myqcloud.com"
	intranetFormat   = "https://%s-%s.cos.%s.myqcloud.com"
)

var (
	errOverObjectLimit       = errors.New("over object limit")
	errNotOK                 = errors.New("cos not return ok")
	errInvalidProcessResults = errors.New("cos invalid process results")
)

type Client struct {
	conf          *Config
	translateCli  *v20180321.Client
	cosCli        *cos.Client
	stsCli        *sts.Client
	httpCli       *http.Client
	baseBucketUrl string
	env           string
}

func New(ctx context.Context, env string, conf *Config) (*Client, error) {
	comCre := common.NewCredential(conf.SecretID, conf.SecretKey)
	trClient, err := v20180321.NewClient(comCre, conf.Region, profile.NewClientProfile())
	if err != nil {
		mylogger.Error(ctx, "new tencent translate client fail", zap.Error(err))
		return nil, err
	}

	//baseBucketUrl := fmt.Sprintf(accelerateFormat, conf.Bucket, conf.AppID)
	//if env == environment.EnvLive {
	baseBucketUrl := fmt.Sprintf(intranetFormat, conf.Bucket, conf.AppID, conf.Region)
	//}

	bucketUrl, err := url.Parse(baseBucketUrl)
	if err != nil {
		mylogger.Error(ctx, "parse dev url fail", zap.Error(err))
		return nil, err
	}

	baseUrl := &cos.BaseURL{BucketURL: bucketUrl}
	cosCli := cos.NewClient(baseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.SecretID,
			SecretKey: conf.SecretKey,
		},
	})

	cli := &Client{
		conf:          conf,
		translateCli:  trClient,
		cosCli:        cosCli,
		stsCli:        sts.NewClient(conf.SecretID, conf.SecretKey, nil),
		httpCli:       &http.Client{Timeout: time.Second * 60},
		baseBucketUrl: baseBucketUrl,
		env:           env,
	}

	return cli, nil
}

func (c *Client) GetConf() *Config {
	return c.conf
}
