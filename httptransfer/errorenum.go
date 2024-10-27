package httptransfer

import "encoding/json"

type ErrorCode struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (err ErrorCode) Error() string {
	data, _ := json.Marshal(err)
	return string(data)
}

func (err ErrorCode) HasError() bool {
	return err.Code != 0
}

var (
	// common
	ErrorInvalidArgument   = ErrorCode{Code: 100000, Msg: "输入参数错误"}
	ErrorFrequentOperation = ErrorCode{Code: 100001, Msg: "频繁操作，稍后重试"}
	ErrorTryLater          = ErrorCode{Code: 100002, Msg: "请稍后重试"}

	// login
	ErrorLoginForbid = ErrorCode{Code: 101000, Msg: "账号禁止登陆，请联系我们"}

	// gold
	ErrorGoldInsufficient     = ErrorCode{Code: 102000, Msg: "能量不足，去攒能量吧！"}
	ErrorGoldShareFriendLimit = ErrorCode{Code: 102001, Msg: "分享超过限制，明天再来试试！"}
	ErrorCreditInsufficient   = ErrorCode{Code: 102002, Msg: "积分不足，去充值吧！"}
)
