package model

type AdminListReq struct {
	Range  string `form:"range"`
	Filter string `form:"filter"`
	Sort   string `form:"sort"`
	// parse from Range
	Limit  int
	Offset int
	// parse from Sort
	SortStr string
}
