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

// DeleteDriveFile 删除用户在云空间内的文件或者文件夹。文件或者文件夹被删除后, 会进入用户回收站里。
//
// 要删除文件需要确保应用具有下述两种权限之一:
// 1. 该应用是文件所有者并且具有该文件所在父文件夹的编辑权限。
// 2. 该应用并非文件所有者, 但是是该文件所在父文件夹的所有者或者拥有该父文件夹的所有权限（full access）。
// 该接口不支持并发调用, 且调用频率上限为5QPS。删除文件夹会异步执行并返回一个task_id, 可以使用[task_check](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/task_check)接口查询任务执行状态。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/delete
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/file/delete
func (r *DriveService) DeleteDriveFile(ctx context.Context, request *DeleteDriveFileReq, options ...MethodOptionFunc) (*DeleteDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteDriveFile != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#DeleteDriveFile mock enable")
		return r.cli.mock.mockDriveDeleteDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteDriveFile",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteDriveFile mock DriveDeleteDriveFile method
func (r *Mock) MockDriveDeleteDriveFile(f func(ctx context.Context, request *DeleteDriveFileReq, options ...MethodOptionFunc) (*DeleteDriveFileResp, *Response, error)) {
	r.mockDriveDeleteDriveFile = f
}

// UnMockDriveDeleteDriveFile un-mock DriveDeleteDriveFile method
func (r *Mock) UnMockDriveDeleteDriveFile() {
	r.mockDriveDeleteDriveFile = nil
}

// DeleteDriveFileReq ...
type DeleteDriveFileReq struct {
	FileToken string `path:"file_token" json:"-"` // 需要删除的文件token, 示例值: "boxcnrHpsg1QDqXAAAyachabcef"
	Type      string `query:"type" json:"-"`      // 被删除文件的类型, 示例值: file, 可选值有: file: 文件类型, docx: 新版文档类型, bitable: 多维表格类型, folder: 文件夹类型, doc: 文档类型, sheet: 电子表格类型, mindnote: 思维笔记类型, shortcut: 快捷方式类型, slides: 幻灯片
}

// DeleteDriveFileResp ...
type DeleteDriveFileResp struct {
	TaskID string `json:"task_id,omitempty"` // 异步任务id, 删除文件夹时返回
}

// deleteDriveFileResp ...
type deleteDriveFileResp struct {
	Code  int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string               `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteDriveFileResp `json:"data,omitempty"`
	Error *ErrorDetail         `json:"error,omitempty"`
}
