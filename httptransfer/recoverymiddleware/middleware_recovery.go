package recoverymiddleware

import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/li-zeyuan/common-go/mylogger"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				goErr := errors.Wrap(err, 0)
				reset := string([]byte{27, 91, 48, 109})
				mylogger.Errorf(c.Request.Context(), "[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s",
					httpRequest, goErr.Error(), goErr.Stack(), reset)
			}
		}()
		c.Next()
	}
}
