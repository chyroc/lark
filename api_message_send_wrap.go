package lark

import (
	"context"
	"fmt"
)

type MessageSendAPI struct {
	cli           *Lark
	msgAPI        *MessageAPI
	receiveID     string
	receiveIDType IDType
}

func (r *MessageAPI) Send() *MessageSendAPI {
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

func (r *MessageSendAPI) SendPost(ctx context.Context, card string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypePost, card)
}

func (r *MessageSendAPI) SendCard(ctx context.Context, card string) (*SendRawMessageResp, *Response, error) {
	return r.send(ctx, MsgTypeInteractive, card)
}

func (r *MessageSendAPI) SendShareChat(ctx context.Context, chatID string) (*SendRawMessageResp, *Response, error) {
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

func (r *MessageSendAPI) send(ctx context.Context, msgType MsgType, format string, args ...interface{}) (*SendRawMessageResp, *Response, error) {
	return r.msgAPI.SendRawMessage(ctx, &SendRawMessageReq{
		ReceiveIDType: &r.receiveIDType,
		ReceiveID:     &r.receiveID,
		Content:       fmt.Sprintf(format, args...),
		MsgType:       msgType,
	})
}
