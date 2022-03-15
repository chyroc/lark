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
package card_test

import (
	"testing"

	"github.com/chyroc/lark/card"
	"github.com/stretchr/testify/assert"
)

func Test_Object(t *testing.T) {
	as := assert.New(t)

	// confirm
	as.Equal(`{"title":{"tag":"plain_text","content":"title"},"text":{"tag":"plain_text","content":"txt"}}`, jsonString(card.Confirm("title", "txt")))
	as.Equal(`{"title":{"tag":"plain_text","content":"hi"},"text":{"tag":"plain_text","content":"txt"}}`, jsonString(card.Confirm("title", "txt").SetTitle(card.Text("hi"))))
	as.Equal(`{"title":{"tag":"plain_text","content":"title"},"text":{"tag":"lark_md","content":"hi"}}`, jsonString(card.Confirm("title", "txt").SetText(card.MarkdownText("hi"))))

	// field
	as.Equal(`{"text":{"tag":"plain_text","content":"field"}}`, jsonString(card.FieldText("field")))

	// option
	as.Equal(`{"text":{"tag":"plain_text","content":"txt"},"value":"val"}`, jsonString(card.SelectOption("txt", "val")))
	as.Equal(`{"value":"id"}`, jsonString(card.PersonOption("id")))
	as.Equal(`{"text":{"tag":"plain_text","content":"txt"},"url":"url"}`, jsonString(card.LinkOption("txt", "url")))

	// text
	as.Equal(`{"tag":"plain_text","content":"txt"}`, jsonString(card.Text("txt")))
	as.Equal(`{"tag":"lark_md","content":"md"}`, jsonString(card.MarkdownText("md")))
	as.Equal(`{"tag":"lark_md","content":"val"}`, jsonString(card.Text("txt").SetMd("val")))
	as.Equal(`{"tag":"plain_text","content":"val"}`, jsonString(card.MarkdownText("md").SetText("val")))

	// url
	as.Equal(`{"url":"url"}`, jsonString(card.URL("url")))
	as.Equal(`{"url":"url","ios_url":"val"}`, jsonString(card.URL("url").SetIOSURL("val")))
	as.Equal(`{"url":"url","android_url":"val"}`, jsonString(card.URL("url").SetAndroidURL("val")))
	as.Equal(`{"url":"url","pc_url":"val"}`, jsonString(card.URL("url").SetPCURL("val")))
}
