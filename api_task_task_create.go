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

// CreateTask 该接口可以创建一个任务（基本信息）, 如果需要绑定协作者等需要调用别的资源管理接口。其中查询字段 user_id_type 是用于控制返回体中 creator_id 的类型, 不传时默认返回 open_id。当使用tenant_access_token 调用接口时, 如果user_id_type为user_id, 则不会返回creator_id。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/create
func (r *TaskService) CreateTask(ctx context.Context, request *CreateTaskReq, options ...MethodOptionFunc) (*CreateTaskResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#CreateTask mock enable")
		return r.cli.mock.mockTaskCreateTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTask",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTask mock TaskCreateTask method
func (r *Mock) MockTaskCreateTask(f func(ctx context.Context, request *CreateTaskReq, options ...MethodOptionFunc) (*CreateTaskResp, *Response, error)) {
	r.mockTaskCreateTask = f
}

// UnMockTaskCreateTask un-mock TaskCreateTask method
func (r *Mock) UnMockTaskCreateTask() {
	r.mockTaskCreateTask = nil
}

// CreateTaskReq ...
type CreateTaskReq struct {
	UserIDType      *IDType              `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Summary         *string              `json:"summary,omitempty"`          // 任务标题。创建任务时, 标题和富文本标题不能同时为空, 需要至少填充其中一个字段, <md-alert>, 任务标题和任务富文本标题同时存在时只使用富文本标题, </md-alert>, 示例值: "每天喝八杯水, 保持身心愉悦", 长度范围: `0` ～ `256` 字符
	Description     *string              `json:"description,omitempty"`      // 任务备注, <md-alert>, 任务备注和任务富文本备注同时存在时只使用富文本备注, </md-alert>, 示例值: "多吃水果, 多运动, 健康生活, 快乐工作。", 长度范围: `0` ～ `65536` 字符
	Extra           *string              `json:"extra,omitempty"`            // 接入方可以自定义的附属信息二进制格式, 采用 base64 编码, 解析方式由接入方自己决定, 示例值: "dGVzdA[", 长度范围: `0` ～ `65536` 字符
	Due             *CreateTaskReqDue    `json:"due,omitempty"`              // 任务的截止时间设置
	Origin          *CreateTaskReqOrigin `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息
	CanEdit         *bool                `json:"can_edit,omitempty"`         // 此字段用于控制该任务在飞书任务中心是否可编辑, 默认为false, 若为true则第三方需考虑是否需要接入事件来接收任务在任务中心的变更信息, （即将废弃）, 示例值: true, 默认值: `false`
	Custom          *string              `json:"custom,omitempty"`           // 此字段用于存储第三方需透传到端上的自定义数据, Json格式。取值举例中custom_complete字段存储「完成」按钮的跳转链接（href）或提示信息（tip）, pc、ios、android三端均可自定义, 其中tip字段的key为语言类型, value为提示信息, 可自行增加或减少语言类型, 支持的各地区语言名: it_it, th_th, ko_kr, es_es, ja_jp, zh_cn, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, en_us, hi_in, vi_vn。href的优先级高于tip, href和tip同时不为空时只跳转不提示。链接和提示信息可自定义, 其余的key需按举例中的结构传递, 示例值: "{\"custom_complete\":{\"android\":{\"href\":\"https://www.feishu.cn/\", \"tip\":{\"zh_cn\":\"你好\", \"en_us\":\"hello\"}}, \"ios\":{\"href\":\"https://www.feishu.cn/\", \"tip\":{\"zh_cn\":\"你好\", \"en_us\":\"hello\"}}, \"pc\":{\"href\":\"https://www.feishu.cn/\", \"tip\":{\"zh_cn\":\"你好\", \"en_us\":\"hello\"}}}}", 长度范围: `0` ～ `65536` 字符
	CollaboratorIDs []string             `json:"collaborator_ids,omitempty"` // 创建任务时添加的执行者用户id列表, 示例值: ["ou_1400208f15333e20e11339d39067844b", "ou_84ed6312949945c8ae6168f10829e9e6"], 最大长度: `100`
	FollowerIDs     []string             `json:"follower_ids,omitempty"`     // 创建任务时添加的关注者用户id列表, 示例值: ["ou_1400208f15333e20e11339d39067844b", "ou_84ed6312949945c8ae6168f10829e9e6"], 最大长度: `100`
	RepeatRule      *string              `json:"repeat_rule,omitempty"`      // 重复任务重复规则。语法格式参见[RRule语法规范](https://www.ietf.org/rfc/rfc2445.txt) 4.3.10小节, 示例值: "FREQ=WEEKLY;INTERVAL=1;BYDAY=MO, TU, WE, TH, FR"
	RichSummary     *string              `json:"rich_summary,omitempty"`     // 富文本任务标题。创建任务时, 如果没有标题填充, 飞书服务器会将其视为无主题的任务。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module), 示例值: "每天喝八杯水, 保持身心愉悦\[飞书开放平台\](https://open.feishu.cn/)", 长度范围: `0` ～ `256` 字符
	RichDescription *string              `json:"rich_description,omitempty"` // 富文本任务备注。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module), 示例值: "多吃水果, 多运动, 健康生活, 快乐工作。\[飞书开放平台](https://open.feishu.cn/)", 长度范围: `0` ～ `65536` 字符
}

// CreateTaskReqDue ...
type CreateTaskReqDue struct {
	Time     *string `json:"time,omitempty"`       // 截止时间的时间戳（单位为秒）, 示例值: "1623124318"
	Timezone *string `json:"timezone,omitempty"`   // 截止时间对应的时区, 使用IANA Time Zone Database标准, 如Asia/Shanghai, 示例值: "Asia/Shanghai", 默认值: `Asia/Shanghai`
	IsAllDay *bool   `json:"is_all_day,omitempty"` // 标记任务是否为全天任务（全天任务的截止时间为当天 UTC 时间的 0 点）, 示例值: false, 默认值: `false`
}

// CreateTaskReqOrigin ...
type CreateTaskReqOrigin struct {
	PlatformI18nName string                   `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称, 用于在任务中心详情页展示。请提供一个字典, 多种语言名称映射。支持的各地区语言名: it_it, th_th, ko_kr, es_es, ja_jp, zh_cn, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, en_us, hi_in, vi_vn, 示例值: "{\"zh_cn\": \"IT 工作台\", \"en_us\": \"IT Workspace\"}", 长度范围: `0` ～ `1024` 字符
	Href             *CreateTaskReqOriginHref `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// CreateTaskReqOriginHref ...
type CreateTaskReqOriginHref struct {
	URL   *string `json:"url,omitempty"`   // 具体链接地址, 示例值: "https://support.feishu.com/internal/foo-bar", 长度范围: `0` ～ `1024` 字符
	Title *string `json:"title,omitempty"` // 链接对应的标题, 示例值: "反馈一个问题, 需要协助排查", 长度范围: `0` ～ `512` 字符
}

// CreateTaskResp ...
type CreateTaskResp struct {
	Task *CreateTaskRespTask `json:"task,omitempty"` // 返回创建好的任务
}

// CreateTaskRespTask ...
type CreateTaskRespTask struct {
	ID              string                            `json:"id,omitempty"`               // 任务 ID, 由飞书任务服务器发号
	Summary         string                            `json:"summary,omitempty"`          // 任务标题。创建任务时, 标题和富文本标题不能同时为空, 需要至少填充其中一个字段, <md-alert>, 任务标题和任务富文本标题同时存在时只使用富文本标题, </md-alert>
	Description     string                            `json:"description,omitempty"`      // 任务备注, <md-alert>, 任务备注和任务富文本备注同时存在时只使用富文本备注, </md-alert>
	CompleteTime    string                            `json:"complete_time,omitempty"`    // 任务的完成时间戳（单位为秒）, 如果完成时间为 0, 则表示任务尚未完成
	CreatorID       string                            `json:"creator_id,omitempty"`       // 任务的创建者 ID。在创建任务时无需填充该字段
	Extra           string                            `json:"extra,omitempty"`            // 接入方可以自定义的附属信息二进制格式, 采用 base64 编码, 解析方式由接入方自己决定
	CreateTime      string                            `json:"create_time,omitempty"`      // 任务的创建时间戳（单位为秒）
	UpdateTime      string                            `json:"update_time,omitempty"`      // 任务的更新时间戳（单位为秒）
	Due             *CreateTaskRespTaskDue            `json:"due,omitempty"`              // 任务的截止时间设置
	Origin          *CreateTaskRespTaskOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息
	CanEdit         bool                              `json:"can_edit,omitempty"`         // 此字段用于控制该任务在飞书任务中心是否可编辑, 默认为false, 若为true则第三方需考虑是否需要接入事件来接收任务在任务中心的变更信息, （即将废弃）
	Custom          string                            `json:"custom,omitempty"`           // 此字段用于存储第三方需透传到端上的自定义数据, Json格式。取值举例中custom_complete字段存储「完成」按钮的跳转链接（href）或提示信息（tip）, pc、ios、android三端均可自定义, 其中tip字段的key为语言类型, value为提示信息, 可自行增加或减少语言类型, 支持的各地区语言名: it_it, th_th, ko_kr, es_es, ja_jp, zh_cn, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, en_us, hi_in, vi_vn。href的优先级高于tip, href和tip同时不为空时只跳转不提示。链接和提示信息可自定义, 其余的key需按举例中的结构传递
	Source          int64                             `json:"source,omitempty"`           // 任务创建的来源, 可选值有: 0: 未知类型, 1: 来源任务中心创建, 2: 来源消息转任务, 3: 来源云文档, 4: 来源文档单品, 5: 来源PANO, 6: 来源tenant_access_token创建的任务, 7: 来源user_access_token创建的任务, 8: 来源新版云文档
	Followers       []*CreateTaskRespTaskFollower     `json:"followers,omitempty"`        // 任务的关注者
	Collaborators   []*CreateTaskRespTaskCollaborator `json:"collaborators,omitempty"`    // 任务的执行者
	CollaboratorIDs []string                          `json:"collaborator_ids,omitempty"` // 创建任务时添加的执行者用户id列表
	FollowerIDs     []string                          `json:"follower_ids,omitempty"`     // 创建任务时添加的关注者用户id列表
	RepeatRule      string                            `json:"repeat_rule,omitempty"`      // 重复任务重复规则。语法格式参见[RRule语法规范](https://www.ietf.org/rfc/rfc2445.txt) 4.3.10小节
	RichSummary     string                            `json:"rich_summary,omitempty"`     // 富文本任务标题。创建任务时, 如果没有标题填充, 飞书服务器会将其视为无主题的任务。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module)
	RichDescription string                            `json:"rich_description,omitempty"` // 富文本任务备注。语法格式参见[Markdown模块](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/markdown-module)
}

// CreateTaskRespTaskCollaborator ...
type CreateTaskRespTaskCollaborator struct {
	ID     string   `json:"id,omitempty"`      // 任务执行者的 ID
	IDList []string `json:"id_list,omitempty"` // 执行者的用户ID列表。
}

// CreateTaskRespTaskDue ...
type CreateTaskRespTaskDue struct {
	Time     string `json:"time,omitempty"`       // 截止时间的时间戳（单位为秒）
	Timezone string `json:"timezone,omitempty"`   // 截止时间对应的时区, 使用IANA Time Zone Database标准, 如Asia/Shanghai
	IsAllDay bool   `json:"is_all_day,omitempty"` // 标记任务是否为全天任务（全天任务的截止时间为当天 UTC 时间的 0 点）
}

// CreateTaskRespTaskFollower ...
type CreateTaskRespTaskFollower struct {
	ID     string   `json:"id,omitempty"`      // 任务关注人 ID
	IDList []string `json:"id_list,omitempty"` // 要添加的关注人ID列表
}

// CreateTaskRespTaskOrigin ...
type CreateTaskRespTaskOrigin struct {
	PlatformI18nName string                        `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称, 用于在任务中心详情页展示。请提供一个字典, 多种语言名称映射。支持的各地区语言名: it_it, th_th, ko_kr, es_es, ja_jp, zh_cn, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, en_us, hi_in, vi_vn
	Href             *CreateTaskRespTaskOriginHref `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// CreateTaskRespTaskOriginHref ...
type CreateTaskRespTaskOriginHref struct {
	URL   string `json:"url,omitempty"`   // 具体链接地址
	Title string `json:"title,omitempty"` // 链接对应的标题
}

// createTaskResp ...
type createTaskResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateTaskResp `json:"data,omitempty"`
}
