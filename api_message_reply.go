package lark

import (
	"context"
)

// ReplyRawMessage 回复指定消息，支持文本、富文本、卡片、群名片、个人名片、图片、视频、文件等多种消息类型。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 回复私聊消息，需要机器人对用户有可见性
// - 回复群组消息，需要机器人在群中
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/reply
func (r *MessageAPI) ReplyRawMessage(ctx context.Context, request *ReplyRawMessageReq) (*ReplyRawMessageResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages/:message_id/reply",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(replyRawMessageResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Message", "ReplyRawMessage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type ReplyRawMessageReq struct {
	MessageID string  `path:"message_id" json:"-"` // 待回复的消息的ID, 示例值："om_dc13264520392913993dd051dba21dcf"
	Content   string  `json:"content,omitempty"`   // 消息内容 json 格式，格式说明参考: [发送消息content说明](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/im-v1/message/create_json), 示例值："{\"text\":\"<at user_id=\"ou_155184d1e73cbfb8973e5a9e698e74f2\">Tom</at> test content\"}"
	MsgType   MsgType `json:"msg_type,omitempty"`  // 消息类型，包括：text、post、image、file、audio、media、sticker、interactive、share_card、share_user, 示例值："text"
}

type replyRawMessageResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *ReplyRawMessageResp `json:"data,omitempty"` //
}

type ReplyRawMessageResp struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id open_message_id
	RootID         string       `json:"root_id,omitempty"`          // 根消息id open_message_id
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id open_message_id
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 text post card image等等
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳(毫秒)
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者，可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容，json结构，格式说明参考： [消息content说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被艾特的人或应用的id
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 上一层级的消息id open_message_id
}
