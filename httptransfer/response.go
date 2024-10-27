package httptransfer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccJSONResp(c *gin.Context, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	resp := JsonResponse{
		Data: data,
	}

	c.JSON(http.StatusOK, resp)
}

func ErrJSONResp(c *gin.Context, httpCode int, err error) {
	if err == nil {
		return
	}

	resp := JsonResponse{}
	resp.Code = -1
	resp.Msg = "Internal Server Error"
	if errEnum, ok := err.(ErrorCode); ok {
		resp.Code = errEnum.Code
		resp.Msg = errEnum.Msg
	}

	c.JSON(httpCode, resp)
	c.Abort()
}
