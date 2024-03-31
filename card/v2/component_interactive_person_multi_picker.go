package card

import "encoding/json"

// PersonMultiPicker 创建人员选择-多选组件
func PersonMultiPicker(name string, ids ...string) *ComponentPersonMultiPicker {
	opts := make([]*PersonPickerOption, 0)
	for _, id := range ids {
		opts = append(opts, &PersonPickerOption{
			Value: id,
		})
	}
	return &ComponentPersonMultiPicker{
		Name:    name,
		Options: opts,
	}
}

// ComponentPersonMultiPicker 人员选择-多选
//
// docs:https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/multi-select-user-picker
//
// 人员选择-多选组件支持添加指定人员作为多选选项。多选组件是一种交互组件, 需嵌入在表单容器中使用。本文档介绍人员选择-多选组件的 JSON 结构和相关属性。
//
// 注意事项
// 人员选择-多选组件支持飞书 V7.4 及以上版本的客户端。在低于该版本的飞书客户端上, 人员选择-多选组件的内容将展示为“请升级至最新版客户端以查看操作”。
// 人员选择-多选组件仅支持通过撰写卡片 JSON 代码的方式使用, 暂不支持在卡片搭建工具上构建使用。
//
// 嵌套规则
// 人员选择-多选组件仅支持嵌入在表单容器中使用, 通过表单容器的“提交”按钮提交选择的内容。了解表单容器及其交互配置, 参考表单容器。
//
//go:generate generate_set_attrs -type=ComponentPersonMultiPicker
//go:generate generate_to_map -type=ComponentPersonMultiPicker
type ComponentPersonMultiPicker struct {
	componentBase

	// 组件边框样式。可选值：
	//
	// default：带边框样式
	// text：不带边框的纯文本样式
	Type BorderType `json:"type,omitempty"`

	//
	// 表单容器中组件的唯一标识。用于识别用户提交的数据属于哪个组件。在同一张卡片内, 该字段的值全局唯一。
	Name string `json:"name,omitempty"`

	// 多选组件的内容是否必选。当组件内嵌在表单容器中时, 该属性生效。可取值：
	//
	// true：多选组件必选。当用户点击表单容器的“提交”时, 未填写多选组件, 则前端提示“有必填项未填写”, 不会向开发者的服务端发起回传请求。
	// false：多选组件选填。当用户点击表单容器的“提交”时, 未填写多选组件, 仍提交表单容器中的数据。
	Required bool `json:"required,omitempty"`

	// 是否禁用该多选组件。可选值：
	//
	// true：禁用多选组件
	// false：多选组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 人员选择组件内的占位文本。
	Placeholder *ObjectText `json:"placeholder,omitempty"`

	// 人员选择组件的宽度。支持以下枚举值：
	//
	// default：默认宽度
	// fill：卡片最大支持宽度
	// [100,∞)px：自定义宽度。超出卡片宽度时将按最大支持宽度展示
	Width Width `json:"width,omitempty"`

	// 多选组件默认选中的选项。数组项的值需要和 options.value 对应。
	SelectedValues []any `json:"selected_values,omitempty"`

	// 选项值配置。按选项数组的顺序展示选项内容。
	Options []*PersonPickerOption `json:"options,omitempty"`
}

// MarshalJSON ...
func (r ComponentPersonMultiPicker) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "PersonMultiPicker"
	return json.Marshal(m)
}

// SetType set ComponentPersonMultiPicker.Type attribute
func (r *ComponentPersonMultiPicker) SetType(val BorderType) *ComponentPersonMultiPicker {
	r.Type = val
	return r
}

// SetTypeDefault set ComponentPersonMultiPicker.Type attribute to BorderTypeDefault
func (r *ComponentPersonMultiPicker) SetTypeDefault() *ComponentPersonMultiPicker {
	r.Type = BorderTypeDefault
	return r
}

// SetTypeText set ComponentPersonMultiPicker.Type attribute to BorderTypeText
func (r *ComponentPersonMultiPicker) SetTypeText() *ComponentPersonMultiPicker {
	r.Type = BorderTypeText
	return r
}

// SetName set ComponentPersonMultiPicker.Name attribute
func (r *ComponentPersonMultiPicker) SetName(val string) *ComponentPersonMultiPicker {
	r.Name = val
	return r
}

// SetRequired set ComponentPersonMultiPicker.Required attribute
func (r *ComponentPersonMultiPicker) SetRequired(val bool) *ComponentPersonMultiPicker {
	r.Required = val
	return r
}

// SetDisabled set ComponentPersonMultiPicker.Disabled attribute
func (r *ComponentPersonMultiPicker) SetDisabled(val bool) *ComponentPersonMultiPicker {
	r.Disabled = val
	return r
}

// SetPlaceholder set ComponentPersonMultiPicker.Placeholder attribute
func (r *ComponentPersonMultiPicker) SetPlaceholder(val *ObjectText) *ComponentPersonMultiPicker {
	r.Placeholder = val
	return r
}

// SetWidth set ComponentPersonMultiPicker.Width attribute
func (r *ComponentPersonMultiPicker) SetWidth(val Width) *ComponentPersonMultiPicker {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentPersonMultiPicker.Width attribute to WidthDefault
func (r *ComponentPersonMultiPicker) SetWidthDefault() *ComponentPersonMultiPicker {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentPersonMultiPicker.Width attribute to WidthFill
func (r *ComponentPersonMultiPicker) SetWidthFill() *ComponentPersonMultiPicker {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentPersonMultiPicker.Width attribute to WidthAuto
func (r *ComponentPersonMultiPicker) SetWidthAuto() *ComponentPersonMultiPicker {
	r.Width = WidthAuto
	return r
}

// SetSelectedValues set ComponentPersonMultiPicker.SelectedValues attribute
func (r *ComponentPersonMultiPicker) SetSelectedValues(val ...any) *ComponentPersonMultiPicker {
	r.SelectedValues = val
	return r
}

// SetOptions set ComponentPersonMultiPicker.Options attribute
func (r *ComponentPersonMultiPicker) SetOptions(val ...*PersonPickerOption) *ComponentPersonMultiPicker {
	r.Options = val
	return r
}

// toMap conv ComponentPersonMultiPicker to map
func (r *ComponentPersonMultiPicker) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 8)
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
	if r.Placeholder != nil {
		res["placeholder"] = r.Placeholder
	}
	if r.Width != "" {
		res["width"] = r.Width
	}
	if len(r.SelectedValues) != 0 {
		res["selected_values"] = r.SelectedValues
	}
	if len(r.Options) != 0 {
		res["options"] = r.Options
	}
	return res
}
