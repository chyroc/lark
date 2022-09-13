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

func Example9() {
	content := card.Card(
		card.Markdown("**个人审批效率总览**"),
		card.ColumnSet(
			card.Column(
				card.Markdown("已审批单量\n**29单**\n"+card.GreenText("领先团队59%")).SetCenterTextAlign(),
			).SetWeight(1),
			card.Column(
				card.Markdown("平均审批耗时\n**0.9小时**\n"+card.GreenText("领先团队100%")).SetCenterTextAlign(),
			).SetWeight(1),
			card.Column(
				card.Markdown("待批率\n**25%**\n"+card.RedText("落后团队29%")).SetCenterTextAlign(),
			).SetWeight(1),
		).SetFlexMode("bisect").SetBackgroundStyle("grey").SetHorizontalSpacing("default"),
		card.HR(),
		card.Markdown("**团队审批效率总览**"),
		card.ColumnSet(
			card.Column(
				card.Markdown("**审批人**\n王大明\n张军\n李小方"),
			).SetWeight(1),
			card.Column(
				card.Markdown("**审批时长**\n小于1小时\n2小时\n3小时"),
			).SetWeight(1),
			card.Column(
				card.Markdown("**对比上周变化**\n"+card.GreenText("↓12%")+"\n"+card.RedText("↑5%")+"\n"+card.GreenText("↓25%")),
			).SetWeight(1),
		).SetFlexMode("none").SetBackgroundStyle("default").SetHorizontalSpacing("default"),
	).
		SetHeader(
			card.Header("Title").SetGreen(),
		)
	fmt.Println(content.String())
}
