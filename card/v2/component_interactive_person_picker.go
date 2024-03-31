package card

import "encoding/json"

func PersonPicker(ids ...string) *ComponentPersonPicker {
	opts := make([]*PersonPickerOption, 0)
	for _, id := range ids {
		opts = append(opts, &PersonPickerOption{
			Value: id,
		})
	}
	return &ComponentPersonPicker{
		Options: opts,
	}
}

// ComponentPersonPicker 人员选择-单选
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/single-select-user-picker
//
// 人员选择-单选组件支持添加指定人员作为单选选项, 是一种交互组件。本文档介绍人员选择-单选组件的 JSON 结构和相关属性。
//
// 注意事项
// 在卡片 JSON 代码中, 若人员选择-单选组件直接位于卡片根节点, 而非嵌套在其它组件中, 你需将其 JSON 数据配置在交互模块（"tag": "action"）的 actions 字段中使用。
//
// 嵌套规则
// 人员选择-单选组件支持嵌套在分栏、表单容器、折叠面板、循环容器中使用。
//
//go:generate generate_set_attrs -type=ComponentPersonPicker
//go:generate generate_to_map -type=ComponentPersonPicker
type ComponentPersonPicker struct {
	componentBase

	// 组件边框样式。可选值：
	//
	// default：带边框样式
	// text：不带边框的纯文本样式
	Type BorderType `json:"type,omitempty"`

	// 单选组件的内容是否必选。当组件内嵌在表单容器中时, 该属性生效。可取值：
	//
	// true：单选组件必选。当用户点击表单容器的“提交”时, 未填写单选组件, 则前端提示“有必填项未填写”, 不会向开发者的服务端发起回传请求。
	// false：单选组件选填。当用户点击表单容器的“提交”时, 未填写单选组件, 仍提交表单容器中的数据。
	Required bool `json:"required,omitempty"`

	// 是否禁用该单选组件。可选值：
	//
	// true：禁用单选组件组件
	// false：单选组件组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 人员选择组件内的占位文本。
	Placeholder *ObjectText `json:"placeholder,omitempty"`

	// 人员选择组件的宽度。支持以下枚举值：
	//
	// default：默认宽度
	// fill：卡片最大支持宽度
	// [100,∞)px：自定义宽度。超出卡片宽度时将按最大支持宽度展示
	Width Width `json:"width,omitempty"`

	// 选项值配置。按选项数组的顺序展示选项内容。
	Options []*PersonPickerOption `json:"options,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意：confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`
}

type PersonPickerOption struct {
	// 选项配置, 可添加候选用户的 user_id、union_id 或 open_id。了解更多, 参考如何选择使用哪种 ID 和 如何获取不同的用户 ID。
	//
	// 注意：当 options 数组为空, 或 value 的值全部无效时, 候选项展示为卡片所在会话中所有成员选项。
	Value string `json:"value,omitempty"`
}

// MarshalJSON ...
func (r ComponentPersonPicker) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "select_person"
	return json.Marshal(m)
}

// SetType set ComponentPersonPicker.Type attribute
func (r *ComponentPersonPicker) SetType(val BorderType) *ComponentPersonPicker {
	r.Type = val
	return r
}

// SetTypeDefault set ComponentPersonPicker.Type attribute to BorderTypeDefault
func (r *ComponentPersonPicker) SetTypeDefault() *ComponentPersonPicker {
	r.Type = BorderTypeDefault
	return r
}

// SetTypeText set ComponentPersonPicker.Type attribute to BorderTypeText
func (r *ComponentPersonPicker) SetTypeText() *ComponentPersonPicker {
	r.Type = BorderTypeText
	return r
}

// SetRequired set ComponentPersonPicker.Required attribute
func (r *ComponentPersonPicker) SetRequired(val bool) *ComponentPersonPicker {
	r.Required = val
	return r
}

// SetDisabled set ComponentPersonPicker.Disabled attribute
func (r *ComponentPersonPicker) SetDisabled(val bool) *ComponentPersonPicker {
	r.Disabled = val
	return r
}

// SetPlaceholder set ComponentPersonPicker.Placeholder attribute
func (r *ComponentPersonPicker) SetPlaceholder(val *ObjectText) *ComponentPersonPicker {
	r.Placeholder = val
	return r
}

// SetWidth set ComponentPersonPicker.Width attribute
func (r *ComponentPersonPicker) SetWidth(val Width) *ComponentPersonPicker {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentPersonPicker.Width attribute to WidthDefault
func (r *ComponentPersonPicker) SetWidthDefault() *ComponentPersonPicker {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentPersonPicker.Width attribute to WidthFill
func (r *ComponentPersonPicker) SetWidthFill() *ComponentPersonPicker {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentPersonPicker.Width attribute to WidthAuto
func (r *ComponentPersonPicker) SetWidthAuto() *ComponentPersonPicker {
	r.Width = WidthAuto
	return r
}

// SetOptions set ComponentPersonPicker.Options attribute
func (r *ComponentPersonPicker) SetOptions(val ...*PersonPickerOption) *ComponentPersonPicker {
	r.Options = val
	return r
}

// SetConfirm set ComponentPersonPicker.Confirm attribute
func (r *ComponentPersonPicker) SetConfirm(val *ObjectConfirm) *ComponentPersonPicker {
	r.Confirm = val
	return r
}

// toMap conv ComponentPersonPicker to map
func (r *ComponentPersonPicker) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 7)
	if r.Type != "" {
		res["type"] = r.Type
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
	if len(r.Options) != 0 {
		res["options"] = r.Options
	}
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	return res
}
