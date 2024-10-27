package model

type IdReq struct {
	Id int64 `json:"id" form:"id"`
}

type StartAndLimit struct {
	Start int `form:"start" json:"start"`
	Limit int `form:"limit" json:"limit"`
}

type BaseListResp struct {
	Total int64         `json:"total"`
	List  []interface{} `json:"list"`
}
