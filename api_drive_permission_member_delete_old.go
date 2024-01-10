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

// DeleteDriveMemberPermissionOld 该接口用于根据 filetoken 移除文档协作者的权限。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYTN3UjL2UzN14iN1cTN
func (r *DriveService) DeleteDriveMemberPermissionOld(ctx context.Context, request *DeleteDriveMemberPermissionOldReq, options ...MethodOptionFunc) (*DeleteDriveMemberPermissionOldResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteDriveMemberPermissionOld != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#DeleteDriveMemberPermissionOld mock enable")
		return r.cli.mock.mockDriveDeleteDriveMemberPermissionOld(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteDriveMemberPermissionOld",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/permission/member/delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteDriveMemberPermissionOldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteDriveMemberPermissionOld mock DriveDeleteDriveMemberPermissionOld method
func (r *Mock) MockDriveDeleteDriveMemberPermissionOld(f func(ctx context.Context, request *DeleteDriveMemberPermissionOldReq, options ...MethodOptionFunc) (*DeleteDriveMemberPermissionOldResp, *Response, error)) {
	r.mockDriveDeleteDriveMemberPermissionOld = f
}

// UnMockDriveDeleteDriveMemberPermissionOld un-mock DriveDeleteDriveMemberPermissionOld method
func (r *Mock) UnMockDriveDeleteDriveMemberPermissionOld() {
	r.mockDriveDeleteDriveMemberPermissionOld = nil
}

// DeleteDriveMemberPermissionOldReq ...
type DeleteDriveMemberPermissionOldReq struct {
	Token      string `json:"token,omitempty"`       // 文件的 token, 获取方式见 [对接前说明](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type       string `json:"type,omitempty"`        // 文档类型 "doc"  or  "sheet" or "bitable"  or "file"
	MemberType string `json:"member_type,omitempty"` // 用户类型, 可选 "openid"、"openchat"、"userid"
	MemberID   string `json:"member_id,omitempty"`   // 用户类型下的值
}

// DeleteDriveMemberPermissionOldResp ...
type DeleteDriveMemberPermissionOldResp struct {
	IsSuccess bool `json:"is_success,omitempty"` // 是否操作成功
}

// deleteDriveMemberPermissionOldResp ...
type deleteDriveMemberPermissionOldResp struct {
	Code int64                               `json:"code,omitempty"`
	Msg  string                              `json:"msg,omitempty"`
	Data *DeleteDriveMemberPermissionOldResp `json:"data,omitempty"`
}
