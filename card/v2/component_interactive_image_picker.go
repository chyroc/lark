package card

import "encoding/json"

func ImagePicker(name string, imgKeys ...string) *ComponentImagePicker {
	imageOptions := make([]*ImagePickerOption, 0, len(imgKeys))
	for _, imgKey := range imgKeys {
		imageOptions = append(imageOptions, &ImagePickerOption{
			ImgKey: imgKey,
			Value:  imgKey,
		})
	}
	return &ComponentImagePicker{
		Name:    name,
		Options: imageOptions,
	}
}

// ComponentImagePicker 多图选择
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/image-picker
//
// 多图选择组件是用于提供图片选项的交互组件, 支持单选、多选图片。多图选择组件适用于以图片为主要选项的场景, 如在卡片中展示商品图、模板图、AI生成的图片等供用户选择。本文档介绍多图选择组件的 JSON 结构和相关属性。
//
// 注意事项
// 多图选择组件仅支持通过撰写卡片 JSON 代码的方式使用, 暂不支持在卡片搭建工具上构建使用。
// 多图选择组件支持飞书 V7.6 及以上版本的客户端。在低于该版本的飞书客户端上, 多图选择的内容将展示为“请升级至最新版本客户端, 以查看内容”的占位图。
//
// 嵌套规则
// 多图选择组件可嵌套在卡片的根节点、多列布局、表单容器组件中。在不同的嵌套关系中, 多图选择组件支持的交互形态不同：
//
// 当多图选择组件不内嵌在表单容器中时, 多图选择组件仅支持单选图片, 且终端用户点击图片选项后立即提交, 触发回传交互, 不支持多选和异步提交。
// 当多图选择组件内嵌在表单容器中时, 多图选择组件支持单选、多选交互并支持异步提交, 即终端用户需选择图片后, 点击表单容器的提交按钮, 将本地缓存的表单内容一次回调至开发者服务端, 实现异步提交。
//
// 回调结构
// 为组件成功配置交互后, 用户基于组件进行交互时, 你在开发者后台配置的请求地址将会收到回调数据。
// 如果你添加的是新版卡片回传交互回调(card.action.trigger), 可参考卡片回传交互了解回调结构。
// 如果你添加的是旧版卡片回传交互回调(card.action.trigger_v1), 可参考消息卡片回传交互（旧）了解回调结构。
//
//go:generate generate_set_attrs -type=ComponentImagePicker
//go:generate generate_to_map -type=ComponentImagePicker
type ComponentImagePicker struct {
	componentBase

	// 图片加载等状态时的组件风格样式。可取值：
	//
	// default：默认灰色样式
	// laser：彩色渐变样式, 建议 AI 场景使用
	Style string `json:"style,omitempty"`

	// 图片是否多选。可选值：
	//
	// true：多选, 仅支持异步提交。多图选择组件需内嵌在表单容器中, 否则卡片 JSON 报错。
	// false：单选。
	// 组件在表单容器内时, 图片选项展示为带单选按钮（radio button）的异步提交样式。
	// 组件不在表单容器内时, 图片选项展示为不带单选按钮（radio button）的同步提交样式。
	MultiSelect bool `json:"multi_select,omitempty"`

	// 图片选项的布局方式。可选值：
	//
	// stretch：每个选项的图片宽度撑满父容器宽度, 高度按图片大小等比例缩放。
	// bisect：二等分排布, 每个选项图片宽度占父容器的 1/2, 高度按图片大小等比例缩放。
	// trisect：三等分排布, 每个选项图片宽度占父容器的 1/3, 高度按图片大小等比例缩放。
	Layout string `json:"layout,omitempty"`

	// 该日期选择器组件的唯一标识。当日期选择器内嵌在表单容器时, 该属性生效, 用于识别用户提交的数据属于哪个组件。
	//
	// 注意: 当日期选择器组件嵌套在表单容器中时, 该字段必填且需在卡片全局内唯一。
	Name string `json:"name,omitempty"`

	// 日期的内容是否必选。当组件内嵌在表单容器中时, 该属性可用。其它情况将报错或不生效。可取值：
	//
	// true：日期必选。当用户点击表单容器的“提交”时, 未填写日期, 则前端提示“有必填项未填写”, 不会向开发者的服务端发起回传请求。
	// false：日期选填。当用户点击表单容器的“提交”时, 未填写日期, 仍提交表单容器中的数据。
	Required bool `json:"required,omitempty"`

	// 点击图片选项后是否弹窗放大图片。当多图选择组件嵌套在表单容器中时, 该属性生效。
	//
	// true：点击图片后, 弹出图片查看器放大查看当前点击的图片。
	// false：点击图片后, 响应卡片本身的交互事件, 不弹出图片查看器。
	CanPreview bool `json:"can_preview,omitempty"`

	// 选项中图片的宽高比。图片按最短边撑满图片渲染容器, 按照居中裁剪的方式自适应裁剪。可取值：
	//
	// 1:1
	// 16:9
	// 4:3
	AspectRatio AspectRatio `json:"aspect_ratio,omitempty"`

	// 是否禁用该日期选择器。该属性仅支持飞书 V7.4 及以上版本的客户端。可选值：
	//
	// true：禁用日期选择器组件
	// false：日期选择器组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 禁用整个组件后, 用户将光标悬浮在整个组件上时展示的禁用提示文案。
	DisabledTips *ObjectText `json:"disabled_tips,omitempty"`

	// 你可在交互事件中自定义回传参数, 支持回传字符串, 或 "key":"value" 构成的对象结构体。
	Value any `json:"value,omitempty"`

	// 选项数组, 用于配置多图选择组件中每个图片选项的属性。
	Options []*ImagePickerOption `json:"options,omitempty"`
}

type ImagePickerOption struct {
	// 图片资源的 Key。你可以调用上传图片接口或在搭建工具中上传图片, 获取图片的 key。
	ImgKey string `json:"img_key,omitempty"`

	// 自定义每个图片选项的回传参数。在回传交互中指定的回传参数将透传至开发者的服务端。
	Value any `json:"value,omitempty"`

	// 是否禁用某个图片选项。可选值：
	//
	// true：禁用该选项
	// false：选项保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 禁用某个图片选项后, 用户将光标悬浮在选项上或点击选项时展示的禁用提示文案。
	DisabledTips *ObjectText `json:"disabled_tips,omitempty"`

	// 用户在 PC 端将光标悬浮在多图选择上方时的文案提醒。默认为空。
	HoverTips string `json:"hover_tips,omitempty"`
}

// MarshalJSON ...
func (r ComponentImagePicker) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "select_img"
	return json.Marshal(m)
}

// SetStyle set ComponentImagePicker.Style attribute
func (r *ComponentImagePicker) SetStyle(val string) *ComponentImagePicker {
	r.Style = val
	return r
}

// SetMultiSelect set ComponentImagePicker.MultiSelect attribute
func (r *ComponentImagePicker) SetMultiSelect(val bool) *ComponentImagePicker {
	r.MultiSelect = val
	return r
}

// SetLayout set ComponentImagePicker.Layout attribute
func (r *ComponentImagePicker) SetLayout(val string) *ComponentImagePicker {
	r.Layout = val
	return r
}

// SetName set ComponentImagePicker.Name attribute
func (r *ComponentImagePicker) SetName(val string) *ComponentImagePicker {
	r.Name = val
	return r
}

// SetRequired set ComponentImagePicker.Required attribute
func (r *ComponentImagePicker) SetRequired(val bool) *ComponentImagePicker {
	r.Required = val
	return r
}

// SetCanPreview set ComponentImagePicker.CanPreview attribute
func (r *ComponentImagePicker) SetCanPreview(val bool) *ComponentImagePicker {
	r.CanPreview = val
	return r
}

// SetAspectRatio set ComponentImagePicker.AspectRatio attribute
func (r *ComponentImagePicker) SetAspectRatio(val AspectRatio) *ComponentImagePicker {
	r.AspectRatio = val
	return r
}

// SetAspectRatioOneOne set ComponentImagePicker.AspectRatio attribute to AspectRatioOneOne
func (r *ComponentImagePicker) SetAspectRatioOneOne() *ComponentImagePicker {
	r.AspectRatio = AspectRatioOneOne
	return r
}

// SetAspectRatioTwoOne set ComponentImagePicker.AspectRatio attribute to AspectRatioTwoOne
func (r *ComponentImagePicker) SetAspectRatioTwoOne() *ComponentImagePicker {
	r.AspectRatio = AspectRatioTwoOne
	return r
}

// SetAspectRatioFourThree set ComponentImagePicker.AspectRatio attribute to AspectRatioFourThree
func (r *ComponentImagePicker) SetAspectRatioFourThree() *ComponentImagePicker {
	r.AspectRatio = AspectRatioFourThree
	return r
}

// SetAspectRatioSixteenNine set ComponentImagePicker.AspectRatio attribute to AspectRatioSixteenNine
func (r *ComponentImagePicker) SetAspectRatioSixteenNine() *ComponentImagePicker {
	r.AspectRatio = AspectRatioSixteenNine
	return r
}

// SetDisabled set ComponentImagePicker.Disabled attribute
func (r *ComponentImagePicker) SetDisabled(val bool) *ComponentImagePicker {
	r.Disabled = val
	return r
}

// SetDisabledTips set ComponentImagePicker.DisabledTips attribute
func (r *ComponentImagePicker) SetDisabledTips(val *ObjectText) *ComponentImagePicker {
	r.DisabledTips = val
	return r
}

// SetValue set ComponentImagePicker.Value attribute
func (r *ComponentImagePicker) SetValue(val any) *ComponentImagePicker {
	r.Value = val
	return r
}

// SetOptions set ComponentImagePicker.Options attribute
func (r *ComponentImagePicker) SetOptions(val ...*ImagePickerOption) *ComponentImagePicker {
	r.Options = val
	return r
}

// toMap conv ComponentImagePicker to map
func (r *ComponentImagePicker) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 11)
	if r.Style != "" {
		res["style"] = r.Style
	}
	if r.MultiSelect != false {
		res["multi_select"] = r.MultiSelect
	}
	if r.Layout != "" {
		res["layout"] = r.Layout
	}
	if r.Name != "" {
		res["name"] = r.Name
	}
	if r.Required != false {
		res["required"] = r.Required
	}
	if r.CanPreview != false {
		res["can_preview"] = r.CanPreview
	}
	if r.AspectRatio != "" {
		res["aspect_ratio"] = r.AspectRatio
	}
	if r.Disabled != false {
		res["disabled"] = r.Disabled
	}
	if r.DisabledTips != nil {
		res["disabled_tips"] = r.DisabledTips
	}
	if r.Value != 0 {
		res["value"] = r.Value
	}
	if len(r.Options) != 0 {
		res["options"] = r.Options
	}
	return res
}
