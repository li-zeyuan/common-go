package httptransfer

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
)

var (
	validate = validator.New()
)

func ParseBody(c *gin.Context, reqPointer interface{}) error {
	if reflect.ValueOf(reqPointer).Kind() != reflect.Ptr {
		return errors.New("request params must pointer")
	}

	err := c.BindJSON(reqPointer)
	if err != nil {
		mylogger.Error(c.Request.Context(), "bind json error", zap.Error(err))
		return err
	}

	err = validate.Struct(reqPointer)
	if err != nil {
		mylogger.Error(c.Request.Context(), "validate body fail", zap.Error(err))
		return err
	}

	return nil
}

func ParseQuery(c *gin.Context, reqPointer interface{}) error {
	if reflect.ValueOf(reqPointer).Kind() != reflect.Ptr {
		return errors.New("request params must pointer")
	}

	err := c.BindQuery(reqPointer)
	if err != nil {
		mylogger.Error(c.Request.Context(), "bind query error", zap.Error(err))
		return err
	}

	err = validate.Struct(reqPointer)
	if err != nil {
		return err
	}

	return nil
}
