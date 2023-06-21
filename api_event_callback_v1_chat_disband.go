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

// EventV1ChatDisband 如果你希望订阅机器人进出群、群内有人@机器人事件, 请前往[机器人进群](https://open.feishu.cn/document/ukTMukTMukTM/ugzMugzMugzM/event/bot-added-to-group) 或 [机器人退群](https://open.feishu.cn/document/ukTMukTMukTM/ucDO04yN4QjL3gDN)。
//
// 为了更好地提升该事件的安全性, 我们对其进行了升级, 请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/disbanded)
// 群聊被解散后触发此事件。
// * 特殊说明: 只有开启机器人能力并且机器人所在的群被解散时才能触发此事件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQDOwUjL0gDM14CN4ATN/event/group-closed
// new doc:
func (r *EventCallbackService) HandlerEventV1ChatDisband(f EventV1ChatDisbandHandler) {
	r.cli.eventHandler.eventV1ChatDisbandHandler = f
}

// EventV1ChatDisbandHandler event EventV1ChatDisband handler
type EventV1ChatDisbandHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1ChatDisband) (string, error)

// EventV1ChatDisband ...
type EventV1ChatDisband struct {
	AppID     string                           `json:"app_id,omitempty"`  // 如: cli_9c8609450f78d102
	ChatID    string                           `json:"chat_id,omitempty"` // 如: oc_9f2df3c095c9395334bb6e70ced0fa83
	Operator  *EventV1ChatDisbandEventOperator `json:"operator,omitempty"`
	TenantKey string                           `json:"tenant_key,omitempty"` // 如: 736588c9260f175d
	Type      string                           `json:"type,omitempty"`       // 如: chat_disband
}

// EventV1ChatDisbandEventOperator ...
type EventV1ChatDisbandEventOperator struct {
	OpenID string `json:"open_id,omitempty"` // 如: ou_18eac85d35a26f989317ad4f02e8bbbb
	UserID string `json:"user_id,omitempty"` // 如: ca51d83b
}
