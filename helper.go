/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package lark

import (
	"encoding/json"
	"fmt"
)

// Error ...
type Error struct {
	Scope    string
	FuncName string
	Code     int64
	Msg      string
}

// Error ...
func (r *Error) Error() string {
	if r.Code == 0 {
		return ""
	}
	return fmt.Sprintf("request %s#%s failed: code: %d, msg: %s", r.Scope, r.FuncName, r.Code, r.Msg)
}

// NewError ...
func NewError(scope, funcName string, code int64, msg string) error {
	return &Error{
		Scope:    scope,
		FuncName: funcName,
		Code:     code,
		Msg:      msg,
	}
}

// GetErrorCode ...
func GetErrorCode(err error) int64 {
	if err != nil {
		if e, ok := err.(*Error); ok {
			return e.Code
		}
	}
	return 0
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
		Post:               new(MessageContentPost),
	}
	switch msgType {
	case MsgTypeText:
		return msg, unmarshalMessageContent(msgType, content, msg.Text)
	case MsgTypePost:
		return msg, unmarshalMessageContent(msgType, content, msg.Post)
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
	// case MsgTypeInteractive:
	// 	return msg, unmarshalMessageContent(msgType, content, msg./)
	case MsgTypeRedBag:
		return msg, unmarshalMessageContent(msgType, content, msg.RedBag)
	case MsgTypeShareCalendarEvent:
		return msg, unmarshalMessageContent(msgType, content, msg.ShareCalendarEvent)
	case MsgTypeShareChat:
		return msg, unmarshalMessageContent(msgType, content, msg.ShareChat)
	case MsgTypeShareUser:
		return msg, unmarshalMessageContent(msgType, content, msg.ShareUser)
	case MsgTypeSystem:
		return msg, unmarshalMessageContent(msgType, content, msg.System)
	case MsgTypeLocation:
		return msg, unmarshalMessageContent(msgType, content, msg.Location)
	case MsgTypeVideoChat:
		return msg, unmarshalMessageContent(msgType, content, msg.VideoChat)
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

// mask key
const (
	httpRequestHeaderAuthorization = "Authorization"
	httpRequestHeaderHelpdeskAuth  = "X-Lark-Helpdesk-Authorization"
)

func jsonHeader(headers map[string]string) string {
	val := make(map[string]string, len(headers))
	for k, v := range headers {
		if k == httpRequestHeaderAuthorization || k == httpRequestHeaderHelpdeskAuth {
			val[k] = maskString(v, 9, '*') // `Bearer xx******`
		} else {
			val[k] = v
		}
	}
	bs, _ := json.Marshal(val)
	return string(bs)
}

func maskString(s string, prefixCount int, mask rune) string {
	ss := []rune(s)
	res := make([]rune, len(ss))
	for i, v := range ss {
		if i < prefixCount {
			res[i] = v
		} else {
			res[i] = mask
		}
	}
	return string(res)
}
