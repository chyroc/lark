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

// GetCoreHrLeaveBalanceList 批量获取员工各个假期的余额数据。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/leave/leave_balances
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/leave/leave_balances
func (r *CoreHrService) GetCoreHrLeaveBalanceList(ctx context.Context, request *GetCoreHrLeaveBalanceListReq, options ...MethodOptionFunc) (*GetCoreHrLeaveBalanceListResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrLeaveBalanceList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrLeaveBalanceList mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrLeaveBalanceList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrLeaveBalanceList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/leaves/leave_balances",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrLeaveBalanceListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrLeaveBalanceList mock CoreHrGetCoreHrLeaveBalanceList method
func (r *Mock) MockCoreHrGetCoreHrLeaveBalanceList(f func(ctx context.Context, request *GetCoreHrLeaveBalanceListReq, options ...MethodOptionFunc) (*GetCoreHrLeaveBalanceListResp, *Response, error)) {
	r.mockCoreHrGetCoreHrLeaveBalanceList = f
}

// UnMockCoreHrGetCoreHrLeaveBalanceList un-mock CoreHrGetCoreHrLeaveBalanceList method
func (r *Mock) UnMockCoreHrGetCoreHrLeaveBalanceList() {
	r.mockCoreHrGetCoreHrLeaveBalanceList = nil
}

// GetCoreHrLeaveBalanceListReq ...
type GetCoreHrLeaveBalanceListReq struct {
	PageToken        *string  `query:"page_token" json:"-"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
	PageSize         int64    `query:"page_size" json:"-"`          // 分页大小, 示例值: 20
	AsOfDate         *string  `query:"as_of_date" json:"-"`         // 查询截止日期, 即截止到某天余额数据的日期（不传则默认为当天）, 示例值: 2022-07-29
	EmploymentIDList []string `query:"employment_id_list" json:"-"` // 员工 ID 列表, 最大 100 个（不传则默认查询全部员工）, 示例值: 6919733291281024526
	UserIDType       *IDType  `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetCoreHrLeaveBalanceListResp ...
type GetCoreHrLeaveBalanceListResp struct {
	EmploymentLeaveBalanceList []*GetCoreHrLeaveBalanceListRespEmploymentLeaveBalance `json:"employment_leave_balance_list,omitempty"` // 员工假期余额信息列表
	HasMore                    bool                                                   `json:"has_more,omitempty"`                      // 是否还有更多项
	PageToken                  string                                                 `json:"page_token,omitempty"`                    // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetCoreHrLeaveBalanceListRespEmploymentLeaveBalance ...
type GetCoreHrLeaveBalanceListRespEmploymentLeaveBalance struct {
	EmploymentID     string                                                               `json:"employment_id,omitempty"`      // 雇佣信息ID
	EmploymentName   []*GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceEmploymentName `json:"employment_name,omitempty"`    // 员工姓名
	AsOfDate         string                                                               `json:"as_of_date,omitempty"`         // 截止日期, 即查询截止到某天余额数据的日期
	LeaveBalanceList []*GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceLeaveBalance   `json:"leave_balance_list,omitempty"` // 假期余额列表
}

// GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceEmploymentName ...
type GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceEmploymentName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceLeaveBalance ...
type GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceLeaveBalance struct {
	LeaveTypeID          string                                                                          `json:"leave_type_id,omitempty"`          // 假期类型ID
	LeaveTypeName        []*GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceLeaveBalanceLeaveTypeName `json:"leave_type_name,omitempty"`        // 假期类型名称
	HistoricalCyclesLeft string                                                                          `json:"historical_cycles_left,omitempty"` // 结转的历史周期授予时长
	ThisCycleTotal       int64                                                                           `json:"this_cycle_total,omitempty"`       // 本周期授予时长
	ThisCycleTaken       string                                                                          `json:"this_cycle_taken,omitempty"`       // 本周期已休时长
	LeaveBalance         string                                                                          `json:"leave_balance,omitempty"`          // 假期余额
	LeaveDurationUnit    int64                                                                           `json:"leave_duration_unit,omitempty"`    // 假期时长单位, 可选值有: 1: 天, 2: 小时
}

// GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceLeaveBalanceLeaveTypeName ...
type GetCoreHrLeaveBalanceListRespEmploymentLeaveBalanceLeaveBalanceLeaveTypeName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHrLeaveBalanceListResp ...
type getCoreHrLeaveBalanceListResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrLeaveBalanceListResp `json:"data,omitempty"`
}
