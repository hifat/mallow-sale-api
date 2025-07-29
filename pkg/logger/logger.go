package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Errorf(template string, args ...interface{})
	Sync()
}

type zapLogger struct {
	sugar *zap.SugaredLogger
}

func New() Logger {
	z, _ := zap.NewProduction()
	return &zapLogger{
		sugar: z.Sugar().WithOptions(zap.AddCallerSkip(1)),
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
