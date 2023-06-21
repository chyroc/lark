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

// DeleteChatTopNotice 撤销会话中的置顶。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 机器人或授权用户必须在群组中
// - 撤销内部群置顶时, 操作者须与群组在同一租户下
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-top_notice/delete_top_notice
// new doc: https://open.feishu.cn/document/server-docs/group/chat/delete_top_notice
func (r *ChatService) DeleteChatTopNotice(ctx context.Context, request *DeleteChatTopNoticeReq, options ...MethodOptionFunc) (*DeleteChatTopNoticeResp, *Response, error) {
	if r.cli.mock.mockChatDeleteChatTopNotice != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#DeleteChatTopNotice mock enable")
		return r.cli.mock.mockChatDeleteChatTopNotice(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "DeleteChatTopNotice",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/top_notice/delete_top_notice",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteChatTopNoticeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatDeleteChatTopNotice mock ChatDeleteChatTopNotice method
func (r *Mock) MockChatDeleteChatTopNotice(f func(ctx context.Context, request *DeleteChatTopNoticeReq, options ...MethodOptionFunc) (*DeleteChatTopNoticeResp, *Response, error)) {
	r.mockChatDeleteChatTopNotice = f
}

// UnMockChatDeleteChatTopNotice un-mock ChatDeleteChatTopNotice method
func (r *Mock) UnMockChatDeleteChatTopNotice() {
	r.mockChatDeleteChatTopNotice = nil
}

// DeleteChatTopNoticeReq ...
type DeleteChatTopNoticeReq struct {
	ChatID string `path:"chat_id" json:"-"` // 待撤销置顶的群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: "oc_5ad11d72b830411d72b836c20"
}

// DeleteChatTopNoticeResp ...
type DeleteChatTopNoticeResp struct {
}

// deleteChatTopNoticeResp ...
type deleteChatTopNoticeResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *DeleteChatTopNoticeResp `json:"data,omitempty"`
}
