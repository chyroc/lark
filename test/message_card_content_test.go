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

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_CardContent(t *testing.T) {
	as := assert.New(t)

	action := lark.MessageContentCardModuleAction{
		Actions: []lark.MessageContentCardElement{
			&lark.MessageContentCardElementButton{
				Text: &lark.MessageContentCardObjectText{
					Tag:     "lark_md",
					Content: "click-me",
					Lines:   1,
				},
				URL:  "https://www.baidu.com",
				Type: "default",
			},
		},
		Layout: "bisected",
	}
	bs, err := action.MarshalJSON()
	as.Nil(err)
	s := string(bs)
	as.Contains(s, `{"actions":[{`)
	as.Contains(s, `"tag":"button"`)
	as.Contains(s, `"text":{`)
	as.Contains(s, `"tag":"lark_md"`)
	as.Contains(s, `"content":"click-me"`)
	as.Contains(s, `"lines":1`)
	as.Contains(s, `"type":"default"`)
	as.Contains(s, `"url":"https://www.baidu.com"`)
	as.Contains(s, `"tag":"action"`)
}
