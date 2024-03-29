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

// GetPerformanceStageTaskByUser 获取指定周期的指定用户的任务信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/performance-v1/stage_task/find_by_user_list
func (r *PerformanceService) GetPerformanceStageTaskByUser(ctx context.Context, request *GetPerformanceStageTaskByUserReq, options ...MethodOptionFunc) (*GetPerformanceStageTaskByUserResp, *Response, error) {
	if r.cli.mock.mockPerformanceGetPerformanceStageTaskByUser != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Performance#GetPerformanceStageTaskByUser mock enable")
		return r.cli.mock.mockPerformanceGetPerformanceStageTaskByUser(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Performance",
		API:                   "GetPerformanceStageTaskByUser",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/performance/v1/stage_tasks/find_by_user_list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getPerformanceStageTaskByUserResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockPerformanceGetPerformanceStageTaskByUser mock PerformanceGetPerformanceStageTaskByUser method
func (r *Mock) MockPerformanceGetPerformanceStageTaskByUser(f func(ctx context.Context, request *GetPerformanceStageTaskByUserReq, options ...MethodOptionFunc) (*GetPerformanceStageTaskByUserResp, *Response, error)) {
	r.mockPerformanceGetPerformanceStageTaskByUser = f
}

// UnMockPerformanceGetPerformanceStageTaskByUser un-mock PerformanceGetPerformanceStageTaskByUser method
func (r *Mock) UnMockPerformanceGetPerformanceStageTaskByUser() {
	r.mockPerformanceGetPerformanceStageTaskByUser = nil
}

// GetPerformanceStageTaskByUserReq ...
type GetPerformanceStageTaskByUserReq struct {
	UserIDType      *IDType  `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	SemesterID      string   `json:"semester_id,omitempty"`       // 周期 ID。1 次只允许查询 1 个周期。可以通过「获取周期」接口获得周期 ID, 示例值: "7033710017401751071"
	UserIDLists     []string `json:"user_id_lists,omitempty"`     // 用户 ID 列表, 如果以用户身份访问, 该值仅能填写本人 ID, 示例值: ["bega29ca"], 最大长度: `50`
	TaskOptionLists []int64  `json:"task_option_lists,omitempty"` // 任务分类（不传默认包含所有）, 1. 待完成, 2. 已完成, 3. 已逾期（仅当租户设置不允许逾期提交时才有此分类）, 示例值: [1, 2, 3], 最大长度: `3`
	AfterTime       *string  `json:"after_time,omitempty"`        // 查询在此时间之后截止的环节, 示例值: "1630425599999"
	BeforeTime      *string  `json:"before_time,omitempty"`       // 查询在此时间之前截止的环节, 示例值: "1630425599999"
}

// GetPerformanceStageTaskByUserResp ...
type GetPerformanceStageTaskByUserResp struct {
	Base  *GetPerformanceStageTaskByUserRespBase   `json:"base,omitempty"`  // 周期基本信息
	Items []*GetPerformanceStageTaskByUserRespItem `json:"items,omitempty"` // 周期任务
}

// GetPerformanceStageTaskByUserRespBase ...
type GetPerformanceStageTaskByUserRespBase struct {
	SemesterID   string                                             `json:"semester_id,omitempty"`   // 周期 ID
	SemesterName *GetPerformanceStageTaskByUserRespBaseSemesterName `json:"semester_name,omitempty"` // 周期名称
	StartTime    string                                             `json:"start_time,omitempty"`    // 开始时间
	EndTime      string                                             `json:"end_time,omitempty"`      // 结束时间
}

// GetPerformanceStageTaskByUserRespBaseSemesterName ...
type GetPerformanceStageTaskByUserRespBaseSemesterName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// GetPerformanceStageTaskByUserRespItem ...
type GetPerformanceStageTaskByUserRespItem struct {
	UserID             string                                                    `json:"user_id,omitempty"`               // 用户 ID
	StageNumLists      []*GetPerformanceStageTaskByUserRespItemStageNumList      `json:"stage_num_lists,omitempty"`       // 各分类的任务数
	StageTaskInfoLists []*GetPerformanceStageTaskByUserRespItemStageTaskInfoList `json:"stage_task_info_lists,omitempty"` // 环节任务信息
}

// GetPerformanceStageTaskByUserRespItemStageNumList ...
type GetPerformanceStageTaskByUserRespItemStageNumList struct {
	TaskOptionID int64 `json:"task_option_id,omitempty"` // 任务分类, 可选值有: 1: 待完成, 2: 已完成, 3: 已逾期（此分类仅在租户系统设置为不允许逾期提交时存在）
	StageNum     int64 `json:"stage_num,omitempty"`      // 环节任务数量
}

// GetPerformanceStageTaskByUserRespItemStageTaskInfoList ...
type GetPerformanceStageTaskByUserRespItemStageTaskInfoList struct {
	StageID         string                                                      `json:"stage_id,omitempty"`          // 环节 ID
	Name            *GetPerformanceStageTaskByUserRespItemStageTaskInfoListName `json:"name,omitempty"`              // 环节名称
	Deadline        string                                                      `json:"deadline,omitempty"`          // 环节截止时间
	NeedTodoCount   int64                                                       `json:"need_todo_count,omitempty"`   // 未完成的任务数量
	JumpURL         string                                                      `json:"jump_url,omitempty"`          // 处理任务的系统页面链接
	StageTaskStatus string                                                      `json:"stage_task_status,omitempty"` // 环节任务状态, 可选值有: need_todo: 还有待完成的任务, overdue: 剩余未完成的任务均已逾期, all_done: 全部任务均已完成, stage_pause: 环节被暂停
	TaskOptionID    int64                                                       `json:"task_option_id,omitempty"`    // 任务分类, 可选值有: 1: 待完成, 2: 已完成, 3: 已逾期（此分类仅在租户系统设置为不允许逾期提交时存在）
}

// GetPerformanceStageTaskByUserRespItemStageTaskInfoListName ...
type GetPerformanceStageTaskByUserRespItemStageTaskInfoListName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// getPerformanceStageTaskByUserResp ...
type getPerformanceStageTaskByUserResp struct {
	Code  int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                             `json:"msg,omitempty"`  // 错误描述
	Data  *GetPerformanceStageTaskByUserResp `json:"data,omitempty"`
	Error *ErrorDetail                       `json:"error,omitempty"`
}
