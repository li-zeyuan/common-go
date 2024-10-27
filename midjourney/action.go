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

// https://www.openai-hk.com/docs/midjourney/taskapi.html#%E6%89%A7%E8%A1%8C%E5%8A%A8%E4%BD%9C

func (c *Client) Action(ctx context.Context, params *model.ActionReq) (*model.MjBaseResponse, error) {
	b, err := json.Marshal(params)
	if err != nil {
		mylogger.Error(ctx, "json marshal action params fail", zap.Error(err))
		return nil, err
	}
	req, err := http.NewRequest("POST", c.conf.Address+c.conf.Pattern+"/mj/submit/action", bytes.NewReader(b))
	if err != nil {
		mylogger.Error(ctx, "new action request fail", zap.Error(err))
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
