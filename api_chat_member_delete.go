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

// DeleteChatMember 将用户或机器人移出群聊。
//
// 注意事项:
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 用户或机器人在任何条件下均可移除自己出群（即主动退群）
// - 仅有群主/管理员 或 创建群组并且具备 [更新应用所创建群的群信息] 权限的机器人, 可以移除其他用户或者机器人
// - 每次请求, 最多移除50个用户或者5个机器人
// - 操作内部群时, 操作者须与群组在同一租户下
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/delete
// new doc: https://open.feishu.cn/document/server-docs/group/chat-member/delete
func (r *ChatService) DeleteChatMember(ctx context.Context, request *DeleteChatMemberReq, options ...MethodOptionFunc) (*DeleteChatMemberResp, *Response, error) {
	if r.cli.mock.mockChatDeleteChatMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#DeleteChatMember mock enable")
		return r.cli.mock.mockChatDeleteChatMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "DeleteChatMember",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/members",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteChatMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatDeleteChatMember mock ChatDeleteChatMember method
func (r *Mock) MockChatDeleteChatMember(f func(ctx context.Context, request *DeleteChatMemberReq, options ...MethodOptionFunc) (*DeleteChatMemberResp, *Response, error)) {
	r.mockChatDeleteChatMember = f
}

// UnMockChatDeleteChatMember un-mock ChatDeleteChatMember method
func (r *Mock) UnMockChatDeleteChatMember() {
	r.mockChatDeleteChatMember = nil
}

// DeleteChatMemberReq ...
type DeleteChatMemberReq struct {
	ChatID       string   `path:"chat_id" json:"-"`         // 群 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 注意: 仅支持群模式为`group`、`topic`的群组ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
	MemberIDType *IDType  `query:"member_id_type" json:"-"` // 出群成员ID类型, 注意: 移除机器人请使用 [app_id], 示例值: "open_id", 可选值有: user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), app_id: 飞书开放平台应用的唯一标识。在创建应用时, 由系统自动生成, 用户不能自行修改。[了解更多: 如何获取应用的 App ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id), 默认值: `open_id`
	IDList       []string `json:"id_list,omitempty"`        // 成员列表。移除群内的用户时推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid)；移除群内的机器人时需填写应用的App ID, 请参考[如何获取应用的 App ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-app-id), 注意: 成员列表不可为空, 列表中填写的成员ID类型应与 [member_id_type] 参数中选择的类型相对应, 示例值: ["4d7a3c6g"]
}

// DeleteChatMemberResp ...
type DeleteChatMemberResp struct {
	InvalidIDList []string `json:"invalid_id_list,omitempty"` // 无效成员列表
}

// deleteChatMemberResp ...
type deleteChatMemberResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *DeleteChatMemberResp `json:"data,omitempty"`
}
