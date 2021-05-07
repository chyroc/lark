package lark

import (
	"context"
)

// EventV2IMChatMemberBotDeletedV1
//
// 机器人被移出群聊时触发此事件。
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 [即时通讯] 分类下的 [机器人被移出群] 事件
// - 事件会向被移出群的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-bot/events/deleted
func (r *EventCallbackAPI) HandlerEventV2IMChatMemberBotDeletedV1(f eventV2IMChatMemberBotDeletedV1Handler) {
	r.cli.eventHandler.eventV2IMChatMemberBotDeletedV1Handler = f
}

type eventV2IMChatMemberBotDeletedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMChatMemberBotDeletedV1) (string, error)

type EventV2IMChatMemberBotDeletedV1 struct {
	ChatID     string                                     `json:"chat_id,omitempty"`     // 群组 ID
	OperatorID *EventV2IMChatMemberBotDeletedV1OperatorID `json:"operator_id,omitempty"` // 用户 ID
}

type EventV2IMChatMemberBotDeletedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
