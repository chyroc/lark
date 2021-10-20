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

type MessageSendAPI struct {
	cli           *Lark
	msgAPI        *MessageService
	receiveID     string
	receiveIDType IDType
}

func (r *MessageService) Send() *MessageSendAPI {
	return &MessageSendAPI{
		cli:    r.cli,
		msgAPI: r,
	}
}

func (r *MessageSendAPI) ToUserID(id string) *MessageSendAPI {
	return r.to(id, IDTypeUserID)
}

func (r *MessageSendAPI) ToUnionID(id string) *MessageSendAPI {
	return r.to(id, IDTypeUnionID)
}

func (r *MessageSendAPI) ToOpenID(id string) *MessageSendAPI {
	return r.to(id, IDTypeOpenID)
}

func (r *MessageSendAPI) ToAppID(id string) *MessageSendAPI {
	return r.to(id, IDTypeAppID)
}

func (r *MessageSendAPI) ToChatID(id string) *MessageSendAPI {
	return r.to(id, IDTypeChatID)
}

func (r *MessageSendAPI) ToEmail(id string) *MessageSendAPI {
	return r.to(id, IDTypeEmail)
}

func (r *MessageSendAPI) to(receiveID string, receiveIDType IDType) *MessageSendAPI {
	r.receiveID = receiveID
	r.receiveIDType = receiveIDType
	return r
}

func (r *MessageSendAPI) SendText(ctx context.Context, text string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeText, `{"text":%q}`, text)
}

func (r *MessageSendAPI) SendImage(ctx context.Context, imageKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeImage, `{"image_key":%q}`, imageKey)
}

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

func (r *MessageSendAPI) SendShareChat(ctx context.Context, chatID string) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotWebHookURL != "" {
		return r.send(ctx, MsgTypeShareChat, `{"share_chat_id":%q}`, chatID)
	}
	return r.send(ctx, MsgTypeShareChat, `{"chat_id":%q}`, chatID)
}

func (r *MessageSendAPI) SendShareUser(ctx context.Context, userID string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeShareUser, `{"user_id":%q}`, userID)
}

func (r *MessageSendAPI) SendAudio(ctx context.Context, fileKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeAudio, `{"file_key":%q}`, fileKey)
}

func (r *MessageSendAPI) SendMedia(ctx context.Context, imageKey, fileKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeMedia, `{"image_key":%q,"file_key":%q}`, imageKey, fileKey)
}

func (r *MessageSendAPI) SendFile(ctx context.Context, fileKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeFile, `{"file_key":%q}`, fileKey)
}

func (r *MessageSendAPI) SendSticker(ctx context.Context, fileKey string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeSticker, `{"file_key":%q}`, fileKey)
}

type MessageReplyAPI struct {
	cli       *Lark
	msgAPI    *MessageService
	messageID string
}

func (r *MessageService) Reply(messageID string) *MessageReplyAPI {
	return &MessageReplyAPI{
		cli:       r.cli,
		msgAPI:    r,
		messageID: messageID,
	}
}

func (r *MessageReplyAPI) SendText(ctx context.Context, text string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeText, `{"text":%q}`, text)
}

func (r *MessageReplyAPI) SendImage(ctx context.Context, imageKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeImage, `{"image_key":%q}`, imageKey)
}

func (r *MessageReplyAPI) SendPost(ctx context.Context, card string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypePost, card)
}

func (r *MessageReplyAPI) SendCard(ctx context.Context, card string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeInteractive, card)
}

func (r *MessageReplyAPI) SendShareChat(ctx context.Context, chatID string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeShareChat, `{"chat_id":%q}`, chatID)
}

func (r *MessageReplyAPI) SendShareUser(ctx context.Context, userID string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeShareUser, `{"user_id":%q}`, userID)
}

func (r *MessageReplyAPI) SendAudio(ctx context.Context, fileKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeAudio, `{"file_key":%q}`, fileKey)
}

func (r *MessageReplyAPI) SendMedia(ctx context.Context, imageKey, fileKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeMedia, `{"image_key":%q,"file_key":%q}`, imageKey, fileKey)
}

func (r *MessageReplyAPI) SendFile(ctx context.Context, fileKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeFile, `{"file_key":%q}`, fileKey)
}

func (r *MessageReplyAPI) SendSticker(ctx context.Context, fileKey string) (*ReplyRawMessageResp, *Response, error) {
	return r.reply(ctx, MsgTypeSticker, `{"file_key":%q}`, fileKey)
}

func (r *MessageSendAPI) send(ctx context.Context, msgType MsgType, format string, args ...interface{}) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotWebHookURL != "" {
		return r.msgAPI.sendCustomBotMessage(ctx, &sendCustomBotMessageReq{
			MsgType: msgType,
			Content: fmt.Sprintf(format, args...),
		})
	}
	return r.msgAPI.SendRawMessage(ctx, &SendRawMessageReq{
		ReceiveIDType: r.receiveIDType,
		ReceiveID:     r.receiveID,
		Content:       fmt.Sprintf(format, args...),
		MsgType:       msgType,
	})
}

func (r *MessageReplyAPI) reply(ctx context.Context, msgType MsgType, format string, args ...interface{}) (*ReplyRawMessageResp, *Response, error) {
	return r.msgAPI.ReplyRawMessage(ctx, &ReplyRawMessageReq{
		MessageID: r.messageID,
		Content:   fmt.Sprintf(format, args...),
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
