package mysqlstore

import (
	comModel "github.com/li-zeyuan/common-go/model"
)

type Filter struct {
	Query interface{}
	Args  []interface{}
}

func NewFilter(query string, args []interface{}) *Filter {
	return &Filter{
		Query: query,
		Args:  args,
	}
}

func NewFilterCommon(query string, args interface{}) *Filter {
	return &Filter{
		Query: query,
		Args:  []interface{}{args},
	}
}

func NewDelAtFilter(isEnable bool) *Filter {
	query := comModel.DefaultDelCondition
	if !isEnable {
		query = "deleted_at > 0"
	}

	return NewFilter(query, nil)
}

func NewWhereIDFilter(id int64) *Filter {
	return NewFilter(comModel.WhereIdCondition, []interface{}{id})
}

func NewWhereIDsFilter(ids []int64) *Filter {
	return NewFilter("id in (?)", []interface{}{ids})
}

func NewWhereUidFilter(uid int64) *Filter {
	return NewFilter("uid = ? ", []interface{}{uid})
}
