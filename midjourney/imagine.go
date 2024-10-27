package midjourney

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/li-zeyuan/common-go/httptransfer"
	"github.com/li-zeyuan/common-go/model"
	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
)

// https://www.openai-hk.com/docs/midjourney/taskapi.html#%E6%8F%90%E4%BA%A4imagine%E4%BB%BB%E5%8A%A1
func (c *Client) Imagine(ctx context.Context, params *model.ImagineReq) (*model.MjBaseResponse, error) {
	b, err := json.Marshal(params)
	if err != nil {
		mylogger.Error(ctx, "json marshal imagine params fail", zap.Error(err))
		return nil, err
	}
	req, err := http.NewRequest("POST", c.conf.Address+c.conf.Pattern+"/mj/submit/imagine", bytes.NewReader(b))
	if err != nil {
		mylogger.Error(ctx, "new imagine request fail", zap.Error(err))
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.conf.Key)
	req.Header.Set("Content-Type", "application/json")

	httpResp, err := c.client.Do(req)
	if err != nil {
		mylogger.Error(ctx, "do img2img fail", zap.Error(err))
		return nil, err
	}

	defer httpResp.Body.Close()
	resp := new(model.MjBaseResponse)
	if err = checkResp(ctx, httpResp, resp); err != nil {
		return nil, err
	}

	if len(resp.Result) == 0 {
		return nil, httptransfer.ErrorTryLater
	}

	return resp, nil
}
