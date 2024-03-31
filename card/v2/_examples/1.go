package _examples

import (
	"github.com/chyroc/lark/card/v2"
)

func Example1() *card.ComponentCard {
	return card.Card(
		card.Image("img_v3_027m_c5c3ec75-cd3c-4092-8faf-9f6ae3697d4g").
			SetPreview(true).
			SetScaleTypeFitHorizontal().
			SetCornerRadiusPixel(10),
		example1BuildPost(
			"img_v3_027s_851c5059-00b5-4cf7-adeb-95ddc7e6d95g",
			"**影视飓风**",
			"过去一年你参与了哪些项目、最长出差记录是多久、最远出差目的地是哪里......带你走进百大 up 主的年会抽奖 & 年度报告，看他们是如何玩转飞书开放能力，办了一场炫酷又暖心的公司年会！",
			"img_v3_0289_cead13d0-bfd7-4fbf-8f8d-54dbb341e04g",
			"https://mp.weixin.qq.com/s/X98sAGRAVprBt8-IxZrsSg",
		),
		example1BuildPost(
			"img_v3_0289_979245ad-2480-4fc5-ada7-42238b73f69g",
			"**亿咖通科技**",
			"借助飞书集成平台、飞书应用引擎等工具，结合飞书开放能力，对外加强沟通交流，对内推动高效协作，实现企业数智化的“降本增效”！",
			"img_v3_0289_5e2b5dec-e916-4abb-adb4-c72745cfa97g",
			"https://mp.weixin.qq.com/s/X98sAGRAVprBt8-IxZrsSg",
		),
		example1BuildPost(
			"img_v3_0289_c0f7cd0c-4388-4c53-9078-8996ef0a6e6g",
			"**巴奴火锅**",
			"巴奴的数字化团队，不仅服务和支持着内外部用户，还要做业务创新和提效的主导者。飞书开放能力与飞书应用引擎助力着越来越多的低代码应用服务着巴奴的员工们，实现了关键业务在手边。",
			"img_v3_0289_18073e19-b8c5-49e4-84e1-e4c93b1f3d5g",
			"https://mp.weixin.qq.com/s/X98sAGRAVprBt8-IxZrsSg",
		),
		card.Divider(),
		card.Action(
			card.Button("一键直达客户案例").
				SetTypePrimaryFilled().
				SetSizeMedium().
				SetIcon(card.StandardIcon("cursor_outlined").SetColorWhite()).
				SetMultiURL(card.MultiURL("https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzkzOTUwNDgwNQ==&action=getalbum&album_id=3013784104276000776&scene=173&subscene=227&sessionid=1706868434&enterid=1706868439&from_msgid=2247487830&from_itemidx=1&count=3&nolastread=1#wechat_redirect")),
			card.Button("查看更多社区精选").
				SetTypePrimary().
				SetSizeMedium().
				SetIcon(card.StandardIcon("thumbsup_filled").SetColorBlue()).
				SetMultiURL(card.MultiURL("https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzkzOTUwNDgwNQ==&action=getalbum&album_id=3093555415403134979&scene=173&subscene=227&sessionid=1706868379&enterid=1706868382&from_msgid=2247489622&from_itemidx=1&count=3&nolastread=1#wechat_redirect")),
		),
	).SetHeader(
		card.Header("开工大吉！快来「飞书开发者」看看近期有哪些更新吧").
			SetTextTagList(
				card.TextTag("微信公众号").SetColorGreen(),
				card.TextTag("不定期抽奖").SetColorOrange(),
			).
			SetTemplateIndigo().
			SetIcon(card.CustomIcon("img_v3_027m_91fb8fc2-e6a3-406c-aa9d-04569e29bccg")),
	)
}

func example1BuildPost(icon, title, content, image, url string) *card.ComponentColumnSet {
	return card.ColumnSet(
		card.Column(
			card.Markdown(title).
				SetTextAlignLeft().
				SetTextSizeNormal().
				SetIcon(card.CustomIcon(icon)),
			card.Markdown(content).
				SetTextAlignLeft().
				SetTextSizeNotation(),
		).
			SetWidthWeighted().
			SetVerticalAlignTop().
			SetWeight(1),
		card.Column(
			card.ColumnSet(
				card.Column(
					card.PlainText("").SetText(
						card.Text("").
							SetTextSizeNormal().
							SetTextAlignLeft().
							SetTextColorDefault(),
					),
				).
					SetWidthWeighted().
					SetWeight(1),
			).
				SetFlexModeNone().
				SetHorizontalSpacingDefault().
				SetBackgroundStyleDefault(),
			card.Image(image).
				SetScaleTypeCropCenter().
				SetSizeMedium().
				SetCornerRadiusPixel(10),
			// TODO: "complex_interaction": true,
			card.Button("查看详情").
				SetTypePrimaryText().
				SetSizeSmall().
				SetIcon(card.StandardIcon("laser_outlined").SetColorBlue()).
				SetBehaviors(card.OpenURLBehavior(url, "", "", "")),
		).
			SetWidthAuto().
			SetVerticalAlignTop(),
	).
		SetFlexModeNone().
		SetHorizontalAlignLeft().
		SetMargin(card.MarginTop(16))
}
