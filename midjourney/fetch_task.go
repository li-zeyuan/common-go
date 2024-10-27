package midjourney

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/li-zeyuan/common-go/model"
	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
)

const (
	FetchStatusSucc = "SUCCESS"
	FetchStatusFail = "FAILURE"
)

// free
// https://www.openai-hk.com/docs/midjourney/taskapi.html#%E6%A0%B9%E6%8D%AEid%E5%88%97%E8%A1%A8%E6%9F%A5%E8%AF%A2%E4%BB%BB%E5%8A%A1

func (c *Client) FetchTasks(ctx context.Context, req *model.FetchTasksReq) (map[string]*model.FetchTasksResp, error) {
	if len(req.Ids) == 0 {
		return map[string]*model.FetchTasksResp{}, nil
	}

	b, err := json.Marshal(req)
	if err != nil {
		mylogger.Error(ctx, "json marshal fetch_tasks request fail", zap.Error(err))
		return nil, err
	}
	httpReq, err := http.NewRequest(http.MethodPost, c.conf.Address+"mj/task/list-by-condition", bytes.NewReader(b))
	if err != nil {
		mylogger.Error(ctx, "new fetch_tasks request fail", zap.Error(err))
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.conf.Key)
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		mylogger.Error(ctx, "do fetch_tasks fail", zap.Error(err))
		return nil, err
	}

	defer httpResp.Body.Close()

	resp := make([]*model.FetchTasksResp, 0)
	if err = checkResp(ctx, httpResp, &resp); err != nil {
		return nil, err
	}

	respMap := make(map[string]*model.FetchTasksResp, len(req.Ids))
	for _, item := range resp {
		respMap[item.Id] = item
	}

	return respMap, nil
}

func (c *Client) ProgressToInt(ctx context.Context, progress string) int {
	if len(progress) == 0 {
		return 0
	}

	if !strings.Contains(progress, "%") {
		return 0
	}

	intStr := progress[:len(progress)-1]
	if len(intStr) == 0 {
		return 0
	}

	i, err := strconv.Atoi(intStr)
	if err != nil {
		mylogger.Error(ctx, "conv progress fail", zap.Error(err), zap.String("progress", progress))
		return 0
	}

	return i
}
