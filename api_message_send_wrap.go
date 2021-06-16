package lark

import (
	"context"
	"encoding/json"
	"fmt"
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

func (r *MessageSendAPI) send(ctx context.Context, msgType MsgType, format string, args ...interface{}) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotWebHookURL != "" {
		return r.msgAPI.sendCustomBotMessage(ctx, &sendCustomBotMessageReq{
			MsgType: msgType,
			Content: fmt.Sprintf(format, args...),
		})
	}
	return r.msgAPI.SendRawMessage(ctx, &SendRawMessageReq{
		ReceiveIDType: r.receiveIDType,
		ReceiveID:     &r.receiveID,
		Content:       fmt.Sprintf(format, args...),
		MsgType:       msgType,
	})
}
