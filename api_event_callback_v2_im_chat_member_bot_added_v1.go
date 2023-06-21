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

// EventV2IMChatMemberBotAddedV1 机器人被用户添加至群聊时触发此事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=im&version=v1&resource=chat.member.bot&event=added)
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 需要订阅 [消息与群组] 分类下的 [机器人进群] 事件
// - 事件会向进群的机器人进行推送
// - 机器人邀请机器人不会触发事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-bot/events/added
// new doc: https://open.feishu.cn/document/server-docs/group/chat-member/event/added-2
func (r *EventCallbackService) HandlerEventV2IMChatMemberBotAddedV1(f EventV2IMChatMemberBotAddedV1Handler) {
	r.cli.eventHandler.eventV2IMChatMemberBotAddedV1Handler = f
}

// EventV2IMChatMemberBotAddedV1Handler event EventV2IMChatMemberBotAddedV1 handler
type EventV2IMChatMemberBotAddedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMChatMemberBotAddedV1) (string, error)

// EventV2IMChatMemberBotAddedV1 ...
type EventV2IMChatMemberBotAddedV1 struct {
	ChatID            string                                   `json:"chat_id,omitempty"`             // 群组 ID, 详情参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description)
	OperatorID        *EventV2IMChatMemberBotAddedV1OperatorID `json:"operator_id,omitempty"`         // 用户 ID
	External          bool                                     `json:"external,omitempty"`            // 是否是外部群
	OperatorTenantKey string                                   `json:"operator_tenant_key,omitempty"` // 操作者的租户Key, 为租户在飞书上的唯一标识, 用来换取对应的tenant_access_token, 也可以用作租户在应用中的唯一标识
	Name              string                                   `json:"name,omitempty"`                // 群名称
	I18nNames         *I18nNames                               `json:"i18n_names,omitempty"`          // 群国际化名称
}

// EventV2IMChatMemberBotAddedV1OperatorID ...
type EventV2IMChatMemberBotAddedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
