package lark

import (
	"context"
)

// EventIMChatMemberUserDeletedV1
//
// 用户主动退群或被移出群聊时推送事件
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)并且机器人所在群发生上述变化
// - 机器人需要订阅 [即时通讯] 分类下的 [用户主动退群或被移出群聊] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/deleted
func (r *EventCallbackAPI) HandlerEventIMChatMemberUserDeletedV1(f eventIMChatMemberUserDeletedV1Handler) {
	r.cli.eventHandler.eventIMChatMemberUserDeletedV1Handler = f
}

type eventIMChatMemberUserDeletedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventIMChatMemberUserDeletedV1) (string, error)

type EventIMChatMemberUserDeletedV1 struct {
	ChatID     string                                    `json:"chat_id,omitempty"`     // 群组 ID
	OperatorID *EventIMChatMemberUserDeletedV1OperatorID `json:"operator_id,omitempty"` // 用户 ID
	Users      []*EventIMChatMemberUserDeletedV1User     `json:"users,omitempty"`       // 被移除用户列表
}

type EventIMChatMemberUserDeletedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventIMChatMemberUserDeletedV1User struct {
	Name   string                                    `json:"name,omitempty"`    // 用户名字
	UserID *EventIMChatMemberUserDeletedV1UserUserID `json:"user_id,omitempty"` // 用户 ID
}

type EventIMChatMemberUserDeletedV1UserUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 userid
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
