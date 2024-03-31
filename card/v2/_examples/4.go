package _examples

import "github.com/chyroc/lark/card/v2"

func Example4() *card.ComponentCard {
	return card.Card(
		card.ColumnSet(
			card.Column(
				buildExample4Post(
					"img_v2_136e4af5-9893-4f96-a52a-60bb9b6347cg",
					"2023 年 11 月 11 日 10:00",
				),
				buildExample4Post(
					"img_v2_f9d63232-61ec-44f0-9fb7-bf8ba82ea92g",
					"北京学清嘉创大厦 B 座 F2-26",
				),
			).SetWidthWeighted().SetVerticalAlignTop().SetWeight(1),
			card.Column(
				card.Button("🔥 立即报名").SetTypePrimary(),
			).SetWidthAuto().SetVerticalAlignCenter(),
		).SetFlexModeNone().SetHorizontalAlignLeft(),
		card.Image("img_v2_609930d7-21cc-475a-baba-3de5dafe079g").SetScaleTypeFitHorizontal(),
	).
		SetHeader(
			card.Header("十周年庆典").SetIcon(card.CustomIcon("img_v2_1bfeb6af-dd10-4e49-8fd2-5a392bc86ccg")),
		)
}

func buildExample4Post(icon, content string) *card.ComponentColumnSet {
	return card.ColumnSet(
		card.Column(
			card.Image(icon).SetScaleTypeCropCenter().SetSizeTiny(),
		).SetWidthAuto().SetVerticalAlignCenter(),
		card.Column(
			card.ColumnSet(
				card.Column(
					card.PlainText(content),
				).SetWidthWeighted().SetWeight(1),
			).SetFlexModeNone(),
		).SetWidthWeighted().SetVerticalAlignTop(),
	).SetFlexModeNone().SetHorizontalAlignLeft()
}
