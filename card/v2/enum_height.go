package card

import (
	"strconv"
)

// Height 高度
//
// auto: 自适应高度
// [10,999]px: 自定义高度, 如 "20px"
type Height = string

const (
	HeightAuto Height = "auto"
)

func PixelHeight(px int64) Height {
	return strconv.FormatInt(px, 10) + "px"
}
