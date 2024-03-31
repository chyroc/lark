package card

import "encoding/json"

func DatetimePicker(name string) *ComponentDatetimePicker {
	return &ComponentDatetimePicker{
		Name: name,
	}
}

// ComponentDatetimePicker 日期时间选择器
//
// docs:https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/date-time-picker
// 日期时间选择器组件是用于提供时间和日期选项的交互组件。本文档介绍日期时间选择器组件的 JSON 结构和相关属性。
//
// 注意事项
// 日期时间选择器组件支持飞书 V3.11.0 及以上版本的客户端。在低于该版本的飞书客户端上, 表单容器的内容将展示为“当前版本不支持查看此消息”。
// 在使用日期时间选择器时, 你需提醒用户与当前选择场景相对应的时区信息。例如, 在预定海外酒店的场景下, 一般使用酒店所在地时区; 设置日程场景下, 一般使用用户当前所在地的时区。开放平台会返回用户当前的时区作为参考, 但并不代表用户选择了该时区。
// 在卡片 JSON 代码中, 若日期时间选择器组件直接位于卡片根节点, 而非嵌套在其它组件中, 你需将其 JSON 数据配置在交互模块（"tag": "action"）的 actions 字段中使用。
//
// 嵌套规则
// 日期时间选择器组件支持嵌套在分栏、表单容器、折叠面板、循环容器中使用。
//
// 回调结构
// 为组件成功配置交互后, 用户基于组件进行交互时, 你在开发者后台配置的请求地址将会收到回调数据。
// 如果你添加的是新版卡片回传交互回调(card.action.trigger), 可参考卡片回传交互了解回调结构。
// 如果你添加的是旧版卡片回传交互回调(card.action.trigger_v1), 可参考消息卡片回传交互（旧）了解回调结构。
//
//go:generate generate_set_attrs -type=ComponentDatetimePicker
//go:generate generate_to_map -type=ComponentDatetimePicker
type ComponentDatetimePicker struct {
	componentBase

	// 该日期选择器组件的唯一标识。当日期选择器内嵌在表单容器时, 该属性生效, 用于识别用户提交的数据属于哪个组件。
	//
	// 注意: 当日期选择器组件嵌套在表单容器中时, 该字段必填且需在卡片全局内唯一。
	Name string `json:"name,omitempty"`

	// 日期的内容是否必选。当组件内嵌在表单容器中时, 该属性可用。其它情况将报错或不生效。可取值：
	//
	// true：日期必选。当用户点击表单容器的“提交”时, 未填写日期, 则前端提示“有必填项未填写”, 不会向开发者的服务端发起回传请求。
	// false：日期选填。当用户点击表单容器的“提交”时, 未填写日期, 仍提交表单容器中的数据。
	Required bool `json:"required,omitempty"`

	// 是否禁用该日期选择器。该属性仅支持飞书 V7.4 及以上版本的客户端。可选值：
	//
	// true：禁用日期选择器组件
	// false：日期选择器组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 日期选择器组件内的占位文本。
	Placeholder *ObjectText `json:"placeholder,omitempty"`

	// 日期选择器组件的宽度。支持以下枚举值：
	//
	// default：默认宽度
	// fill：卡片最大支持宽度
	// [100,∞)px：自定义宽度。超出卡片宽度时将按最大支持宽度展示
	Width Width `json:"width,omitempty"`

	// 日期日期时间选择器组件的初始选项值。格式为 yyyy-MM-dd HH:mm。该配置将会覆盖 placeholder 配置的占位文本。
	InitialDatetime string `json:"initial_datetime,omitempty"`

	// 设置交互的回传数据, 当用户点击交互组件的选项后, 会将 value 的值返回给接收回调数据的服务器。后续你可以通过服务器接收的 value 值进行业务处理。该字段值仅支持 key-value 形式的 JSON 结构, 且 key 为 String 类型。示例值：
	Value map[string]any `json:"value,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意：confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`
}

// MarshalJSON ...
func (r ComponentDatetimePicker) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "picker_datetime"
	return json.Marshal(m)
}

// SetName set ComponentDatetimePicker.Name attribute
func (r *ComponentDatetimePicker) SetName(val string) *ComponentDatetimePicker {
	r.Name = val
	return r
}

// SetRequired set ComponentDatetimePicker.Required attribute
func (r *ComponentDatetimePicker) SetRequired(val bool) *ComponentDatetimePicker {
	r.Required = val
	return r
}

// SetDisabled set ComponentDatetimePicker.Disabled attribute
func (r *ComponentDatetimePicker) SetDisabled(val bool) *ComponentDatetimePicker {
	r.Disabled = val
	return r
}

// SetPlaceholder set ComponentDatetimePicker.Placeholder attribute
func (r *ComponentDatetimePicker) SetPlaceholder(val *ObjectText) *ComponentDatetimePicker {
	r.Placeholder = val
	return r
}

// SetWidth set ComponentDatetimePicker.Width attribute
func (r *ComponentDatetimePicker) SetWidth(val Width) *ComponentDatetimePicker {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentDatetimePicker.Width attribute to WidthDefault
func (r *ComponentDatetimePicker) SetWidthDefault() *ComponentDatetimePicker {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentDatetimePicker.Width attribute to WidthFill
func (r *ComponentDatetimePicker) SetWidthFill() *ComponentDatetimePicker {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentDatetimePicker.Width attribute to WidthAuto
func (r *ComponentDatetimePicker) SetWidthAuto() *ComponentDatetimePicker {
	r.Width = WidthAuto
	return r
}

// SetInitialDatetime set ComponentDatetimePicker.InitialDatetime attribute
func (r *ComponentDatetimePicker) SetInitialDatetime(val string) *ComponentDatetimePicker {
	r.InitialDatetime = val
	return r
}

// SetValue set ComponentDatetimePicker.Value attribute
func (r *ComponentDatetimePicker) SetValue(val map[string]any) *ComponentDatetimePicker {
	r.Value = val
	return r
}

// SetConfirm set ComponentDatetimePicker.Confirm attribute
func (r *ComponentDatetimePicker) SetConfirm(val *ObjectConfirm) *ComponentDatetimePicker {
	r.Confirm = val
	return r
}

// toMap conv ComponentDatetimePicker to map
func (r *ComponentDatetimePicker) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 8)
	if r.Name != "" {
		res["name"] = r.Name
	}
	if r.Required != false {
		res["required"] = r.Required
	}
	if r.Disabled != false {
		res["disabled"] = r.Disabled
	}
	if r.Placeholder != nil {
		res["placeholder"] = r.Placeholder
	}
	if r.Width != "" {
		res["width"] = r.Width
	}
	if r.InitialDatetime != "" {
		res["initial_datetime"] = r.InitialDatetime
	}
	if len(r.Value) != 0 {
		res["value"] = r.Value
	}
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	return res
}
