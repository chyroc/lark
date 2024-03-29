// Code generated by lark_sdk_gen. DO NOT EDIT.
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

import (
	"context"
)

// GetDriveFileTask 查询删除文件夹等异步任务的状态信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/task_check
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/file/async-task/task_check
func (r *DriveService) GetDriveFileTask(ctx context.Context, request *GetDriveFileTaskReq, options ...MethodOptionFunc) (*GetDriveFileTaskResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveFileTask != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#GetDriveFileTask mock enable")
		return r.cli.mock.mockDriveGetDriveFileTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveFileTask",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/task_check",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveFileTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveFileTask mock DriveGetDriveFileTask method
func (r *Mock) MockDriveGetDriveFileTask(f func(ctx context.Context, request *GetDriveFileTaskReq, options ...MethodOptionFunc) (*GetDriveFileTaskResp, *Response, error)) {
	r.mockDriveGetDriveFileTask = f
}

// UnMockDriveGetDriveFileTask un-mock DriveGetDriveFileTask method
func (r *Mock) UnMockDriveGetDriveFileTask() {
	r.mockDriveGetDriveFileTask = nil
}

// GetDriveFileTaskReq ...
type GetDriveFileTaskReq struct {
	TaskID string `query:"task_id" json:"-"` // 文件相关异步任务id, 示例值: 12345
}

// GetDriveFileTaskResp ...
type GetDriveFileTaskResp struct {
	Status string `json:"status,omitempty"` // 异步任务的执行状态, 如果任务执行成功则返回success, 如果任务执行失败则返回fail, 如果任务还在执行中则返回process。
}

// getDriveFileTaskResp ...
type getDriveFileTaskResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *GetDriveFileTaskResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
