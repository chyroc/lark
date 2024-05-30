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
//
// 具体影响参考 LogLevel 注释
type Logger interface {
	Log(ctx context.Context, level LogLevel, msg string, args ...interface{})
}

// LogLevel ...
//
// 正常情况下:
// trace: 大于两条日志: 1,2. Token 获取日志, 3,4: 同 debug 的两条日志
// debug: 两条日志: 1. xx start, body=<body>, 2. xx success, log_id=<logid>, body=<body>
// info: 两条日志: 1. xx start 2. xx success, log_id=<logid>
// warn,error: 没有日志
//
// 出错的情况:
// 各日志级别都会输出 error 级别的日志, 如果是<=debug, 会有 body, 否则没有
// 如果不希望输出 error 日志, 可以通过 WithDisableErrorLog 将 error 降级到 warn 日志
//
//  1. 开启了 mock 的情况下, mock 链路打印一条 debug 日志
//
//  2. req:
//     2.1 trace/debug: [debug] 日志 + 请求体
//     2.2 info: [info] 日志 + 无请求体
//     2.3 warn/error: 没有日志
//
//  3. resp:
//     3.1 trace/debug: 无论有无错误: [debug] 日志 + 返回体
//     3.2 info: 无论有无错误: [info] 日志 + 无返回体
//     3.3 warn/error: 有错误: [error] 日志 + 无返回体; 无错误: 无日志
type LogLevel int

// LogLevelTrace ...
const (
	LogLevelTrace LogLevel = iota + 1
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

func (r *Lark) getLogLevel(scope, api string) LogLevel {
	switch api {
	case "GetAppAccessToken", "GetTenantAccessToken", "GetAppTicket":
		// 针对内部自动调用的 3 个 token 相关接口
		// 如果 log 级别被调用方设置为 trace, 则表示调用方需要非常详细的日志, 仍旧输出 token 接口详细日志
		// 否则, 认为 token 的 req 和 resp 不需要日志, 只需要 error 日志
		if r.logLevel == LogLevelTrace {
			return LogLevelDebug
		}
		return LogLevelError
	}
	if r.logLevel <= LogLevelDebug {
		return LogLevelDebug
	} else if r.logLevel <= LogLevelInfo {
		return LogLevelInfo
	}
	return LogLevelError
}

func (r *Lark) Log(ctx context.Context, level LogLevel, msg string, args ...interface{}) {
	if level == LogLevelError && r.disableErrorLog {
		level = LogLevelWarn
	}
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
