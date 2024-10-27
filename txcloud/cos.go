package txcloud

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/li-zeyuan/common-go/mylogger"
	"github.com/tencentyun/cos-go-sdk-v5"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"go.uber.org/zap"
)

func (c *Client) Key2Url(key string) string {
	return c.baseBucketUrl + "/" + key
}

func (c *Client) GenObjectName(uid int64, suffix string) string {
	return fmt.Sprintf("/%s/%s-%d-%s%s", c.env, time.Now().Format("20060102150405"), uid, uuid.New().String()[:8], suffix)
}

func (c *Client) GetBaseBucketUrl() string {
	return fmt.Sprintf(intranetFormat, c.conf.Bucket, c.conf.AppID, c.conf.Region)
}

func (c *Client) PutObject(ctx context.Context, name string, fileBytes []byte) (string, string, error) {
	if len(fileBytes) > c.conf.ObjectLimitSizeByte*2 {
		mylogger.Error(ctx, "too large object size", zap.Error(errOverObjectLimit), zap.Int("size", len(fileBytes)))
		return "", "", errOverObjectLimit
	}

	opt := &cos.ObjectPutOptions{
		ACLHeaderOptions: nil,
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			XOptionHeader: &http.Header{},
		},
	}

	format := "webp"
	pic := &cos.PicOperations{
		IsPicInfo: 1,
		Rules: []cos.PicOperationsRules{
			{
				FileId: strings.ReplaceAll(name, path.Ext(name), fmt.Sprintf(".%s", format)),
				Rule:   fmt.Sprintf("imageMogr2/format/%s", format),
			},
		},
	}
	opt.XOptionHeader.Add("Pic-Operations", cos.EncodePicOperations(pic))

	proResult, resp, err := c.cosCli.CI.Put(ctx, name, bytes.NewReader(fileBytes), opt)
	if err != nil {
		mylogger.Error(ctx, "cos put fail", zap.Error(err), zap.String("name", name))
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		mylogger.Error(ctx, "cos invalid status code", zap.Int("status_code", resp.StatusCode), zap.String("name", name))
		return "", "", errNotOK
	}
	if len(proResult.ProcessResults) == 0 {
		mylogger.Error(ctx, "cos invalid process results", zap.Int("status_code", resp.StatusCode), zap.String("name", name))
		return "", "", errInvalidProcessResults
	}

	return proResult.OriginalInfo.Key, proResult.ProcessResults[0].Key, nil
}

func (c *Client) GetObject(ctx context.Context, name string) ([]byte, error) {
	resp, err := c.cosCli.Object.Get(ctx, name, nil)
	if err != nil {
		mylogger.Error(ctx, "cos get fail", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		mylogger.Error(ctx, "cos get invalid status code", zap.Int("status_code", resp.StatusCode), zap.String("name", name))
		return nil, errNotOK
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		mylogger.Error(ctx, "read object fail", zap.Error(err))
		return nil, err
	}

	return bs, err
}

func (c *Client) GetExternalObject(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.httpCli.Get(url)
	if err != nil {
		mylogger.Error(ctx, "get open object fail", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		mylogger.Error(ctx, "get invalid status code", zap.Int("status_code", resp.StatusCode), zap.String("url", url))
		return nil, errNotOK
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		mylogger.Error(ctx, "read open object fail", zap.Error(err))
		return nil, err
	}

	return bs, err
}

// https://cloud.tencent.com/document/product/436/14048
func (c *Client) UserPutObjectTempKey(ctx context.Context, uid int64) (*sts.CredentialResult, error) {
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(1 * time.Minute.Seconds()),
		Region:          c.conf.Region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					// https://cloud.tencent.com/document/product/436/31923
					Action: []string{
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						fmt.Sprintf("qcs::cos:%s:uid/%s:%s-%s/%s/*", c.conf.Region, c.conf.AppID, c.conf.Bucket, c.conf.AppID, c.env),
					},
					// https://cloud.tencent.com/document/product/436/71306
					Condition: map[string]map[string]interface{}{
						"numeric_less_than_equal": {
							"cos:content-length": c.conf.ObjectLimitSizeByte,
						},
						"ForAllValues:StringEquals": {
							"cos:content-type": c.conf.AllowContentType,
						},
					},
				},
			},
		},
	}

	res, err := c.stsCli.GetCredential(opt)
	if err != nil {
		mylogger.Error(ctx, "get temp key fail", zap.Int64("uid", uid), zap.Error(err))
		return nil, err
	}

	return res, nil
}
