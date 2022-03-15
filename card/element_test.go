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

func Test_Element(t *testing.T) {
	as := assert.New(t)

	// button
	as.Equal(`{"tag":"button","text":{"tag":"plain_text","content":"title"},"value":{"val":"1"}}`, jsonString(card.Button("title", map[string]interface{}{"val": "1"})))
	as.Equal(`{"tag":"button","text":{"tag":"plain_text","content":"title"},"url":"url"}`, jsonString(card.LinkButton("title", "url")))

	// time picker
	as.Equal(`{"tag":"date_picker","initial_date":"ini","value":{"val":"1"}}`, jsonString(card.DatePicker("ini", map[string]interface{}{"val": "1"})))
	as.Equal(`{"tag":"picker_time","initial_time":"ini","value":{"val":"1"}}`, jsonString(card.TimePicker("ini", map[string]interface{}{"val": "1"})))
	as.Equal(`{"tag":"picker_datetime","initial_datetime":"ini","value":{"val":"1"}}`, jsonString(card.DateTimePicker("ini", map[string]interface{}{"val": "1"})))

	// image
	as.Equal(`{"img_key":"img","tag":"img"}`, jsonString(card.ElementImage("img")))

	// overflow
	as.Equal(`{"options":[{"text":{"tag":"plain_text","content":"title"},"value":"value"}],"tag":"overflow"}`, jsonString(card.Overflow(
		card.SelectOption("title", "value"),
	)))

	// select menu: static
	as.Equal(`{"tag":"select_static","options":[{"text":{"tag":"plain_text","content":"title"},"value":"value"}]}`, jsonString(card.StaticSelectMenu(
		card.SelectOption("title", "value"),
	)))
	as.Equal(`{"tag":"select_static","placeholder":{"tag":"plain_text","content":"placeholder"},"options":[{"text":{"tag":"plain_text","content":"title"},"value":"value"}]}`, jsonString(card.StaticSelectMenu(
		card.SelectOption("title", "value"),
	).SetPlaceholder(card.Text("placeholder"))))

	// select menu: person
	as.Equal(`{"tag":"select_person","placeholder":{"tag":"plain_text","content":"placeholder"},"options":[{"value":"1"},{"value":"2"}],"confirm":{"title":{"tag":"plain_text","content":"t"},"text":{"tag":"plain_text","content":"v"}}}`, jsonString(card.PersonSelectMenuForIDs("1", "2").SetPlaceholder(card.Text("placeholder")).SetConfirm(card.Confirm("t", "v"))))
}
