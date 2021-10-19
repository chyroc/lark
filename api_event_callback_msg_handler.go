package lark

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/chyroc/lark/internal"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b

func (r *EventCallbackService) listenCallback(ctx context.Context, isSecurity bool, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	bs, err := ioutil.ReadAll(reader)
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

	// 只有类型不是 url_verification 且 设置了校验header 的情况下，才发起校验
	if req.Type != eventTypeURLVerification && isSecurity {
		if err := r.checkSecurity(bs, header); err != nil {
			writer.WriteHeader(500)
			_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
			return
		}
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

func (r *EventCallbackService) checkSecurity(body []byte, header http.Header) error {
	timestamp := header.Get("X-Lark-Request-Timestamp")
	nonce := header.Get("X-Lark-Request-Nonce")
	expectSignature := header.Get("X-Lark-Signature")
	realSignature := internal.CalculateLarkCallbackSignature(timestamp, nonce, r.cli.encryptKey, body)
	if expectSignature != realSignature {
		return fmt.Errorf("need check security, but security check invalid")
	}
	return nil
}

func (r *EventCallbackService) parserReq(ctx context.Context, body []byte) (*eventReq, error) {
	req := new(eventReq)
	req.eventBody = new(eventBody)

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

	switch {
	case req.Type == "url_verification":
		return req, nil
	case req.Schema == "2.0":
		if err := r.parserEventV2(req); err != nil {
			return req, err
		}
	case req.UUID != "":
		if err := r.parserEventV1(req); err != nil {
			return req, err
		}
	}

	return req, nil
}

func (r *EventCallbackService) handlerReq(ctx context.Context, writer io.Writer, req *eventReq) (string, error) {
	if r.cli.verificationToken == "" {
		return "", fmt.Errorf("must set verification token")
	}

	if req.getToken() != r.cli.verificationToken {
		return "", fmt.Errorf("verification token check failed")
	}

	if req.Type == eventTypeURLVerification {
		return fmt.Sprintf(`{"challenge":%q}`, req.Challenge), nil
	}

	handled, s, err := r.handlerEvent(ctx, req)
	if handled {
		return s, err
	}

	return "{}", nil
}

const eventTypeURLVerification = "url_verification"
