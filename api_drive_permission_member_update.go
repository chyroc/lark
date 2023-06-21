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

// UpdateDriveMemberPermission 该接口用于根据 filetoken 更新文档协作者的权限。
//
// 该接口要求文档协作者已存在, 如还未对文档协作者授权请先调用[「增加权限」 ](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/create)接口进行授权。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/update
// new doc: https://open.feishu.cn/document/server-docs/docs/permission/permission-member/update
func (r *DriveService) UpdateDriveMemberPermission(ctx context.Context, request *UpdateDriveMemberPermissionReq, options ...MethodOptionFunc) (*UpdateDriveMemberPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDriveMemberPermission != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDriveMemberPermission mock enable")
		return r.cli.mock.mockDriveUpdateDriveMemberPermission(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDriveMemberPermission",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/permissions/:token/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDriveMemberPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateDriveMemberPermission mock DriveUpdateDriveMemberPermission method
func (r *Mock) MockDriveUpdateDriveMemberPermission(f func(ctx context.Context, request *UpdateDriveMemberPermissionReq, options ...MethodOptionFunc) (*UpdateDriveMemberPermissionResp, *Response, error)) {
	r.mockDriveUpdateDriveMemberPermission = f
}

// UnMockDriveUpdateDriveMemberPermission un-mock DriveUpdateDriveMemberPermission method
func (r *Mock) UnMockDriveUpdateDriveMemberPermission() {
	r.mockDriveUpdateDriveMemberPermission = nil
}

// UpdateDriveMemberPermissionReq ...
type UpdateDriveMemberPermissionReq struct {
	Token            string `path:"token" json:"-"`              // 文件的 token, 获取方式见 [如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "doccnBKgoMyY5OMbUG6FioTXuBe"
	MemberID         string `path:"member_id" json:"-"`          // 协作者 ID, 与协作者 ID 类型需要对应, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	NeedNotification *bool  `query:"need_notification" json:"-"` // 更新权限后是否通知对方, 注意: 使用`tenant_access_token`访问不支持该参数, 示例值: false, 默认值: `false`
	Type             string `query:"type" json:"-"`              // 文件类型, 需要与文件的 token 相匹配, 示例值: "doc", 可选值有: doc: 文档, sheet: 电子表格, file: 云空间文件, wiki: 知识库节点, bitable: 多维表格, docx: 新版文档, mindnote: 思维笔记, minutes: 妙记
	MemberType       string `json:"member_type,omitempty"`       // 协作者 ID 类型, 与协作者 ID 需要对应, 示例值: "openid", 可选值有: email: 飞书邮箱, openid: 开放平台ID, openchat: 开放平台群组ID, opendepartmentid: 开放平台部门ID, userid: 用户自定义ID
	Perm             string `json:"perm,omitempty"`              // 协作者对应的权限角色, 注意: 妙记还不支持可管理角色, 示例值: "view", 可选值有: view: 可阅读角色, edit: 可编辑角色, full_access: 可管理角色
}

// UpdateDriveMemberPermissionResp ...
type UpdateDriveMemberPermissionResp struct {
	Member *UpdateDriveMemberPermissionRespMember `json:"member,omitempty"` // 本次更新权限的用户信息
}

// UpdateDriveMemberPermissionRespMember ...
type UpdateDriveMemberPermissionRespMember struct {
	MemberType string `json:"member_type,omitempty"` // 协作者 ID 类型, 与协作者 ID 需要对应, 可选值有: email: 飞书邮箱, openid: 开放平台ID, openchat: 开放平台群组ID, opendepartmentid: 开放平台部门ID, userid: 用户自定义ID
	MemberID   string `json:"member_id,omitempty"`   // 协作者 ID, 与协作者 ID 类型需要对应
	Perm       string `json:"perm,omitempty"`        // 协作者对应的权限角色, 注意: 妙记还不支持可管理角色, 可选值有: view: 可阅读角色, edit: 可编辑角色, full_access: 可管理角色
}

// updateDriveMemberPermissionResp ...
type updateDriveMemberPermissionResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *UpdateDriveMemberPermissionResp `json:"data,omitempty"`
}
