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
