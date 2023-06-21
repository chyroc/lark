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

// UpdateVCReserveConfigForm 更新会议室预定表单。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve_config-form/patch
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/scope_config/patch-2
func (r *VCService) UpdateVCReserveConfigForm(ctx context.Context, request *UpdateVCReserveConfigFormReq, options ...MethodOptionFunc) (*UpdateVCReserveConfigFormResp, *Response, error) {
	if r.cli.mock.mockVCUpdateVCReserveConfigForm != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateVCReserveConfigForm mock enable")
		return r.cli.mock.mockVCUpdateVCReserveConfigForm(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "UpdateVCReserveConfigForm",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/reserve_configs/:reserve_config_id/form",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateVCReserveConfigFormResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCUpdateVCReserveConfigForm mock VCUpdateVCReserveConfigForm method
func (r *Mock) MockVCUpdateVCReserveConfigForm(f func(ctx context.Context, request *UpdateVCReserveConfigFormReq, options ...MethodOptionFunc) (*UpdateVCReserveConfigFormResp, *Response, error)) {
	r.mockVCUpdateVCReserveConfigForm = f
}

// UnMockVCUpdateVCReserveConfigForm un-mock VCUpdateVCReserveConfigForm method
func (r *Mock) UnMockVCUpdateVCReserveConfigForm() {
	r.mockVCUpdateVCReserveConfigForm = nil
}

// UpdateVCReserveConfigFormReq ...
type UpdateVCReserveConfigFormReq struct {
	ReserveConfigID   string                                         `path:"reserve_config_id" json:"-"`    // 会议室或层级ID, 示例值: "omm_3c5dd7e09bac0c1758fcf9511bd1a771"
	UserIDType        *IDType                                        `query:"user_id_type" json:"-"`        // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ScopeType         int64                                          `json:"scope_type,omitempty"`          // 1代表层级, 2代表会议室, 示例值: 2, 取值范围: `1` ～ `2`
	ReserveFormConfig *UpdateVCReserveConfigFormReqReserveFormConfig `json:"reserve_form_config,omitempty"` // 预定表单设置
}

// UpdateVCReserveConfigFormReqReserveFormConfig ...
type UpdateVCReserveConfigFormReqReserveFormConfig struct {
	ReserveForm   bool                                                         `json:"reserve_form,omitempty"`   // 预定表单开关, true表示打开, false表示关闭, 示例值: false, 默认值: `false`
	NotifiedUsers []*UpdateVCReserveConfigFormReqReserveFormConfigNotifiedUser `json:"notified_users,omitempty"` // 通知人列表
	NotifiedTime  *int64                                                       `json:"notified_time,omitempty"`  // 最晚于会议开始前 notified_time收到通知（单位: 分/时/天）, 示例值: 3
	TimeUnit      *int64                                                       `json:"time_unit,omitempty"`      // 时间单位, 1为分钟；2为小时；3为天, 默认为天, 示例值: 3, 取值范围: `1` ～ `3`
}

// UpdateVCReserveConfigFormReqReserveFormConfigNotifiedUser ...
type UpdateVCReserveConfigFormReqReserveFormConfigNotifiedUser struct {
	UserID string `json:"user_id,omitempty"` // 预定管理员ID, 示例值: "ou_a27b07a9071d90577c0177bcec98f856"
}

// UpdateVCReserveConfigFormResp ...
type UpdateVCReserveConfigFormResp struct {
}

// updateVCReserveConfigFormResp ...
type updateVCReserveConfigFormResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *UpdateVCReserveConfigFormResp `json:"data,omitempty"`
}
