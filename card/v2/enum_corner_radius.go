package card

import (
	"strconv"
)

// CornerRadius 圆角半径, 单位是像素（px）或百分比（%）。取值遵循以下格式：
//
// [0,∞]px, 如 "10px"
// [0,100]%, 如 "30%"
type CornerRadius = string

func PixelCornerRadius(px int64) CornerRadius {
	return strconv.FormatInt(px, 10) + "px"
}

func PercentCornerRadius(percent int64) CornerRadius {
	return strconv.FormatInt(percent, 10) + "%"
}
