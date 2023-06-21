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

// UpdateDriveMemberPermissionOld 该接口用于根据 filetoken 更新文档协作者的权限。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ucTN3UjL3UzN14yN1cTN
// new doc: https://open.feishu.cn/document
func (r *DriveService) UpdateDriveMemberPermissionOld(ctx context.Context, request *UpdateDriveMemberPermissionOldReq, options ...MethodOptionFunc) (*UpdateDriveMemberPermissionOldResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDriveMemberPermissionOld != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDriveMemberPermissionOld mock enable")
		return r.cli.mock.mockDriveUpdateDriveMemberPermissionOld(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDriveMemberPermissionOld",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/permission/member/update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDriveMemberPermissionOldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateDriveMemberPermissionOld mock DriveUpdateDriveMemberPermissionOld method
func (r *Mock) MockDriveUpdateDriveMemberPermissionOld(f func(ctx context.Context, request *UpdateDriveMemberPermissionOldReq, options ...MethodOptionFunc) (*UpdateDriveMemberPermissionOldResp, *Response, error)) {
	r.mockDriveUpdateDriveMemberPermissionOld = f
}

// UnMockDriveUpdateDriveMemberPermissionOld un-mock DriveUpdateDriveMemberPermissionOld method
func (r *Mock) UnMockDriveUpdateDriveMemberPermissionOld() {
	r.mockDriveUpdateDriveMemberPermissionOld = nil
}

// UpdateDriveMemberPermissionOldReq ...
type UpdateDriveMemberPermissionOldReq struct {
	Token      string `json:"token,omitempty"`       // 文件的 token, 获取方式见 [对接前说明](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type       string `json:"type,omitempty"`        // 文档类型  "doc"  or  "sheet" or "file"
	MemberType string `json:"member_type,omitempty"` // 用户类型, 可选 "openid"、"openchat"、"userid"
	MemberID   string `json:"member_id,omitempty"`   // 用户类型下的值
	Perm       string `json:"perm,omitempty"`        // 权限, "view" or "edit"
	NotifyLark bool   `json:"notify_lark,omitempty"` // 修改权限后是否飞书/lark通知对方 true 通知 or false 不通知
}

// UpdateDriveMemberPermissionOldResp ...
type UpdateDriveMemberPermissionOldResp struct {
	IsSuccess bool `json:"is_success,omitempty"` // 是否操作成功
}

// updateDriveMemberPermissionOldResp ...
type updateDriveMemberPermissionOldResp struct {
	Code int64                               `json:"code,omitempty"`
	Msg  string                              `json:"msg,omitempty"`
	Data *UpdateDriveMemberPermissionOldResp `json:"data,omitempty"`
}
