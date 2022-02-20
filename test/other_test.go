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
package test

import (
	"testing"

	"github.com/chyroc/lark/internal"
	"github.com/stretchr/testify/assert"
)

func Test_LarkCallbackSignature(t *testing.T) {
	as := assert.New(t)

	timestamp := "1626947018"
	nonce := "64607134"
	encryptKey := "fake-xx"
	body := []byte(`{fake}`)
	signature := "7db78fdbc2fb617c46e713c3ec3f4b4658d94166682a80434fe557f23afcfd06"
	res := internal.CalculateLarkCallbackSignature(timestamp, nonce, encryptKey, body)
	as.Equal(signature, res)
}
