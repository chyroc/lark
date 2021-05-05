package lark

import (
	"context"
)

// EventIMMessageReceiveV1
//
// 机器人接收到用户发送的消息后触发此事件。
//
//
//
//
//
//
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)  ，具备==获取单聊、群组消息== 或 ==获取与发送单聊、群组消息==权限，并订阅 ==即时通讯== 分类下的 ==接收消息== 事件才可接收推送
// - 同时，将根据应用具备的权限，判断可推送的信息：
// - 当具备==获取用户发给机器人的单聊消息== 权限时，可接收与机器人单聊会话中，发送给机器人的所有消息
// - 当具备==获取群组中用户@机器人的消息== 权限，可接收机器人所在群聊中 @ 机器人的消息
// - 当具备==获取群组中所有的消息== 权限，可接收机器人所在群聊中所有的消息
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive
func (r *EventCallbackAPI) HandlerEventIMMessageReceiveV1(f eventIMMessageReceiveV1Handler) {
	r.cli.eventHandler.eventIMMessageReceiveV1Handler = f
}

type eventIMMessageReceiveV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventIMMessageReceiveV1) (string, error)

type EventIMMessageReceiveV1 struct {
	Sender  *EventIMMessageReceiveV1Sender  `json:"sender,omitempty"`  // 事件的发送者
	Message *EventIMMessageReceiveV1Message `json:"message,omitempty"` // 事件中包含的消息内容
}

type EventIMMessageReceiveV1Sender struct {
	SenderID   *EventIMMessageReceiveV1SenderSenderID `json:"sender_id,omitempty"`   // 用户 ID
	SenderType string                                 `json:"sender_type,omitempty"` // 消息发送者类型。目前只支持用户发送的消息
}

type EventIMMessageReceiveV1SenderSenderID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventIMMessageReceiveV1Message struct {
	MessageID   string                                   `json:"message_id,omitempty"`   // 消息的 open_message_id
	RootID      string                                   `json:"root_id,omitempty"`      // 回复消息 根 id
	ParentID    string                                   `json:"parent_id,omitempty"`    // 回复消息 父 id
	CreateTime  string                                   `json:"create_time,omitempty"`  // 消息发送时间 毫秒
	ChatID      string                                   `json:"chat_id,omitempty"`      // 消息所在的群组 id
	ChatType    string                                   `json:"chat_type,omitempty"`    // 消息所在的群组类型,单聊或群聊
	MessageType string                                   `json:"message_type,omitempty"` // 消息类型
	Content     string                                   `json:"content,omitempty"`      // 消息内容, json 格式 ,[各类型消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	Mentions    []*EventIMMessageReceiveV1MessageMention `json:"mentions,omitempty"`     // 被提及用户的信息
}

type EventIMMessageReceiveV1MessageMention struct {
	Key  string                                   `json:"key,omitempty"`  // mention key
	ID   *EventIMMessageReceiveV1MessageMentionID `json:"id,omitempty"`   // 用户 ID
	Name string                                   `json:"name,omitempty"` // 用户姓名
}

type EventIMMessageReceiveV1MessageMentionID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
