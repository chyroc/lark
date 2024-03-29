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

// ForwardThreadMessage 向用户、群聊或话题转发一个话题。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人须在要转发话题所在的群中且可见该话题
// - 向用户转发话题, 需要机器人对用户有[可用性](https://open.feishu.cn/document/home/introduction-to-scope-and-authorization/availability)
// - 向群组中转发话题, 需要机器人在群组中
// - 转发话题生成的新消息的消息内容为固定值[Merged and Forwarded Message], 其子消息可以使用[获取指定消息的内容](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/get)接口获取
// - 为避免对用户造成打扰, 向同一用户发送消息的限频为 [5 QPS], 向同一群组发送消息的限频为群内机器人共享 [5 QPS]
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/thread/forward
func (r *MessageService) ForwardThreadMessage(ctx context.Context, request *ForwardThreadMessageReq, options ...MethodOptionFunc) (*ForwardThreadMessageResp, *Response, error) {
	if r.cli.mock.mockMessageForwardThreadMessage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#ForwardThreadMessage mock enable")
		return r.cli.mock.mockMessageForwardThreadMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "ForwardThreadMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/threads/:thread_id/forward",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(forwardThreadMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageForwardThreadMessage mock MessageForwardThreadMessage method
func (r *Mock) MockMessageForwardThreadMessage(f func(ctx context.Context, request *ForwardThreadMessageReq, options ...MethodOptionFunc) (*ForwardThreadMessageResp, *Response, error)) {
	r.mockMessageForwardThreadMessage = f
}

// UnMockMessageForwardThreadMessage un-mock MessageForwardThreadMessage method
func (r *Mock) UnMockMessageForwardThreadMessage() {
	r.mockMessageForwardThreadMessage = nil
}

// ForwardThreadMessageReq ...
type ForwardThreadMessageReq struct {
	ThreadID      string  `path:"thread_id" json:"-"`        // 要转发的话题ID, 详见: [话题介绍](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/thread-introduction), </md-enum-item>, 示例值: "omt_dc132645203"
	ReceiveIDType IDType  `query:"receive_id_type" json:"-"` // 消息接收者id类型 open_id/user_id/union_id/email/chat_id/thread_id, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。, user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。, union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。, email: 以用户的真实邮箱来标识用户。, chat_id: 以群ID来标识群聊。[了解更多: 如何获取群ID ](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), thread_id: 以话题ID来标识话题。了解更多: [话题介绍](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/thread-introduction),
	UUID          *string `query:"uuid" json:"-"`            // 由开发者生成的唯一字符串序列, 用于转发消息请求去重；持有相同uuid的请求在1小时内向同一个目标的转发只可成功一次, 示例值: b13g2t38-1jd2-458b-8djf-dtbca5104204, 最大长度: `50` 字符
	ReceiveID     string  `json:"receive_id,omitempty"`      // 依据receive_id_type的值, 填写对应的转发目标的ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
}

// ForwardThreadMessageResp ...
type ForwardThreadMessageResp struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	RootID         string       `json:"root_id,omitempty"`          // 根消息id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id, 用于回复消息场景, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
	ThreadID       string       `json:"thread_id,omitempty"`        // 消息所属的话题 ID（不返回说明该消息非话题消息）, 说明参见: [话题介绍](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/thread-introduction)
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 包括: text、post、image、file、audio、media、sticker、interactive、share_chat、share_user等, 类型定义请参考[接收消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳（毫秒）
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳（毫秒）
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者, 可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被@的用户或机器人的id列表
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 合并转发消息中, 上一层级的消息id message_id, 说明参见: [消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2)
}

// forwardThreadMessageResp ...
type forwardThreadMessageResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *ForwardThreadMessageResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
