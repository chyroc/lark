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

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/card"
	"github.com/stretchr/testify/assert"
)

func Test_Module(t *testing.T) {
	as := assert.New(t)

	// action
	as.Equal(`{"actions":[{"action_type":"request","tag":"button","text":{"tag":"plain_text","content":"button"},"type":"default"}],"tag":"action"}`, jsonString(card.Action(card.Button("button", nil))))

	// div
	as.Equal(`{"tag":"div"}`, jsonString(card.Div()))
	as.Equal(`{"tag":"div","text":{"tag":"plain_text","content":"hi"}}`, jsonString(card.Div().SetText(card.Text("hi"))))
	as.Equal(`{"fields":[{"is_short":false,"text":{"tag":"plain_text","content":"field1"}}],"tag":"div"}`, jsonString(card.Div().SetFields(card.FieldText("field1"))))
	as.Equal(`{"fields":[{"is_short":false,"text":{"tag":"plain_text","content":"field1"}},{"is_short":false,"text":{"tag":"plain_text","content":"field2"}}],"tag":"div"}`, jsonString(card.Div().SetFields(card.FieldText("field1"), card.FieldText("field2"))))
	as.Equal(`{"extra":{"tag":"plain_text","content":"hi"},"tag":"div"}`, jsonString(card.Div().SetExtra(card.Text("hi"))))

	// hr
	as.Equal(`{"tag":"hr"}`, jsonString(card.HR()))

	// image
	as.Equal(`{"alt":{"tag":"plain_text","content":""},"img_key":"img","tag":"img"}`, jsonString(card.ModuleImage("img")))

	// markdown
	as.Equal(`{"content":"- 1","tag":"markdown"}`, jsonString(card.Markdown("- 1")))
	as.Equal(`{"content":"- 1","href":{"a":{"url":"hi"}},"tag":"markdown"}`, jsonString(card.Markdown("- 1").SetHref(map[string]*lark.MessageContentCardObjectURL{
		"a": card.URL("hi"),
	})))

	// note
	as.Equal(`{"elements":[{"tag":"plain_text","content":"hi"}],"tag":"note"}`, jsonString(card.Note(card.Text("hi"))))
	as.Equal(`{"elements":[{"img_key":"image","tag":"img"}],"tag":"note"}`, jsonString(card.Note(card.ElementImage("image"))))
}
