package midjourney

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/li-zeyuan/common-go/httptransfer"
	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
)

// prompt:https://docs.midjourney.com/docs/prompts
// 探索prompt：https://docs.midjourney.com/docs/explore-prompting
// /imagine <描述> | <模型> | <数量> | <选项>：根据描述和指定模型生成指定数量（最多16张）图片，并应用指定选项（如放大、裁剪、合成等）。

type Client struct {
	client *http.Client

	conf *Config
}

func NewClient(cfg *Config) *Client {
	return &Client{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		conf: cfg,
	}
}

func checkResp(ctx context.Context, httpResp *http.Response, resp interface{}) error {
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		mylogger.Error(ctx, "read img2img body fail", zap.Error(err))
		return err
	}

	if httpResp.StatusCode != http.StatusOK {
		mylogger.Error(ctx, fmt.Sprintf("mj response status_code: %d, body: %s", httpResp.StatusCode, string(body)))
		return httptransfer.ErrorTryLater
	}

	if err = json.Unmarshal(body, resp); err != nil {
		mylogger.Error(ctx, "unmarshal img2img body fail", zap.Error(err), zap.String("body", string(body)))
		return err
	}

	return nil
}
