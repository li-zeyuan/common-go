package testdata

import (
	"log"

	"github.com/li-zeyuan/common-go/mysqlstore"
)

func InitMysql() {
	if err := mysqlstore.New(&mysqlstore.Config{
		DSN: "root:root@tcp(localhost:3306)/ai?charset=utf8mb4&parseTime=True&loc=UTC",
	}); err != nil {
		log.Fatal(err)
	}
}
