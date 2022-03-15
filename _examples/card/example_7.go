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
			card.FieldMarkdown("**🗒 提醒标题：**\n机器人和消息卡片使用规范优化评审"),
			card.FieldMarkdown("**⏰ 提醒时间：**\\n2021-06-02 15:30（GMT+08:00)"),
			card.FieldMarkdown(""),
			card.FieldMarkdown("**👤 被提醒人：**\\n<at email=test@email.com></at>"),
		),
		card.HR(),

		card.Div().
			SetText(card.MarkdownText("[修改提醒](https://www.feishu.cn)")),
	).
		SetHeader(card.Header("李健已设置提醒 🎉"))
	fmt.Println(content.String())
}
