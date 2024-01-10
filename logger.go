/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

// 关于本 lib 中所打印的日志说明
// 当日志级别设置为 Trace 后, event数据, api 的 request 和 response 数据均会打印(app_secret, helpdesk_token, access_token, encrypt_key 4者会被 mask 为 *)
// 当日志级别设置为 Debug 后, mock-enable
// 当日志级别设置为 Debug 后, 所有接口请求的时候，会打印一个日志，包含接口名称，不包含请求体

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

func (r *Lark) Log(ctx context.Context, level LogLevel, msg string, args ...interface{}) {
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
