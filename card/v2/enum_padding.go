package card

import (
	"strconv"
)

// Padding 列的内边距
//
// 单值, 如 "10px", 表示列的四个外边距都为 10 px。
// 多值, 如 "4px 12px 4px 12px", 表示列的上、右、下、左的外边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔。
type Padding = string

func ConvPadding(padding int64) string {
	return strconv.FormatInt(padding, 10) + "px"
}

func Paddings(top, right, bottom, left int64) string {
	return ConvPadding(top) + " " + ConvPadding(right) + " " + ConvPadding(bottom) + " " + ConvPadding(left)
}

func PaddingTop(padding int64) string {
	return Paddings(padding, 0, 0, 0)
}

func PaddingRight(padding int64) string {
	return Paddings(0, padding, 0, 0)
}

func PaddingBottom(padding int64) string {
	return Paddings(0, 0, padding, 0)
}

func PaddingLeft(padding int64) string {
	return Paddings(0, 0, 0, padding)
}

func PaddingTopBottom(padding int64) string {
	return Paddings(padding, 0, padding, 0)
}

func PaddingLeftRight(padding int64) string {
	return Paddings(0, padding, 0, padding)
}

func PaddingAll(padding int64) string {
	return Paddings(padding, padding, padding, padding)
}
