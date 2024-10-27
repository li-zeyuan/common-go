package httptransfer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common-go/model"
	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap"
)

var (
	errArrLen  = errors.New("arr len error")
	errPointer = errors.New("filter param must pointer")
)

func ParseAdminQuery(c *gin.Context, params *model.AdminListReq) error {
	err := ParseQuery(c, params)
	if err != nil {
		return err
	}

	params.Offset, params.Limit, err = RangeParser(c.Request.Context(), params.Range)
	if err != nil {
		return err
	}

	params.SortStr, err = SortParser(c.Request.Context(), params.Sort)
	if err != nil {
		return err
	}

	return nil
}

func RangeParser(ctx context.Context, r string) (int, int, error) {
	var offset, limit int
	if len(r) > 0 {
		var arr []int
		err := json.Unmarshal([]byte(r), &arr)
		if err != nil {
			mylogger.Error(ctx, "unmarshal range error: ", zap.Error(err))
			return 0, 0, err
		}

		if len(arr) != 2 {
			mylogger.Error(ctx, errArrLen.Error(), zap.Any("arr", arr))
			return 0, 0, errArrLen
		}

		offset = arr[0]
		limit = arr[1] - arr[0] + 1
	}

	return offset, limit, nil
}

func SortParser(ctx context.Context, sort string) (string, error) {
	var sortStr string
	if len(sort) > 0 {
		var arr []string
		err := json.Unmarshal([]byte(sort), &arr)
		if err != nil {
			mylogger.Error(ctx, "unmarshal sort error: ", zap.Error(err))
			return "", err
		}

		if len(arr) != 2 {
			mylogger.Error(ctx, errArrLen.Error(), zap.Any("arr", arr))
			return "", errArrLen
		}

		sortStr = arr[0] + " " + arr[1]
	}

	return sortStr, nil
}

func FilterParser(ctx context.Context, f string, filterPointer interface{}) error {
	if reflect.ValueOf(filterPointer).Kind() != reflect.Ptr {
		mylogger.Error(ctx, errPointer.Error())
		return errPointer
	}

	if len(f) == 0 {
		return nil
	}

	err := json.Unmarshal([]byte(f), filterPointer)
	if err != nil {
		mylogger.Error(ctx, "unmarshal filter error: ", zap.Error(err))
		return err
	}

	return nil
}

func SetContentRangeHeader(c *gin.Context, offset, respLen, total int) {
	c.Header("Content-Range", fmt.Sprintf("subjects %d-%d/%d", offset, offset+respLen-1, total))
}
