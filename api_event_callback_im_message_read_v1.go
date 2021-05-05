package lark

import (
	"context"
)

// EventIMMessageReadV1
//
// 用户阅读机器人发送的单聊消息后触发此事件。
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 ==即时通讯== 分类下的 ==消息已读== 事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/message_read
func (r *EventCallbackAPI) HandlerEventIMMessageReadV1(f eventIMMessageReadV1Handler) {
	r.cli.eventHandler.eventIMMessageReadV1Handler = f
}

type eventIMMessageReadV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventIMMessageReadV1) (string, error)

type EventIMMessageReadV1 struct {
	Reader        *EventIMMessageReadV1Reader `json:"reader,omitempty"`          // -
	MessageIDList []string                    `json:"message_id_list,omitempty"` // 消息列表
}

type EventIMMessageReadV1Reader struct {
	ReaderID *EventIMMessageReadV1ReaderReaderID `json:"reader_id,omitempty"` // 用户 ID
	ReadTime string                              `json:"read_time,omitempty"` // 阅读时间
}

type EventIMMessageReadV1ReaderReaderID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
