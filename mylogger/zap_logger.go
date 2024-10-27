package mylogger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	log           *zap.Logger
	level         zapcore.Level
	sugaredLogger *zap.SugaredLogger
}

func NewZap() (*Logger, error) {
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		log.Println("parse level error", err)
		return nil, err
	}

	aLogger := &Logger{
		level: level,
	}

	debugFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.DebugLevel && level >= aLogger.level
	})

	infoFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel && level >= aLogger.level
	})

	errFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel && level >= aLogger.level
	})

	consoleFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= aLogger.level
	})

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	cores := make([]zapcore.Core, 0)
	if config.IsConsole {
		cores = append(
			cores,
			zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), consoleFunc))
	}

	cores = append(cores,
		zapcore.NewCore(encoder, zapcore.AddSync(GetWriter(config.LoggingDir+"/debug.log")), debugFunc),
		zapcore.NewCore(encoder, zapcore.AddSync(GetWriter(config.LoggingDir+"/info.log")), infoFunc),
		zapcore.NewCore(encoder, zapcore.AddSync(GetWriter(config.LoggingDir+"/error.log")), errFunc),
	)

	zapOptions := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zap.PanicLevel),
		zap.AddCallerSkip(2),
	}
	aLogger.log = zap.New(zapcore.NewTee(cores...), zapOptions...)
	aLogger.sugaredLogger = aLogger.log.Sugar()

	return aLogger, nil
}

func (m *Logger) debugEnable() bool {
	return m.level >= zapcore.DebugLevel
}

func (m *Logger) Debug(msg string, fields ...zap.Field) {
	if !m.debugEnable() {
		return
	}

	m.log.Debug(msg, fields...)
}

func (m *Logger) Info(msg string, fields ...zap.Field) {
	m.log.Info(msg, fields...)
}

func (m *Logger) Warn(msg string, fields ...zap.Field) {
	m.log.Warn(msg, fields...)
}

func (m *Logger) Error(msg string, fields ...zap.Field) {
	m.log.Error(msg, fields...)
}

func (m *Logger) Fatal(msg string, fields ...zap.Field) {
	m.log.Fatal(msg, fields...)
}

func (m *Logger) Debugf(template string, args ...interface{}) {
	m.sugaredLogger.Debugf(template, args...)
}

func (m *Logger) Infof(template string, args ...interface{}) {
	m.sugaredLogger.Infof(template, args...)
}

func (m *Logger) Warnf(template string, args ...interface{}) {
	m.sugaredLogger.Warnf(template, args...)
}

func (m *Logger) Errorf(template string, args ...interface{}) {
	m.sugaredLogger.Errorf(template, args...)
}

func (m *Logger) Fatalf(template string, args ...interface{}) {
	m.sugaredLogger.Fatalf(template, args...)
}
