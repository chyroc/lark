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

func (r *EventCallbackService) listenCardCallback(ctx context.Context, isSecurity bool, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	// check security
	if err := r.checkCardSecurity(bs, isSecurity, header); err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	req, err := r.parserCardCallbackReq(ctx, header, bs)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	s, err := r.handlerCardCallbackReq(ctx, writer, req)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	_, _ = writer.Write([]byte(s))
	return
}

func (r *EventCallbackService) checkCardSecurity(body []byte, isSecurity bool, header http.Header) error {
	if !isSecurity {
		return nil
	}

	timestamp := header.Get("X-Lark-Request-Timestamp")
	nonce := header.Get("X-Lark-Request-Nonce")
	expectSignature := header.Get("X-Lark-Signature")
	realSignature := internal.CalculateLarkCallbackSignature(timestamp, nonce, r.cli.encryptKey, body)
	if expectSignature != realSignature {
		return fmt.Errorf("need check security, but security check invalid")
	}
	return nil
}

type cardCallbackReq struct {
	*EventCardCallback
	Challenge string `json:"challenge"` // 应用需要原样返回的值
	Type      string `json:"type"`      // url_verification
	Token     string `json:"token"`     // Token的使用可参考文档“通过Token验证事件来源”
	Encrypt   string `json:"encrypt"`
}

func (r *EventCallbackService) parserCardCallbackReq(ctx context.Context, header http.Header, body []byte) (*cardCallbackReq, error) {
	req := new(cardCallbackReq)

	if err := json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("lark card event unmarshal event_req %s failed", body)
	}

	if req.Encrypt != "" {
		if r.cli.encryptKey == "" {
			return nil, fmt.Errorf("receive encrypt string, but encrypt key not set")
		}

		decrypted, err := decryptEncryptString(r.cli.encryptKey, req.Encrypt)
		if err != nil {
			return nil, err
		}
		req, err = r.parserCardCallbackReq(ctx, header, []byte(decrypted))
		if err != nil {
			return nil, err
		}
	}
	if req.EventCardCallback == nil {
		req.EventCardCallback = new(EventCardCallback)
	}
	req.EventCardCallback.RefreshToken = header.Get("X-Refresh-Token")

	return req, nil
}

func (r *EventCallbackService) handlerCardCallbackReq(ctx context.Context, writer io.Writer, req *cardCallbackReq) (string, error) {
	if req.Token != "" {
		if r.cli.verificationToken == "" {
			return "", fmt.Errorf("must set verification token")
		}

		if req.Token != r.cli.verificationToken {
			return "", fmt.Errorf("verification token check failed")
		}
	}

	if req.Type == "url_verification" {
		return fmt.Sprintf(`{"challenge":%q}`, req.Challenge), nil
	}

	handled, s, err := r.handlerCardEvent(ctx, req.EventCardCallback)
	if handled {
		return s, err
	}

	return "{}", nil
}

func (r *EventCallbackService) handlerCardEvent(ctx context.Context, req *EventCardCallback) (handled bool, s string, err error) {
	if r.cli.eventHandler.eventCardHandler != nil {
		s, err := r.cli.eventHandler.eventCardHandler(ctx, r.cli, req)
		return true, s, err
	}
	return false, "", nil
}
