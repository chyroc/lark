package card

import (
	"fmt"

	"github.com/chyroc/lark/card"
)

func Example2() {
	content := card.Card(
		card.ModuleImage("img_v2_0ccf88fb-09f3-42fb-8861-5d283001db5g"),
		card.Div().SetText(card.MarkdownText("佳节又中秋，月儿照当头。中秋月圆，佳节思亲🏡 \n我们为同学们准备了线下趣味小游戏，只要参与就有机会抽取惊喜中秋礼物～")),
		card.Div().SetText(card.MarkdownText("**🎄 圣诞派对时间：**12月24日 14:00-17:00\n\n**🎁 礼物交换方式：**请各位小伙伴们在包装好你准备的礼物之后，附上卡片祝福语，在 12 月 23 日下午 14:00 前交给前台，我们会统一交到圣诞老人手里～")),
		card.Div().SetFields(
			card.FieldMarkdown(""),
			card.FieldMarkdown("**📌 活动地点：**\n4F 餐厅"),
		),
		card.Action(
			card.LinkButton("查看活动规则", "https://open.feishu.cn/").SetPrimary(),
		),
	).
		SetHeader(card.Header("🎑 趣味中秋，活动预告🎉").SetYellow())
	fmt.Println(content.String())
}
