package card

type VerticalAlign = string

const (
	VerticalAlignTop    VerticalAlign = "top"    // 上对齐
	VerticalAlignCenter VerticalAlign = "center" // 居中对齐
	VerticalAlignBottom VerticalAlign = "bottom" // 下对齐
)

type TextAlign = string

const (
	TextAlignLeft   TextAlign = "left"   // 左对齐
	TextAlignCenter TextAlign = "center" // 居中对齐
	TextAlignRight  TextAlign = "right"  // 右对齐
)

type TextSize = string

const (
	TextSizeHeading0  TextSize = "heading-0"  // 特大标题（30px）
	TextSizeHeading1  TextSize = "heading-1"  // 一级标题（24px）
	TextSizeHeading2  TextSize = "heading-2"  // 二级标题（20 px）
	TextSizeHeading3  TextSize = "heading-3"  // 三级标题（18px）
	TextSizeHeading4  TextSize = "heading-4"  // 四级标题（16px）
	TextSizeHeading   TextSize = "heading"    // 标题（16px）
	TextSizeNormal    TextSize = "normal"     // 正文（14px）
	TextSizeNotation  TextSize = "notation"   // 辅助信息（12px）
	TextSizeXxxxLarge TextSize = "xxxx-large" // 30px
	TextSizeXxxLarge  TextSize = "xxx-large"  // 24px
	TextSizeXxLarge   TextSize = "xx-large"   // 20px
	TextSizeXLarge    TextSize = "x-large"    // 18px
	TextSizeLarge     TextSize = "large"      // 16px
	TextSizeMedium    TextSize = "medium"     // 14px
	TextSizeSmall     TextSize = "small"      // 12px
	TextSizeXSmall    TextSize = "x-small"    // 10px
)

// BorderType 组件边框样式
type BorderType = string

const (
	BorderTypeDefault BorderType = "default" // 带边框样式
	BorderTypeText    BorderType = "text"    // 不带边框的纯文本样式
)
