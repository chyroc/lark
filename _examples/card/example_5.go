package card

import (
	"fmt"

	"github.com/chyroc/lark/card"
)

func Example5() {
	content := card.Card(
		card.ModuleImage("img_v2_cb03ec35-a638-4b93-9e6f-5e2d0e549deg"),
		card.Div().
			SetText(card.MarkdownText("关于服务台**动态答案**及**多轮对话**功能的小讨论\\n[点击查看 >>](https://open.feishu.cn)")),
		card.HR(),

		card.Div().
			SetText(card.MarkdownText("表格的**权限能否更精细化**？——快来分享你的使用场景\\n[点击查看 >>](https://open.feishu.cn)")).
			SetExtra(card.ElementImage("img_v2_c3903791-3f53-40b0-8ecd-b457288d2d6g")),
		card.HR(),

		card.Div().
			SetText(card.MarkdownText("**企业百科**怎么玩？大家的经验分享，你要的宝藏贴都在这里！\\n[点击查看 >>](https://open.feishu.cn)")).
			SetExtra(card.ElementImage("img_v2_6e73ff00-a379-4d33-9e9f-1bea36190d3g")),
		card.HR(),

		card.Note(
			card.MarkdownText("💡本栏目每天早上为你带来回顾，你想在这里看到什么样的帖子？给我们留言吧 😘 "),
		),
	).
		SetHeader(card.Header("📊 如果表格权限更细化，你希望怎么用它?").SetGreen())
	fmt.Println(content.String())
}
