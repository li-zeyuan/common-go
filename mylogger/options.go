package mylogger

var config = &LoggerCfg{}

type LoggerCfg struct {
	// debug可打印所以sql
	Level      string `yaml:"level"`
	LoggingDir string `yaml:"dir"`
	IsCompress bool   `yaml:"is_compress"`
	IsConsole  bool   `yaml:"is_console"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackup  int    `yaml:"max_backup"`
}

func DefaultCfg() *LoggerCfg {
	return &LoggerCfg{
		Level:      "debug",
		LoggingDir: "logs",
		IsCompress: true,
		IsConsole:  true,
		MaxSize:    10,
		MaxAge:     5,
		MaxBackup:  5,
	}
}

type optionFun = func(o *LoggerCfg)

func WhitLevel(level string) optionFun {
	return func(o *LoggerCfg) {
		o.Level = level
	}
}

func WhitLoggingDir(dir string) optionFun {
	return func(o *LoggerCfg) {
		o.LoggingDir = dir
	}
}

func WhitMaxSize(maxSize int) optionFun {
	return func(o *LoggerCfg) {
		o.MaxSize = maxSize
	}
}

func WhitMaxAge(maxAge int) optionFun {
	return func(o *LoggerCfg) {
		o.MaxAge = maxAge
	}
}

func WhitMaxBackup(maxBackup int) optionFun {
	return func(o *LoggerCfg) {
		o.MaxBackup = maxBackup
	}
}

func WhitIsCompress(isCompress bool) optionFun {
	return func(o *LoggerCfg) {
		o.IsCompress = isCompress
	}
}

func WhitIsConsole(isConsole bool) optionFun {
	return func(o *LoggerCfg) {
		o.IsConsole = isConsole
	}
}
