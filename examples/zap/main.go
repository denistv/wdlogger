package main

import (
	"errors"
	"github.com/denistv/wdlogger"
	"github.com/denistv/wdlogger/wrappers/zapwrap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	zapWrapper, err := zapwrap.NewGELF(os.Stdout, zapcore.DebugLevel)
	if err != nil {
		panic(err)
	}

	zapWrapper.Debug(
		"debug message",
		wdlogger.NewStringField("string_field", "string field value"),
		wdlogger.NewInt64Field("int64_field", 12345),
		wdlogger.NewErrorField("error_field", errors.New("unexpected error")),
	)

	zapWrapper.Info(
		"info message",
		wdlogger.NewStringField("string_field", "string field value"),
		wdlogger.NewInt64Field("int64_field", 12345),
	)

	zapWrapper.Warn(
		"warn message",
		wdlogger.NewStringField("string_field", "string field value"),
		wdlogger.NewInt64Field("int64_field", 12345),
	)

	zapWrapper.Error(
		"error message",
		wdlogger.NewStringField("string_field", "string field value"),
		wdlogger.NewInt64Field("int64_field", 12345),
		wdlogger.NewErrorField("error", errors.New("unexpected error")),
	)

	zapWrapper.Fatal(
		"fatal message",
		wdlogger.NewStringField("string_field", "string field value"),
		wdlogger.NewInt64Field("int64_field", 12345),
		wdlogger.NewErrorField("error", errors.New("unexpected error")),
	)
}
