package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

const (
	DefaultJsonField = "{}"

	DefaultDelCondition = "deleted_at = 0"
	WhereIdCondition    = "id = ?"

	OrderUpdatedAtDesc = "updated_at desc"
	OrderCreatedAtDesc = "created_at desc"
	OrderCreatedAtAsc  = "created_at asc"
	OrderIDAsc         = "id asc"
	OrderIDDesc        = "id desc"

	ColumnUpdatedAt = "updated_at"
	ColumnCreatedAt = "created_at"
	ColumnDeletedAt = "deleted_at"
)

type BaseModel struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int64     `json:"deleted_at"`
}

type StringArray []string

func (urls StringArray) Value() (driver.Value, error) {
	if len(urls) == 0 {
		return []byte("[]"), nil
	}

	return json.Marshal(urls)
}

func (urls *StringArray) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("urls no byte type")
	}
	return json.Unmarshal(b, urls)
}

type IntArray []int64

func (of *IntArray) Value() (driver.Value, error) {
	if of == nil {
		return []byte(`[]`), nil
	}

	return json.Marshal(of)
}

func (of *IntArray) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("no byte type")
	}

	return json.Unmarshal(b, &of)
}
