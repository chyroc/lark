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

// SendEphemeralMessage 用于机器人在群会话中发送仅指定用户可见的消息卡片。卡片上将展示"仅对你可见"标识。
//
// ![image.png](//sf3-cn.feishucdn.com/obj/open-platform-opendoc/b0ec0ce45942463381457edc7b62e144_RXYCFtfUtb.png?lazyload=true&width=1592&height=486)
// ## 使用场景
// 临时消息卡片多用于群聊中用户与机器人交互的中间态。例如在群聊中用户需要使用待办事项类bot创建一条提醒, bot 发送了可设置提醒日期和提醒内容的一张可交互的消息卡片, 此卡片在没有设置为临时卡片的情况下为群内全员可见, 即群内可看见该用户与 bot 交互的过程。而设置为临时卡片后, 交互过程仅该用户可见, 群内其他成员只会看到最终设置完成的提醒卡片。
// 通过临时消息卡片, 可以减少消息对群聊中不相关用户的打扰, 有效降低群消息的噪声。
// 需要启用[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)；需要机器人在会话群里。
// - 仅触发临时卡片的用户自己可见。
// - 不支持转发。
// - 只能在群聊使用。
// - 仅在用户处于在线状态的飞书客户端上可见。
// - 临时消息卡片的[呈现能力](https://open.feishu.cn/document/ukTMukTMukTM/uEjNwUjLxYDM14SM2ATN)、[交互能力](https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN)与消息卡片一致。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uETOyYjLxkjM24SM5IjN
// new doc: https://open.feishu.cn/document/server-docs/im-v1/message-card/send-message-cards-that-are-only-visible-to-certain-people
func (r *MessageService) SendEphemeralMessage(ctx context.Context, request *SendEphemeralMessageReq, options ...MethodOptionFunc) (*SendEphemeralMessageResp, *Response, error) {
	if r.cli.mock.mockMessageSendEphemeralMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#SendEphemeralMessage mock enable")
		return r.cli.mock.mockMessageSendEphemeralMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "SendEphemeralMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/ephemeral/v1/send",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(sendEphemeralMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageSendEphemeralMessage mock MessageSendEphemeralMessage method
func (r *Mock) MockMessageSendEphemeralMessage(f func(ctx context.Context, request *SendEphemeralMessageReq, options ...MethodOptionFunc) (*SendEphemeralMessageResp, *Response, error)) {
	r.mockMessageSendEphemeralMessage = f
}

// UnMockMessageSendEphemeralMessage un-mock MessageSendEphemeralMessage method
func (r *Mock) UnMockMessageSendEphemeralMessage() {
	r.mockMessageSendEphemeralMessage = nil
}

// SendEphemeralMessageReq ...
type SendEphemeralMessageReq struct {
	ChatID  string              `json:"chat_id,omitempty"`  // 发送临时消息的群ID可通过事件推送获取
	OpenID  string              `json:"open_id,omitempty"`  // 指定发送临时消息卡片的用户, 其他人将无法看到临时消息卡片；只需要填 open_id、email、user_id中的一个即可, 推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid) （服务端依次读取字段的顺序为 open_id > user_id > email）
	UserID  string              `json:"user_id,omitempty"`  // 指定发送临时消息卡片的用户, 其他人将无法看到临时消息卡片；只需要填 open_id、email、user_id中的一个即可, 推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid) （服务端依次读取字段的顺序为 open_id > user_id > email）
	Email   string              `json:"email,omitempty"`    // 指定发送临时消息卡片的用户, 其他人将无法看到临时消息卡片；只需要填 open_id、email、user_id中的一个即可, 推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid) （服务端依次读取字段的顺序为 open_id > user_id > email）
	MsgType MsgType             `json:"msg_type,omitempty"` // 消息的类型, 此处固定填 "interactive"
	Card    *MessageContentCard `json:"card,omitempty"`     // 消息卡片的描述内容, 具体参考 [基础结构](https://open.feishu.cn/document/ukTMukTMukTM/uEjNwUjLxYDM14SM2ATN)
}

// SendEphemeralMessageResp ...
type SendEphemeralMessageResp struct {
	MessageID string `json:"message_id,omitempty"` // 消息 ID
}

// sendEphemeralMessageResp ...
type sendEphemeralMessageResp struct {
	Code int64                     `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 返回码描述
	Data *SendEphemeralMessageResp `json:"data,omitempty"`
}
