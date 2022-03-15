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
