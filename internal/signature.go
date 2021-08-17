package internal

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// 对于一般的事件回调，encryptKey 是加密秘钥： https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-security-verification
// 对于消息卡片回调，encryptKey是verify-token： https://open.feishu.cn/document/ukTMukTMukTM/uYzMxEjL2MTMx4iNzETM
func CalculateLarkCallbackSignature(timestamp, nonce, encryptKey string, body []byte) string {
	buf := new(bytes.Buffer)
	buf.WriteString(timestamp)
	buf.WriteString(nonce)
	buf.WriteString(encryptKey)
	buf.Write(body)

	h := sha256.New()
	h.Write(buf.Bytes())
	return fmt.Sprintf("%x", h.Sum(nil))
}
