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

// UpdateChatTopNotice 更新会话中的群置顶信息, 可以将群中的某一条消息, 或者群公告置顶显示。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人或授权用户必须在群组中
// - 更新内部群置顶时, 操作者须与群组在同一租户下
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-top_notice/put_top_notice
// new doc: https://open.feishu.cn/document/server-docs/group/chat/put_top_notice
func (r *ChatService) UpdateChatTopNotice(ctx context.Context, request *UpdateChatTopNoticeReq, options ...MethodOptionFunc) (*UpdateChatTopNoticeResp, *Response, error) {
	if r.cli.mock.mockChatUpdateChatTopNotice != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#UpdateChatTopNotice mock enable")
		return r.cli.mock.mockChatUpdateChatTopNotice(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "UpdateChatTopNotice",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/top_notice/put_top_notice",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateChatTopNoticeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatUpdateChatTopNotice mock ChatUpdateChatTopNotice method
func (r *Mock) MockChatUpdateChatTopNotice(f func(ctx context.Context, request *UpdateChatTopNoticeReq, options ...MethodOptionFunc) (*UpdateChatTopNoticeResp, *Response, error)) {
	r.mockChatUpdateChatTopNotice = f
}

// UnMockChatUpdateChatTopNotice un-mock ChatUpdateChatTopNotice method
func (r *Mock) UnMockChatUpdateChatTopNotice() {
	r.mockChatUpdateChatTopNotice = nil
}

// UpdateChatTopNoticeReq ...
type UpdateChatTopNoticeReq struct {
	ChatID        string                                 `path:"chat_id" json:"-"`          // 待修改置顶的群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_5ad11d72b830411d72b836c20"
	ChatTopNotice []*UpdateChatTopNoticeReqChatTopNotice `json:"chat_top_notice,omitempty"` // 要进行发布的群置顶
}

// UpdateChatTopNoticeReqChatTopNotice ...
type UpdateChatTopNoticeReqChatTopNotice struct {
	ActionType *string `json:"action_type,omitempty"` // 置顶的类型, 注意: 选择 [消息类型] 时必须填写`message_id`字段, 选择 [群公告类型] 时填写的`message_id`将被忽略, 示例值: "2", 可选值有: 1: 消息类型, 2: 群公告类型, 默认值: `2`
	MessageID  *string `json:"message_id,omitempty"`  // 消息ID, 详情参见[消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2), 示例值: "om_dc13264520392913993dd051dba21dcf"
}

// UpdateChatTopNoticeResp ...
type UpdateChatTopNoticeResp struct {
}

// updateChatTopNoticeResp ...
type updateChatTopNoticeResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateChatTopNoticeResp `json:"data,omitempty"`
}
