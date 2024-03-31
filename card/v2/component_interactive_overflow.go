package card

import "encoding/json"

// Overflow 折叠按钮组
func Overflow(options ...*ObjectOverflowOption) *ComponentOverflow {
	return &ComponentOverflow{
		Options: options,
	}
}

// ComponentOverflow 折叠按钮组
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/overflow
//
// 折叠按钮组支持添加多个按钮并将其折叠。点击按钮组将会展示组内所有按钮。本文档介绍折叠按钮组组件的 JSON 结构和相关属性。
//
// 注意事项
// 折叠按钮组（overflow）最低支持的飞书版本为 V3.7.0, 如果低于该版本, 用户使用该组件时会提示 当前版本不支持查看此消息。
// 在卡片 JSON 代码中, 若折叠按钮组组件直接位于卡片根节点, 而非嵌套在其它组件中, 你需将其 JSON 数据配置在交互模块（"tag": "action"）的 actions 字段中使用。
//
// 嵌套规则
// 按钮组件支持嵌套在表单容器、折叠面板、循环容器中使用。
//
//go:generate generate_set_attrs -type=ComponentOverflow
//go:generate generate_to_map -type=ComponentOverflow
type ComponentOverflow struct {
	componentBase

	// 折叠按钮组的宽度。支持以下枚举值：
	//
	// default: 默认宽度
	// fill: 卡片最大支持宽度
	// [100,∞)px: 自定义宽度。超出卡片宽度时将按最大支持宽度展示
	Width Width `json:"width,omitempty"`

	// 折叠按钮组当中的选项按钮。详见下文 options 字段说明。
	Options []*ObjectOverflowOption `json:"options,omitempty"`

	// 组件整体的回调数据。当用户点击折叠按钮组的折叠按钮后, 会将 value 的值返回给接收回调数据的服务器。后续你可以通过服务器接收的 value 值进行业务处理。
	//
	// 该字段值仅支持 key-value 形式的 JSON 结构, 且 key 为 String 类型。示例值如下:
	Value map[string]interface{} `json:"value,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意：confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`
}

// MarshalJSON ...
func (r ComponentOverflow) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "overflow"
	return json.Marshal(m)
}

// SetWidth set ComponentOverflow.Width attribute
func (r *ComponentOverflow) SetWidth(val Width) *ComponentOverflow {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentOverflow.Width attribute to WidthDefault
func (r *ComponentOverflow) SetWidthDefault() *ComponentOverflow {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentOverflow.Width attribute to WidthFill
func (r *ComponentOverflow) SetWidthFill() *ComponentOverflow {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentOverflow.Width attribute to WidthAuto
func (r *ComponentOverflow) SetWidthAuto() *ComponentOverflow {
	r.Width = WidthAuto
	return r
}

// SetOptions set ComponentOverflow.Options attribute
func (r *ComponentOverflow) SetOptions(val ...*ObjectOverflowOption) *ComponentOverflow {
	r.Options = val
	return r
}

// SetValue set ComponentOverflow.Value attribute
func (r *ComponentOverflow) SetValue(val map[string]interface{}) *ComponentOverflow {
	r.Value = val
	return r
}

// SetConfirm set ComponentOverflow.Confirm attribute
func (r *ComponentOverflow) SetConfirm(val *ObjectConfirm) *ComponentOverflow {
	r.Confirm = val
	return r
}

// toMap conv ComponentOverflow to map
func (r *ComponentOverflow) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 4)
	if r.Width != "" {
		res["width"] = r.Width
	}
	if len(r.Options) != 0 {
		res["options"] = r.Options
	}
	if len(r.Value) != 0 {
		res["value"] = r.Value
	}
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	return res
}
