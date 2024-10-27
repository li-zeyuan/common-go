package model

const (
	NotifyFail    = "FAIL"
	NotifySuccess = "SUCCESS"
)

type BaseWechatResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WeChatLoginReq struct {
	Code string `json:"code" validate:"min=1"`
}

type WeChatLoginResp struct {
	Token string `json:"token"`
}

type WeChatPhoneLoginReq struct {
	PhoneCode   string `json:"phone_code" validate:"min=1"`
	SessionCode string `json:"session_code" validate:"min=1"`
}

type WXSessionRet struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	BaseWechatResp
}

type WeChatGetUserPhoneResp struct {
	BaseWechatResp
	PhoneInfo *WeChatGetUserPhoneRespPhone `json:"phone_info"`
}

type WeChatGetUserPhoneRespPhone struct {
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}

type GetAccessTokenResp struct {
	BaseWechatResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type PrepaySign struct {
	TimeStamp string `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}

type NotifyResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type WeComRobotReq struct {
	MsgType string             `json:"msgtype"`
	Text    *WeComRobotReqText `json:"text"`
}

type WeComRobotReqText struct {
	Content             string   `json:"content"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type WeComRobotContent struct {
	Title     string `json:"title,omitempty"`
	Uid       string `json:"uid"`
	RequestId string `json:"request_id,omitempty"`
	Message   string `json:"message,omitempty"`
}
