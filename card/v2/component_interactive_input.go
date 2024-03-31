package card

import "encoding/json"

func Input(name string) *ComponentInput {
	return &ComponentInput{
		Name: name,
	}
}

// ComponentInput 输入框
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/input
//
// 在使用卡片进行内容收集的场景下, 你可能需要同时获取用户的主观内容, 如原因、评价、备注等。在这种情况下, 你可以使用输入框组件, 实现简单的文本内容收集的场景。本文档介绍输入框组件的 JSON 结构和相关属性。
//
// 注意事项
// 输入框仅支持飞书 V6.8 及以上版本的客户端。在低于该版本的飞书客户端上, 输入框的内容将默认展示为一句“请升级至最新版本客户端, 以查看内容”的占位图。你也可通过卡片的 fallback 字段自定义组件的降级展示方式。
// 输入框组件仅支持单行纯文本格式, 不支持多行文本和富文本。
// 在卡片 JSON 代码中, 若输入框组件直接位于卡片根节点, 而非嵌套在其它组件中, 你需将其 JSON 数据配置在交互模块（"tag": "action"）的 actions 字段中使用, 并删除 required 字段, 否则将报错。
//
// 嵌套规则
// 输入框组件支持嵌套在分栏、表单容器、折叠面板、循环容器中使用。在表单容器中, 输入框组件的数据为异步提交的形式, 即用户填写完所有表单项后, 点击表单容器中绑定提交事件的按钮, 才会将包括输入框组件的所有数据一次回调至开发者的服务端。
//
//go:generate generate_set_attrs -type=ComponentInput
//go:generate generate_to_map -type=ComponentInput
type ComponentInput struct {
	componentBase

	// 输入框的唯一标识。当输入框内嵌在表单容器时, 该属性生效, 用于识别用户提交的文本属于哪个输入框。
	//
	// 注意：当输入框组件嵌套在表单容器中时, 该字段必填且需在卡片全局内唯一。
	Name string `json:"name,omitempty"`

	// 输入框的内容是否必填。当输入框内嵌在表单容器时, 该属性可用。其它情况将报错或不生效。可取值：
	//
	// true：输入框必填。当用户点击表单容器的“提交”时, 未填写输入框, 则前端提示“有必填项未填写”, 不会向开发者的服务端发起回传请求。
	// false：输入框选填。当用户点击表单容器的“提交”时, 未填写输入框, 仍提交表单容器中的数据。
	Required bool `json:"required,omitempty"`

	// 是否禁用该输入框。该属性仅支持飞书 V7.4 及以上版本的客户端。可选值：
	//
	// true：禁用输入框组件
	// false：输入框组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 输入框中的占位文本。
	Placeholder *ObjectText `json:"placeholder,omitempty"`

	// 输入框中为用户预填写的内容。展示为用户在输入框中输入文本后待提交的样式。
	DefaultValue string `json:"default_value,omitempty"`

	// 输入框的宽度。支持以下枚举值：
	//
	// default: 默认宽度
	// fill: 卡片最大支持宽度
	// [100,∞)px: 自定义宽度。超出卡片宽度时将按最大支持宽度展示
	Width Width `json:"width,omitempty"`

	// 输入框可容纳的最大文本长度, 可取 1~1,000 范围内的整数。
	// 当用户输入的文本字符数超过最大文本长度, 组件将报错提示。
	MaxLength int64 `json:"max_length,omitempty"`

	// 文本标签, 即对输入框的描述, 用于提示用户要填写的内容。多用于表单容器中内嵌的输入框组件。
	Label *ObjectText `json:"label,omitempty"`

	// 文本标签的位置。可取值：
	//
	// top: 文本标签位于输入框上方
	// left: 文本标签位于输入框左边
	// 注意：在移动端等窄屏幕场景下, 文本标签将自适应固定展示在输入框上方
	LabelPosition LabelPosition `json:"label_position,omitempty"`

	// 你可在交互事件中自定义回传数据, 支持 string 或 object 数据类型。
	Value interface{} `json:"value,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意：confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`

	// 设置输入框组件的降级文案。由于输入框仅支持飞书 V6.8 及以上版本的客户端, 你需选择在低于此版本的客户端上, 该组件的降级展示方式：
	//
	// 不填写该字段, 使用系统默认的降级文案：“请升级至最新版本客户端, 以查看内容”
	// "drop": 填写 "drop", 在旧版本客户端上直接丢弃该输入框组件
	// 使用 text 文本对象自定义降级文案
	Fallback *ObjectFallback `json:"fallback,omitempty"`
}

type LabelPosition = string

const (
	LabelPositionTop  LabelPosition = "top"  // 文本标签位于输入框上方
	LabelPositionLeft LabelPosition = "left" // 文本标签位于输入框左边
)

// MarshalJSON ...
func (r ComponentInput) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "input"
	return json.Marshal(m)
}

// SetName set ComponentInput.Name attribute
func (r *ComponentInput) SetName(val string) *ComponentInput {
	r.Name = val
	return r
}

// SetRequired set ComponentInput.Required attribute
func (r *ComponentInput) SetRequired(val bool) *ComponentInput {
	r.Required = val
	return r
}

// SetDisabled set ComponentInput.Disabled attribute
func (r *ComponentInput) SetDisabled(val bool) *ComponentInput {
	r.Disabled = val
	return r
}

// SetPlaceholder set ComponentInput.Placeholder attribute
func (r *ComponentInput) SetPlaceholder(val *ObjectText) *ComponentInput {
	r.Placeholder = val
	return r
}

// SetDefaultValue set ComponentInput.DefaultValue attribute
func (r *ComponentInput) SetDefaultValue(val string) *ComponentInput {
	r.DefaultValue = val
	return r
}

// SetWidth set ComponentInput.Width attribute
func (r *ComponentInput) SetWidth(val Width) *ComponentInput {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentInput.Width attribute to WidthDefault
func (r *ComponentInput) SetWidthDefault() *ComponentInput {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentInput.Width attribute to WidthFill
func (r *ComponentInput) SetWidthFill() *ComponentInput {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentInput.Width attribute to WidthAuto
func (r *ComponentInput) SetWidthAuto() *ComponentInput {
	r.Width = WidthAuto
	return r
}

// SetMaxLength set ComponentInput.MaxLength attribute
func (r *ComponentInput) SetMaxLength(val int64) *ComponentInput {
	r.MaxLength = val
	return r
}

// SetLabel set ComponentInput.Label attribute
func (r *ComponentInput) SetLabel(val *ObjectText) *ComponentInput {
	r.Label = val
	return r
}

// SetLabelPosition set ComponentInput.LabelPosition attribute
func (r *ComponentInput) SetLabelPosition(val LabelPosition) *ComponentInput {
	r.LabelPosition = val
	return r
}

// SetLabelPositionTop set ComponentInput.LabelPosition attribute to LabelPositionTop
func (r *ComponentInput) SetLabelPositionTop() *ComponentInput {
	r.LabelPosition = LabelPositionTop
	return r
}

// SetLabelPositionLeft set ComponentInput.LabelPosition attribute to LabelPositionLeft
func (r *ComponentInput) SetLabelPositionLeft() *ComponentInput {
	r.LabelPosition = LabelPositionLeft
	return r
}

// SetValue set ComponentInput.Value attribute
func (r *ComponentInput) SetValue(val interface{}) *ComponentInput {
	r.Value = val
	return r
}

// SetConfirm set ComponentInput.Confirm attribute
func (r *ComponentInput) SetConfirm(val *ObjectConfirm) *ComponentInput {
	r.Confirm = val
	return r
}

// SetFallback set ComponentInput.Fallback attribute
func (r *ComponentInput) SetFallback(val *ObjectFallback) *ComponentInput {
	r.Fallback = val
	return r
}

// toMap conv ComponentInput to map
func (r *ComponentInput) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 12)
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
	if r.DefaultValue != "" {
		res["default_value"] = r.DefaultValue
	}
	if r.Width != "" {
		res["width"] = r.Width
	}
	if r.MaxLength != 0 {
		res["max_length"] = r.MaxLength
	}
	if r.Label != nil {
		res["label"] = r.Label
	}
	if r.LabelPosition != "" {
		res["label_position"] = r.LabelPosition
	}
	if r.Value != 0 {
		res["value"] = r.Value
	}
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	if r.Fallback != nil {
		res["fallback"] = r.Fallback
	}
	return res
}
