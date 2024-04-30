package card

import "encoding/json"

// StaticPicker 创建下拉选择-单选组件
func StaticPicker(name string, options ...*StaticPickerOption) *ComponentStaticPicker {
	return &ComponentStaticPicker{
		Name:    name,
		Options: options,
	}
}

// ComponentStaticPicker 下拉选择-单选
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/single-select-dropdown-menu
//
// 下拉选择-单选组件支持自定义单选菜单的选项文本、图标和回传参数, 是一种交互组件。本文档介绍下拉选择-单选组件的 JSON 结构和相关属性。
//
// 注意事项
// 下拉选择-单选组件最低支持的飞书版本为 V3.7.0。如果低于该版本, 用户使用该组件时会提示 当前版本不支持查看此消息。
// 在卡片 JSON 代码中, 若下拉选择-单选组件直接位于卡片根节点, 而非嵌套在其它组件中, 你需将其 JSON 数据配置在交互模块（"tag": "action"）的 actions 字段中使用。
//
// 嵌套规则
// 下拉选择-单选组件支持嵌套在分栏、表单容器、折叠面板、循环容器中使用。
//
// 回调结构
// 为组件成功配置交互后, 用户基于组件进行交互时, 你在开发者后台配置的请求地址将会收到回调数据。
//
// 如果你添加的是新版卡片回传交互回调(card.action.trigger), 可参考卡片回传交互了解回调结构。
// 如果你添加的是旧版卡片回传交互回调(card.action.trigger_v1), 可参考消息卡片回传交互（旧）了解回调结构。
//
//go:generate generate_set_attrs -type=ComponentStaticPicker
//go:generate generate_to_map -type=ComponentStaticPicker
type ComponentStaticPicker struct {
	componentBase

	// 组件边框样式。可选值:
	//
	// default: 带边框样式
	// text: 不带边框的纯文本样式
	Type string `json:"type,omitempty"`

	// 单选组件的唯一标识。当单选组件内嵌在表单容器时, 该属性生效, 用于识别用户提交的文本属于哪个单选组件。
	//
	// 注意: 当单选组件嵌套在表单容器中时, 该字段必填且需在卡片全局内唯一。
	Name string `json:"name,omitempty"`

	// 单选组件的内容是否必选。当组件内嵌在表单容器中时, 该属性可用。其它情况将报错或不生效。可取值:
	//
	// true: 单选组件必选。当用户点击表单容器的“提交”时, 未填写单选组件, 则前端提示“有必填项未填写”, 不会向开发者的服务端发起回传请求。
	// false: 单选组件选填。当用户点击表单容器的“提交”时, 未填写单选组件, 仍提交表单容器中的数据。
	Required bool `json:"required,omitempty"`

	// 是否禁用该单选组件。该属性仅支持飞书 V7.4 及以上版本的客户端。可选值:
	//
	// true: 禁用单选组件组件
	// false: 单选组件组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 下拉选择组件的初始选项值。取值上限为选项的数量。该配置将会覆盖 placeholder 配置的占位文本。
	InitialIndex int64 `json:"initial_index,omitempty"`

	// 下拉选择组件内的占位文本。
	Placeholder *ObjectText `json:"placeholder,omitempty"`

	// 单选组件的宽度。支持以下枚举值:
	//
	// default: 默认宽度
	// fill: 卡片最大支持宽度
	// [100,∞)px: 自定义宽度。超出卡片宽度时将按最大支持宽度展示
	Width Width `json:"width,omitempty"`

	// 选项的配置。
	Options []*StaticPickerOption `json:"options,omitempty"`

	// 添加图标作为文本前缀图标。支持自定义或使用图标库中的图标。
	Icon *ObjectIcon `json:"icon,omitempty"`

	// 自定义选项回调值。当用户点击交互组件的选项后, 会将 value 的值返回给接收回调数据的服务器。后续你可以通过服务器接收的 value 值进行业务处理。
	//
	// 该字段值仅支持 key-value 形式的 JSON 结构, 且 key 为 String 类型。示例值:
	Value any `json:"value,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意: confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`
}

//go:generate generate_set_attrs -type=StaticPickerOption
//go:generate generate_to_map -type=StaticPickerOption
type StaticPickerOption struct {
	// 选项的名称。
	Text *ObjectText `json:"text,omitempty"`

	// 添加图标作为文本前缀图标。支持自定义或使用图标库中的图标。
	Icon *ObjectIcon `json:"icon,omitempty"`

	// 自定义选项回调值。当用户点击交互组件的选项后，会将 value 的值返回给接收回调数据的服务器。后续你可以通过服务器接收的 value 值进行业务处理。
	//
	// 该字段值仅支持 key-value 形式的 JSON 结构，且 key 为 String 类型。示例值：
	Value string `json:"value,omitempty"`
}

// MarshalJSON ...
func (r ComponentStaticPicker) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "select_static"
	return json.Marshal(m)
}

// SetType set ComponentStaticPicker.Type attribute
func (r *ComponentStaticPicker) SetType(val string) *ComponentStaticPicker {
	r.Type = val
	return r
}

// SetName set ComponentStaticPicker.Name attribute
func (r *ComponentStaticPicker) SetName(val string) *ComponentStaticPicker {
	r.Name = val
	return r
}

// SetRequired set ComponentStaticPicker.Required attribute
func (r *ComponentStaticPicker) SetRequired(val bool) *ComponentStaticPicker {
	r.Required = val
	return r
}

// SetDisabled set ComponentStaticPicker.Disabled attribute
func (r *ComponentStaticPicker) SetDisabled(val bool) *ComponentStaticPicker {
	r.Disabled = val
	return r
}

// SetInitialIndex set ComponentStaticPicker.InitialIndex attribute
func (r *ComponentStaticPicker) SetInitialIndex(val int64) *ComponentStaticPicker {
	r.InitialIndex = val
	return r
}

// SetPlaceholder set ComponentStaticPicker.Placeholder attribute
func (r *ComponentStaticPicker) SetPlaceholder(val *ObjectText) *ComponentStaticPicker {
	r.Placeholder = val
	return r
}

// SetWidth set ComponentStaticPicker.Width attribute
func (r *ComponentStaticPicker) SetWidth(val Width) *ComponentStaticPicker {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentStaticPicker.Width attribute to WidthDefault
func (r *ComponentStaticPicker) SetWidthDefault() *ComponentStaticPicker {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentStaticPicker.Width attribute to WidthFill
func (r *ComponentStaticPicker) SetWidthFill() *ComponentStaticPicker {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentStaticPicker.Width attribute to WidthAuto
func (r *ComponentStaticPicker) SetWidthAuto() *ComponentStaticPicker {
	r.Width = WidthAuto
	return r
}

// SetOptions set ComponentStaticPicker.Options attribute
func (r *ComponentStaticPicker) SetOptions(val ...*StaticPickerOption) *ComponentStaticPicker {
	r.Options = val
	return r
}

// SetIcon set ComponentStaticPicker.Icon attribute
func (r *ComponentStaticPicker) SetIcon(val *ObjectIcon) *ComponentStaticPicker {
	r.Icon = val
	return r
}

// SetValue set ComponentStaticPicker.Value attribute
func (r *ComponentStaticPicker) SetValue(val any) *ComponentStaticPicker {
	r.Value = val
	return r
}

// SetConfirm set ComponentStaticPicker.Confirm attribute
func (r *ComponentStaticPicker) SetConfirm(val *ObjectConfirm) *ComponentStaticPicker {
	r.Confirm = val
	return r
}

// toMap conv ComponentStaticPicker to map
func (r *ComponentStaticPicker) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 11)
	if r.Type != "" {
		res["type"] = r.Type
	}
	if r.Name != "" {
		res["name"] = r.Name
	}
	if r.Required != false {
		res["required"] = r.Required
	}
	if r.Disabled != false {
		res["disabled"] = r.Disabled
	}
	if r.InitialIndex != 0 {
		res["initial_index"] = r.InitialIndex
	}
	if r.Placeholder != nil {
		res["placeholder"] = r.Placeholder
	}
	if r.Width != "" {
		res["width"] = r.Width
	}
	if len(r.Options) != 0 {
		res["options"] = r.Options
	}
	if r.Icon != nil {
		res["icon"] = r.Icon
	}
	if r.Value != 0 {
		res["value"] = r.Value
	}
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	return res
}

// SetText set StaticPickerOption.Text attribute
func (r *StaticPickerOption) SetText(val *ObjectText) *StaticPickerOption {
	r.Text = val
	return r
}

// SetIcon set StaticPickerOption.Icon attribute
func (r *StaticPickerOption) SetIcon(val *ObjectIcon) *StaticPickerOption {
	r.Icon = val
	return r
}

// SetValue set StaticPickerOption.Value attribute
func (r *StaticPickerOption) SetValue(val string) *StaticPickerOption {
	r.Value = val
	return r
}

// toMap conv StaticPickerOption to map
func (r *StaticPickerOption) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 3)
	if r.Text != nil {
		res["text"] = r.Text
	}
	if r.Icon != nil {
		res["icon"] = r.Icon
	}
	if r.Value != "" {
		res["value"] = r.Value
	}
	return res
}
