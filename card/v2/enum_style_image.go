package card

type ScaleType = string

const (
	ScaleTypeCropCenter    ScaleType = "crop_center"    // 居中裁剪
	ScaleTypeCropTop       ScaleType = "crop_top"       // 顶部裁剪
	ScaleTypeFitHorizontal ScaleType = "fit_horizontal" // 完整展示不裁剪
)

type ImageSize = string

const (
	ImageSizeLarge                 ImageSize = "large"                   // 大图, 尺寸为 160 × 160, 适用于多图混排。
	ImageSizeMedium                ImageSize = "medium"                  // 中图, 尺寸为 80 × 80, 适用于图文混排的封面图。
	ImageSizeSmall                 ImageSize = "small"                   // 小图, 尺寸为 40 × 40, 适用于人员头像。
	ImageSizeTiny                  ImageSize = "tiny"                    // 超小图, 尺寸为 16 × 16, 适用于图标、备注。
	ImageSizeStretchWithoutPadding ImageSize = "stretch_without_padding" // 通栏图, 适用于高宽比小于 16:9 的图片, 图片的宽度将撑满卡片宽度。
	// ImageSize[1,999] [1,999]px = "[1,999] [1,999]px"// 自定义图片尺寸, 单位为像素, 中间用空格分隔。
)
