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

// PartUploadDriveFile 上传对应的文件块。
//
// 该接口不支持太高的并发, 且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_part
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/upload/multipart-upload-file-/upload_part
func (r *DriveService) PartUploadDriveFile(ctx context.Context, request *PartUploadDriveFileReq, options ...MethodOptionFunc) (*PartUploadDriveFileResp, *Response, error) {
	if r.cli.mock.mockDrivePartUploadDriveFile != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#PartUploadDriveFile mock enable")
		return r.cli.mock.mockDrivePartUploadDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "PartUploadDriveFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/upload_part",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
		IsFile:                true,
	}
	resp := new(partUploadDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDrivePartUploadDriveFile mock DrivePartUploadDriveFile method
func (r *Mock) MockDrivePartUploadDriveFile(f func(ctx context.Context, request *PartUploadDriveFileReq, options ...MethodOptionFunc) (*PartUploadDriveFileResp, *Response, error)) {
	r.mockDrivePartUploadDriveFile = f
}

// UnMockDrivePartUploadDriveFile un-mock DrivePartUploadDriveFile method
func (r *Mock) UnMockDrivePartUploadDriveFile() {
	r.mockDrivePartUploadDriveFile = nil
}

// PartUploadDriveFileReq ...
type PartUploadDriveFileReq struct {
	UploadID string    `json:"upload_id,omitempty"` // 分片上传事务ID, 示例值: "7111211691345512356"
	Seq      int64     `json:"seq,omitempty"`       // 块号, 从0开始计数, 示例值: 0
	Size     int64     `json:"size,omitempty"`      // 块大小（以字节为单位）, 示例值: 4194304
	Checksum *string   `json:"checksum,omitempty"`  // 文件分块adler32校验和(可选), 示例值: "12342388237783212356"
	File     io.Reader `json:"file,omitempty"`      // 文件分片二进制内容, 示例值: file binary
}

// PartUploadDriveFileResp ...
type PartUploadDriveFileResp struct {
}

// partUploadDriveFileResp ...
type partUploadDriveFileResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *PartUploadDriveFileResp `json:"data,omitempty"`
}
