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

// CreateCoreHRDepartment 创建部门。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/department/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/department/create
func (r *CoreHRService) CreateCoreHRDepartment(ctx context.Context, request *CreateCoreHRDepartmentReq, options ...MethodOptionFunc) (*CreateCoreHRDepartmentResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRDepartment != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRDepartment mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRDepartment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRDepartment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/departments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRDepartmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRDepartment mock CoreHRCreateCoreHRDepartment method
func (r *Mock) MockCoreHRCreateCoreHRDepartment(f func(ctx context.Context, request *CreateCoreHRDepartmentReq, options ...MethodOptionFunc) (*CreateCoreHRDepartmentResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRDepartment = f
}

// UnMockCoreHRCreateCoreHRDepartment un-mock CoreHRCreateCoreHRDepartment method
func (r *Mock) UnMockCoreHRCreateCoreHRDepartment() {
	r.mockCoreHRCreateCoreHRDepartment = nil
}

// CreateCoreHRDepartmentReq ...
type CreateCoreHRDepartmentReq struct {
	ClientToken      *string                                    `query:"client_token" json:"-"`       // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	UserIDType       *IDType                                    `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: people_corehr_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType                          `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 类型, 示例值: people_corehr_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `people_corehr_department_id`
	SubType          *CreateCoreHRDepartmentReqSubType          `json:"sub_type,omitempty"`           // 子类型
	Manager          *string                                    `json:"manager,omitempty"`            // 部门负责人, 示例值: "6893013238632416776"
	IsConfidential   bool                                       `json:"is_confidential,omitempty"`    // 是否保密, 示例值: true
	HiberarchyCommon *CreateCoreHRDepartmentReqHiberarchyCommon `json:"hiberarchy_common,omitempty"`  // 层级关系, 内层字段见实体
	EffectiveTime    string                                     `json:"effective_time,omitempty"`     // 生效时间, 示例值: "2020-05-01 00:00:00"
	ExpirationTime   *string                                    `json:"expiration_time,omitempty"`    // 失效时间, 示例值: "2020-05-02 00:00:00"
	CustomFields     []*CreateCoreHRDepartmentReqCustomField    `json:"custom_fields,omitempty"`      // 自定义字段
	CostCenterID     *string                                    `json:"cost_center_id,omitempty"`     // 成本中心id, 示例值: "7142384817131652652"
}

// CreateCoreHRDepartmentReqCustomField ...
type CreateCoreHRDepartmentReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// CreateCoreHRDepartmentReqHiberarchyCommon ...
type CreateCoreHRDepartmentReqHiberarchyCommon struct {
	ParentID       *string                                                 `json:"parent_id,omitempty"`       // 上级组织 ID, 示例值: "4719168654814483759"
	Name           []*CreateCoreHRDepartmentReqHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *CreateCoreHRDepartmentReqHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得
	Active         bool                                                    `json:"active,omitempty"`          // 是否启用, 示例值: true
	ExpirationTime *string                                                 `json:"expiration_time,omitempty"` // 失效时间, 示例值: "2020-05-02 00:00:00"
	Code           *string                                                 `json:"code,omitempty"`            // 编码, 示例值: "12456"
	Description    []*CreateCoreHRDepartmentReqHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	CustomFields   []*CreateCoreHRDepartmentReqHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// CreateCoreHRDepartmentReqHiberarchyCommonCustomField ...
type CreateCoreHRDepartmentReqHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// CreateCoreHRDepartmentReqHiberarchyCommonDescription ...
type CreateCoreHRDepartmentReqHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// CreateCoreHRDepartmentReqHiberarchyCommonName ...
type CreateCoreHRDepartmentReqHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// CreateCoreHRDepartmentReqHiberarchyCommonType ...
type CreateCoreHRDepartmentReqHiberarchyCommonType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRDepartmentReqSubType ...
type CreateCoreHRDepartmentReqSubType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRDepartmentResp ...
type CreateCoreHRDepartmentResp struct {
	Department *CreateCoreHRDepartmentRespDepartment `json:"department,omitempty"` // 创建成功的部门信息
}

// CreateCoreHRDepartmentRespDepartment ...
type CreateCoreHRDepartmentRespDepartment struct {
	ID               string                                                `json:"id,omitempty"`                // 实体在CoreHR内部的唯一键
	SubType          *CreateCoreHRDepartmentRespDepartmentSubType          `json:"sub_type,omitempty"`          // 子类型
	Manager          string                                                `json:"manager,omitempty"`           // 部门负责人
	IsConfidential   bool                                                  `json:"is_confidential,omitempty"`   // 是否保密
	HiberarchyCommon *CreateCoreHRDepartmentRespDepartmentHiberarchyCommon `json:"hiberarchy_common,omitempty"` // 层级关系, 内层字段见实体
	EffectiveTime    string                                                `json:"effective_time,omitempty"`    // 生效时间
	ExpirationTime   string                                                `json:"expiration_time,omitempty"`   // 失效时间
	CustomFields     []*CreateCoreHRDepartmentRespDepartmentCustomField    `json:"custom_fields,omitempty"`     // 自定义字段
	CostCenterID     string                                                `json:"cost_center_id,omitempty"`    // 成本中心id
}

// CreateCoreHRDepartmentRespDepartmentCustomField ...
type CreateCoreHRDepartmentRespDepartmentCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHRDepartmentRespDepartmentHiberarchyCommon ...
type CreateCoreHRDepartmentRespDepartmentHiberarchyCommon struct {
	ParentID       string                                                             `json:"parent_id,omitempty"`       // 上级组织 ID
	Name           []*CreateCoreHRDepartmentRespDepartmentHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *CreateCoreHRDepartmentRespDepartmentHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得
	Active         bool                                                               `json:"active,omitempty"`          // 是否启用
	EffectiveTime  string                                                             `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                             `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                             `json:"code,omitempty"`            // 编码
	Description    []*CreateCoreHRDepartmentRespDepartmentHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                             `json:"tree_order,omitempty"`      // 树形排序, 代表同层级的部门排序序号
	ListOrder      string                                                             `json:"list_order,omitempty"`      // 列表排序, 代表所有部门的混排序号
	CustomFields   []*CreateCoreHRDepartmentRespDepartmentHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// CreateCoreHRDepartmentRespDepartmentHiberarchyCommonCustomField ...
type CreateCoreHRDepartmentRespDepartmentHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHRDepartmentRespDepartmentHiberarchyCommonDescription ...
type CreateCoreHRDepartmentRespDepartmentHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHRDepartmentRespDepartmentHiberarchyCommonName ...
type CreateCoreHRDepartmentRespDepartmentHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHRDepartmentRespDepartmentHiberarchyCommonType ...
type CreateCoreHRDepartmentRespDepartmentHiberarchyCommonType struct {
	EnumName string                                                             `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRDepartmentRespDepartmentHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRDepartmentRespDepartmentHiberarchyCommonTypeDisplay ...
type CreateCoreHRDepartmentRespDepartmentHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHRDepartmentRespDepartmentSubType ...
type CreateCoreHRDepartmentRespDepartmentSubType struct {
	EnumName string                                                `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRDepartmentRespDepartmentSubTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRDepartmentRespDepartmentSubTypeDisplay ...
type CreateCoreHRDepartmentRespDepartmentSubTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// createCoreHRDepartmentResp ...
type createCoreHRDepartmentResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CreateCoreHRDepartmentResp `json:"data,omitempty"`
}
