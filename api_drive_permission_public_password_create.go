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

// CreateDrivePermissionPublicPassword 该接口用于根据 filetoken 开启云文档的密码。
//
// 注意: 开启密码, 需要先通过”云文档“-“权限”-“设置”-“更新云文档权限设置”的接口更新元文档为互联网上获得链接的任何人可阅读/编辑
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-public-password/create
// new doc: https://open.feishu.cn/document/server-docs/docs/permission/permission-public/permission-public-password/create
func (r *DriveService) CreateDrivePermissionPublicPassword(ctx context.Context, request *CreateDrivePermissionPublicPasswordReq, options ...MethodOptionFunc) (*CreateDrivePermissionPublicPasswordResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDrivePermissionPublicPassword != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#CreateDrivePermissionPublicPassword mock enable")
		return r.cli.mock.mockDriveCreateDrivePermissionPublicPassword(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDrivePermissionPublicPassword",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/permissions/:token/public/password",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDrivePermissionPublicPasswordResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDrivePermissionPublicPassword mock DriveCreateDrivePermissionPublicPassword method
func (r *Mock) MockDriveCreateDrivePermissionPublicPassword(f func(ctx context.Context, request *CreateDrivePermissionPublicPasswordReq, options ...MethodOptionFunc) (*CreateDrivePermissionPublicPasswordResp, *Response, error)) {
	r.mockDriveCreateDrivePermissionPublicPassword = f
}

// UnMockDriveCreateDrivePermissionPublicPassword un-mock DriveCreateDrivePermissionPublicPassword method
func (r *Mock) UnMockDriveCreateDrivePermissionPublicPassword() {
	r.mockDriveCreateDrivePermissionPublicPassword = nil
}

// CreateDrivePermissionPublicPasswordReq ...
type CreateDrivePermissionPublicPasswordReq struct {
	Token string `path:"token" json:"-"` // 文件的 token, 获取方式见 [如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doccnBKgoMyY5OMbUG6FioTXuBe"
	Type  string `query:"type" json:"-"` // 文件类型, 需要与文件的 token 相匹配, 示例值: doc, 可选值有: doc: 文档, sheet: 电子表格, file: 云空间文件, wiki: 知识库节点, bitable: 多维表格, docx: 新版文档, mindnote: 思维笔记, minutes: 妙计（暂不支持）
}

// CreateDrivePermissionPublicPasswordResp ...
type CreateDrivePermissionPublicPasswordResp struct {
	Password string `json:"password,omitempty"` // 密码
}

// createDrivePermissionPublicPasswordResp ...
type createDrivePermissionPublicPasswordResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *CreateDrivePermissionPublicPasswordResp `json:"data,omitempty"`
}
