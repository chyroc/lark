package card

import "encoding/json"

// Button 按钮
func Button(text string) *ComponentButton {
	return &ComponentButton{
		Type:  ButtonTypeDefault,
		Width: WidthDefault,
		Text:  Text(text),
	}
}

// FormButton 表单中的按钮
func FormButton(name, text string) *ComponentButton {
	return &ComponentButton{
		Type:       ButtonTypeDefault,
		Width:      WidthDefault,
		Text:       Text(text),
		Name:       name,
		ActionType: ButtonActionTypeFormSubmit,
	}
}

// ComponentButton 按钮
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/button
//
// 按钮组件是一种交互组件, 支持多种样式和尺寸, 并支持添加图标作为前缀图标。本文档介绍按钮组件的 JSON 结构和相关属性。
//
// 嵌套规则
// 按钮组件支持嵌套在分栏、表单容器、折叠面板、循环容器中使用。
//
//go:generate generate_set_attrs -type=ComponentButton
//go:generate generate_to_map -type=ComponentButton
type ComponentButton struct {
	componentBase

	// 按钮的类型。可选值：
	//
	// default: 黑色字体按钮, 有边框
	// primary: 蓝色字体按钮, 有边框
	// danger: 红色字体按钮, 有边框
	// text: 黑色字体按钮, 无边框
	// primary_text: 蓝色字体按钮, 无边框
	// danger_text: 红色字体按钮, 无边框
	// primary_filled: 蓝底白字按钮
	// danger_filled: 红底白字按钮
	// laser: 镭射按钮
	Type ButtonType `json:"type,omitempty"`

	// 按钮的尺寸。可选值:
	//
	// tiny: 超小尺寸, PC 端为 24 px; 移动端为 28 px
	// small: 小尺寸, PC 端为 28 px; 移动端为 28 px
	// medium: 中尺寸, PC 端为 32 px; 移动端为 36 px
	// large: 大尺寸, PC 端为 40 px; 移动端为 48 px
	Size ButtonSize `json:"size,omitempty"`

	// 按钮的宽度。支持以下枚举值:
	//
	// default: 默认宽度
	// fill: 卡片最大支持宽度
	// [100,∞)px: 自定义宽度, 如 120px。超出卡片宽度时将按最大支持宽度展示
	Width Width `json:"width,omitempty"`

	// 按钮上的文本。
	Text *ObjectText `json:"text,omitempty"`

	// 添加图标作为文本前缀图标。支持自定义或使用图标库中的图标。
	Icon *ObjectIcon `json:"icon,omitempty"`

	// 用户在 PC 端将光标悬浮在交互容器上方时的文案提醒。默认为空。
	HoverTips string `json:"hover_tips,omitempty"`

	// 是否禁按钮。可选值：
	//
	// true：禁用按钮
	// false：按钮组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 禁用按钮后, 用户触发交互时的弹窗文案提醒。默认为空, 即不弹窗。
	DisabledTips *ObjectText `json:"disabled_tips,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意：confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`

	// 基于 url 元素配置多端跳转链接，详情参见旧版文档url 元素。该字段与 url 字段不可同时设置。
	MultiURL *ObjectMultiURL `json:"multi_url,omitempty"`

	// 配置交互类型和具体交互行为。详情参考配置卡片交互。
	Behaviors []*ObjectBehavior `json:"behaviors,omitempty"`

	// 该字段用于配置回传交互。当用户点击交互组件后，会将 value 的值返回给接收回调数据的服务器。后续你可以通过服务器接收的 value 值进行业务处理。
	//
	// 该字段值仅支持 key-value 形式的 JSON 结构，且 key 为 String 类型。示例值：
	Value map[string]any `json:"value,omitempty"`

	// 用于提交表单的按钮组件的唯一标识，用于识别用户在交互后，点击的是哪个按钮。在表单容器中所有的交互组件中，该字段必填，否则数据会发送失败。
	Name string `json:"name,omitempty"`

	// 用于提交表单的按钮组件的交互类型。固定取值 form_submit，表示提交表单。
	ActionType ButtonActionType `json:"action_type,omitempty"`
}

type ButtonType = string

const (
	ButtonTypeDefault       ButtonType = "default"        // 黑色字体按钮, 有边框
	ButtonTypePrimary       ButtonType = "primary"        // 蓝色字体按钮, 有边框
	ButtonTypeDanger        ButtonType = "danger"         // 红色字体按钮, 有边框
	ButtonTypeText          ButtonType = "text"           // 黑色字体按钮, 无边框
	ButtonTypePrimaryText   ButtonType = "primary_text"   // 蓝色字体按钮, 无边框
	ButtonTypeDangerText    ButtonType = "danger_text"    // 红色字体按钮, 无边框
	ButtonTypePrimaryFilled ButtonType = "primary_filled" // 蓝底白字按钮
	ButtonTypeDangerFilled  ButtonType = "danger_filled"  // 红底白字按钮
	ButtonTypeLaser         ButtonType = "laser"          // 镭射按钮
)

type ButtonSize = string

const (
	ButtonSizeTiny   ButtonSize = "tiny"   // 超小尺寸, PC 端为 24 px; 移动端为 28 px
	ButtonSizeSmall  ButtonSize = "small"  // 小尺寸, PC 端为 28 px; 移动端为 28 px
	ButtonSizeMedium ButtonSize = "medium" // 中尺寸, PC 端为 32 px; 移动端为 36 px
	ButtonSizeLarge  ButtonSize = "large"  // 大尺寸, PC 端为 40 px; 移动端为 48 px
)

type ButtonActionType = string

const (
	ButtonActionTypeFormSubmit ButtonActionType = "form_submit"
)

// MarshalJSON ...
func (r ComponentButton) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "button"
	if len(r.Value) > 0 {
		m["complex_interaction"] = true // 是否同时生效上述历史字段配置的跳转链接交互和回传交互。默认仅生效跳转链接交互。
	}
	return json.Marshal(m)
}

// SetType set ComponentButton.Type attribute
func (r *ComponentButton) SetType(val ButtonType) *ComponentButton {
	r.Type = val
	return r
}

// SetTypeDefault set ComponentButton.Type attribute to ButtonTypeDefault
func (r *ComponentButton) SetTypeDefault() *ComponentButton {
	r.Type = ButtonTypeDefault
	return r
}

// SetTypePrimary set ComponentButton.Type attribute to ButtonTypePrimary
func (r *ComponentButton) SetTypePrimary() *ComponentButton {
	r.Type = ButtonTypePrimary
	return r
}

// SetTypeDanger set ComponentButton.Type attribute to ButtonTypeDanger
func (r *ComponentButton) SetTypeDanger() *ComponentButton {
	r.Type = ButtonTypeDanger
	return r
}

// SetTypeText set ComponentButton.Type attribute to ButtonTypeText
func (r *ComponentButton) SetTypeText() *ComponentButton {
	r.Type = ButtonTypeText
	return r
}

// SetTypePrimaryText set ComponentButton.Type attribute to ButtonTypePrimaryText
func (r *ComponentButton) SetTypePrimaryText() *ComponentButton {
	r.Type = ButtonTypePrimaryText
	return r
}

// SetTypeDangerText set ComponentButton.Type attribute to ButtonTypeDangerText
func (r *ComponentButton) SetTypeDangerText() *ComponentButton {
	r.Type = ButtonTypeDangerText
	return r
}

// SetTypePrimaryFilled set ComponentButton.Type attribute to ButtonTypePrimaryFilled
func (r *ComponentButton) SetTypePrimaryFilled() *ComponentButton {
	r.Type = ButtonTypePrimaryFilled
	return r
}

// SetTypeDangerFilled set ComponentButton.Type attribute to ButtonTypeDangerFilled
func (r *ComponentButton) SetTypeDangerFilled() *ComponentButton {
	r.Type = ButtonTypeDangerFilled
	return r
}

// SetTypeLaser set ComponentButton.Type attribute to ButtonTypeLaser
func (r *ComponentButton) SetTypeLaser() *ComponentButton {
	r.Type = ButtonTypeLaser
	return r
}

// SetSize set ComponentButton.Size attribute
func (r *ComponentButton) SetSize(val ButtonSize) *ComponentButton {
	r.Size = val
	return r
}

// SetSizeTiny set ComponentButton.Size attribute to ButtonSizeTiny
func (r *ComponentButton) SetSizeTiny() *ComponentButton {
	r.Size = ButtonSizeTiny
	return r
}

// SetSizeSmall set ComponentButton.Size attribute to ButtonSizeSmall
func (r *ComponentButton) SetSizeSmall() *ComponentButton {
	r.Size = ButtonSizeSmall
	return r
}

// SetSizeMedium set ComponentButton.Size attribute to ButtonSizeMedium
func (r *ComponentButton) SetSizeMedium() *ComponentButton {
	r.Size = ButtonSizeMedium
	return r
}

// SetSizeLarge set ComponentButton.Size attribute to ButtonSizeLarge
func (r *ComponentButton) SetSizeLarge() *ComponentButton {
	r.Size = ButtonSizeLarge
	return r
}

// SetWidth set ComponentButton.Width attribute
func (r *ComponentButton) SetWidth(val Width) *ComponentButton {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentButton.Width attribute to WidthDefault
func (r *ComponentButton) SetWidthDefault() *ComponentButton {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentButton.Width attribute to WidthFill
func (r *ComponentButton) SetWidthFill() *ComponentButton {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentButton.Width attribute to WidthAuto
func (r *ComponentButton) SetWidthAuto() *ComponentButton {
	r.Width = WidthAuto
	return r
}

// SetText set ComponentButton.Text attribute
func (r *ComponentButton) SetText(val *ObjectText) *ComponentButton {
	r.Text = val
	return r
}

// SetIcon set ComponentButton.Icon attribute
func (r *ComponentButton) SetIcon(val *ObjectIcon) *ComponentButton {
	r.Icon = val
	return r
}

// SetHoverTips set ComponentButton.HoverTips attribute
func (r *ComponentButton) SetHoverTips(val string) *ComponentButton {
	r.HoverTips = val
	return r
}

// SetDisabled set ComponentButton.Disabled attribute
func (r *ComponentButton) SetDisabled(val bool) *ComponentButton {
	r.Disabled = val
	return r
}

// SetDisabledTips set ComponentButton.DisabledTips attribute
func (r *ComponentButton) SetDisabledTips(val *ObjectText) *ComponentButton {
	r.DisabledTips = val
	return r
}

// SetConfirm set ComponentButton.Confirm attribute
func (r *ComponentButton) SetConfirm(val *ObjectConfirm) *ComponentButton {
	r.Confirm = val
	return r
}

// SetMultiURL set ComponentButton.MultiURL attribute
func (r *ComponentButton) SetMultiURL(val *ObjectMultiURL) *ComponentButton {
	r.MultiURL = val
	return r
}

// SetBehaviors set ComponentButton.Behaviors attribute
func (r *ComponentButton) SetBehaviors(val ...*ObjectBehavior) *ComponentButton {
	r.Behaviors = val
	return r
}

// SetValue set ComponentButton.Value attribute
func (r *ComponentButton) SetValue(val map[string]any) *ComponentButton {
	r.Value = val
	return r
}

// SetName set ComponentButton.Name attribute
func (r *ComponentButton) SetName(val string) *ComponentButton {
	r.Name = val
	return r
}

// SetActionType set ComponentButton.ActionType attribute
func (r *ComponentButton) SetActionType(val ButtonActionType) *ComponentButton {
	r.ActionType = val
	return r
}

// SetActionTypeFormSubmit set ComponentButton.ActionType attribute to ButtonActionTypeFormSubmit
func (r *ComponentButton) SetActionTypeFormSubmit() *ComponentButton {
	r.ActionType = ButtonActionTypeFormSubmit
	return r
}

// toMap conv ComponentButton to map
func (r *ComponentButton) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 14)
	if r.Type != "" {
		res["type"] = r.Type
	}
	if r.Size != "" {
		res["size"] = r.Size
	}
	if r.Width != "" {
		res["width"] = r.Width
	}
	if r.Text != nil {
		res["text"] = r.Text
	}
	if r.Icon != nil {
		res["icon"] = r.Icon
	}
	if r.HoverTips != "" {
		res["hover_tips"] = r.HoverTips
	}
	if r.Disabled != false {
		res["disabled"] = r.Disabled
	}
	if r.DisabledTips != nil {
		res["disabled_tips"] = r.DisabledTips
	}
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	if r.MultiURL != nil {
		res["multi_url"] = r.MultiURL
	}
	if len(r.Behaviors) != 0 {
		res["behaviors"] = r.Behaviors
	}
	if len(r.Value) != 0 {
		res["value"] = r.Value
	}
	if r.Name != "" {
		res["name"] = r.Name
	}
	if r.ActionType != "" {
		res["action_type"] = r.ActionType
	}
	return res
}
