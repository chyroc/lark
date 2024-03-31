package _examples

import "github.com/chyroc/lark/card/v2"

func Example2() *card.ComponentCard {
	return card.Card(
		card.Image("img_v3_0289_87ebfc6a-221c-4f56-986e-0ebcf098123g").
			SetScaleTypeCropCenter().
			SetSizeStretchWithoutPadding(),
		card.Markdown("**:FIREWORKS:热烈祝贺华南区再下一城！其它区域的小伙伴们也要加油哦~~~** ").
			SetTextAlignLeft().
			SetTextSizeNormal(),
		card.ColumnSet(
			buildExampleVsContent(
				"**订单金额**",
				"**<font color='red'>20,000</font>** ",
			),
			buildExampleVsContent(
				"**售前区域**",
				"**<font color='green'> 华南区</font>**",
			),
			buildExampleVsContent(
				"**赢单售前**",
				"**<at id=all></at>**",
			),
		).
			SetFlexModeNone().
			SetHorizontalAlignLeft().
			SetMargin(card.MarginTop(16)),
		card.Markdown("**:Trophy: 每一个客户的成功都离不开各团队的紧密配合！\n:LOVE: 衷心感谢在此项目中付出的各位同事和幕后英雄们！**").SetTextAlignLeft().SetTextSizeNormal(),
		card.Note(
			card.StandardIcon("lowerhand_filled"),
			card.Text("[机密] 请注意信息安全，严禁外传"),
		),
	).
		SetHeader(
			card.Header("恭喜飞书科技有限公司正式签约 🔥🔥🔥").
				SetTemplateRed().
				SetIcon(card.StandardIcon("premium-gleam_filled")),
		)
}

func buildExampleVsContent(title, content string) *card.ComponentColumn {
	return card.Column(
		card.ColumnSet(
			card.Column(
				card.Markdown(title).SetTextAlignCenter().SetTextSizeHeading(),
				card.Markdown(content).SetTextAlignCenter().SetTextSizeNormal(),
			).SetWidthWeighted().SetVerticalAlignTop().SetVerticalSpacingMedium().SetBackgroundStyleGrey().SetWeight(1),
		).SetFlexModeNone().SetHorizontalSpacingSmall().SetHorizontalAlignLeft(),
	).SetWidthWeighted().SetVerticalAlignTop().SetWeight(1)
}
