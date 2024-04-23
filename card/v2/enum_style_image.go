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

type ImageMode = string

const (
	ImageModeCropCenter    ImageMode = "crop_center"    // 居中裁剪模式，对长图会限高，并居中裁剪后展示。
	ImageModeFitHorizontal ImageMode = "fit_horizontal" // 平铺模式，宽度撑满卡片完整展示上传的图片。
	ImageModeStretch       ImageMode = "stretch"        // 自适应。图片宽度撑满卡片宽度，当图片 高:宽 小于 16:9 时，完整展示原图。当图片 高:宽 大于 16:9 时，顶部对齐裁剪图片，并在图片底部展示 长图 脚标。
	ImageModeLarge         ImageMode = "large"          // 大图，尺寸为 160 × 160，适用于多图混排。
	ImageModeMedium        ImageMode = "medium"         // 中图，尺寸为 80 × 80，适用于图文混排的封面图。
	ImageModeSmall         ImageMode = "small"          // 小图，尺寸为 40 × 40，适用于人员头像。
	ImageModeTiny          ImageMode = "tiny"           // 超小图，尺寸为 16 × 16，适用于图标、备注。
)
