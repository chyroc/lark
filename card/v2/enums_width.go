package card

import (
	"strconv"
)

type Width = string

const (
	WidthDefault Width = "default" // 默认宽度
	WidthFill    Width = "fill"    // 卡片最大支持宽度
	WidthAuto    Width = "auto"    // 自适应内容宽度
	// 自定义宽度, 如 120px。超出卡片宽度时将按最大支持宽度展示
	// 自定义宽度百分比：自定义列宽度占当前表格画布宽度的百分比（表格画布宽度 = 卡片宽度-卡片左右内边距）, 如 25%。取值范围是 [1%,100%]
)

func PixelWidth(px int64) Width {
	return strconv.FormatInt(px, 10) + "px"
}

func PercentWidth(percent int64) Width {
	return strconv.FormatInt(percent, 10) + "%"
}
