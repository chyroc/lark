package lark

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/chyroc/lark/internal"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b
func (r *EventCallbackService) ListenCallback(ctx context.Context, reader io.Reader, writer http.ResponseWriter) {
	r.listenCallback(ctx, false, http.Header{}, reader, writer)
}

// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b
func (r *EventCallbackService) ListenSecurityCallback(ctx context.Context, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	r.listenCallback(ctx, true, header, reader, writer)
}

// ListenCardCallback 卡片消息回调监听
//
// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b
func (r *EventCallbackService) ListenCardCallback(ctx context.Context, checkSecurity bool, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	r.listenCardCallback(ctx, checkSecurity, header, reader, writer)
}

// HandlerEventCard
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN
func (r *EventCallbackService) HandlerEventCard(f EventCardHandler) {
	r.cli.eventHandler.eventCardHandler = f
}

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

type EventCardHandler func(ctx context.Context, cli *Lark, event *EventCardCallback) (string, error)

type EventCardCallback struct {
	RefreshToken string `json:"refresh_token"` // header: X-Refresh-Token

	OpenID        string                   `json:"open_id"`
	UserID        string                   `json:"user_id"`
	OpenMessageID string                   `json:"open_message_id"`
	TenantKey     string                   `json:"tenant_key"`
	Token         string                   `json:"token"` // 更新卡片用的token(凭证)
	Action        *EventCardCallbackAction `json:"action"`
}

type EventCardCallbackAction struct {
	Value  json.RawMessage `json:"value"`  // 交互元素的value字段值
	Tag    string          `json:"tag"`    // 交互元素的tag字段值
	Option string          `json:"option"` // 选中option的value（button元素不适用）
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
	realSignature := internal.CalculateLarkCallbackSignature(timestamp, nonce, r.cli.verificationToken, body)
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

// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b

type eventReq struct {
	Encrypt string `json:"encrypt"` // 加密

	// v2 字段
	Schema string         `json:"schema"` // 2.0
	Header *EventHeaderV2 `json:"header"`

	// v1 字段
	UUID      string `json:"uuid"`      // 事件的唯一标识
	Token     string `json:"token"`     // 即Verification Token
	TS        string `json:"ts"`        // 事件发送的时间，一般近似于事件发生的时间。
	Type      string `json:"type"`      // url_verification, event_callback,
	Challenge string `json:"challenge"` // 配合 url_verification

	// 通用字段
	Event json.RawMessage `json:"event"`

	// v2 解析后字段
	*eventBody
}

func (r *eventReq) unmarshalEvent(e interface{}) error {
	err := json.Unmarshal(r.Event, e)
	if err != nil {
		return fmt.Errorf("lark event unmarshal event %s failed", err)
	}
	return nil
}

func (r *eventReq) getToken() string {
	if r.Header != nil && r.Header.Token != "" {
		return r.Header.Token
	}
	return r.Token
}

func (r *eventReq) headerV1(eventType EventType) *EventHeaderV1 {
	return &EventHeaderV1{
		UUID:      r.UUID,
		EventType: eventType,
		TS:        r.TS,
		Token:     r.Token,
	}
}

func decryptEncryptString(encryptKey string, cryptoText string) (string, error) {
	var key []byte
	{
		h := sha256.New()
		h.Write([]byte(encryptKey))
		key = h.Sum(nil)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCBCDecrypter(block, iv)
	stream.CryptBlocks(ciphertext, ciphertext)

	return string(ciphertext[:len(ciphertext)-int(ciphertext[len(ciphertext)-1])]), nil
}
