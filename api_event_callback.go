/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
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

// ListenCallback ...
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM
func (r *EventCallbackService) ListenCallback(ctx context.Context, reader io.Reader, writer http.ResponseWriter) {
	r.listenAllCallback(ctx, false, http.Header{}, reader, writer)
}

// ListenSecurityCallback ...
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM
func (r *EventCallbackService) ListenSecurityCallback(ctx context.Context, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	r.listenAllCallback(ctx, true, header, reader, writer)
}

// ListenCardCallback 卡片消息回调监听
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM
//
// Deprecated: please use ListenCallback instead
func (r *EventCallbackService) ListenCardCallback(ctx context.Context, checkSecurity bool, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	r.listenAllCallback(ctx, checkSecurity, header, reader, writer)
}

// EventCardCallback ...
type EventCardCallback struct {
	RefreshToken string `json:"refresh_token"` // header: X-Refresh-Token

	OpenID        string                   `json:"open_id"`
	UserID        string                   `json:"user_id"`
	OpenMessageID string                   `json:"open_message_id"`
	TenantKey     string                   `json:"tenant_key"` // 租户标识
	Token         string                   `json:"token"`      // 刷新凭证，业务方凭此在30分钟内最多可刷新两次消息卡片
	Action        *EventCardCallbackAction `json:"action"`     // 交互信息
}

// EventCardCallbackAction ...
type EventCardCallbackAction struct {
	Value  json.RawMessage `json:"value"`  // 交互元素的value字段值
	Tag    string          `json:"tag"`    // 交互元素的tag字段值
	Option string          `json:"option"` // 选中option的value（button元素不适用）
}

// EventCardHandler ...
type EventCardHandler func(ctx context.Context, cli *Lark, event *EventCardCallback) (string, error)

// HandlerEventCard
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN
func (r *EventCallbackService) HandlerEventCard(f EventCardHandler) {
	r.cli.eventHandler.eventCardHandler = f
}

const eventTypeURLVerification = "url_verification"

// 统一卡片和消息的回调
//
// - 读取 body
// - 解析数据
// - 安全校验
// - 处理逻辑
func (r *EventCallbackService) listenAllCallback(ctx context.Context, isSecurity bool, header http.Header, reader io.Reader, writer http.ResponseWriter) {
	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	req, err := r.parserAllCallbackRequest(ctx, header, bs)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	// 只有类型不是 url_verification 且 设置了校验的情况下，才发起校验
	if err := r.checkCallbackSecurity(isSecurity, header, bs, req); err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	s, err := r.handlerAllEvent(ctx, writer, req)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"err":%q}`, err)))
		return
	}

	_, _ = writer.Write([]byte(s))
	return
}

func (r *EventCallbackService) parserAllCallbackRequest(ctx context.Context, header http.Header, body []byte) (*eventReq, error) {
	req := &eventReq{
		eventBody: &eventBody{},
		EventCardCallback: &EventCardCallback{
			RefreshToken: header.Get("X-Refresh-Token"),
		},
	}
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("lark event unmarshal req %s failed", body)
	}

	if req.Encrypt != "" {
		if r.cli.encryptKey == "" {
			return nil, fmt.Errorf("receive encrypt string, but encrypt key not set")
		}

		decrypted, err := decryptEncryptString(r.cli.encryptKey, req.Encrypt)
		if err != nil {
			return nil, err
		}

		req, err = r.parserAllCallbackRequest(ctx, header, []byte(decrypted))
		if err != nil {
			return nil, err
		}
	}

	switch {
	case req.Type == eventTypeURLVerification:
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

// 通用事件：https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-security-verification
// 卡片事件：https://open.feishu.cn/document/ukTMukTMukTM/uYzMxEjL2MTMx4iNzETM
func (r *EventCallbackService) checkCallbackSecurity(isSecurity bool, header http.Header, body []byte, req *eventReq) error {
	if req.Type == eventTypeURLVerification || !isSecurity {
		return nil
	}

	if req.isNormalEvent() {
		if r.cli.verificationToken == "" {
			return fmt.Errorf("must set verification token")
		}

		if req.getToken() != r.cli.verificationToken {
			return fmt.Errorf("verification token check failed")
		}
	}

	encryptKey := r.cli.verificationToken
	if req.isNormalEvent() {
		encryptKey = r.cli.encryptKey
	}

	timestamp := header.Get("X-Lark-Request-Timestamp")
	nonce := header.Get("X-Lark-Request-Nonce")
	expectSignature := header.Get("X-Lark-Signature")
	realSignature := internal.CalculateLarkCallbackSignature(timestamp, nonce, encryptKey, body)
	if expectSignature != realSignature {
		return fmt.Errorf("need check security, but security check invalid")
	}
	return nil
}

func (r *EventCallbackService) handlerAllEvent(ctx context.Context, writer io.Writer, req *eventReq) (string, error) {
	if req.Type == eventTypeURLVerification {
		return fmt.Sprintf(`{"challenge":%q}`, req.Challenge), nil
	}

	if req.isNormalEvent() {
		handled, s, err := r.handlerEvent(ctx, req)
		if handled {
			return s, err
		}
	} else {
		handled, s, err := r.handlerCardEvent(ctx, req.EventCardCallback)
		if handled {
			return s, err
		}
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

	// v1 v2 通用字段
	Event json.RawMessage `json:"event"`

	// card event
	*EventCardCallback

	// v2 event 解析后字段
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

func (r *eventReq) isNormalEvent() bool {
	return r.Schema == "2.0" || r.UUID != ""
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
