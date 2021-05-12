package lark

import (
	"context"
	"fmt"
)

type Logger interface {
	Debug(ctx context.Context, msg string, args ...interface{}) // debug消息，多
	Info(ctx context.Context, msg string, args ...interface{})  // 常见消息，正常
	Warn(ctx context.Context, msg string, args ...interface{})  // 警示消息，少
	Error(ctx context.Context, msg string, args ...interface{}) // 错误消息，出错的时候才有
}

type LogLevel int

const (
	LogLevelDebug LogLevel = iota + 1
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

func (r *Lark) logDebug(ctx context.Context, msg string, args ...interface{}) {
	if r.logger != nil && LogLevelDebug >= r.logLevel {
		r.logger.Debug(ctx, msg, args...)
	}
}

func (r *Lark) logInfo(ctx context.Context, msg string, args ...interface{}) {
	if r.logger != nil && LogLevelInfo >= r.logLevel {
		r.logger.Info(ctx, msg, args...)
	}
}

func (r *Lark) logWarn(ctx context.Context, msg string, args ...interface{}) {
	if r.logger != nil && LogLevelWarn >= r.logLevel {
		r.logger.Warn(ctx, msg, args...)
	}
}

func (r *Lark) logError(ctx context.Context, msg string, args ...interface{}) {
	if r.logger != nil && LogLevelError >= r.logLevel {
		r.logger.Error(ctx, msg, args...)
	}
}

type LoggerStdout struct{}

func NewLoggerStdout() Logger {
	return &LoggerStdout{}
}

func (l *LoggerStdout) Debug(ctx context.Context, msg string, args ...interface{}) {
	fmt.Printf("[debug] "+msg+"\n", args...)
}

func (l *LoggerStdout) Info(ctx context.Context, msg string, args ...interface{}) {
	fmt.Printf("[info] "+msg+"\n", args...)
}

func (l *LoggerStdout) Warn(ctx context.Context, msg string, args ...interface{}) {
	fmt.Printf("[wran] "+msg+"\n", args...)
}

func (l *LoggerStdout) Error(ctx context.Context, msg string, args ...interface{}) {
	fmt.Printf("[error] "+msg+"\n", args...)
}
