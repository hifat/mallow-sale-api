package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//go:generate mockgen -source=./logger.go -destination=./mock/logger.go -package=mockLogger
type ILogger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Errorf(template string, args ...interface{})
	Sync()
}

var log *zap.Logger

type zapLogger struct {
	sugar *zap.SugaredLogger
}

func init() {
	// object key support cloud run logging
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "receiveTimestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	config.EncoderConfig.LevelKey = "severity"

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func New() ILogger {
	return &zapLogger{
		sugar: log.Sugar(),
	}
}

func (l *zapLogger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugar.Warn(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

// Sync flushes any buffered log entries.
// The result of l.sugar.Sync() is assigned to the blank identifier (_) to ignore any error returned,
// because in many cases, the error is not critical (e.g., "sync /dev/stderr: inappropriate ioctl for device").
// This avoids unused variable warnings if the error is intentionally ignored.
func (l *zapLogger) Sync() {
	_ = l.sugar.Sync()
}
