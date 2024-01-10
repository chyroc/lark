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

// GetApplicationFeedbackList 查询应用的反馈数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-feedback/list
// new doc: https://open.feishu.cn/document/server-docs/application-v6/application-feedback/list
func (r *ApplicationService) GetApplicationFeedbackList(ctx context.Context, request *GetApplicationFeedbackListReq, options ...MethodOptionFunc) (*GetApplicationFeedbackListResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationFeedbackList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#GetApplicationFeedbackList mock enable")
		return r.cli.mock.mockApplicationGetApplicationFeedbackList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationFeedbackList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id/feedbacks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationFeedbackListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationFeedbackList mock ApplicationGetApplicationFeedbackList method
func (r *Mock) MockApplicationGetApplicationFeedbackList(f func(ctx context.Context, request *GetApplicationFeedbackListReq, options ...MethodOptionFunc) (*GetApplicationFeedbackListResp, *Response, error)) {
	r.mockApplicationGetApplicationFeedbackList = f
}

// UnMockApplicationGetApplicationFeedbackList un-mock ApplicationGetApplicationFeedbackList method
func (r *Mock) UnMockApplicationGetApplicationFeedbackList() {
	r.mockApplicationGetApplicationFeedbackList = nil
}

// GetApplicationFeedbackListReq ...
type GetApplicationFeedbackListReq struct {
	AppID        string  `path:"app_id" json:"-"`         // 目标应用 ID（本租户创建的所有应用）, 示例值: "cli_9f115af860f7901b"
	FromDate     *string `query:"from_date" json:"-"`     // 查询的起始日期, 格式为yyyy-mm-dd。不填则默认为当前日期减去180天, 示例值: "2022-01-30"
	ToDate       *string `query:"to_date" json:"-"`       // 查询的结束日期, 格式为yyyy-mm-dd。不填默认为当前日期, 只能查询 180 天内的数据, 示例值: "2022-01-30"
	FeedbackType *int64  `query:"feedback_type" json:"-"` // 反馈类型, 不填写则表示查询所有反馈类型, 示例值: 1, 可选值有: 1: 故障反馈, 2: 产品建议
	Status       *int64  `query:"status" json:"-"`        // 反馈处理状态, 不填写则表示查询所有处理类型, 示例值: 0, 可选值有: 0: 反馈未处理, 1: 反馈已处理, 2: 反馈处理中, 3: 反馈已关闭
	UserIDType   *IDType `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	PageToken    *string `query:"page_token" json:"-"`    // 分页拉取反馈列表起始位置标示, 不填表示从头开始, 示例值: ""7064688334618378259""
	PageSize     *int64  `query:"page_size" json:"-"`     // 本次拉取反馈列表最大个数, 示例值: 100, 默认值: `100`, 取值范围: `1` ～ `100`
}

// GetApplicationFeedbackListResp ...
type GetApplicationFeedbackListResp struct {
	FeedbackList []*GetApplicationFeedbackListRespFeedback `json:"feedback_list,omitempty"` // 应用的反馈列表
	HasMore      bool                                      `json:"has_more,omitempty"`      // 是否还有更多用户反馈列表, true: 是, false: 否
	PageToken    string                                    `json:"page_token,omitempty"`    // 拉取下一页应用反馈列表时使用的 page_token
}

// GetApplicationFeedbackListRespFeedback ...
type GetApplicationFeedbackListRespFeedback struct {
	FeedbackID   string   `json:"feedback_id,omitempty"`   // 应用反馈 ID, 应用反馈记录唯一标识
	AppID        string   `json:"app_id,omitempty"`        // 被反馈应用ID
	FeedbackTime string   `json:"feedback_time,omitempty"` // 反馈提交时间, 格式为yyyy-mm-dd hh:mm:ss
	TenantName   string   `json:"tenant_name,omitempty"`   // 反馈用户的租户名, 查询 isv 应用时返回
	FeedbackType int64    `json:"feedback_type,omitempty"` // 反馈类型, 可选值有: 1: 故障反馈, 2: 产品建议
	Status       int64    `json:"status,omitempty"`        // 反馈处理状态, 可选值有: 0: 反馈未处理, 1: 反馈已处理, 2: 反馈处理中, 3: 反馈已关闭
	FaultType    []int64  `json:"fault_type,omitempty"`    // 故障类型列表: 1: 黑屏 2: 白屏 3: 无法打开小程序  4: 卡顿 5: 小程序闪退 6: 页面加载慢 7: 死机 8: 其他异常
	FaultTime    string   `json:"fault_time,omitempty"`    // 故障时间, 格式为yyyy-mm-dd hh:mm:ss
	Source       int64    `json:"source,omitempty"`        // 反馈来源: 1: 小程序 2: 网页应用 3: 机器人 4: webSDK, 可选值有: 1: 小程序, 2: 网页应用, 3: 机器人, 4: WebSDK
	Contact      string   `json:"contact,omitempty"`       // 用户联系方式, 只有用户填写联系方式后返回, 字段权限要求（满足任一）: 获取用户邮箱信息, 获取用户手机号
	UpdateTime   string   `json:"update_time,omitempty"`   // 反馈处理时间, 格式为yyyy-mm-dd hh:mm:ss
	Description  string   `json:"description,omitempty"`   // 反馈问题描述
	UserID       string   `json:"user_id,omitempty"`       // 反馈用户id, 租户内用户的唯一标识, ID值与查询参数中的user_id_type对应
	OperatorID   string   `json:"operator_id,omitempty"`   // 操作者id, 租户内用户的唯一标识, ID值与查询参数中的user_id_type 对应
	Images       []string `json:"images,omitempty"`        // 反馈图片url列表, url 过期时间三天
	FeedbackPath string   `json:"feedback_path,omitempty"` // 反馈页面路径, 如触发反馈的应用类型为小程序, 则上报小程序当前页面的path信息, 如触发反馈的应用类型为网页或网页应用, 则上报当前网页的url信息, 如为其他应用类型, 则字段返回值为空
}

// getApplicationFeedbackListResp ...
type getApplicationFeedbackListResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetApplicationFeedbackListResp `json:"data,omitempty"`
}
