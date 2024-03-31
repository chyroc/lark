package _examples

import (
	"fmt"

	"github.com/chyroc/lark/card/v2"
)

func Example5() *card.ComponentCard {
	datas := []cardData{
		{
			Avatar:          "img_v2_a4b2d72a-211b-440b-8950-880066cfbabg",
			Name:            "李梅",
			Prize:           "金牛奖得主",
			Description:     "销售部门精钢零件销售冠军🥇",
			Performance:     "1198 万",
			DiffMonthColor:  "red",
			DiffMonthRate:   "+11.2%",
			DiffSeasonColor: "red",
			DiffSeasonRate:  "+34.7%",
		},
		{
			Avatar:          "img_v2_0920e890-b6ec-4e06-8889-8abd2690843g",
			Name:            "张锐",
			Prize:           "实力开拓者",
			Description:     "销售部门精钢零件销售亚军🥈",
			Performance:     " 975 万",
			DiffMonthColor:  "green",
			DiffMonthRate:   "-0.52%",
			DiffSeasonColor: "red",
			DiffSeasonRate:  "+13.9%",
		},
		{
			Avatar:          "img_v2_5e220251-1dac-42cb-8901-3031e8bd38cg",
			Name:            "李健",
			Prize:           "业务坚守标兵",
			Description:     "销售部门精钢零件销售季军🥉",
			Performance:     "798 万",
			DiffMonthColor:  "red",
			DiffMonthRate:   "+8.2%",
			DiffSeasonColor: "red",
			DiffSeasonRate:  "+5.8%",
		},
	}

	return card.Card(
		append([]card.Component{
			card.Image("img_v2_196bac64-1731-4bb1-b7c6-d8090482c3bg").SetScaleTypeFitHorizontal(),
		}, buildExample5PostList(datas)...)...,
	).SetHeader(
		card.Header("2024 年 05 月销售部门业绩排行榜来啦 🎉🎉🎉").SetTemplateRed(),
	)
}

type cardData struct {
	Avatar          string `json:"avatar"`
	Name            string `json:"name"`
	Prize           string `json:"prize"`
	Description     string `json:"description"`
	Performance     string `json:"performance"`
	DiffMonthColor  string `json:"diff_month_color"`
	DiffMonthRate   string `json:"diff_month_rate"`
	DiffSeasonColor string `json:"diff_season_color"`
	DiffSeasonRate  string `json:"diff_season_rate"`
}

func buildExample5PostList(data []cardData) []card.Component {
	res := []card.Component{}
	for _, v := range data {
		res = append(res, buildExample5Post(v))
	}
	return res
}

func buildExample5Post(data cardData) *card.ComponentColumnSet {
	return card.ColumnSet(
		card.Column(
			card.Image(data.Avatar).SetScaleTypeCropCenter().SetCornerRadiusPixel(60),
		).SetWidthWeighted().SetVerticalAlignTop().SetWeight(1),

		card.Column(
			card.Markdown(fmt.Sprintf("%s**｜%s**\n<font color='grey'>%s </font>", data.Name, data.Prize, data.Description)).SetTextAlignLeft(),
			card.ColumnSet(
				buildExample5Rate(
					data.Performance,
					"red",
					"本月业绩",
					"grey",
				),
				buildExample5Rate(
					data.DiffMonthRate,
					data.DiffMonthColor,
					"较上月份",
					"grey",
				),
				buildExample5Rate(
					data.DiffSeasonRate,
					data.DiffSeasonColor,
					"较上季度",
					"grey",
				),
			).SetFlexModeNone().SetBackgroundStyleGrey().SetHorizontalAlignLeft(),
		).SetWidthWeighted().SetVerticalAlignTop().SetWeight(4),
	).SetFlexModeNone().SetHorizontalAlignLeft()
}

func buildExample5Rate(title, titleColor, desc, descColor string) *card.ComponentColumn {
	return card.Column(
		card.Markdown(fmt.Sprintf("**<font color='%s'> %s</font>**\n<font color='%s'>%s</font>", titleColor, title, descColor, desc)).SetTextAlignCenter(),
	).SetWidthWeighted().SetVerticalAlignTop().SetWeight(1)
}
