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
	"io"
)

// DownloadDriveExportTask 根据[查询导出任务结果](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/get)返回的导出产物`token`, 下载导出产物文件到本地。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/download
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/export_task/download
func (r *DriveService) DownloadDriveExportTask(ctx context.Context, request *DownloadDriveExportTaskReq, options ...MethodOptionFunc) (*DownloadDriveExportTaskResp, *Response, error) {
	if r.cli.mock.mockDriveDownloadDriveExportTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DownloadDriveExportTask mock enable")
		return r.cli.mock.mockDriveDownloadDriveExportTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DownloadDriveExportTask",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/export_tasks/file/:file_token/download",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(downloadDriveExportTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDownloadDriveExportTask mock DriveDownloadDriveExportTask method
func (r *Mock) MockDriveDownloadDriveExportTask(f func(ctx context.Context, request *DownloadDriveExportTaskReq, options ...MethodOptionFunc) (*DownloadDriveExportTaskResp, *Response, error)) {
	r.mockDriveDownloadDriveExportTask = f
}

// UnMockDriveDownloadDriveExportTask un-mock DriveDownloadDriveExportTask method
func (r *Mock) UnMockDriveDownloadDriveExportTask() {
	r.mockDriveDownloadDriveExportTask = nil
}

// DownloadDriveExportTaskReq ...
type DownloadDriveExportTaskReq struct {
	FileToken string `path:"file_token" json:"-"` // 导出文档 token, 示例值: "boxcnNAlfwHxxxxxxxxxxSaLSec"
}

// downloadDriveExportTaskResp ...
type downloadDriveExportTaskResp struct {
	Code int64                        `json:"code,omitempty"`
	Msg  string                       `json:"msg,omitempty"`
	Data *DownloadDriveExportTaskResp `json:"data,omitempty"`
}

func (r *downloadDriveExportTaskResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &DownloadDriveExportTaskResp{}
	}
	r.Data.File = file
}

// DownloadDriveExportTaskResp ...
type DownloadDriveExportTaskResp struct {
	File io.Reader `json:"file,omitempty"`
}
