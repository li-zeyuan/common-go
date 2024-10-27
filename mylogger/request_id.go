package mylogger

import (
	"context"
)

const (
	XRequestIDKey = "X-Request-ID"
	RequestIdKey  = "request_id"
)

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	v := ctx.Value(XRequestIDKey)
	rid, ok := v.(string)
	if !ok {
		return ""
	}

	return rid
}
