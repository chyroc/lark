package _examples

import (
	"github.com/chyroc/lark/card/v2"
)

func Example6() *card.ComponentCard {
	return card.Card(
		card.Markdown("2024年2月8日/9日前往拉萨，招募同行伙伴，目前 3️⃣ = 1️⃣").SetTextAlignLeft().SetTextSizeNormal(),
		card.MultiImage(
			"img_v3_027d_feaf588d-b31b-43b9-bd61-c5595406d40g",
			"img_v3_027d_8bd00882-7eb6-48df-8999-a79cd90d88ag",
			"img_v3_027d_1abba963-fd3e-4666-ade2-198c8d0813bg",
			"img_v3_027d_e361c8eb-28d4-406f-a44d-cfa9d858f44g",
			"img_v3_027d_e5dcd820-bfa7-42e7-9009-82aea1133b3g",
			"img_v3_027d_867abe97-d32a-4955-854a-6550cdfeb1bg",
			"img_v3_027d_d644f50c-510a-4fcb-98f8-ed70632b3d8g",
			"img_v3_027d_b932d7e7-be81-4be7-8222-2663d83dc58g",
			"img_v3_027d_b6c77830-6fad-41bf-ac3b-4906e7d83f0g",
		).SetCombinationModeTrisect().SetCornerRadiusPercent(5),
		card.Note(
			card.StandardIcon("thumbsup_outlined"),
			card.Text("11"),
		),
	).SetHeader(
		card.Header("李天天发布了 1 条公司圈动态").
			SetTemplateBlue().
			SetIconStandard("larkcommunity_colorful"),
	)
}
