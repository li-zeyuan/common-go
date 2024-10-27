package logmiddleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
)

const maxLogBodyLen = 512

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		reqBytes, err := c.GetRawData()
		if err != nil {
			mylogger.Error(c.Request.Context(), "get raw request body fail", zap.Error(err))
			//return err
		}

		if len(reqBytes) > 0 {
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBytes))
		}
		if len(reqBytes) > maxLogBodyLen {
			reqBytes = reqBytes[:maxLogBodyLen]
		}

		crw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = crw

		c.Next()

		respBytes := crw.body.Bytes()
		if len(respBytes) > maxLogBodyLen {
			respBytes = respBytes[:maxLogBodyLen]
		}

		latency := time.Now().Sub(start)
		fields := []zap.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("request_body", string(reqBytes)),
			zap.String("response_body", string(respBytes)),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if latency > time.Millisecond*200 {
			mylogger.Warn(c.Request.Context(), "slow request(>200ms)", fields...)
		}

		if c.Writer.Status() >= http.StatusInternalServerError {
			mylogger.Error(c.Request.Context(), "http_status gte 500", fields...)
		}

		mylogger.Debug(c.Request.Context(), "debug request", fields...)
	}
}
