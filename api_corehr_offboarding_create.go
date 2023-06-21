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

// CreateCoreHrOffboarding 操作员工直接离职。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/offboarding/submit
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/offboarding/submit
func (r *CoreHrService) CreateCoreHrOffboarding(ctx context.Context, request *CreateCoreHrOffboardingReq, options ...MethodOptionFunc) (*CreateCoreHrOffboardingResp, *Response, error) {
	if r.cli.mock.mockCoreHrCreateCoreHrOffboarding != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#CreateCoreHrOffboarding mock enable")
		return r.cli.mock.mockCoreHrCreateCoreHrOffboarding(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "CreateCoreHrOffboarding",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/offboardings/submit",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHrOffboardingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrCreateCoreHrOffboarding mock CoreHrCreateCoreHrOffboarding method
func (r *Mock) MockCoreHrCreateCoreHrOffboarding(f func(ctx context.Context, request *CreateCoreHrOffboardingReq, options ...MethodOptionFunc) (*CreateCoreHrOffboardingResp, *Response, error)) {
	r.mockCoreHrCreateCoreHrOffboarding = f
}

// UnMockCoreHrCreateCoreHrOffboarding un-mock CoreHrCreateCoreHrOffboarding method
func (r *Mock) UnMockCoreHrCreateCoreHrOffboarding() {
	r.mockCoreHrCreateCoreHrOffboarding = nil
}

// CreateCoreHrOffboardingReq ...
type CreateCoreHrOffboardingReq struct {
	UserIDType                        *IDType                                  `query:"user_id_type" json:"-"`                         // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	OffboardingMode                   int64                                    `json:"offboarding_mode,omitempty"`                     // 离职方式, 示例值: 1, 可选值有: 1: 直接离职
	EmploymentID                      string                                   `json:"employment_id,omitempty"`                        // 雇员 id, 示例值: "6982509313466189342"
	OffboardingDate                   string                                   `json:"offboarding_date,omitempty"`                     // 离职日期, 示例值: "2022-05-18"
	OffboardingReasonUniqueIdentifier string                                   `json:"offboarding_reason_unique_identifier,omitempty"` // 离职原因, 可通过接口, [【查询员工离职原因列表】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/offboarding/query)获取, 示例值: "reason_for_offboarding_option8"
	OffboardingReasonExplanation      *string                                  `json:"offboarding_reason_explanation,omitempty"`       // 离职原因说明, 长度限制6000, 示例值: "离职原因说明"
	CustomFields                      []*CreateCoreHrOffboardingReqCustomField `json:"custom_fields,omitempty"`                        // 自定义字段
}

// CreateCoreHrOffboardingReqCustomField ...
type CreateCoreHrOffboardingReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// CreateCoreHrOffboardingResp ...
type CreateCoreHrOffboardingResp struct {
	OffboardingID                     string `json:"offboarding_id,omitempty"`                       // 离职记录 id
	EmploymentID                      string `json:"employment_id,omitempty"`                        // 雇员 id
	OffboardingReasonUniqueIdentifier string `json:"offboarding_reason_unique_identifier,omitempty"` // 离职原因
	OffboardingDate                   string `json:"offboarding_date,omitempty"`                     // 离职日期
	OffboardingReasonExplanation      string `json:"offboarding_reason_explanation,omitempty"`       // 离职原因说明
	CreatedTime                       string `json:"created_time,omitempty"`                         // 创建时间
}

// createCoreHrOffboardingResp ...
type createCoreHrOffboardingResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *CreateCoreHrOffboardingResp `json:"data,omitempty"`
}
