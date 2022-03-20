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
	"testing"

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_post(t *testing.T) {
	as := assert.New(t)
	post := `{
    "title":"我是一个标题",
    "content":[
        [
            {
                "tag":"text",
                "text":"第一行 :"
            },
            {
                "tag":"a",
                "href":"http://www.feishu.cn",
                "text":"超链接"
            },
            {
                "tag":"at",
                "user_id":"@_user_1",
                "user_name":""
            }
        ],
        [
            {
                "tag":"img",
                "image_key":"img_47354fbc-a159-40ed-86ab-2ad0f1acb42g",
                "width":300,
                "height":300
            }
        ],
        [
            {
                "tag":"text",
                "text":"第二行:"
            },
            {
                "tag":"text",
                "text":"文本测试"
            }
        ],
        [
            {
                "tag":"img",
                "image_key":"img_47354fbc-a159-40ed-86ab-2ad0f1acb42g",
                "width":300,
                "height":300
            }
        ]
    ]
}`
	var postMessage lark.MessageContentPost
	err := json.Unmarshal([]byte(post), &postMessage)
	as.Nil(err)

	as.Equal("我是一个标题", postMessage.Title)

	as.Equal("第一行 :", postMessage.Content[0][0].(lark.MessageContentPostText).Text)

	as.Equal("超链接", postMessage.Content[0][1].(lark.MessageContentPostLink).Text)
	as.Equal("http://www.feishu.cn", postMessage.Content[0][1].(lark.MessageContentPostLink).Href)

	as.Equal("@_user_1", postMessage.Content[0][2].(lark.MessageContentPostAt).UserID)

	as.Equal("img_47354fbc-a159-40ed-86ab-2ad0f1acb42g", postMessage.Content[1][0].(lark.MessageContentPostImage).ImageKey)

	as.Equal("第二行:", postMessage.Content[2][0].(lark.MessageContentPostText).Text)

	as.Equal("文本测试", postMessage.Content[2][1].(lark.MessageContentPostText).Text)

	as.Equal("img_47354fbc-a159-40ed-86ab-2ad0f1acb42g", postMessage.Content[3][0].(lark.MessageContentPostImage).ImageKey)
}
