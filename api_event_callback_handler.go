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
		return req, nil
	case req.Schema == "2.0":
		if err := r.parserEventV2(req); err != nil {
			return req, err
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

	if req.Type == "url_verification" {
		return fmt.Sprintf(`{"challenge":%q}`, req.Challenge), nil
	}

	handled, s, err := r.handlerEventV2(ctx, req)
	if handled {
		return s, err
	}

	return "{}", nil
}
