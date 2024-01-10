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

// GetPerformanceStageTaskByPage 可按分页获取周期下所有用户的任务信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/performance-v1/stage_task/find_by_page
func (r *PerformanceService) GetPerformanceStageTaskByPage(ctx context.Context, request *GetPerformanceStageTaskByPageReq, options ...MethodOptionFunc) (*GetPerformanceStageTaskByPageResp, *Response, error) {
	if r.cli.mock.mockPerformanceGetPerformanceStageTaskByPage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Performance#GetPerformanceStageTaskByPage mock enable")
		return r.cli.mock.mockPerformanceGetPerformanceStageTaskByPage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Performance",
		API:                   "GetPerformanceStageTaskByPage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/performance/v1/stage_tasks/find_by_page",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getPerformanceStageTaskByPageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockPerformanceGetPerformanceStageTaskByPage mock PerformanceGetPerformanceStageTaskByPage method
func (r *Mock) MockPerformanceGetPerformanceStageTaskByPage(f func(ctx context.Context, request *GetPerformanceStageTaskByPageReq, options ...MethodOptionFunc) (*GetPerformanceStageTaskByPageResp, *Response, error)) {
	r.mockPerformanceGetPerformanceStageTaskByPage = f
}

// UnMockPerformanceGetPerformanceStageTaskByPage un-mock PerformanceGetPerformanceStageTaskByPage method
func (r *Mock) UnMockPerformanceGetPerformanceStageTaskByPage() {
	r.mockPerformanceGetPerformanceStageTaskByPage = nil
}

// GetPerformanceStageTaskByPageReq ...
type GetPerformanceStageTaskByPageReq struct {
	UserIDType      *IDType `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	SemesterID      string  `json:"semester_id,omitempty"`       // 周期 ID。1 次只允许查询 1 个周期。可以通过「获取周期」接口获得周期 ID, 示例值: "7033710017401751071"
	TaskOptionLists []int64 `json:"task_option_lists,omitempty"` // 任务分类（不传默认包含所有）, 1. 待完成, 2. 已完成, 3. 已逾期（仅当租户设置不允许逾期提交时才有此分类）, 示例值: [1, 2, 3], 最大长度: `3`
	AfterTime       *string `json:"after_time,omitempty"`        // 查询在此时间之后截止的环节, 示例值: "1630425599999"
	BeforeTime      *string `json:"before_time,omitempty"`       // 查询早于当前时间截止的环节, 示例值: "1630425599999"
	PageToken       *string `json:"page_token,omitempty"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw="
	PageSize        *int64  `json:"page_size,omitempty"`         // 分页大小, 示例值: 30, 默认值: `20`, 最大值: `50`
}

// GetPerformanceStageTaskByPageResp ...
type GetPerformanceStageTaskByPageResp struct {
	Base      *GetPerformanceStageTaskByPageRespBase   `json:"base,omitempty"`       // 周期基本信息
	Items     []*GetPerformanceStageTaskByPageRespItem `json:"items,omitempty"`      // 周期任务
	HasMore   bool                                     `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                   `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// GetPerformanceStageTaskByPageRespBase ...
type GetPerformanceStageTaskByPageRespBase struct {
	SemesterID   string                                             `json:"semester_id,omitempty"`   // 周期 ID
	SemesterName *GetPerformanceStageTaskByPageRespBaseSemesterName `json:"semester_name,omitempty"` // 周期名称
	StartTime    string                                             `json:"start_time,omitempty"`    // 开始时间
	EndTime      string                                             `json:"end_time,omitempty"`      // 结束时间
}

// GetPerformanceStageTaskByPageRespBaseSemesterName ...
type GetPerformanceStageTaskByPageRespBaseSemesterName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// GetPerformanceStageTaskByPageRespItem ...
type GetPerformanceStageTaskByPageRespItem struct {
	UserID             string                                                    `json:"user_id,omitempty"`               // 用户 ID
	StageNumLists      []*GetPerformanceStageTaskByPageRespItemStageNumList      `json:"stage_num_lists,omitempty"`       // 各分类的任务数
	StageTaskInfoLists []*GetPerformanceStageTaskByPageRespItemStageTaskInfoList `json:"stage_task_info_lists,omitempty"` // 环节任务信息
}

// GetPerformanceStageTaskByPageRespItemStageNumList ...
type GetPerformanceStageTaskByPageRespItemStageNumList struct {
	TaskOptionID int64 `json:"task_option_id,omitempty"` // 任务分类, 可选值有: 1: 待完成, 2: 已完成, 3: 已逾期（此分类仅在租户系统设置为不允许逾期提交时存在）
	StageNum     int64 `json:"stage_num,omitempty"`      // 环节任务数量
}

// GetPerformanceStageTaskByPageRespItemStageTaskInfoList ...
type GetPerformanceStageTaskByPageRespItemStageTaskInfoList struct {
	StageID         string                                                      `json:"stage_id,omitempty"`          // 环节 ID
	Name            *GetPerformanceStageTaskByPageRespItemStageTaskInfoListName `json:"name,omitempty"`              // 环节名称
	Deadline        string                                                      `json:"deadline,omitempty"`          // 环节截止时间
	NeedTodoCount   int64                                                       `json:"need_todo_count,omitempty"`   // 未完成的任务数量
	JumpURL         string                                                      `json:"jump_url,omitempty"`          // 处理任务的系统页面链接
	StageTaskStatus string                                                      `json:"stage_task_status,omitempty"` // 环节任务状态, 可选值有: need_todo: 还有待完成的任务, overdue: 剩余未完成的任务均已逾期, all_done: 全部任务均已完成, stage_pause: 环节被暂停
	TaskOptionID    int64                                                       `json:"task_option_id,omitempty"`    // 任务分类, 可选值有: 1: 待完成, 2: 已完成, 3: 已逾期（此分类仅在租户系统设置为不允许逾期提交时存在）
}

// GetPerformanceStageTaskByPageRespItemStageTaskInfoListName ...
type GetPerformanceStageTaskByPageRespItemStageTaskInfoListName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// getPerformanceStageTaskByPageResp ...
type getPerformanceStageTaskByPageResp struct {
	Code int64                              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                             `json:"msg,omitempty"`  // 错误描述
	Data *GetPerformanceStageTaskByPageResp `json:"data,omitempty"`
}
