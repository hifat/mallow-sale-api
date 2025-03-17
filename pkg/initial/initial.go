package initial

import (
	"github.com/go-playground/validator/v10"
	"github.com/hifat/goroger-core/rules"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Validator *validator.Validate
var Logger *zap.Logger

func init() {
	var err error

	Validator, err = rules.Register()
	if err != nil {
		panic(err)
	}

	zconfig := zap.NewProductionConfig()
	zconfig.EncoderConfig.TimeKey = "receiveTimestamp"
	zconfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zconfig.EncoderConfig.StacktraceKey = ""
	zconfig.EncoderConfig.LevelKey = "severity"

	Logger, err = zconfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}
