package lark

import (
	"context"
)

// EventIMChatMemberUserWithdrawnV1
//
// 撤销拉用户进群后触发此事件。
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 [即时通讯] 分类下的 [撤销拉用户进群] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/withdrawn
func (r *EventCallbackAPI) HandlerEventIMChatMemberUserWithdrawnV1(f eventIMChatMemberUserWithdrawnV1Handler) {
	r.cli.eventHandler.eventIMChatMemberUserWithdrawnV1Handler = f
}

type eventIMChatMemberUserWithdrawnV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventIMChatMemberUserWithdrawnV1) (string, error)

type EventIMChatMemberUserWithdrawnV1 struct {
	ChatID     string                                      `json:"chat_id,omitempty"`     // 群组 ID
	OperatorID *EventIMChatMemberUserWithdrawnV1OperatorID `json:"operator_id,omitempty"` // 用户 ID
	Users      []*EventIMChatMemberUserWithdrawnV1User     `json:"users,omitempty"`       // 被撤销加群的用户列表
}

type EventIMChatMemberUserWithdrawnV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventIMChatMemberUserWithdrawnV1User struct {
	Name   string                                      `json:"name,omitempty"`    // 用户名字
	UserID *EventIMChatMemberUserWithdrawnV1UserUserID `json:"user_id,omitempty"` // 用户 ID
}

type EventIMChatMemberUserWithdrawnV1UserUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
