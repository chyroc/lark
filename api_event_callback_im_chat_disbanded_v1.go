package lark

import (
	"context"
)

// EventIMChatDisbandedV1
//
// 群组被解散后触发此事件。
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 ==即时通讯== 分类下的 ==解散群== 事件
// - 事件会向群内订阅了该事件的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/events/disbanded
func (r *EventCallbackAPI) HandlerEventIMChatDisbandedV1(f eventIMChatDisbandedV1Handler) {
	r.cli.eventHandler.eventIMChatDisbandedV1Handler = f
}

type eventIMChatDisbandedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventIMChatDisbandedV1) (string, error)

type EventIMChatDisbandedV1 struct {
	ChatID     string                            `json:"chat_id,omitempty"`     // 群组 ID
	OperatorID *EventIMChatDisbandedV1OperatorID `json:"operator_id,omitempty"` // 用户 ID
}

type EventIMChatDisbandedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
