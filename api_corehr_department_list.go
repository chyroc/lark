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

// GetCoreHRDepartmentList 批量查询部门。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/department/list
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/department/list
func (r *CoreHRService) GetCoreHRDepartmentList(ctx context.Context, request *GetCoreHRDepartmentListReq, options ...MethodOptionFunc) (*GetCoreHRDepartmentListResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRDepartmentList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRDepartmentList mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRDepartmentList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRDepartmentList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/departments",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRDepartmentListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRDepartmentList mock CoreHRGetCoreHRDepartmentList method
func (r *Mock) MockCoreHRGetCoreHRDepartmentList(f func(ctx context.Context, request *GetCoreHRDepartmentListReq, options ...MethodOptionFunc) (*GetCoreHRDepartmentListResp, *Response, error)) {
	r.mockCoreHRGetCoreHRDepartmentList = f
}

// UnMockCoreHRGetCoreHRDepartmentList un-mock CoreHRGetCoreHRDepartmentList method
func (r *Mock) UnMockCoreHRGetCoreHRDepartmentList() {
	r.mockCoreHRGetCoreHRDepartmentList = nil
}

// GetCoreHRDepartmentListReq ...
type GetCoreHRDepartmentListReq struct {
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "6966234786251671053"
	PageSize         int64             `query:"page_size" json:"-"`          // 分页大小, 示例值: 100
	DepartmentIDList []string          `query:"department_id_list" json:"-"` // 部门ID列表, 示例值: 7094136522860922111
	NameList         []string          `query:"name_list" json:"-"`          // 部门名称列表, 需精确匹配, 示例值: 后端研发部
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: people_corehr_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 类型, 示例值: people_corehr_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `people_corehr_department_id`
}

// GetCoreHRDepartmentListResp ...
type GetCoreHRDepartmentListResp struct {
	Items     []*GetCoreHRDepartmentListRespItem `json:"items,omitempty"`      // 查询的部门信息
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCoreHRDepartmentListRespItem ...
type GetCoreHRDepartmentListRespItem struct {
	ID               string                                           `json:"id,omitempty"`                // 部门 ID
	SubType          *GetCoreHRDepartmentListRespItemSubType          `json:"sub_type,omitempty"`          // 部门子类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)部门子类型（department_sub_type）枚举定义部分获得
	Manager          string                                           `json:"manager,omitempty"`           // 部门负责人 ID, 枚举值及详细信息可通过【批量查询雇佣信息】接口查询获得
	IsConfidential   bool                                             `json:"is_confidential,omitempty"`   // 是否保密
	HiberarchyCommon *GetCoreHRDepartmentListRespItemHiberarchyCommon `json:"hiberarchy_common,omitempty"` // 层级关系, 内层字段见实体
	EffectiveTime    string                                           `json:"effective_time,omitempty"`    // 生效时间
	ExpirationTime   string                                           `json:"expiration_time,omitempty"`   // 失效时间
	CustomFields     []*GetCoreHRDepartmentListRespItemCustomField    `json:"custom_fields,omitempty"`     // 自定义字段
	CostCenterID     string                                           `json:"cost_center_id,omitempty"`    // 成本中心id
}

// GetCoreHRDepartmentListRespItemCustomField ...
type GetCoreHRDepartmentListRespItemCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRDepartmentListRespItemHiberarchyCommon ...
type GetCoreHRDepartmentListRespItemHiberarchyCommon struct {
	ParentID       string                                                        `json:"parent_id,omitempty"`       // 上级组织 ID
	Name           []*GetCoreHRDepartmentListRespItemHiberarchyCommonName        `json:"name,omitempty"`            // 名称
	Type           *GetCoreHRDepartmentListRespItemHiberarchyCommonType          `json:"type,omitempty"`            // 组织类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)组织类型（organization_type）枚举定义部分获得
	Active         bool                                                          `json:"active,omitempty"`          // 是否启用
	EffectiveTime  string                                                        `json:"effective_time,omitempty"`  // 生效时间
	ExpirationTime string                                                        `json:"expiration_time,omitempty"` // 失效时间
	Code           string                                                        `json:"code,omitempty"`            // 编码
	Description    []*GetCoreHRDepartmentListRespItemHiberarchyCommonDescription `json:"description,omitempty"`     // 描述
	TreeOrder      string                                                        `json:"tree_order,omitempty"`      // 树形排序, 代表同层级的部门排序序号。新建的部门, 该字段默认为空, 有两种情况会自动写入值: 管理员在部门管理页面上拖动排序；, 定时任务更新该字段为空的数据, 3分钟/次
	ListOrder      string                                                        `json:"list_order,omitempty"`      // 列表排序, 代表所有部门的混排序号。新建的部门, 该字段默认为空, 有两种情况会自动写入值: 管理员在部门管理页面上拖动排序；, 定时任务更新该字段为空的数据, 3分钟/次
	CustomFields   []*GetCoreHRDepartmentListRespItemHiberarchyCommonCustomField `json:"custom_fields,omitempty"`   // 自定义字段
}

// GetCoreHRDepartmentListRespItemHiberarchyCommonCustomField ...
type GetCoreHRDepartmentListRespItemHiberarchyCommonCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRDepartmentListRespItemHiberarchyCommonDescription ...
type GetCoreHRDepartmentListRespItemHiberarchyCommonDescription struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRDepartmentListRespItemHiberarchyCommonName ...
type GetCoreHRDepartmentListRespItemHiberarchyCommonName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRDepartmentListRespItemHiberarchyCommonType ...
type GetCoreHRDepartmentListRespItemHiberarchyCommonType struct {
	EnumName string                                                        `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRDepartmentListRespItemHiberarchyCommonTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRDepartmentListRespItemHiberarchyCommonTypeDisplay ...
type GetCoreHRDepartmentListRespItemHiberarchyCommonTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRDepartmentListRespItemSubType ...
type GetCoreHRDepartmentListRespItemSubType struct {
	EnumName string                                           `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRDepartmentListRespItemSubTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRDepartmentListRespItemSubTypeDisplay ...
type GetCoreHRDepartmentListRespItemSubTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRDepartmentListResp ...
type getCoreHRDepartmentListResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHRDepartmentListResp `json:"data,omitempty"`
}
