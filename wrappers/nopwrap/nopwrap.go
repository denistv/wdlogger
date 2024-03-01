package nopwrap

import "github.com/denistv/wdlogger"

func NewNopWrapper() *NopWrapper {
	return &NopWrapper{}
}

// NopWrapper Реализация логгера для тестов, которая ничего не логирует
type NopWrapper struct{}

func (n NopWrapper) Debug(_ string, _ ...wdlogger.Field) {}

func (n NopWrapper) Info(_ string, _ ...wdlogger.Field) {}

func (n NopWrapper) Warn(_ string, _ ...wdlogger.Field) {}

func (n NopWrapper) Error(_ string, _ ...wdlogger.Field) {}

func (n NopWrapper) Panic(_ string, _ ...wdlogger.Field) {}

func (n NopWrapper) Fatal(_ string, _ ...wdlogger.Field) {}
