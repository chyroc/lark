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

// DeleteChat 解散群组。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 如果使用tenant_access_token, 需要机器人符合以下任一情况才可解散群:
// - 机器人是群主
// - 机器人是群的创建者且具备[更新应用所创建群的群信息]权限
// - 如果使用user_access_token, 需要对应的用户是群主才可解散群
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/delete
// new doc: https://open.feishu.cn/document/server-docs/group/chat/delete
func (r *ChatService) DeleteChat(ctx context.Context, request *DeleteChatReq, options ...MethodOptionFunc) (*DeleteChatResp, *Response, error) {
	if r.cli.mock.mockChatDeleteChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#DeleteChat mock enable")
		return r.cli.mock.mockChatDeleteChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "DeleteChat",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatDeleteChat mock ChatDeleteChat method
func (r *Mock) MockChatDeleteChat(f func(ctx context.Context, request *DeleteChatReq, options ...MethodOptionFunc) (*DeleteChatResp, *Response, error)) {
	r.mockChatDeleteChat = f
}

// UnMockChatDeleteChat un-mock ChatDeleteChat method
func (r *Mock) UnMockChatDeleteChat() {
	r.mockChatDeleteChat = nil
}

// DeleteChatReq ...
type DeleteChatReq struct {
	ChatID string `path:"chat_id" json:"-"` // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 仅支持群模式为`group`的群组ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
}

// DeleteChatResp ...
type DeleteChatResp struct {
}

// deleteChatResp ...
type deleteChatResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteChatResp `json:"data,omitempty"`
}
