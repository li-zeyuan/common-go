package mysqlstore

import (
	"time"

	comModel "github.com/li-zeyuan/common-go/model"
	"gorm.io/gorm"
)

func Wheres(db *gorm.DB, filters ...*Filter) *gorm.DB {
	if len(filters) == 0 {
		return db
	}

	for _, f := range filters {
		db = db.Where(f.Query, f.Args...)
	}

	return db
}

func NewDelUpdateMap() map[string]interface{} {
	return map[string]interface{}{
		comModel.ColumnDeletedAt: time.Now().Unix(),
	}
}
