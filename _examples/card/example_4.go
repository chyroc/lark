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

func Example4() {
	content := card.Card(
		card.ModuleImage("img_200f82a4-918e-46a3-8862-fdfb006ff95g"),
		card.Div().
			SetText(card.MarkdownText("**#投票#文档到底要不要支持字体设置？**\\n\\n[“不要，要消灭工作环境中的‘被迫格式’指令。”](https://open.feishu.cn)")).
			SetExtra(card.ElementImage("img_200f82a4-918e-46a3-8862-fdfb006ff95g")),
		card.HR(),

		card.Div().
			SetText(card.MarkdownText("**飞书'便签'这么好用，存在感却这么低?**\\n\\n[“飞书可以推出一款独立的便签应用。”](https://open.feishu.cn)")).
			SetExtra(
				card.ElementImage("img_2fe6e807-addf-4753-8e35-2052dcd9118g"),
			),
		card.HR(),

		card.Div().
			SetText(card.MarkdownText("**飞书应该如何应对竞争？**\\n\\n[“飞书应该是一个生态，一个体系，一个航母作战群作战。”](https://open.feishu.cn)")).
			SetExtra(
				card.ElementImage("img_2f2930e2-6e71-48a1-a484-6be2709b4beg"),
			),
	).
		SetHeader(card.Header("🎉 #一周精选#「创刊号」来了，看看你的发帖上榜没？").SetBlue())
	fmt.Println(content.String())
}
