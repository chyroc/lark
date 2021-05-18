package lark

import (
	"context"
	"fmt"
)

type Logger interface {
	Log(ctx context.Context, level LogLevel, msg string, args ...interface{})
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota + 1
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

func (r LogLevel) String() string {
	switch r {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	default:
		return ""
	}
}

func (r *Lark) log(ctx context.Context, level LogLevel, msg string, args ...interface{}) {
	if r.logger != nil && level >= r.logLevel {
		r.logger.Log(ctx, level, msg, args...)
	}
}

type LoggerStdout struct{}

func NewLoggerStdout() Logger {
	return &LoggerStdout{}
}

func (l *LoggerStdout) Log(ctx context.Context, level LogLevel, msg string, args ...interface{}) {
	fmt.Printf("["+level.String()+"] "+msg+"\n", args...)
}
