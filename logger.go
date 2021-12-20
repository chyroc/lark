package lark

import (
	"context"
	"fmt"
)

// Logger ...
type Logger interface {
	Log(ctx context.Context, level LogLevel, msg string, args ...interface{})
}

// LogLevel ...
type LogLevel int

// LogLevelTrace ...
const (
	LogLevelTrace LogLevel = iota + 1 // 只有两个 log req 和 resp 的 文本内容
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// String ...
func (r LogLevel) String() string {
	switch r {
	case LogLevelTrace:
		return "TRACE"
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
	if r.logger != nil && r.logLevel <= level {
		r.logger.Log(ctx, level, msg, args...)
	}
}

// LoggerStdout ...
type LoggerStdout struct{}

// NewLoggerStdout ...
func NewLoggerStdout() Logger {
	return &LoggerStdout{}
}

// Log ...
func (l *LoggerStdout) Log(ctx context.Context, level LogLevel, msg string, args ...interface{}) {
	fmt.Printf("["+level.String()+"] "+msg+"\n", args...)
}
