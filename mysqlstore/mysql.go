package mysqlstore

import (
	"context"
	"time"

	"github.com/li-zeyuan/common-go/mylogger"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

var Db *gorm.DB

func New(conf *Config) error {
	if Db != nil {
		return nil
	}

	var err error
	Db, err = gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: buildLogger(),
	})
	if err != nil {
		return err
	}
	sqlDb, err := Db.DB()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(conf.MaxConn)
	sqlDb.SetMaxOpenConns(conf.MaxOpen)
	sqlDb.SetConnMaxIdleTime(time.Duration(conf.Timeout))
	return nil
}

func buildLogger() zapgorm2.Logger {
	log := zapgorm2.New(mylogger.GetZapLogger())
	log.SlowThreshold = time.Second
	log.IgnoreRecordNotFoundError = true
	log.LogLevel = logger.Info
	log.Context = func(ctx context.Context) []zapcore.Field {
		return []zapcore.Field{zapcore.Field{Key: mylogger.RequestIdKey, Type: zapcore.StringType, String: mylogger.GetRequestID(ctx)}}
	}
	//logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	return log
}

func Close() {
	sqlDb, err := Db.DB()
	if err != nil {
		return
	}

	sqlDb.Close()
}
