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

// DeleteMessage 机器人撤回机器人自己发送的消息或群主撤回群内消息。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability), 撤回消息时机器人仍需要在会话内
// - 机器人可以撤回单聊和群组内, 自己发送 且 发送时间不超过租户管理员配置的可撤回时限的消息（默认为24小时）
// - 若机器人要撤回群内他人发送的消息, 则机器人必须是该群的群主、管理员 或者 创建者, 且消息发送时间不超过1年
// - 无法撤回通过「[批量发送消息](https://open.feishu.cn/document/ukTMukTMukTM/ucDO1EjL3gTNx4yN4UTM)」接口发送的消息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/delete
// new doc: https://open.feishu.cn/document/server-docs/im-v1/message/delete
func (r *MessageService) DeleteMessage(ctx context.Context, request *DeleteMessageReq, options ...MethodOptionFunc) (*DeleteMessageResp, *Response, error) {
	if r.cli.mock.mockMessageDeleteMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#DeleteMessage mock enable")
		return r.cli.mock.mockMessageDeleteMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "DeleteMessage",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageDeleteMessage mock MessageDeleteMessage method
func (r *Mock) MockMessageDeleteMessage(f func(ctx context.Context, request *DeleteMessageReq, options ...MethodOptionFunc) (*DeleteMessageResp, *Response, error)) {
	r.mockMessageDeleteMessage = f
}

// UnMockMessageDeleteMessage un-mock MessageDeleteMessage method
func (r *Mock) UnMockMessageDeleteMessage() {
	r.mockMessageDeleteMessage = nil
}

// DeleteMessageReq ...
type DeleteMessageReq struct {
	MessageID string `path:"message_id" json:"-"` // 待撤回的消息的ID, 示例值: "om_dc13264520392913993dd051dba21dcf"
}

// DeleteMessageResp ...
type DeleteMessageResp struct {
}

// deleteMessageResp ...
type deleteMessageResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMessageResp `json:"data,omitempty"`
}
