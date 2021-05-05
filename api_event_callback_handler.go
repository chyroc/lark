package lark

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b

func (r *EventCallbackAPI) listenCallback(ctx context.Context, reader io.Reader, writer http.ResponseWriter) {
	bs, err := io.ReadAll(reader)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	req, err := r.parserReq(ctx, bs)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	s, err := r.handlerReq(ctx, writer, req)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	_, _ = writer.Write([]byte(s))
	return
}

func (r *EventCallbackAPI) parserReq(ctx context.Context, body []byte) (*eventReq, error) {
	fmt.Println("bs", string(body))
	req := new(eventReq)

	if err := json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("lark event unmarshal event_req %s failed", body)
	}

	if req.Encrypt != "" {
		if r.cli.encryptKey == "" {
			return nil, fmt.Errorf("receive encrypt string, but encrypt key not set")
		}

		decrypted, err := decryptEncryptString(r.cli.encryptKey, req.Encrypt)
		if err != nil {
			return nil, err
		}
		req, err = r.parserReq(ctx, []byte(decrypted))
		if err != nil {
			return nil, err
		}
	}

	spew.Dump(req)

	switch {
	case req.Type == "url_verification":
		// _, err = writer.Write([]byte(fmt.Sprintf(`{"challenge":%q}`, req.Challenge)))
		// return err
		return req, nil
	case req.Schema == "2.0":
		if req.Header == nil {
			return req, fmt.Errorf("get schema=2.0, but header is nil")
		}
		switch req.Header.EventType {
		case EventTypeMessageReceive:
			event := new(EventMessageReceive)
			if err := req.unmarshalEvent(event); err != nil {
				return req, err
			}
			req.eventMessageReceive = event
			return req, nil
		}
	}

	return req, nil
}

func (r *EventCallbackAPI) handlerReq(ctx context.Context, writer io.Writer, req *eventReq) (string, error) {
	if r.cli.verificationToken == "" {
		return "", fmt.Errorf("must set verification token")
	}

	if req.getToken() != r.cli.verificationToken {
		return "", fmt.Errorf("verification token check failed")
	}

	switch {
	case req.Type == "url_verification":
		return fmt.Sprintf(`{"challenge":%q}`, req.Challenge), nil
	case req.eventMessageReceive != nil:
		if r.cli.eventHandler.eventTypeImageReceiveHandler != nil {
			return r.cli.eventHandler.eventTypeImageReceiveHandler(ctx, r.cli, req.Schema, req.Header, req.eventMessageReceive)
		}
	}
	return "{}", nil
}
