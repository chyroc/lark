package lark

import (
	"context"
	"fmt"
)

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

func (r *MessageReplyAPI) reply(ctx context.Context, msgType MsgType, format string, args ...interface{}) (*ReplyRawMessageResp, *Response, error) {
	return r.msgAPI.ReplyRawMessage(ctx, &ReplyRawMessageReq{
		MessageID: r.messageID,
		Content:   fmt.Sprintf(format, args...),
		MsgType:   msgType,
	})
}
