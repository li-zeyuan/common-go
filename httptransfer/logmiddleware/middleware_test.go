package logmiddleware

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common-go/mylogger"
	"github.com/stretchr/testify/assert"
)

func TestLogMiddleware(t *testing.T) {
	cfg := mylogger.DefaultCfg()
	cfg.Level = "debug"
	mylogger.Init(cfg)

	r := gin.New()
	r.Use(LogMiddleware())
	r.POST("/ping", func(c *gin.Context) {
		reqBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, string(reqBody), "test")

		time.Sleep(time.Millisecond * 200)
		c.JSON(http.StatusInternalServerError, "response body")
	})

	res1 := httptest.NewRecorder()
	req1, _ := http.NewRequestWithContext(context.Background(), "POST", "/ping", strings.NewReader("test"))
	r.ServeHTTP(res1, req1)
	t.Log(res1.Body.String())
	assert.Equal(t, res1.Body.String(), `"response body"`)
}
