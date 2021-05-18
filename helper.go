package lark

import (
	"encoding/json"
	"fmt"
)

func GetErrorCode(err error) int64 {
	if err != nil {
		if e, ok := err.(*Error); ok {
			return e.Code
		}
	}
	return -1
}

// UnwrapMessageContent 接收消息解析
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content
func UnwrapMessageContent(msgType MsgType, content string) (*MessageContent, error) {
	msg := &MessageContent{
		Text:               new(MessageContentText),
		Image:              new(MessageContentImage),
		File:               new(MessageContentFile),
		Audio:              new(MessageContentAudio),
		Media:              new(MessageContentMedia),
		Sticker:            new(MessageContentSticker),
		RedBag:             new(MessageContentRedBag),
		ShareCalendarEvent: new(MessageContentShareCalendarEvent),
		ShareChat:          new(MessageContentShareChat),
		ShareUser:          new(MessageContentShareUser),
		System:             new(MessageContentSystem),
		Location:           new(MessageContentLocation),
		VideoChat:          new(MessageContentVideoChat),
	}
	switch msgType {
	case MsgTypeText:
		return msg, unmarshalMessageContent(msgType, content, msg.Text)
	case MsgTypeImage:
		return msg, unmarshalMessageContent(msgType, content, msg.Image)
	case MsgTypeFile:
		return msg, unmarshalMessageContent(msgType, content, msg.File)
	case MsgTypeAudio:
		return msg, unmarshalMessageContent(msgType, content, msg.Audio)
	case MsgTypeMedia:
		return msg, unmarshalMessageContent(msgType, content, msg.Media)
	case MsgTypeSticker:
		return msg, unmarshalMessageContent(msgType, content, msg.Sticker)
		// case MsgTypeRedBag:
		// 	return msg, unmarshalMessageContent(msgType,content, msg.RedBag)
		// case MsgTypeShareCalendarEvent:
		// 	return msg, unmarshalMessageContent(msgType,content, msg.ShareCalendarEvent)
	case MsgTypeShareChat:
		return msg, unmarshalMessageContent(msgType, content, msg.ShareChat)
	case MsgTypeShareUser:
		return msg, unmarshalMessageContent(msgType, content, msg.ShareUser)
		// case MsgTypeSystem:
		// 	return msg, unmarshalMessageContent(msgType,content, msg.System)
		// case MsgTypeLocation:
		// 	return msg, unmarshalMessageContent(msgType,content, msg.Location)
		// case MsgTypeVideoChat:
		// 	return msg, unmarshalMessageContent(msgType,content, msg.VideoChat)
	}
	return nil, fmt.Errorf("unknown message type: %s", msgType)
}

func unmarshalMessageContent(msgType MsgType, content string, res interface{}) error {
	if err := json.Unmarshal([]byte(content), res); err != nil {
		return fmt.Errorf("invalid content: %s, type: %v", content, msgType)
	}
	return nil
}

func jsonString(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
