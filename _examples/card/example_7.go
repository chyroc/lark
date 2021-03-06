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
package card

import (
	"fmt"

	"github.com/chyroc/lark/card"
)

func Example7() {
	content := card.Card(
		card.Div().SetFields(
			card.FieldMarkdown("**ð æéæ é¢ï¼**\næºå¨äººåæ¶æ¯å¡çä½¿ç¨è§èä¼åè¯å®¡"),
			card.FieldMarkdown("**â° æéæ¶é´ï¼**\\n2021-06-02 15:30ï¼GMT+08:00)"),
			card.FieldMarkdown(""),
			card.FieldMarkdown("**ð¤ è¢«æéäººï¼**\\n<at email=test@email.com></at>"),
		),
		card.HR(),

		card.Div().
			SetText(card.MarkdownText("[ä¿®æ¹æé](https://www.feishu.cn)")),
	).
		SetHeader(card.Header("æå¥å·²è®¾ç½®æé ð"))
	fmt.Println(content.String())
}
