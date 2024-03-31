package card

import (
	"encoding/json"
)

func Image(imgKey string) *ComponentImage {
	return &ComponentImage{
		ImgKey:  imgKey,
		Alt:     Text(""),
		Preview: true,
	}
}

// ComponentImage 图片
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/image
//
// 飞书卡片支持图片组件。
// 你可调用上传图片接口或在搭建工具的图片组件中上传图片, 获取图片的 key 传入图片组件中, 使卡片内容更丰富。
//
// 为保证图片在聊天窗口中呈现的清晰度, 建议你在组件中上传的图片遵从以下规范：
//
// 图片尺寸在 1500 × 3000 px 的范围内。
// 图片大小不超过 10 M。
// 图片的 高度:宽度 不超过 16:9。
//
//go:generate generate_set_attrs -type=ComponentImage
//go:generate generate_to_map -type=ComponentImage
type ComponentImage struct {
	componentBase

	// 图片资源的 Key。你可以调用上传图片接口或在搭建工具中上传图片, 获取图片的 key。
	// 上传图片接口: https://open.larkoffice.com/document/server-docs/im-v1/image/create
	ImgKey string `json:"img_key,omitempty"`

	// 悬浮（hover）在图片上时展示的说明文案
	Alt *ObjectText `json:"alt,omitempty"`

	// 图片标题
	Title *ObjectText `json:"title,omitempty"`

	// 图片的圆角半径, 单位是像素（px）。取值遵循以下格式：
	//
	// [0,∞]px
	// [0,100]%
	CornerRadius CornerRadius `json:"corner_radius,omitempty"`

	// 图片的裁剪模式, 当 size 字段的比例和图片的比例不一致时会触发裁剪。可取值：
	//
	// crop_center: 居中裁剪
	// crop_top: 顶部裁剪
	// fit_horizontal: 完整展示不裁剪
	ScaleType ScaleType `json:"scale_type,omitempty"`

	// 图片尺寸。仅在 scale_type 字段为 crop_center 或 crop_top 时生效。可取值：
	//
	// large：大图, 尺寸为 160 × 160, 适用于多图混排。
	// medium：中图, 尺寸为 80 × 80, 适用于图文混排的封面图。
	// small：小图, 尺寸为 40 × 40, 适用于人员头像。
	// tiny：超小图, 尺寸为 16 × 16, 适用于图标、备注。
	// stretch_without_padding：通栏图, 适用于高宽比小于 16:9 的图片, 图片的宽度将撑满卡片宽度。
	// [1,999]px [1,999]px：自定义图片尺寸, 单位为像素, 中间用空格分隔。
	Size ImageSize `json:"size,omitempty"`

	// 是否为透明底色。默认为 false, 即图片为白色底色。
	Transparent bool `json:"transparent,omitempty"`

	// 点击后是否放大图片。
	//
	// true：点击图片后, 弹出图片查看器放大查看当前点击的图片。
	// false：点击图片后, 响应卡片本身的交互事件, 不弹出图片查看器。
	// 提示：如果你为卡片配置了跳转链接card_link参数, 可将该参数设置为 false, 后续用户点击卡片上的图片也能响应 card_link 链接跳转。
	Preview bool `json:"preview,omitempty"`
}

// MarshalJSON ...
func (r ComponentImage) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "img"
	return json.Marshal(m)
}

func (r *ComponentImage) SetCornerRadiusPixel(val int64) *ComponentImage {
	r.CornerRadius = PixelCornerRadius(val)
	return r
}

func (r *ComponentImage) SetCornerRadiusPercent(val int64) *ComponentImage {
	r.CornerRadius = PercentCornerRadius(val)
	return r
}

// SetImgKey set ComponentImage.ImgKey attribute
func (r *ComponentImage) SetImgKey(val string) *ComponentImage {
	r.ImgKey = val
	return r
}

// SetAlt set ComponentImage.Alt attribute
func (r *ComponentImage) SetAlt(val *ObjectText) *ComponentImage {
	r.Alt = val
	return r
}

// SetTitle set ComponentImage.Title attribute
func (r *ComponentImage) SetTitle(val *ObjectText) *ComponentImage {
	r.Title = val
	return r
}

// SetCornerRadius set ComponentImage.CornerRadius attribute
func (r *ComponentImage) SetCornerRadius(val CornerRadius) *ComponentImage {
	r.CornerRadius = val
	return r
}

// SetScaleType set ComponentImage.ScaleType attribute
func (r *ComponentImage) SetScaleType(val ScaleType) *ComponentImage {
	r.ScaleType = val
	return r
}

// SetScaleTypeCropCenter set ComponentImage.ScaleType attribute to ScaleTypeCropCenter
func (r *ComponentImage) SetScaleTypeCropCenter() *ComponentImage {
	r.ScaleType = ScaleTypeCropCenter
	return r
}

// SetScaleTypeCropTop set ComponentImage.ScaleType attribute to ScaleTypeCropTop
func (r *ComponentImage) SetScaleTypeCropTop() *ComponentImage {
	r.ScaleType = ScaleTypeCropTop
	return r
}

// SetScaleTypeFitHorizontal set ComponentImage.ScaleType attribute to ScaleTypeFitHorizontal
func (r *ComponentImage) SetScaleTypeFitHorizontal() *ComponentImage {
	r.ScaleType = ScaleTypeFitHorizontal
	return r
}

// SetSize set ComponentImage.Size attribute
func (r *ComponentImage) SetSize(val ImageSize) *ComponentImage {
	r.Size = val
	return r
}

// SetSizeLarge set ComponentImage.Size attribute to ImageSizeLarge
func (r *ComponentImage) SetSizeLarge() *ComponentImage {
	r.Size = ImageSizeLarge
	return r
}

// SetSizeMedium set ComponentImage.Size attribute to ImageSizeMedium
func (r *ComponentImage) SetSizeMedium() *ComponentImage {
	r.Size = ImageSizeMedium
	return r
}

// SetSizeSmall set ComponentImage.Size attribute to ImageSizeSmall
func (r *ComponentImage) SetSizeSmall() *ComponentImage {
	r.Size = ImageSizeSmall
	return r
}

// SetSizeTiny set ComponentImage.Size attribute to ImageSizeTiny
func (r *ComponentImage) SetSizeTiny() *ComponentImage {
	r.Size = ImageSizeTiny
	return r
}

// SetSizeStretchWithoutPadding set ComponentImage.Size attribute to ImageSizeStretchWithoutPadding
func (r *ComponentImage) SetSizeStretchWithoutPadding() *ComponentImage {
	r.Size = ImageSizeStretchWithoutPadding
	return r
}

// SetTransparent set ComponentImage.Transparent attribute
func (r *ComponentImage) SetTransparent(val bool) *ComponentImage {
	r.Transparent = val
	return r
}

// SetPreview set ComponentImage.Preview attribute
func (r *ComponentImage) SetPreview(val bool) *ComponentImage {
	r.Preview = val
	return r
}

// toMap conv ComponentImage to map
func (r *ComponentImage) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 8)
	if r.ImgKey != "" {
		res["img_key"] = r.ImgKey
	}
	if r.Alt != nil {
		res["alt"] = r.Alt
	}
	if r.Title != nil {
		res["title"] = r.Title
	}
	if r.CornerRadius != "" {
		res["corner_radius"] = r.CornerRadius
	}
	if r.ScaleType != "" {
		res["scale_type"] = r.ScaleType
	}
	if r.Size != "" {
		res["size"] = r.Size
	}
	if r.Transparent != false {
		res["transparent"] = r.Transparent
	}
	if r.Preview != false {
		res["preview"] = r.Preview
	}
	return res
}
