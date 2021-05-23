package lark

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

func (r *MessageService) sendCustomBotMessage(ctx context.Context, request *sendCustomBotMessageReq) (*SendRawMessageResp, *Response, error) {
	if r.cli.customBotSecret != "" {
		now := strconv.FormatInt(time.Now().Unix(), 10)
		secret, err := genSign(r.cli.customBotSecret, now)
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

func genSign(secret string, timestamp string) (string, error) {
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
