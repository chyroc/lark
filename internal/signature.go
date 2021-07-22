package internal

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

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
