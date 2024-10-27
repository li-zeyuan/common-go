package mylogger

import "go.uber.org/zap"

var logger *Logger

func Init(cfg *LoggerCfg) error {
	if cfg == nil {
		cfg = DefaultCfg()
	}

	options := make([]optionFun, 0)
	options = append(options,
		WhitLevel(cfg.Level),
		WhitLoggingDir(cfg.LoggingDir),
		WhitMaxSize(cfg.MaxSize),
		WhitMaxAge(cfg.MaxAge),
		WhitMaxBackup(cfg.MaxBackup),
		WhitIsCompress(cfg.IsCompress),
		WhitIsConsole(cfg.IsConsole),
	)

	for _, o := range options {
		o(config)
	}

	var err error
	logger, err = NewZap()
	return err
}

func GetZapLogger() *zap.Logger {
	return logger.log
}
