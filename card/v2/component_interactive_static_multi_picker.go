package card

import "encoding/json"

// StaticMultiPicker 创建下拉选择-多选组件
func StaticMultiPicker(name string, options ...string) *ComponentStaticMultiPicker {
	opts := make([]*StaticMultiPickerOption, 0)
	for _, option := range options {
		opts = append(opts, &StaticMultiPickerOption{
			Text: Text(option),
		})
	}
	return &ComponentStaticMultiPicker{
		Name:    name,
		Options: opts,
	}
}

// ComponentStaticMultiPicker 下拉选择-多选
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/multi-select-dropdown-menu
//
// 下拉选择-多选组件支持自定义多选菜单的选项文本、图标和回传参数, 是一种交互组件, 需嵌入在表单容器中使用。本文档介绍下拉选择-多选组件的 JSON 结构和相关属性。
//
// 注意事项
// 下拉选择-多选组件支持飞书 V7.4 及以上版本的客户端。在低于该版本的飞书客户端上, 下拉选择-多选组件的内容将展示为“请升级至最新版客户端以查看操作”。
// 下拉选择-多选组件仅支持通过撰写卡片 JSON 代码的方式使用, 暂不支持在卡片搭建工具上构建使用。
//
// 嵌套规则
// 下拉选择-多选组件仅支持嵌入在表单容器中使用, 通过表单容器的“提交”按钮提交选择的内容。了解表单容器及其交互配置, 参考表单容器。
//
//go:generate generate_set_attrs -type=ComponentStaticMultiPicker
//go:generate generate_to_map -type=ComponentStaticMultiPicker
type ComponentStaticMultiPicker struct {
	componentBase

	// 组件边框样式。可选值：
	//
	// default：带边框样式
	// text：不带边框的纯文本样式
	Type BorderType `json:"type,omitempty"`

	// 表单容器中组件的唯一标识。当多选组件内嵌在表单容器时, 该属性生效, 用于识别用户提交的数据属于哪个组件。
	//
	// 注意：当多选组件嵌套在表单容器中时, 该字段必填且需在卡片全局内唯一。
	Name string `json:"name,omitempty"`

	// 用户未选择选项时, 下拉选择组件内的占位文本。
	Placeholder *ObjectText `json:"placeholder,omitempty"`

	// 下拉选择组件的宽度。支持以下枚举值：
	//
	// default：默认宽度：
	// 当组件带边框时（即 "type":"default"）, 默认宽度值固定为 282 px
	// 当组件不带边框时（即 "type":"text"）, 组件宽度自适应选择器的内容宽度
	// fill：组件宽度将撑满父容器宽度
	// [100,∞)px：自定义固定数值宽度, 如 200px。最小值为 100px。超出父容器宽度时, 按撑满父容器宽度展示
	Width Width `json:"width,omitempty"`

	// 多选组件的选项是否必选。当组件内嵌在表单容器中时, 该属性可用。其它情况将报错或不生效。可取值：
	//
	// true：多选组件必选。当用户点击表单容器的“提交”时, 未选择多选选项, 则前端提示“有必填项未填写”, 不会向开发者的服务端发起回传请求。
	// false：多选组件可选。当用户点击表单容器的“提交”时, 未选择多选选项, 仍提交表单容器中的数据。
	Required bool `json:"required,omitempty"`

	// 是否禁用该多选组件。可选值：
	//
	// true：禁用该多选组件, 组件展示自定义的占位文本或选项初始值, 且终端用户不可修改交互
	// false：多选组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 多选组件默认选中的选项。数组项的值需要和 options.value 对应。
	SelectedValues []any `json:"selected_values,omitempty"`

	// 选项值配置。按选项数组的顺序展示选项内容。
	Options []*StaticMultiPickerOption `json:"options,omitempty"`
}

type StaticMultiPickerOption struct {
	// 选项名称。为空时展示空白选项。JSON 结构如下所示, 使用 plain text 对象描述：
	Text *ObjectText `json:"text,omitempty"`

	// 添加图标作为选项前缀图标。支持自定义或使用图标库中的图标。
	Icon *ObjectIcon `json:"icon,omitempty"`

	// 自定义选项的回调值, 用于在回调数据中判断用户选择的是哪些选项。当用户提交选项后, value 的值将返回给接收回调数据的服务器, 你可以基于该值做进一步业务处理。
	//
	// 该字段值仅支持 key-value 形式的 JSON 结构, 且 key 为 String 类型。示例值：
	Value map[string]any `json:"value,omitempty"`
}

// MarshalJSON ...
func (r ComponentStaticMultiPicker) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "multi_select_static"
	return json.Marshal(m)
}

// SetType set ComponentStaticMultiPicker.Type attribute
func (r *ComponentStaticMultiPicker) SetType(val BorderType) *ComponentStaticMultiPicker {
	r.Type = val
	return r
}

// SetTypeDefault set ComponentStaticMultiPicker.Type attribute to BorderTypeDefault
func (r *ComponentStaticMultiPicker) SetTypeDefault() *ComponentStaticMultiPicker {
	r.Type = BorderTypeDefault
	return r
}

// SetTypeText set ComponentStaticMultiPicker.Type attribute to BorderTypeText
func (r *ComponentStaticMultiPicker) SetTypeText() *ComponentStaticMultiPicker {
	r.Type = BorderTypeText
	return r
}

// SetName set ComponentStaticMultiPicker.Name attribute
func (r *ComponentStaticMultiPicker) SetName(val string) *ComponentStaticMultiPicker {
	r.Name = val
	return r
}

// SetPlaceholder set ComponentStaticMultiPicker.Placeholder attribute
func (r *ComponentStaticMultiPicker) SetPlaceholder(val *ObjectText) *ComponentStaticMultiPicker {
	r.Placeholder = val
	return r
}

// SetWidth set ComponentStaticMultiPicker.Width attribute
func (r *ComponentStaticMultiPicker) SetWidth(val Width) *ComponentStaticMultiPicker {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentStaticMultiPicker.Width attribute to WidthDefault
func (r *ComponentStaticMultiPicker) SetWidthDefault() *ComponentStaticMultiPicker {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentStaticMultiPicker.Width attribute to WidthFill
func (r *ComponentStaticMultiPicker) SetWidthFill() *ComponentStaticMultiPicker {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentStaticMultiPicker.Width attribute to WidthAuto
func (r *ComponentStaticMultiPicker) SetWidthAuto() *ComponentStaticMultiPicker {
	r.Width = WidthAuto
	return r
}

// SetRequired set ComponentStaticMultiPicker.Required attribute
func (r *ComponentStaticMultiPicker) SetRequired(val bool) *ComponentStaticMultiPicker {
	r.Required = val
	return r
}

// SetDisabled set ComponentStaticMultiPicker.Disabled attribute
func (r *ComponentStaticMultiPicker) SetDisabled(val bool) *ComponentStaticMultiPicker {
	r.Disabled = val
	return r
}

// SetSelectedValues set ComponentStaticMultiPicker.SelectedValues attribute
func (r *ComponentStaticMultiPicker) SetSelectedValues(val ...any) *ComponentStaticMultiPicker {
	r.SelectedValues = val
	return r
}

// SetOptions set ComponentStaticMultiPicker.Options attribute
func (r *ComponentStaticMultiPicker) SetOptions(val ...*StaticMultiPickerOption) *ComponentStaticMultiPicker {
	r.Options = val
	return r
}

// toMap conv ComponentStaticMultiPicker to map
func (r *ComponentStaticMultiPicker) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 8)
	if r.Type != "" {
		res["type"] = r.Type
	}
	if r.Name != "" {
		res["name"] = r.Name
	}
	if r.Placeholder != nil {
		res["placeholder"] = r.Placeholder
	}
	if r.Width != "" {
		res["width"] = r.Width
	}
	if r.Required != false {
		res["required"] = r.Required
	}
	if r.Disabled != false {
		res["disabled"] = r.Disabled
	}
	if len(r.SelectedValues) != 0 {
		res["selected_values"] = r.SelectedValues
	}
	if len(r.Options) != 0 {
		res["options"] = r.Options
	}
	return res
}
