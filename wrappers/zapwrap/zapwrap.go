package zapwrap

import (
	"context"
	"fmt"
	"os"

	"github.com/denistv/wdlogger"
	"github.com/denistv/wdlogger/gelf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewGELFEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = gelf.ShortMessageField
	encoderConfig.TimeKey = gelf.TimestampField
	encoderConfig.CallerKey = gelf.AdditionalField("caller")
	encoderConfig.FunctionKey = gelf.AdditionalField("function")
	encoderConfig.StacktraceKey = gelf.AdditionalField("stacktrace")

	return encoderConfig
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
	cfg := NewGELFEncoderConfig()
	jsonEncoder := zapcore.NewJSONEncoder(cfg)
	zapCore := zapcore.NewCore(jsonEncoder, ws, level)

	fields, err := newGELFFields()
	if err != nil {
		return nil, fmt.Errorf("creating gelf fields: %w", err)
	}

	zapCore.With(fields)
	zapLogger := zap.New(zapCore)
	zapWrapper := New(zapLogger)

	return zapWrapper, nil
}

func New(zl *zap.Logger) *ZapWrapper {
	return &ZapWrapper{
		zap: zl,
	}
}

// ZapWrapper обертка для Zap
type ZapWrapper struct {
	zap *zap.Logger
}

// standard methods

func (z *ZapWrapper) Debug(msg string, fields ...wdlogger.Field) {
	z.zap.Debug(msg, newZapFields(fields...)...)
}

func (z *ZapWrapper) Info(msg string, fields ...wdlogger.Field) {
	z.zap.Info(msg, newZapFields(fields...)...)
}

func (z *ZapWrapper) Warn(msg string, fields ...wdlogger.Field) {
	z.zap.Warn(msg, newZapFields(fields...)...)
}

func (z *ZapWrapper) Error(msg string, fields ...wdlogger.Field) {
	z.zap.Error(msg, newZapFields(fields...)...)
}

func (z *ZapWrapper) Panic(msg string, fields ...wdlogger.Field) {
	z.zap.Panic(msg, newZapFields(fields...)...)
}

func (z *ZapWrapper) Fatal(msg string, fields ...wdlogger.Field) {
	z.zap.Fatal(msg, newZapFields(fields...)...)
}

// context methods

func (z *ZapWrapper) DebugCtx(ctx context.Context, msg string, fields ...wdlogger.Field) {
	z.Debug(msg, wdlogger.ExtractFields(ctx, fields...)...)
}

func (z *ZapWrapper) InfoCtx(ctx context.Context, msg string, fields ...wdlogger.Field) {
	z.Info(msg, wdlogger.ExtractFields(ctx, fields...)...)
}

func (z *ZapWrapper) WarnCtx(ctx context.Context, msg string, fields ...wdlogger.Field) {
	z.Warn(msg, wdlogger.ExtractFields(ctx, fields...)...)
}

func (z *ZapWrapper) ErrorCtx(ctx context.Context, msg string, fields ...wdlogger.Field) {
	z.Error(msg, wdlogger.ExtractFields(ctx, fields...)...)
}

func (z *ZapWrapper) PanicCtx(ctx context.Context, msg string, fields ...wdlogger.Field) {
	z.Panic(msg, wdlogger.ExtractFields(ctx, fields...)...)
}

func (z *ZapWrapper) FatalCtx(ctx context.Context, msg string, fields ...wdlogger.Field) {
	z.Fatal(msg, wdlogger.ExtractFields(ctx, fields...)...)
}

func (z *ZapWrapper) Sync() error {
	return z.zap.Sync()
}

// NewZapField реализована не самым удачным образом. Когда-нибудь вернусь к этому вопросу (или не вернусь). Сейчас это не принципиальный вопрос.
func newZapField(f wdlogger.Field) zap.Field {
	return zap.Field{
		Key:       f.Key,
		Interface: f.Value,
	}
}

func newZapFields(fs ...wdlogger.Field) []zap.Field {
	out := make([]zap.Field, 0, len(fs))

	for _, f := range fs {
		out = append(out, newZapField(f))
	}

	return out
}
