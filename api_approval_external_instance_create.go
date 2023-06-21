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

// CreateApprovalExternalInstance 审批中心不负责审批的流转, 审批的流转在三方系统, 三方系统在审批流转后生成的审批实例、审批任务、审批抄送数据同步到审批中心。
//
// 用户可以在审批中心中浏览三方系统同步过来的实例、任务、抄送信息, 并且可以跳转回三方系统进行更详细的查看和操作, 其中实例信息在【已发起】列表, 任务信息在【待审批】和【已审批】列表, 抄送信息在【抄送我】列表。
// :::html
// <img src="//sf3-cn.feishucdn.com/obj/open-platform-opendoc/9dff4434afbeb0ef69de7f36b9a6e995_z5iwmTzEgg.png" alt="" style="zoom:17%;" />
// <img src="//sf3-cn.feishucdn.com/obj/open-platform-opendoc/ca6e0e984a7a6d64e1b16a0bac4bf868_tfqjCiaJQM.png" alt="" style="zoom:17%;" />
// <img src="//sf3-cn.feishucdn.com/obj/open-platform-opendoc/529377e238df78d391bbd22e962ad195_T7eefLI1GA.png" alt="" style="zoom:17%;" />
// 对于审批任务, 三方系统也可以配置审批任务的回调接口, 这样审批人可以在审批中心中直接进行审批操作, 审批中心会回调三方系统, 三方系统收到回调后更新任务信息, 并将新的任务信息同步回审批中心, 形成闭环。
// :::html
// <img src="//sf3-cn.feishucdn.com/obj/open-platform-opendoc/721c35428bc1187db3318c572f9979ad_je75QpElcg.png" alt=""  style="zoom:25%;" />
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/external_instance/create
// new doc: https://open.feishu.cn/document/server-docs/approval-v4/external_instance/create
func (r *ApprovalService) CreateApprovalExternalInstance(ctx context.Context, request *CreateApprovalExternalInstanceReq, options ...MethodOptionFunc) (*CreateApprovalExternalInstanceResp, *Response, error) {
	if r.cli.mock.mockApprovalCreateApprovalExternalInstance != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#CreateApprovalExternalInstance mock enable")
		return r.cli.mock.mockApprovalCreateApprovalExternalInstance(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "CreateApprovalExternalInstance",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/external_instances",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createApprovalExternalInstanceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalCreateApprovalExternalInstance mock ApprovalCreateApprovalExternalInstance method
func (r *Mock) MockApprovalCreateApprovalExternalInstance(f func(ctx context.Context, request *CreateApprovalExternalInstanceReq, options ...MethodOptionFunc) (*CreateApprovalExternalInstanceResp, *Response, error)) {
	r.mockApprovalCreateApprovalExternalInstance = f
}

// UnMockApprovalCreateApprovalExternalInstance un-mock ApprovalCreateApprovalExternalInstance method
func (r *Mock) UnMockApprovalCreateApprovalExternalInstance() {
	r.mockApprovalCreateApprovalExternalInstance = nil
}

// CreateApprovalExternalInstanceReq ...
type CreateApprovalExternalInstanceReq struct {
	ApprovalCode          string                                            `json:"approval_code,omitempty"`            // 审批定义 code, 创建审批定义返回的值, 表示该实例属于哪个流程；该字段会影响到列表中该实例的标题, 标题取自对应定义的 name 字段, 示例值: "81D31358-93AF-92D6-7425-01A5D67C4E71"
	Status                string                                            `json:"status,omitempty"`                   // 审批实例状态, 示例值: "PENDING", 可选值有: PENDING: 审批中, APPROVED: 审批流程结束, 结果为同意, REJECTED: 审批流程结束, 结果为拒绝, CANCELED: 审批发起人撤回, DELETED: 审批被删除, HIDDEN: 状态隐藏(不显示状态), TERMINATED: 审批终止
	Extra                 *string                                           `json:"extra,omitempty"`                    // 审批实例扩展 JSON。单据编号通过传business_key字段来实现, 示例值: "{\"xxx\":\"xxx\", \"business_key\":\"xxx\"}"
	InstanceID            string                                            `json:"instance_id,omitempty"`              // 审批实例唯一标识, 用户自定义, 需确保证租户下唯一, 示例值: "24492654"
	Links                 []*CreateApprovalExternalInstanceReqLink          `json:"links,omitempty"`                    // 审批实例链接集合, 用于【已发起】列表的跳转, 跳转回三方系统； pc_link 和 mobile_link 必须填一个, 填写的是哪一端的链接, 即会跳转到该链接, 不受平台影响
	Title                 *string                                           `json:"title,omitempty"`                    // 审批展示名称, 如果填写了该字段, 则审批列表中的审批名称使用该字段, 如果不填该字段, 则审批名称使用审批定义的名称, 示例值: "@i18n@1"
	Form                  []*CreateApprovalExternalInstanceReqForm          `json:"form,omitempty"`                     // 用户提交审批时填写的表单数据, 用于所有审批列表中展示。可传多个值, 但审批中心pc展示前2个, 移动端展示前3个, 长度不超过2048字符, 示例值: [{ "name": "@i18n@2", "value": "@i18n@3" }]
	UserID                *string                                           `json:"user_id,omitempty"`                  // 审批发起人 user_id, 发起人可在【已发起】列表中看到所有已发起的审批; 在【待审批】, 【已审批】【抄送我】列表中, 该字段展示审批是谁发起的。审批发起人 open id, 和 user id 二者至少填一个, 示例值: "a987sf9s"
	UserName              *string                                           `json:"user_name,omitempty"`                // 审批发起人 用户名, 如果发起人不是真实的用户（例如是某个部门）, 没有 user_id, 则可以使用该字段传名称, 示例值: "@i18n@9"
	OpenID                *string                                           `json:"open_id,omitempty"`                  // 审批发起人 open id, 和 user id 二者至少填一个, 示例值: "ou_be73cbc0ee35eb6ca54e9e7cc14998c1"
	DepartmentID          *string                                           `json:"department_id,omitempty"`            // 发起人部门, 用于列表中展示发起人所属部门。不传则不展示。如果用户没加入任何部门, 传 "", 将展示租户名称传 department_name 展示部门名称, 示例值: "od-8ec33278bc2"
	DepartmentName        *string                                           `json:"department_name,omitempty"`          // 审批发起人 部门, 如果发起人不是真实的用户（例如是某个部门）, 没有 department_id, 则可以使用该字段传名称, 示例值: "@i18n@10"
	StartTime             string                                            `json:"start_time,omitempty"`               // 审批发起时间, Unix毫秒时间戳, 示例值: "1556468012678"
	EndTime               string                                            `json:"end_time,omitempty"`                 // 审批实例结束时间: 未结束的审批为 0, Unix毫秒时间戳, 示例值: "1556468012678"
	UpdateTime            string                                            `json:"update_time,omitempty"`              // 审批实例最近更新时间；用于推送数据版本控制如果 update_mode 值为 UPDATE, 则只有传过来的 update_time 有变化时（变大）, 才会更新审批中心中的审批实例信息。使用该字段主要用来避免并发时老的数据更新了新的数据, 示例值: "1556468012678"
	DisplayMethod         *string                                           `json:"display_method,omitempty"`           // 列表页打开审批实例的方式, 示例值: "BROWSER", 可选值有: BROWSER: 跳转系统默认浏览器打开, SIDEBAR: 飞书中侧边抽屉打开, NORMAL: 飞书内嵌页面打开, TRUSTEESHIP: 以托管打开
	UpdateMode            *string                                           `json:"update_mode,omitempty"`              // 更新方式, 当 update_mode=REPLACE时, 每次都以当前推送的数据为最终数据, 会删掉审批中心中多余的任务、抄送数据（不在这次推送的数据中）; 当 update_mode=UPDATE时, 则不会删除审批中心的数据, 而只是进行新增和更新实例、任务数据, 示例值: "UPDATE", 可选值有: REPLACE: 全量替换, 默认值, UPDATE: 增量更新
	TaskList              []*CreateApprovalExternalInstanceReqTask          `json:"task_list,omitempty"`                // 任务列表, 最大长度: `200`
	CcList                []*CreateApprovalExternalInstanceReqCc            `json:"cc_list,omitempty"`                  // 抄送列表, 最大长度: `200`
	I18nResources         []*CreateApprovalExternalInstanceReqI18nResource  `json:"i18n_resources,omitempty"`           // 国际化文案
	TrusteeshipURLToken   *string                                           `json:"trusteeship_url_token,omitempty"`    // 单据托管认证token, 托管回调会附带此token, 帮助业务方认证, 示例值: "788981c886b1c28ac29d1e68efd60683d6d90dfce80938ee9453e2a5f3e9e306"
	TrusteeshipUserIDType *IDType                                           `json:"trusteeship_user_id_type,omitempty"` // 用户的类型, 会影响请求参数用户标识域的选择, 包括加签操作回传的目标用户, 目前仅支持 "user_id", 示例值: "user_id"
	TrusteeshipURLs       *CreateApprovalExternalInstanceReqTrusteeshipURLs `json:"trusteeship_urls,omitempty"`         // 单据托管回调接入方的接口的URL地址
}

// CreateApprovalExternalInstanceReqCc ...
type CreateApprovalExternalInstanceReqCc struct {
	CcID          string                                     `json:"cc_id,omitempty"`          // 审批实例内唯一标识, 示例值: "123456"
	UserID        *string                                    `json:"user_id,omitempty"`        // 抄送人 employee id, 示例值: "12345"
	OpenID        *string                                    `json:"open_id,omitempty"`        // 抄送人 open id, 和user id 二者至少填一个, 示例值: "ou_be73cbc0ee35eb6ca54e9e7cc14998c1"
	Links         []*CreateApprovalExternalInstanceReqCcLink `json:"links,omitempty"`          // 跳转链接, 用于【抄送我的】列表中的跳转pc_link 和 mobile_link 必须填一个, 填写的是哪一端的链接, 即会跳转到该链接, 不受平台影响
	ReadStatus    string                                     `json:"read_status,omitempty"`    // 阅读状态, 空值表示不支持已读未读: 示例值: "READ", 可选值有: READ: 已读, UNREAD: 未读
	Extra         *string                                    `json:"extra,omitempty"`          // 扩展 json, 示例值: "{\"xxx\":\"xxx\"}"
	Title         *string                                    `json:"title,omitempty"`          // 抄送任务名称, 示例值: "xxx"
	CreateTime    string                                     `json:"create_time,omitempty"`    // 抄送发起时间, Unix 毫秒时间戳, 示例值: "1556468012678"
	UpdateTime    string                                     `json:"update_time,omitempty"`    // 抄送最近更新时间, 用于推送数据版本控制更新策略同 instance 的update_time, 示例值: "1556468012678"
	DisplayMethod *string                                    `json:"display_method,omitempty"` // 列表页打开审批任务的方式, 示例值: "BROWSER", 可选值有: BROWSER: 跳转系统默认浏览器打开, SIDEBAR: 飞书中侧边抽屉打开, NORMAL: 飞书内嵌页面打开, TRUSTEESHIP: 以托管模式打开
}

// CreateApprovalExternalInstanceReqCcLink ...
type CreateApprovalExternalInstanceReqCcLink struct {
	PcLink     string  `json:"pc_link,omitempty"`     // pc 端的跳转链接, 当用户使用飞书 pc 端时, 使用该字段进行跳转, 示例值: "https://applink.feishu.cn/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234"
	MobileLink *string `json:"mobile_link,omitempty"` // 移动端 跳转链接, 当用户使用飞书 移动端时, 使用该字段进行跳转, 示例值: "https://applink.feishu.cn/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234"
}

// CreateApprovalExternalInstanceReqForm ...
type CreateApprovalExternalInstanceReqForm struct {
	Name  *string `json:"name,omitempty"`  // 表单字段名称, 示例值: "@i18n@2"
	Value *string `json:"value,omitempty"` // 表单值, 示例值: "@i18n@3"
}

// CreateApprovalExternalInstanceReqI18nResource ...
type CreateApprovalExternalInstanceReqI18nResource struct {
	Locale    string                                               `json:"locale,omitempty"`     // 语言可选值有: zh-CN: 中文 en-US: 英文 ja-JP: 日文, 示例值: "zh-CN", 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
	Texts     []*CreateApprovalExternalInstanceReqI18nResourceText `json:"texts,omitempty"`      // 文案 key, value, i18n key 以 @i18n@ 开头； 该字段主要用于做国际化, 允许用户同时传多个语言的文案, 审批中心会根据用户当前的语音环境使用对应的文案, 如果没有传用户当前的语音环境文案, 则会使用默认的语言文案, 示例值: { "@i18n@1": "权限申请", "@i18n@2": "OA审批", "@i18n@3": "Permission" }
	IsDefault bool                                                 `json:"is_default,omitempty"` // 是否默认语言, 默认语言需要包含所有key, 非默认语言如果key不存在会使用默认语言代替, 示例值: true
}

// CreateApprovalExternalInstanceReqI18nResourceText ...
type CreateApprovalExternalInstanceReqI18nResourceText struct {
	Key   string `json:"key,omitempty"`   // 文案key, 示例值: "@i18n@1"
	Value string `json:"value,omitempty"` // 文案, 示例值: "people"
}

// CreateApprovalExternalInstanceReqLink ...
type CreateApprovalExternalInstanceReqLink struct {
	PcLink     string  `json:"pc_link,omitempty"`     // pc 端的跳转链接, 当用户使用飞书 pc 端时, 使用该字段进行跳转, 托管的链接保持不变, 示例值: "https://applink.feishu.cn/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234"
	MobileLink *string `json:"mobile_link,omitempty"` // 移动端 跳转链接, 当用户使用飞书 移动端时, 使用该字段进行跳转, 托管的链接保持不变, 示例值: "https://applink.feishu.cn/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234"
}

// CreateApprovalExternalInstanceReqTask ...
type CreateApprovalExternalInstanceReqTask struct {
	TaskID            string                                               `json:"task_id,omitempty"`            // 审批实例内的唯一标识, 用于更新审批任务时定位数据, 示例值: "112534"
	UserID            *string                                              `json:"user_id,omitempty"`            // 审批人 user_id, 和 open_id 二者至少填一个。该任务会出现在审批人的【待审批】或【已审批】列表中, 示例值: "a987sf9s"
	OpenID            *string                                              `json:"open_id,omitempty"`            // 审批人 open_id, 和 user_id 二者至少填一个, 示例值: "ou_be73cbc0ee35eb6ca54e9e7cc14998c1"
	Title             *string                                              `json:"title,omitempty"`              // 审批任务名称, 示例值: "i18n1"
	Links             *CreateApprovalExternalInstanceReqTaskLinks          `json:"links,omitempty"`              // 【待审批】或【已审批】中使用的跳转链接, 用于跳转回三方系统pc_link 和 mobile_link 必须填一个, 填写的是哪一端的链接, 即会跳转到该链接, 不受平台影响
	Status            string                                               `json:"status,omitempty"`             // 任务状态, 示例值: "PENDING", 可选值有: PENDING: 待审批, APPROVED: 任务同意, REJECTED: 任务拒绝, TRANSFERRED: 任务转交, DONE: 任务通过但审批人未操作；审批人看不到这个任务, 若想要看到, 可以通过抄送该人.
	Extra             *string                                              `json:"extra,omitempty"`              // 扩展 json, 任务结束原因需传complete_reason字段。枚举值与对应说明: approved: 同意, rejected: 拒绝, node_auto_reject: （因逻辑判断产生的）自动拒绝, specific_rollback: 退回（包括退回到发起人、退回到中间任一审批人）, add: 并加签（添加新审批人, 和我一起审批）, add_pre: 前加签（添加新审批人, 在我之前审批）, add_post: 后加签（添加新审批人, 在我之后审批）, delete_assignee: 减签, forward_resign: 转交（转给其他人审批）, recall: 撤销（撤回单据, 单据失效）, delete : 删除审批单, admin_forward: 管理员在后台操作转交, system_forward: 系统自动转交, auto_skip: 自动通过, manual_skip: 手动跳过, submit_again: 重新提交任务, restart: 重新启动流程, others: 其他（作为兜底）, 示例值: "{\"xxx\":\"xxx\", \"complete_reason\":\"approved\"}"
	CreateTime        string                                               `json:"create_time,omitempty"`        // 任务创建时间, Unix 毫秒时间戳, 示例值: "1556468012678"
	EndTime           string                                               `json:"end_time,omitempty"`           // 任务完成时间: 未结束的审批为 0, Unix 毫秒时间戳, 示例值: "1556468012678"
	UpdateTime        *string                                              `json:"update_time,omitempty"`        // task最近更新时间, 用于推送数据版本控制； 更新策略同 instance 中的 update_time, 示例值: "1556468012678"
	ActionContext     *string                                              `json:"action_context,omitempty"`     // 操作上下文, 当用户操作时, 回调请求中带上该参数, 用于传递该任务的上下文数据, 示例值: "123456"
	ActionConfigs     []*CreateApprovalExternalInstanceReqTaskActionConfig `json:"action_configs,omitempty"`     // 任务级别操作配置, 快捷审批目前支持移动端操作
	DisplayMethod     *string                                              `json:"display_method,omitempty"`     // 列表页打开审批任务的方式, 示例值: "BROWSER", 可选值有: BROWSER: 跳转系统默认浏览器打开, SIDEBAR: 飞书中侧边抽屉打开, NORMAL: 飞书内嵌页面打开, TRUSTEESHIP: 以托管模式打开
	ExcludeStatistics *bool                                                `json:"exclude_statistics,omitempty"` // 三方任务支持不纳入效率统计, false: 纳入效率统计, true: 不纳入效率统计, 示例值: false, 默认值: `false`
	NodeID            *string                                              `json:"node_id,omitempty"`            // 节点id: 必须同时满足, 一个流程内, 每个节点id唯一。如一个流程下「直属上级」、「隔级上级」等每个节点的Node_id均不一样, 同一个流程定义内, 不同审批实例中的相同节点, Node_id要保持不变。例如张三和李四分别发起了请假申请, 这2个审批实例中的「直属上级」节点的node_id应该保持一致, 示例值: "node"
	NodeName          *string                                              `json:"node_name,omitempty"`          // 节点名称, 如「财务审批」「法务审批」, 支持中英日三种语言。示例: i18n@name。需要在i18n_resources中传该名称对应的国际化文案, 示例值: "i18n@name"
}

// CreateApprovalExternalInstanceReqTaskActionConfig ...
type CreateApprovalExternalInstanceReqTaskActionConfig struct {
	ActionType       string  `json:"action_type,omitempty"`        // 操作类型, 每个任务都可以配置2个操作, 会展示审批列表中, 当用户操作时, 回调请求会带上该字段, 表示用户进行了同意操作还是拒绝操作, 可选值有: APPROVE: 同意, REJECT: 拒绝, {KEY}: 任意字符串, 如果使用任意字符串, 则需要提供 action_name, 示例值: "APPROVE"
	ActionName       *string `json:"action_name,omitempty"`        // 操作名称, i18n key 用于前台展示, 如果 action_type 不是 APPROVAL和REJECT, 则必须提供该字段, 用于展示特定的操作名称, 示例值: "@i18n@5"
	IsNeedReason     *bool   `json:"is_need_reason,omitempty"`     // 是否需要意见, 如果为true, 则用户操作时, 会跳转到 意见填写页面, 示例值: false
	IsReasonRequired *bool   `json:"is_reason_required,omitempty"` // 审批意见是否必填, 示例值: false
	IsNeedAttachment *bool   `json:"is_need_attachment,omitempty"` // 意见是否支持上传附件, 示例值: false
}

// CreateApprovalExternalInstanceReqTaskLinks ...
type CreateApprovalExternalInstanceReqTaskLinks struct {
	PcLink     string  `json:"pc_link,omitempty"`     // pc 端的跳转链接, 当用户使用飞书 pc 端时, 使用该字段进行跳转, 示例值: "https://applink.feishu.cn/client/mini_program/open?mode=appCenter&appId=cli_9c90fc38e07a9101&path=pc/pages/detail?id=1234"
	MobileLink *string `json:"mobile_link,omitempty"` // 移动端 跳转链接, 当用户使用飞书 移动端时, 使用该字段进行跳转, 示例值: "https://applink.feishu.cn/client/mini_program/open?appId=cli_9c90fc38e07a9101&path=pages/detail?id=1234"
}

// CreateApprovalExternalInstanceReqTrusteeshipURLs ...
type CreateApprovalExternalInstanceReqTrusteeshipURLs struct {
	FormDetailURL       *string `json:"form_detail_url,omitempty"`        // 获取表单schema相关数据的url地址, 示例值: "https://#{your_domain}/api/form_detail"
	ActionDefinitionURL *string `json:"action_definition_url,omitempty"`  // 表示获取审批操作区数据的url地址, 示例值: "https://#{your_domain}/api/action_definition"
	ApprovalNodeURL     *string `json:"approval_node_url,omitempty"`      // 获取审批记录相关数据的url地址, 示例值: "https://#{your_domain}/api/approval_node"
	ActionCallbackURL   *string `json:"action_callback_url,omitempty"`    // 进行审批操作时回调的url地址, 示例值: "https://#{your_domain}/api/action_callback"
	PullBusinessDataURL *string `json:"pull_business_data_url,omitempty"` // 获取托管动态数据url 地址, 使用该接口时必须要保证历史托管单据的数据中都同步了该接口地址, 如果历史单据中没有该接口需要重新同步历史托管单据的数据来更新该URL。该接口用于飞书审批前端和业务线进行交互使用, 只有使用审批前端的特定组件(由飞书审批前端提供的组件, 并且需要和业务线进行接口交互的组件)才会需要, 示例值: "https://#{your_domain}/api/pull_business_data"
}

// CreateApprovalExternalInstanceResp ...
type CreateApprovalExternalInstanceResp struct {
	Data *CreateApprovalExternalInstanceRespData `json:"data,omitempty"` // 同步的实例数据
}

// CreateApprovalExternalInstanceRespData ...
type CreateApprovalExternalInstanceRespData struct {
	ApprovalCode          string                                                 `json:"approval_code,omitempty"`            // 审批定义 code, 创建审批定义返回的值, 表示该实例属于哪个流程；该字段会影响到列表中该实例的标题, 标题取自对应定义的 name 字段
	Status                string                                                 `json:"status,omitempty"`                   // 审批实例状态, 可选值有: PENDING: 审批中, APPROVED: 审批流程结束, 结果为同意, REJECTED: 审批流程结束, 结果为拒绝, CANCELED: 审批发起人撤回, DELETED: 审批被删除, HIDDEN: 状态隐藏(不显示状态), TERMINATED: 审批终止
	Extra                 string                                                 `json:"extra,omitempty"`                    // 审批实例扩展 JSON。单据编号通过传business_key字段来实现
	InstanceID            string                                                 `json:"instance_id,omitempty"`              // 审批实例唯一标识, 用户自定义, 需确保证租户下唯一
	Links                 []*CreateApprovalExternalInstanceRespDataLink          `json:"links,omitempty"`                    // 审批实例链接集合, 用于【已发起】列表的跳转, 跳转回三方系统； pc_link 和 mobile_link 必须填一个, 填写的是哪一端的链接, 即会跳转到该链接, 不受平台影响
	Title                 string                                                 `json:"title,omitempty"`                    // 审批展示名称, 如果填写了该字段, 则审批列表中的审批名称使用该字段, 如果不填该字段, 则审批名称使用审批定义的名称
	Form                  []*CreateApprovalExternalInstanceRespDataForm          `json:"form,omitempty"`                     // 用户提交审批时填写的表单数据, 用于所有审批列表中展示。可传多个值, 但审批中心pc展示前2个, 移动端展示前3个, 长度不超过2048字符
	UserID                string                                                 `json:"user_id,omitempty"`                  // 审批发起人 user_id, 发起人可在【已发起】列表中看到所有已发起的审批; 在【待审批】, 【已审批】【抄送我】列表中, 该字段展示审批是谁发起的。审批发起人 open id, 和 user id 二者至少填一个。
	UserName              string                                                 `json:"user_name,omitempty"`                // 审批发起人 用户名, 如果发起人不是真实的用户（例如是某个部门）, 没有 user_id, 则可以使用该字段传名称
	OpenID                string                                                 `json:"open_id,omitempty"`                  // 审批发起人 open id, 和 user id 二者至少填一个
	DepartmentID          string                                                 `json:"department_id,omitempty"`            // 发起人部门, 用于列表中展示发起人所属部门。不传则不展示。如果用户没加入任何部门, 传 "", 将展示租户名称传 department_name 展示部门名称
	DepartmentName        string                                                 `json:"department_name,omitempty"`          // 审批发起人 部门, 如果发起人不是真实的用户（例如是某个部门）, 没有 department_id, 则可以使用该字段传名称
	StartTime             string                                                 `json:"start_time,omitempty"`               // 审批发起时间, Unix毫秒时间戳
	EndTime               string                                                 `json:"end_time,omitempty"`                 // 审批实例结束时间: 未结束的审批为 0, Unix毫秒时间戳
	UpdateTime            string                                                 `json:"update_time,omitempty"`              // 审批实例最近更新时间；用于推送数据版本控制如果 update_mode 值为 UPDATE, 则只有传过来的 update_time 有变化时（变大）, 才会更新审批中心中的审批实例信息。使用该字段主要用来避免并发时老的数据更新了新的数据
	DisplayMethod         string                                                 `json:"display_method,omitempty"`           // 列表页打开审批实例的方式, 可选值有: BROWSER: 跳转系统默认浏览器打开, SIDEBAR: 飞书中侧边抽屉打开, NORMAL: 飞书内嵌页面打开, TRUSTEESHIP: 以托管打开
	UpdateMode            string                                                 `json:"update_mode,omitempty"`              // 更新方式, 当 update_mode=REPLACE时, 每次都以当前推送的数据为最终数据, 会删掉审批中心中多余的任务、抄送数据（不在这次推送的数据中）; 当 update_mode=UPDATE时, 则不会删除审批中心的数据, 而只是进行新增和更新实例、任务数据, 可选值有: REPLACE: 全量替换, 默认值, UPDATE: 增量更新
	TaskList              []*CreateApprovalExternalInstanceRespDataTask          `json:"task_list,omitempty"`                // 任务列表
	CcList                []*CreateApprovalExternalInstanceRespDataCc            `json:"cc_list,omitempty"`                  // 抄送列表
	I18nResources         []*CreateApprovalExternalInstanceRespDataI18nResource  `json:"i18n_resources,omitempty"`           // 国际化文案
	TrusteeshipURLToken   string                                                 `json:"trusteeship_url_token,omitempty"`    // 单据托管认证token, 托管回调会附带此token, 帮助业务方认证
	TrusteeshipUserIDType IDType                                                 `json:"trusteeship_user_id_type,omitempty"` // 用户的类型, 会影响请求参数用户标识域的选择, 包括加签操作回传的目标用户, 目前仅支持 "user_id"
	TrusteeshipURLs       *CreateApprovalExternalInstanceRespDataTrusteeshipURLs `json:"trusteeship_urls,omitempty"`         // 单据托管回调接入方的接口的URL地址
}

// CreateApprovalExternalInstanceRespDataCc ...
type CreateApprovalExternalInstanceRespDataCc struct {
	CcID          string                                          `json:"cc_id,omitempty"`          // 审批实例内唯一标识
	UserID        string                                          `json:"user_id,omitempty"`        // 抄送人 employee id
	OpenID        string                                          `json:"open_id,omitempty"`        // 抄送人 open id, 和user id 二者至少填一个
	Links         []*CreateApprovalExternalInstanceRespDataCcLink `json:"links,omitempty"`          // 跳转链接, 用于【抄送我的】列表中的跳转pc_link 和 mobile_link 必须填一个, 填写的是哪一端的链接, 即会跳转到该链接, 不受平台影响
	ReadStatus    string                                          `json:"read_status,omitempty"`    // 阅读状态, 空值表示不支持已读未读: 可选值有: READ: 已读, UNREAD: 未读
	Extra         string                                          `json:"extra,omitempty"`          // 扩展 json
	Title         string                                          `json:"title,omitempty"`          // 抄送任务名称
	CreateTime    string                                          `json:"create_time,omitempty"`    // 抄送发起时间, Unix 毫秒时间戳
	UpdateTime    string                                          `json:"update_time,omitempty"`    // 抄送最近更新时间, 用于推送数据版本控制更新策略同 instance 的update_time
	DisplayMethod string                                          `json:"display_method,omitempty"` // 列表页打开审批任务的方式, 可选值有: BROWSER: 跳转系统默认浏览器打开, SIDEBAR: 飞书中侧边抽屉打开, NORMAL: 飞书内嵌页面打开, TRUSTEESHIP: 以托管模式打开
}

// CreateApprovalExternalInstanceRespDataCcLink ...
type CreateApprovalExternalInstanceRespDataCcLink struct {
	PcLink     string `json:"pc_link,omitempty"`     // pc 端的跳转链接, 当用户使用飞书 pc 端时, 使用该字段进行跳转
	MobileLink string `json:"mobile_link,omitempty"` // 移动端 跳转链接, 当用户使用飞书 移动端时, 使用该字段进行跳转
}

// CreateApprovalExternalInstanceRespDataForm ...
type CreateApprovalExternalInstanceRespDataForm struct {
	Name  string `json:"name,omitempty"`  // 表单字段名称
	Value string `json:"value,omitempty"` // 表单值
}

// CreateApprovalExternalInstanceRespDataI18nResource ...
type CreateApprovalExternalInstanceRespDataI18nResource struct {
	Locale    string                                                    `json:"locale,omitempty"`     // 语言可选值有: zh-CN: 中文 en-US: 英文 ja-JP: 日文, 可选值有: zh-CN: 中文, en-US: 英文, ja-JP: 日文
	Texts     []*CreateApprovalExternalInstanceRespDataI18nResourceText `json:"texts,omitempty"`      // 文案 key, value, i18n key 以 @i18n@ 开头； 该字段主要用于做国际化, 允许用户同时传多个语言的文案, 审批中心会根据用户当前的语音环境使用对应的文案, 如果没有传用户当前的语音环境文案, 则会使用默认的语言文案。
	IsDefault bool                                                      `json:"is_default,omitempty"` // 是否默认语言, 默认语言需要包含所有key, 非默认语言如果key不存在会使用默认语言代替
}

// CreateApprovalExternalInstanceRespDataI18nResourceText ...
type CreateApprovalExternalInstanceRespDataI18nResourceText struct {
	Key   string `json:"key,omitempty"`   // 文案key
	Value string `json:"value,omitempty"` // 文案
}

// CreateApprovalExternalInstanceRespDataLink ...
type CreateApprovalExternalInstanceRespDataLink struct {
	PcLink     string `json:"pc_link,omitempty"`     // pc 端的跳转链接, 当用户使用飞书 pc 端时, 使用该字段进行跳转
	MobileLink string `json:"mobile_link,omitempty"` // 移动端 跳转链接, 当用户使用飞书 移动端时, 使用该字段进行跳转
}

// CreateApprovalExternalInstanceRespDataTask ...
type CreateApprovalExternalInstanceRespDataTask struct {
	TaskID            string                                                    `json:"task_id,omitempty"`            // 审批实例内的唯一标识, 用于更新审批任务时定位数据
	UserID            string                                                    `json:"user_id,omitempty"`            // 审批人 user_id, 和 open_id 二者至少填一个。该任务会出现在审批人的【待审批】或【已审批】列表中
	OpenID            string                                                    `json:"open_id,omitempty"`            // 审批人 open_id, 和 user_id 二者至少填一个
	Title             string                                                    `json:"title,omitempty"`              // 审批任务名称
	Links             *CreateApprovalExternalInstanceRespDataTaskLinks          `json:"links,omitempty"`              // 【待审批】或【已审批】中使用的跳转链接, 用于跳转回三方系统pc_link 和 mobile_link 必须填一个, 填写的是哪一端的链接, 即会跳转到该链接, 不受平台影响
	Status            string                                                    `json:"status,omitempty"`             // 任务状态, 可选值有: PENDING: 待审批, APPROVED: 任务同意, REJECTED: 任务拒绝, TRANSFERRED: 任务转交, DONE: 任务通过但审批人未操作；审批人看不到这个任务, 若想要看到, 可以通过抄送该人.
	Extra             string                                                    `json:"extra,omitempty"`              // 扩展 json, 任务结束原因需传complete_reason字段。枚举值与对应说明: approved: 同意, rejected: 拒绝, node_auto_reject: （因逻辑判断产生的）自动拒绝, specific_rollback: 退回（包括退回到发起人、退回到中间任一审批人）, add: 并加签（添加新审批人, 和我一起审批）, add_pre: 前加签（添加新审批人, 在我之前审批）, add_post: 后加签（添加新审批人, 在我之后审批）, delete_assignee: 减签, forward_resign: 转交（转给其他人审批）, recall: 撤销（撤回单据, 单据失效）, delete : 删除审批单, admin_forward: 管理员在后台操作转交, system_forward: 系统自动转交, auto_skip: 自动通过, manual_skip: 手动跳过, submit_again: 重新提交任务, restart: 重新启动流程, others: 其他（作为兜底）
	CreateTime        string                                                    `json:"create_time,omitempty"`        // 任务创建时间, Unix 毫秒时间戳
	EndTime           string                                                    `json:"end_time,omitempty"`           // 任务完成时间: 未结束的审批为 0, Unix 毫秒时间戳
	UpdateTime        string                                                    `json:"update_time,omitempty"`        // task最近更新时间, 用于推送数据版本控制； 更新策略同 instance 中的 update_time
	ActionContext     string                                                    `json:"action_context,omitempty"`     // 操作上下文, 当用户操作时, 回调请求中带上该参数, 用于传递该任务的上下文数据
	ActionConfigs     []*CreateApprovalExternalInstanceRespDataTaskActionConfig `json:"action_configs,omitempty"`     // 任务级别操作配置, 快捷审批目前支持移动端操作
	DisplayMethod     string                                                    `json:"display_method,omitempty"`     // 列表页打开审批任务的方式, 可选值有: BROWSER: 跳转系统默认浏览器打开, SIDEBAR: 飞书中侧边抽屉打开, NORMAL: 飞书内嵌页面打开, TRUSTEESHIP: 以托管模式打开
	ExcludeStatistics bool                                                      `json:"exclude_statistics,omitempty"` // 三方任务支持不纳入效率统计, false: 纳入效率统计, true: 不纳入效率统计
	NodeID            string                                                    `json:"node_id,omitempty"`            // 节点id: 必须同时满足, 一个流程内, 每个节点id唯一。如一个流程下「直属上级」、「隔级上级」等每个节点的Node_id均不一样, 同一个流程定义内, 不同审批实例中的相同节点, Node_id要保持不变。例如张三和李四分别发起了请假申请, 这2个审批实例中的「直属上级」节点的node_id应该保持一致
	NodeName          string                                                    `json:"node_name,omitempty"`          // 节点名称, 如「财务审批」「法务审批」, 支持中英日三种语言。示例: i18n@name。需要在i18n_resources中传该名称对应的国际化文案
}

// CreateApprovalExternalInstanceRespDataTaskActionConfig ...
type CreateApprovalExternalInstanceRespDataTaskActionConfig struct {
	ActionType       string `json:"action_type,omitempty"`        // 操作类型, 每个任务都可以配置2个操作, 会展示审批列表中, 当用户操作时, 回调请求会带上该字段, 表示用户进行了同意操作还是拒绝操作, 可选值有: APPROVE: 同意, REJECT: 拒绝, {KEY}: 任意字符串, 如果使用任意字符串, 则需要提供 action_name
	ActionName       string `json:"action_name,omitempty"`        // 操作名称, i18n key 用于前台展示, 如果 action_type 不是 APPROVAL和REJECT, 则必须提供该字段, 用于展示特定的操作名称
	IsNeedReason     bool   `json:"is_need_reason,omitempty"`     // 是否需要意见, 如果为true, 则用户操作时, 会跳转到 意见填写页面
	IsReasonRequired bool   `json:"is_reason_required,omitempty"` // 审批意见是否必填
	IsNeedAttachment bool   `json:"is_need_attachment,omitempty"` // 意见是否支持上传附件
}

// CreateApprovalExternalInstanceRespDataTaskLinks ...
type CreateApprovalExternalInstanceRespDataTaskLinks struct {
	PcLink     string `json:"pc_link,omitempty"`     // pc 端的跳转链接, 当用户使用飞书 pc 端时, 使用该字段进行跳转
	MobileLink string `json:"mobile_link,omitempty"` // 移动端 跳转链接, 当用户使用飞书 移动端时, 使用该字段进行跳转
}

// CreateApprovalExternalInstanceRespDataTrusteeshipURLs ...
type CreateApprovalExternalInstanceRespDataTrusteeshipURLs struct {
	FormDetailURL       string `json:"form_detail_url,omitempty"`        // 获取表单schema相关数据的url地址
	ActionDefinitionURL string `json:"action_definition_url,omitempty"`  // 表示获取审批操作区数据的url地址
	ApprovalNodeURL     string `json:"approval_node_url,omitempty"`      // 获取审批记录相关数据的url地址
	ActionCallbackURL   string `json:"action_callback_url,omitempty"`    // 进行审批操作时回调的url地址
	PullBusinessDataURL string `json:"pull_business_data_url,omitempty"` // 获取托管动态数据url 地址, 使用该接口时必须要保证历史托管单据的数据中都同步了该接口地址, 如果历史单据中没有该接口需要重新同步历史托管单据的数据来更新该URL。该接口用于飞书审批前端和业务线进行交互使用, 只有使用审批前端的特定组件(由飞书审批前端提供的组件, 并且需要和业务线进行接口交互的组件)才会需要
}

// createApprovalExternalInstanceResp ...
type createApprovalExternalInstanceResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *CreateApprovalExternalInstanceResp `json:"data,omitempty"`
}
