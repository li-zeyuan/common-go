package txcloud

import (
	"context"

	"github.com/li-zeyuan/common-go/mylogger"
	v20180321 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
	"go.uber.org/zap"
)

var (
	defaultSource          = "auto"
	defaultTarget          = "en"
	defaultProjectId int64 = 0
)

func (c *Client) TextTranslate(ctx context.Context, sourceText string) (string, error) {
	if len(sourceText) == 0 {
		return "", nil
	}

	req := v20180321.NewTextTranslateRequest()
	req.Source = &defaultSource
	req.SourceText = &sourceText
	req.Target = &defaultTarget
	req.ProjectId = &defaultProjectId

	res, err := c.translateCli.TextTranslateWithContext(ctx, req)
	if err != nil {
		mylogger.Error(ctx, "text translate fail", zap.Error(err), zap.String("source_text", sourceText))
		return "", err
	}

	return *res.Response.TargetText, nil
}
