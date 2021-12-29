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
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// MessageSendAPI ...
type MessageSendAPI struct {
	cli           *Lark
	msgAPI        *MessageService
	receiveID     string
	receiveIDType IDType
}

// Send ...
func (r *MessageService) Send() *MessageSendAPI {
	return &MessageSendAPI{
		cli:    r.cli,
		msgAPI: r,
	}
}

// ToUserID ...
func (r *MessageSendAPI) ToUserID(id string) *MessageSendAPI {
	return r.to(id, IDTypeUserID)
}

// ToUnionID ...
func (r *MessageSendAPI) ToUnionID(id string) *MessageSendAPI {
	return r.to(id, IDTypeUnionID)
}

// ToOpenID ...
func (r *MessageSendAPI) ToOpenID(id string) *MessageSendAPI {
	return r.to(id, IDTypeOpenID)
}

// ToAppID ...
func (r *MessageSendAPI) ToAppID(id string) *MessageSendAPI {
	return r.to(id, IDTypeAppID)
}

// ToChatID ...
func (r *MessageSendAPI) ToChatID(id string) *MessageSendAPI {
	return r.to(id, IDTypeChatID)
}

// ToEmail ...
func (r *MessageSendAPI) ToEmail(id string) *MessageSendAPI {
	return r.to(id, IDTypeEmail)
}

func (r *MessageSendAPI) to(receiveID string, receiveIDType IDType) *MessageSendAPI {
	r.receiveID = receiveID
	r.receiveIDType = receiveIDType
	return r
}

// SendText ...
func (r *MessageSendAPI) SendText(ctx context.Context, text string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeText, `{"text":%q}`, text)
}

// SendImage ...
func (r *MessageSendAPI) SendImage(ctx context.Context, imageKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeImage, `{"image_key":%q}`, imageKey)
}

// SendPost ...
func (r *MessageSendAPI) SendPost(ctx context.Context, post string) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotWebHookURL != "" {
		vv := map[string]interface{}{}
		if err := json.Unmarshal([]byte(post), &vv); err != nil {
			return nil, nil, err
		}
		bs, err := json.Marshal(map[string]interface{}{
			"post": vv,
		})
		if err != nil {
			return nil, nil, err
		}
		return r.send(ctx, MsgTypePost, string(bs))
	}
	return r.send(ctx, MsgTypePost, post)
}

// SendCard ...
func (r *MessageSendAPI) SendCard(ctx context.Context, card string) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotWebHookURL != "" {
		vv := map[string]interface{}{}
		if err := json.Unmarshal([]byte(card), &vv); err != nil {
			return nil, nil, err
		}
		return r.msgAPI.sendCustomBotMessage(ctx, &sendCustomBotMessageReq{
			MsgType: MsgTypeInteractive,
			Card:    vv,
		})
	}
	return r.send(ctx, MsgTypeInteractive, card)
}

// SendShareChat ...
func (r *MessageSendAPI) SendShareChat(ctx context.Context, chatID string) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotWebHookURL != "" {
		return r.send(ctx, MsgTypeShareChat, `{"share_chat_id":%q}`, chatID)
	}
	return r.send(ctx, MsgTypeShareChat, `{"chat_id":%q}`, chatID)
}

// SendShareUser ...
func (r *MessageSendAPI) SendShareUser(ctx context.Context, userID string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeShareUser, `{"user_id":%q}`, userID)
}

// SendAudio ...
func (r *MessageSendAPI) SendAudio(ctx context.Context, fileKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeAudio, `{"file_key":%q}`, fileKey)
}

// SendMedia ...
func (r *MessageSendAPI) SendMedia(ctx context.Context, imageKey, fileKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeMedia, `{"image_key":%q,"file_key":%q}`, imageKey, fileKey)
}

// SendFile ...
func (r *MessageSendAPI) SendFile(ctx context.Context, fileKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeFile, `{"file_key":%q}`, fileKey)
}

// SendSticker ...
func (r *MessageSendAPI) SendSticker(ctx context.Context, fileKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeSticker, `{"file_key":%q}`, fileKey)
}

// MessageReplyAPI ...
type MessageReplyAPI struct {
	cli       *Lark
	msgAPI    *MessageService
	messageID string
}

// Reply ...
func (r *MessageService) Reply(messageID string) *MessageReplyAPI {
	return &MessageReplyAPI{
		cli:       r.cli,
		msgAPI:    r,
		messageID: messageID,
	}
}

// SendText ...
func (r *MessageReplyAPI) SendText(ctx context.Context, text string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeText, `{"text":%q}`, text)
}

// SendImage ...
func (r *MessageReplyAPI) SendImage(ctx context.Context, imageKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeImage, `{"image_key":%q}`, imageKey)
}

// SendPost ...
func (r *MessageReplyAPI) SendPost(ctx context.Context, card string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypePost, card)
}

// SendCard ...
func (r *MessageReplyAPI) SendCard(ctx context.Context, card string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeInteractive, card)
}

// SendShareChat ...
func (r *MessageReplyAPI) SendShareChat(ctx context.Context, chatID string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeShareChat, `{"chat_id":%q}`, chatID)
}

// SendShareUser ...
func (r *MessageReplyAPI) SendShareUser(ctx context.Context, userID string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeShareUser, `{"user_id":%q}`, userID)
}

// SendAudio ...
func (r *MessageReplyAPI) SendAudio(ctx context.Context, fileKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeAudio, `{"file_key":%q}`, fileKey)
}

// SendMedia ...
func (r *MessageReplyAPI) SendMedia(ctx context.Context, imageKey, fileKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeMedia, `{"image_key":%q,"file_key":%q}`, imageKey, fileKey)
}

// SendFile ...
func (r *MessageReplyAPI) SendFile(ctx context.Context, fileKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeFile, `{"file_key":%q}`, fileKey)
}

// SendSticker ...
func (r *MessageReplyAPI) SendSticker(ctx context.Context, fileKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeSticker, `{"file_key":%q}`, fileKey)
}

func (r *MessageSendAPI) send(ctx context.Context, msgType MsgType, format string, args ...interface{}) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotWebHookURL != "" {
		return r.msgAPI.sendCustomBotMessage(ctx, &sendCustomBotMessageReq{
			MsgType: msgType,
			Content: formatString(format, args...),
		})
	}
	return r.msgAPI.SendRawMessage(ctx, &SendRawMessageReq{
		ReceiveIDType: r.receiveIDType,
		ReceiveID:     r.receiveID,
		Content:       formatString(format, args...),
		MsgType:       msgType,
	})
}

func (r *MessageReplyAPI) reply(ctx context.Context, msgType MsgType, format string, args ...interface{}) (*ReplyRawMessageResp, *Response, error) {
	return r.msgAPI.ReplyRawMessage(ctx, &ReplyRawMessageReq{
		MessageID: r.messageID,
		Content:   formatString(format, args...),
		MsgType:   msgType,
	})
}

func (r *MessageService) sendCustomBotMessage(ctx context.Context, request *sendCustomBotMessageReq) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotSecret != "" {
		now := strconv.FormatInt(time.Now().Unix(), 10)
		secret, err := generateCustomBotMessageSign(r.cli.customBotSecret, now)
		if err != nil {
			return nil, nil, err
		}
		request.Sign = secret
		request.Timestamp = now
	}
	req := &RawRequestReq{
		Scope:        "Message",
		API:          "SendCustomBotMessage",
		Method:       "POST",
		URL:          r.cli.customBotWebHookURL,
		Body:         request,
		MethodOption: newMethodOption([]MethodOptionFunc{}),
	}
	resp := new(sendRawMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

type sendCustomBotMessageReq struct {
	Timestamp string                 `json:"timestamp,omitempty"`
	Sign      string                 `json:"sign,omitempty"`
	MsgType   MsgType                `json:"msg_type,omitempty"`
	Content   string                 `json:"content,omitempty"`
	Card      map[string]interface{} `json:"card,omitempty"`
}

func generateCustomBotMessageSign(secret string, timestamp string) (string, error) {
	stringToSign := timestamp + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func formatString(format string, args ...interface{}) string {
	if len(args) == 0 {
		return format
	}
	return fmt.Sprintf(format, args...)
}
