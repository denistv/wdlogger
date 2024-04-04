package zapwrap

import (
	"context"

	"github.com/denistv/wdlogger"
	"github.com/denistv/wdlogger/gelf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

// TODO another types
func newZapField(f wdlogger.Field) zap.Field {
	zapField := zap.Field{
		Key:       gelf.AdditionalField(f.Key),
		Interface: f.Value,
	}

	switch f.Value.(type) {
	case error:
		zapField.Type = zapcore.ErrorType
		return zapField
	case string:
		zapField.Type = zapcore.StringType
		zapField.String = f.Value.(string)
		return zapField
	default:
		zapField.Type = zapcore.ReflectType
		return zapField
	}
}

func newZapFields(fs ...wdlogger.Field) []zap.Field {
	out := make([]zap.Field, 0, len(fs))

	for _, f := range fs {
		out = append(out, newZapField(f))
	}

	return out
}
