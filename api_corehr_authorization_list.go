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

// GetCoreHrAuthorizationList 批量查询「飞书人事」-「权限设置」中的用户授权信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/authorization/batch-query-user-authorization
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/authorization/batch-query-user-authorization
func (r *CoreHrService) GetCoreHrAuthorizationList(ctx context.Context, request *GetCoreHrAuthorizationListReq, options ...MethodOptionFunc) (*GetCoreHrAuthorizationListResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrAuthorizationList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrAuthorizationList mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrAuthorizationList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrAuthorizationList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/authorizations/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrAuthorizationListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrAuthorizationList mock CoreHrGetCoreHrAuthorizationList method
func (r *Mock) MockCoreHrGetCoreHrAuthorizationList(f func(ctx context.Context, request *GetCoreHrAuthorizationListReq, options ...MethodOptionFunc) (*GetCoreHrAuthorizationListResp, *Response, error)) {
	r.mockCoreHrGetCoreHrAuthorizationList = f
}

// UnMockCoreHrGetCoreHrAuthorizationList un-mock CoreHrGetCoreHrAuthorizationList method
func (r *Mock) UnMockCoreHrGetCoreHrAuthorizationList() {
	r.mockCoreHrGetCoreHrAuthorizationList = nil
}

// GetCoreHrAuthorizationListReq ...
type GetCoreHrAuthorizationListReq struct {
	EmploymentIDList []string `query:"employment_id_list" json:"-"` // 雇员ID列表, 最大100个（不传则默认查询全部员工）, 示例值: ["6967639606963471902"]
	RoleIDList       []string `query:"role_id_list" json:"-"`       // 角色 ID 列表, 最多 100 个, 示例值: ["6966577636294264356"]
	PageToken        *string  `query:"page_token" json:"-"`         // 页码标识, 获取第一页传空, 每次查询会返回下一页的page_token, 示例值: "6969864184272078374"
	PageSize         *int64   `query:"page_size" json:"-"`          // 每页获取记录数量, 最大20, 示例值: "20"
	UserIDType       *IDType  `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetCoreHrAuthorizationListResp ...
type GetCoreHrAuthorizationListResp struct {
	HasMore   bool                                  `json:"has_more,omitempty"`   // 是否有下一页
	PageToken string                                `json:"page_token,omitempty"` // 下一页页码
	Items     []*GetCoreHrAuthorizationListRespItem `json:"items,omitempty"`      // 查询的用户授权信息
}

// GetCoreHrAuthorizationListRespItem ...
type GetCoreHrAuthorizationListRespItem struct {
	EmploymentID         string                                                `json:"employment_id,omitempty"`          // 雇员 ID
	PermissionDetailList []*GetCoreHrAuthorizationListRespItemPermissionDetail `json:"permission_detail_list,omitempty"` // 授权列表
}

// GetCoreHrAuthorizationListRespItemPermissionDetail ...
type GetCoreHrAuthorizationListRespItemPermissionDetail struct {
	Role                     *GetCoreHrAuthorizationListRespItemPermissionDetailRole                     `json:"role,omitempty"`                       // 角色
	AssignedOrganizationList [][]*GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganization `json:"assigned_organization_list,omitempty"` // 管理的组织对象列表, 当授权为组织类角色, 返回该数据
	GrantorRuleList          []*GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRule            `json:"grantor_rule_list,omitempty"`          // 授权的数据范围, 当授权非组织类角色, 返回该数据
	UpdateTime               string                                                                      `json:"update_time,omitempty"`                // 更新时间
}

// GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganization ...
type GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganization struct {
	AssignedOrganization *GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganizationAssignedOrganization `json:"assigned_organization,omitempty"` // 被授权的组织信息
}

// GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganizationAssignedOrganization ...
type GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganizationAssignedOrganization struct {
	OrgKey    string                                                                                             `json:"org_key,omitempty"`     // 被授权的组织类型, 当前支持的类型有: department: 部门, work_location: 工作地点, company: 公司
	OrgName   *GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganizationAssignedOrganizationOrgName `json:"org_name,omitempty"`    // 管理对象名称, 例如: 部门、工作城市
	OrgIDList []string                                                                                           `json:"org_id_list,omitempty"` // 管理对象id列表, 例如: [, "6966974245293393415", "6967286219029874189", ]
}

// GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganizationAssignedOrganizationOrgName ...
type GetCoreHrAuthorizationListRespItemPermissionDetailAssignedOrganizationAssignedOrganizationOrgName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRule ...
type GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRule struct {
	RuleDimension *GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRuleRuleDimension `json:"rule_dimension,omitempty"` // 数据实体
	RuleType      int64                                                                       `json:"rule_type,omitempty"`      // 管理类型, 枚举: 0 - 无数据权限, 1 - 全部数据权限, 2 - 被授权的用户自己, 3 - 按规则指定范围
}

// GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRuleRuleDimension ...
type GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRuleRuleDimension struct {
	EntityKey  string                                                                                `json:"entity_key,omitempty"`  // user: 员工, department: 部门, location: 地点, company: 公司 , job_level: 职务级别 , job_family: 职务序列 , job: 职务 , entity-onboarding-pre_hire: 待入职 , offboarding_info: 离职 , signature_template: 文件模板 , signature_file: 电子签文件 , cpst_standard: 薪资标准 , cpst_archive: 员工薪资档案 , cpst_grade: 薪资等级
	EntityName *GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRuleRuleDimensionEntityName `json:"entity_name,omitempty"` // 数据实体名称, 例如: 员工、部门
}

// GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRuleRuleDimensionEntityName ...
type GetCoreHrAuthorizationListRespItemPermissionDetailGrantorRuleRuleDimensionEntityName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetCoreHrAuthorizationListRespItemPermissionDetailRole ...
type GetCoreHrAuthorizationListRespItemPermissionDetailRole struct {
	ID          string                                                             `json:"id,omitempty"`          // 角色ID
	Code        string                                                             `json:"code,omitempty"`        // 角色code, 通常用于与其他系统进行交互
	Name        *GetCoreHrAuthorizationListRespItemPermissionDetailRoleName        `json:"name,omitempty"`        // 角色名称
	Description *GetCoreHrAuthorizationListRespItemPermissionDetailRoleDescription `json:"description,omitempty"` // 角色描述
}

// GetCoreHrAuthorizationListRespItemPermissionDetailRoleDescription ...
type GetCoreHrAuthorizationListRespItemPermissionDetailRoleDescription struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// GetCoreHrAuthorizationListRespItemPermissionDetailRoleName ...
type GetCoreHrAuthorizationListRespItemPermissionDetailRoleName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 中文
	EnUs string `json:"en_us,omitempty"` // 英文
}

// getCoreHrAuthorizationListResp ...
type getCoreHrAuthorizationListResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrAuthorizationListResp `json:"data,omitempty"`
}
