package lark

type EventType string

const EventTypeMessageReceive EventType = "im.message.receive_v1"

type EventMessageReceive struct {
	Sender  *EventMessageReceiveSender  `json:"sender,omitempty"`  // 事件的发送者
	Message *EventMessageReceiveMessage `json:"message,omitempty"` // 事件中包含的消息内容
}

type EventMessageReceiveSender struct {
	SenderID   *EventMessageReceiveSenderID `json:"sender_id,omitempty"`   // 用户 ID
	SenderType string                       `json:"sender_type,omitempty"` // 消息发送者类型。目前只支持用户发送的消息
}

type EventMessageReceiveSenderID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventMessageReceiveMessage struct {
	MessageID   string                        `json:"message_id,omitempty"`   // 消息的 open_message_id
	RootID      string                        `json:"root_id,omitempty"`      // 回复消息 根 id
	ParentID    string                        `json:"parent_id,omitempty"`    // 回复消息 父 id
	CreateTime  string                        `json:"create_time,omitempty"`  // 消息发送时间 毫秒
	ChatID      string                        `json:"chat_id,omitempty"`      // 消息所在的群组 id
	ChatType    string                        `json:"chat_type,omitempty"`    // 消息所在的群组类型,单聊或群聊
	MessageType string                        `json:"message_type,omitempty"` // 消息类型
	Content     string                        `json:"content,omitempty"`      // 消息内容, json 格式 ,[各类型消息Content](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	Mentions    []*EventMessageReceiveMention `json:"mentions,omitempty"`     // 被提及用户的信息
}

type EventMessageReceiveMention struct {
	Key  string                        `json:"key,omitempty"`  // mention key
	Id   *EventMessageReceiveMentionID `json:"id,omitempty"`   // 用户 ID
	Name string                        `json:"name,omitempty"` // 用户姓名
}

type EventMessageReceiveMentionID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
