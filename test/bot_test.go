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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Bot(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		resp, _, err := AppAllPermission.Ins().Bot.GetBotInfo(ctx, &lark.GetBotInfoReq{})
		bs, _ := json.Marshal(resp)
		fmt.Println(string(bs))
		// printData(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Equal("lark-sdk", resp.AppName)
	})
}
