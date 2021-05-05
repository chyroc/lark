package lark

import (
	"context"
)

// EventIMChatMemberUserAddedV1
//
// 新用户进群触发此事件。
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 [即时通讯] 分类下的 [用户进群] 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-user/events/added
func (r *EventCallbackAPI) HandlerEventIMChatMemberUserAddedV1(f eventIMChatMemberUserAddedV1Handler) {
	r.cli.eventHandler.eventIMChatMemberUserAddedV1Handler = f
}

type eventIMChatMemberUserAddedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventIMChatMemberUserAddedV1) (string, error)

type EventIMChatMemberUserAddedV1 struct {
	ChatID     string                                  `json:"chat_id,omitempty"`     // 群组 ID
	OperatorID *EventIMChatMemberUserAddedV1OperatorID `json:"operator_id,omitempty"` // 用户 ID
	Users      []*EventIMChatMemberUserAddedV1User     `json:"users,omitempty"`       // 被添加的用户列表
}

type EventIMChatMemberUserAddedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventIMChatMemberUserAddedV1User struct {
	Name   string                                  `json:"name,omitempty"`    // 用户名字
	UserID *EventIMChatMemberUserAddedV1UserUserID `json:"user_id,omitempty"` // 用户 ID
}

type EventIMChatMemberUserAddedV1UserUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
