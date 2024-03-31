package card

import (
	"strconv"
)

// Margin 列的外边距
//
// 单值, 如 "10px", 表示列的四个外边距都为 10 px。
// 多值, 如 "4px 12px 4px 12px", 表示列的上、右、下、左的外边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔。
// 注意：首行列的上外边距强制为 0, 末行列的下外边距强制为 0。
type Margin = string

func ConvMargin(margin int64) string {
	return strconv.FormatInt(margin, 10) + "px"
}

func Margins(top, right, bottom, left int64) string {
	return ConvMargin(top) + " " + ConvMargin(right) + " " + ConvMargin(bottom) + " " + ConvMargin(left)
}

func MarginTop(margin int64) string {
	return Margins(margin, 0, 0, 0)
}

func MarginRight(margin int64) string {
	return Margins(0, margin, 0, 0)
}

func MarginBottom(margin int64) string {
	return Margins(0, 0, margin, 0)
}

func MarginLeft(margin int64) string {
	return Margins(0, 0, 0, margin)
}

func MarginTopBottom(margin int64) string {
	return Margins(margin, 0, margin, 0)
}

func MarginLeftRight(margin int64) string {
	return Margins(0, margin, 0, margin)
}

func MarginAll(margin int64) string {
	return Margins(margin, margin, margin, margin)
}
