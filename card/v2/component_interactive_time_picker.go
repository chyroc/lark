package card

import "encoding/json"

func TimePicker(name string) *ComponentTimePicker {
	return &ComponentTimePicker{
		Name: name,
	}
}

// ComponentTimePicker 时间选择器
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/time-selector
// 时间选择器组件是用于提供时间选项的交互组件。本文档介绍时间选择器组件的 JSON 结构和相关属性。
//
// 注意事项
// 时间选择器组件支持飞书 V3.11.0 及以上版本的客户端。在低于该版本的飞书客户端上, 表单容器的内容将展示为“当前版本不支持查看此消息”。
// 在使用时间选择器时, 你需提醒用户与当前选择时间场景相对应的时区信息。例如, 在预定海外酒店的场景下, 一般使用酒店所在地时区; 设置日程场景下, 一般使用用户当前所在地的时区。开放平台会返回用户当前的时区作为参考, 但并不代表用户选择了该时区。
// 在卡片 JSON 代码中, 若时间选择器组件直接位于卡片根节点, 而非嵌套在其它组件中, 你需将其 JSON 数据配置在交互模块（"tag": "action"）的 actions 字段中使用。
//
// 嵌套规则
// 时间选择器组件支持嵌套在分栏、表单容器、折叠面板、循环容器中使用。
//
// 回调结构
// 为组件成功配置交互后, 用户基于组件进行交互时, 你在开发者后台配置的请求地址将会收到回调数据。
// 如果你添加的是新版卡片回传交互回调(card.action.trigger), 可参考卡片回传交互了解回调结构。
// 如果你添加的是旧版卡片回传交互回调(card.action.trigger_v1), 可参考消息卡片回传交互（旧）了解回调结构。
//
//go:generate generate_set_attrs -type=ComponentTimePicker
//go:generate generate_to_map -type=ComponentTimePicker
type ComponentTimePicker struct {
	componentBase

	// 该时间选择器组件的唯一标识。用于识别用户提交的数据属于哪个组件。
	//
	// 注意: 当时间选择器组件嵌套在表单容器中时, 该字段必填且需在卡片全局内唯一
	Name string `json:"name,omitempty"`

	// 时间否必选。当组件内嵌在表单容器中时, 该属性可用。其它情况将报错或不生效。可取值:
	//
	// true: 时间必选。当用户点击表单容器的“提交”时, 未填写时间, 则前端提示“有必填项未填写”, 不会向开发者的服务端发起回传请求。
	// false: 时间选填。当用户点击表单容器的“提交”时, 未填写时间, 仍提交表单容器中的数据。
	Required bool `json:"required,omitempty"`

	// 是否禁用该时间选择器。该属性仅支持飞书 V7.4 及以上版本的客户端。可选值:
	//
	// true: 禁用时间选择器组件
	// false: 时间选择器组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 时间选择器组件的初始选项值。格式为 HH:mm。该配置将会覆盖 placeholder 配置的占位文本。
	InitialTime string `json:"initial_time,omitempty"`

	// 时间选择器组件内的占位文本。
	//
	// 注意:
	//
	// 未配置 initial_time 字段设置初始选项值时, 该字段必填。
	// 配置 initial_time 字段设置初始选项值后, 该字段不再生效。
	Placeholder *ObjectText `json:"placeholder,omitempty"`

	// 时间选择器组件的宽度。支持以下枚举值:
	//
	// default: 默认宽度
	// fill: 卡片最大支持宽度
	// [100,∞)px: 自定义宽度。超出卡片宽度时将按最大支持宽度展示
	Width Width `json:"width,omitempty"`

	// 设置交互的回传数据, 当用户点击交互组件的选项后, 会将 value 的值返回给接收回调数据的服务器。后续你可以通过服务器接收的 value 值进行业务处理。
	//
	// 该字段值仅支持 key-value 形式的 JSON 结构, 且 key 为 String 类型。示例值:
	Value map[string]any `json:"value,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意: confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`
}

// MarshalJSON ...
func (r ComponentTimePicker) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "Time_picker"
	return json.Marshal(m)
}

// SetName set ComponentTimePicker.Name attribute
func (r *ComponentTimePicker) SetName(val string) *ComponentTimePicker {
	r.Name = val
	return r
}

// SetRequired set ComponentTimePicker.Required attribute
func (r *ComponentTimePicker) SetRequired(val bool) *ComponentTimePicker {
	r.Required = val
	return r
}

// SetDisabled set ComponentTimePicker.Disabled attribute
func (r *ComponentTimePicker) SetDisabled(val bool) *ComponentTimePicker {
	r.Disabled = val
	return r
}

// SetInitialTime set ComponentTimePicker.InitialTime attribute
func (r *ComponentTimePicker) SetInitialTime(val string) *ComponentTimePicker {
	r.InitialTime = val
	return r
}

// SetPlaceholder set ComponentTimePicker.Placeholder attribute
func (r *ComponentTimePicker) SetPlaceholder(val *ObjectText) *ComponentTimePicker {
	r.Placeholder = val
	return r
}

// SetWidth set ComponentTimePicker.Width attribute
func (r *ComponentTimePicker) SetWidth(val Width) *ComponentTimePicker {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentTimePicker.Width attribute to WidthDefault
func (r *ComponentTimePicker) SetWidthDefault() *ComponentTimePicker {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentTimePicker.Width attribute to WidthFill
func (r *ComponentTimePicker) SetWidthFill() *ComponentTimePicker {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentTimePicker.Width attribute to WidthAuto
func (r *ComponentTimePicker) SetWidthAuto() *ComponentTimePicker {
	r.Width = WidthAuto
	return r
}

// SetValue set ComponentTimePicker.Value attribute
func (r *ComponentTimePicker) SetValue(val map[string]any) *ComponentTimePicker {
	r.Value = val
	return r
}

// SetConfirm set ComponentTimePicker.Confirm attribute
func (r *ComponentTimePicker) SetConfirm(val *ObjectConfirm) *ComponentTimePicker {
	r.Confirm = val
	return r
}

// toMap conv ComponentTimePicker to map
func (r *ComponentTimePicker) toMap() map[string]interface{} {
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
	if r.InitialTime != "" {
		res["initial_time"] = r.InitialTime
	}
	if r.Placeholder != nil {
		res["placeholder"] = r.Placeholder
	}
	if r.Width != "" {
		res["width"] = r.Width
	}
	if len(r.Value) != 0 {
		res["value"] = r.Value
	}
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	return res
}
