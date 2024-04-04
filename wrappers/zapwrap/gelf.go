package zapwrap

import (
	"fmt"
	"os"

	"github.com/denistv/wdlogger/gelf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newGELFEncoderConfig() zapcore.EncoderConfig {
	ec := zap.NewProductionEncoderConfig()

	ec.MessageKey = gelf.ShortMessageField
	ec.TimeKey = gelf.TimestampField
	ec.CallerKey = gelf.AdditionalField("caller")
	ec.FunctionKey = gelf.AdditionalField("function")
	ec.StacktraceKey = gelf.AdditionalField("stacktrace")

	return ec
}

func newGELFFields() ([]zap.Field, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("getting hostname: %w", err)
	}

	return []zap.Field{
		zap.String("version", "1.1"),
		zap.String("host", hostname),
	}, nil
}

func NewGELF(ws zapcore.WriteSyncer, level zapcore.Level) (*ZapWrapper, error) {
	cfg := newGELFEncoderConfig()
	jsonEncoder := zapcore.NewJSONEncoder(cfg)

	fields, err := newGELFFields()
	if err != nil {
		return nil, fmt.Errorf("creating gelf fields: %w", err)
	}

	zapCore := zapcore.NewCore(jsonEncoder, ws, level).With(fields)
	zapLogger := zap.New(zapCore).WithOptions(zap.AddStacktrace(zapcore.ErrorLevel))
	zapWrapper := New(zapLogger)

	return zapWrapper, nil
}
