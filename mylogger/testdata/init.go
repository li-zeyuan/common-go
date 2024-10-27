package testdata

import "github.com/li-zeyuan/common-go/mylogger"

func InitLogger() {
	mylogger.Init(&mylogger.LoggerCfg{
		Level:      "debug",
		LoggingDir: "logs",
		IsCompress: true,
		IsConsole:  true,
		MaxSize:    10,
		MaxAge:     5,
		MaxBackup:  5,
	})
}
