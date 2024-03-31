package _examples

import "github.com/chyroc/lark/card/v2"

func Example3() *card.ComponentCard {
	return card.Card(
		card.Image("img_v3_0289_b74726a5-da82-47be-9e2a-52aa9b5a77fg").SetScaleTypeCropCenter().SetSizeStretchWithoutPadding(),
		card.ColumnSet(
			buildExample3Post(
				":WAIL: <font color='red'>企业信息传递痛点</font>",
				"- 各类消息堆积，信息爆炸难聚焦\n- 消息列表刷新，重要信息被淹没\n- 多线处理工作，关键待办找不到",
				"img_v3_0289_6d4bafc3-0a31-4058-804f-bec67584389g",
				"接入前，重要信息易忽略、难触达",
			),
			buildExample3Post(
				":MeMeMe: <font color='green'>信息流开放来帮你</font>",
				"- 自定义添加**按钮**、**标签**，信息更醒目\n- 设置持续**置顶展示**状态，处理不遗漏\n- 消息列表点击**关键按钮**，操作更简便",
				"img_v3_0289_bb92dbb9-fb35-4742-a330-efe07a3cea4g",
				"接入后，重要信息置顶展示、一眼就看到",
			),
		).SetFlexModeStretch().SetHorizontalAlignLeft().SetMargin(card.MarginTop(16)),
		card.Divider(),
		card.Action(
			card.Button("🔮 立即申请开通").
				SetTypePrimaryFilled().
				SetSizeMedium().
				SetMultiURL(card.MultiURL("https://bytedance.larkoffice.com/share/base/form/shrcnoTc4nArIkA0SQzkJS4Zr0e")),
			card.Button("📖 查看接入指南").
				SetSizeMedium().
				SetTypePrimary().
				SetMultiURL(card.MultiURL("https://bytedance.larkoffice.com/wiki/Cqh2wQUgEiY3VmkDkLDc30RqnIh")),
		),
	).SetHeader(
		card.Header("重要信息找不到、易忽略？让「飞书信息流开放」来帮你！").
			SetTextTagList(card.TextTag("🔥火热内测中").SetColorRed()).
			SetTemplateBlue().
			SetIcon(card.StandardIcon("lark-logo_colorful")),
	)
}

func buildExample3Post(title, content, image, alt string) *card.ComponentColumn {
	return card.Column(
		card.Markdown(title).SetTextAlignCenter().SetTextSizeNormal(),
		card.Markdown(content).SetTextAlignLeft().SetTextSizeNormal(),
		card.Image(image).SetScaleTypeFitHorizontal().SetAlt(card.Text(alt)),
	).SetWidthWeighted().SetVerticalAlignTop().SetWeight(1)
}
