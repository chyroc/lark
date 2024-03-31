package card

import "encoding/json"

func MultiImage(imgKeys ...string) *ComponentMultiImage {
	imgList := make([]*ImageResource, len(imgKeys))
	for i, v := range imgKeys {
		imgList[i] = &ImageResource{
			ImgKey: v,
		}
	}
	return &ComponentMultiImage{
		ImgList: imgList,
	}
}

// ComponentMultiImage 多图混排
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/multi-image-laylout
//
// 飞书卡片支持多图混排组件。你可调用上传图片接口或在新版飞书卡片搭建工具中上传图片, 获取图片的 key 传入多图混排组件中, 使卡片内容更丰富。
//
// 在内容推送场景, 你可能需要在卡片内组织编排多张图片。此时你可以使用多图混排组件, 选择图片混排方式, 快速构建多图样式。
//
// 为保证图片在聊天窗口中呈现的清晰度, 建议你在组件中上传的图片遵从以下规范：
//
// 图片尺寸在 1500 × 3000 px 的范围内。
// 图片大小不超过 10 M。
// 图片的 高度:宽度 不超过 16:9。
//
//go:generate generate_set_attrs -type=ComponentMultiImage
//go:generate generate_to_map -type=ComponentMultiImage
type ComponentMultiImage struct {
	componentBase

	// 多图混排的方式, 可取值：
	//
	// double: 双图混排, 最多可排布两张图。
	// triple: 三图混排, 最多可排布三张图。
	// bisect: 等分双列图混排, 每行两个等大的正方形图, 最多可排布三行, 即六张图。
	// trisect: 等分三列图混排, 每行三个等大的正方形图, 最多可排布三行, 即九张图。
	//
	// 若上传的图片数量超过混排方式可容纳的上限, 则系统将根据图片上传的顺序, 优先展示排列顺序中靠前的图片。超出上限的图片将不再显示。
	// 若上传的图片数量未达到混排方式可容纳的上限, 则未排布的部分将保留空白。
	CombinationMode CombinationMode `json:"combination_mode,omitempty"`

	// 多图混排图片的圆角半径, 单位是像素（px）。取值遵循以下格式：
	//
	// [0,∞]px
	// [0,100]%
	CornerRadius CornerRadius `json:"corner_radius,omitempty"`

	// 图片资源的 img_key 数组, 顺序与图片排列顺序一致。
	ImgList []*ImageResource `json:"img_list,omitempty"`
}

func (r ComponentMultiImage) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "img_combination"
	return json.Marshal(m)
}

type CombinationMode = string

const (
	CombinationModeDouble  CombinationMode = "double"  // 双图混排, 最多可排布两张图。
	CombinationModeTriple  CombinationMode = "triple"  // 三图混排, 最多可排布三张图。
	CombinationModeBisect  CombinationMode = "bisect"  // 等分双列图混排, 每行两个等大的正方形图, 最多可排布三行, 即六张图。
	CombinationModeTrisect CombinationMode = "trisect" // 等分三列图混排, 每行三个等大的正方形图, 最多可排布三行, 即九张图。
)

type ImageResource struct {
	ImgKey string `json:"img_key,omitempty"`
}

func (r *ComponentMultiImage) SetCornerRadiusPixel(val int64) *ComponentMultiImage {
	r.CornerRadius = PixelCornerRadius(val)
	return r
}

func (r *ComponentMultiImage) SetCornerRadiusPercent(val int64) *ComponentMultiImage {
	r.CornerRadius = PercentCornerRadius(val)
	return r
}

// SetCombinationMode set ComponentMultiImage.CombinationMode attribute
func (r *ComponentMultiImage) SetCombinationMode(val CombinationMode) *ComponentMultiImage {
	r.CombinationMode = val
	return r
}

// SetCombinationModeDouble set ComponentMultiImage.CombinationMode attribute to CombinationModeDouble
func (r *ComponentMultiImage) SetCombinationModeDouble() *ComponentMultiImage {
	r.CombinationMode = CombinationModeDouble
	return r
}

// SetCombinationModeTriple set ComponentMultiImage.CombinationMode attribute to CombinationModeTriple
func (r *ComponentMultiImage) SetCombinationModeTriple() *ComponentMultiImage {
	r.CombinationMode = CombinationModeTriple
	return r
}

// SetCombinationModeBisect set ComponentMultiImage.CombinationMode attribute to CombinationModeBisect
func (r *ComponentMultiImage) SetCombinationModeBisect() *ComponentMultiImage {
	r.CombinationMode = CombinationModeBisect
	return r
}

// SetCombinationModeTrisect set ComponentMultiImage.CombinationMode attribute to CombinationModeTrisect
func (r *ComponentMultiImage) SetCombinationModeTrisect() *ComponentMultiImage {
	r.CombinationMode = CombinationModeTrisect
	return r
}

// SetCornerRadius set ComponentMultiImage.CornerRadius attribute
func (r *ComponentMultiImage) SetCornerRadius(val CornerRadius) *ComponentMultiImage {
	r.CornerRadius = val
	return r
}

// SetImgList set ComponentMultiImage.ImgList attribute
func (r *ComponentMultiImage) SetImgList(val ...*ImageResource) *ComponentMultiImage {
	r.ImgList = val
	return r
}

// toMap conv ComponentMultiImage to map
func (r *ComponentMultiImage) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 3)
	if r.CombinationMode != "" {
		res["combination_mode"] = r.CombinationMode
	}
	if r.CornerRadius != "" {
		res["corner_radius"] = r.CornerRadius
	}
	if len(r.ImgList) != 0 {
		res["img_list"] = r.ImgList
	}
	return res
}
