package card

import "encoding/json"

func Checker(text string) *ComponentChecker {
	return &ComponentChecker{
		Text: Text(text),
	}
}

// ComponentChecker 勾选器
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/checker
//
// 勾选器是一种交互组件, 支持配置回调响应, 主要用于任务勾选的场景。
//
// 注意事项
// 勾选器仅支持通过撰写卡片 JSON 代码的方式使用, 暂不支持在卡片搭建工具上构建使用。
// 勾选器支持飞书 V7.9 及以上版本的客户端。在低于该版本的飞书客户端上, 勾选器的内容将展示为“请升级至最新版本客户端, 以查看内容”的占位图。
//
// 嵌套规则
// 勾选器组件支持内嵌在所有容器类组件（包括表单容器、交互容器、分栏和折叠面板）中使用。
//
// 回调结构
// 为组件成功配置交互后, 用户基于组件进行交互时, 你在开发者后台配置的请求地址将会收到回调数据。
//
// 如果你添加的是新版卡片回传交互回调(card.action.trigger), 可参考卡片回传交互了解回调结构。
// 如果你添加的是旧版卡片回传交互回调(card.action.trigger_v1), 可参考消息卡片回传交互（旧）了解回调结构。
//
//go:generate generate_set_attrs -type=ComponentChecker
//go:generate generate_to_map -type=ComponentChecker
type ComponentChecker struct {
	componentBase

	// 勾选器组件的唯一标识。用于识别用户提交的数据属于哪个组件。
	//
	// 注意: 当勾选器组件嵌套在表单容器中时, 该字段必填且需在卡片全局内唯一。
	Name string `json:"name,omitempty"`

	// 勾选器的初始勾选状态。可选值:
	//
	// true: 已勾选状态
	// false: 未勾选状态
	Checked bool `json:"checked,omitempty"`

	// 勾选器组件内的普通文本信息。
	Text *ObjectText `json:"text,omitempty"`

	// 当光标悬浮在勾选器上时, 勾选器整体是否有阴影效果。
	//
	// 注意: 要取消阴影效果, 你需确保 overall_checkable 为 false 且 pc_display_rule 不为 on_hover。
	OverallCheckable bool `json:"overall_checkable,omitempty"`

	// 按钮区配置。
	ButtonArea *ButtonArea `json:"button_area,omitempty"`

	// 勾选状态样式。
	CheckedStyle *CheckedStyle `json:"checked_style,omitempty"`

	// 组件整体的外边距, 支持填写单值或多值:
	//
	// 单值: 如 "4px", 表示组件的四个外边距都为 4px
	// 多值: 如 "4px 12px 4px 12px", 表示容器内上、右、下、左的内边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔
	Margin Margin `json:"margin,omitempty"`

	// 组件整体的内边距, 支持填写单值或多值:
	//
	// 单值: 如 "4px", 表示组件内四个内边距都为 4px
	// 多值: 如 "4px 12px 4px 12px", 表示容器内上、右、下、左的内边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔
	Padding Padding `json:"padding,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意: confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`

	// 配置交互类型和具体交互行为。未配置 behaviors 时, 终端用户可勾选, 但仅本地有效。详情参考配置卡片交互。
	Behaviors []any `json:"behaviors,omitempty"`

	// 用户在 PC 端将光标悬浮在勾选器上方时的文案提醒。
	//
	// 注意: 当同时配置 hover_tips 和 disabled_tips 时, disabled_tips 将生效。
	HoverTips *ObjectText `json:"hover_tips,omitempty"`

	// 是否禁用该勾选器。可选值:
	//
	// true: 禁用
	// false: 勾选器组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 禁用勾选器后, 用户在 PC 端将光标悬浮在勾选器上方时的文案提醒。
	DisabledTips *ObjectText `json:"disabled_tips,omitempty"`
}

type ButtonArea struct {
	// PC 端勾选器内按钮的展示规则。移动端始终显示按钮。可取值:
	//
	// always: 按钮始终显示。
	// on_hover: 当光标悬浮在勾选器上时, 按钮显示且勾选器整体有阴影效果。
	PcDisplayRule string `json:"pc_display_rule,omitempty"`

	// 在勾选器中添加并配置按钮。最多可配置三个按钮。详情参考下一小节 buttons 字段说明。
	Buttons []*ComponentButton `json:"buttons,omitempty"`
}

type CheckedStyle struct {
	// 是否展示内容区的贯穿式删除线。
	ShowStrikethrough bool `json:"show_strikethrough,omitempty"`

	// 内容区的不透明度。取值范围为 [0,1] 之间的数字, 不限小数位数。
	Opacity float64 `json:"opacity,omitempty"`
}

// MarshalJSON ...
func (r ComponentChecker) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "checker"
	return json.Marshal(m)
}

// SetName set ComponentChecker.Name attribute
func (r *ComponentChecker) SetName(val string) *ComponentChecker {
	r.Name = val
	return r
}

// SetChecked set ComponentChecker.Checked attribute
func (r *ComponentChecker) SetChecked(val bool) *ComponentChecker {
	r.Checked = val
	return r
}

// SetText set ComponentChecker.Text attribute
func (r *ComponentChecker) SetText(val *ObjectText) *ComponentChecker {
	r.Text = val
	return r
}

// SetOverallCheckable set ComponentChecker.OverallCheckable attribute
func (r *ComponentChecker) SetOverallCheckable(val bool) *ComponentChecker {
	r.OverallCheckable = val
	return r
}

// SetButtonArea set ComponentChecker.ButtonArea attribute
func (r *ComponentChecker) SetButtonArea(val *ButtonArea) *ComponentChecker {
	r.ButtonArea = val
	return r
}

// SetCheckedStyle set ComponentChecker.CheckedStyle attribute
func (r *ComponentChecker) SetCheckedStyle(val *CheckedStyle) *ComponentChecker {
	r.CheckedStyle = val
	return r
}

// SetMargin set ComponentChecker.Margin attribute
func (r *ComponentChecker) SetMargin(val Margin) *ComponentChecker {
	r.Margin = val
	return r
}

// SetPadding set ComponentChecker.Padding attribute
func (r *ComponentChecker) SetPadding(val Padding) *ComponentChecker {
	r.Padding = val
	return r
}

// SetConfirm set ComponentChecker.Confirm attribute
func (r *ComponentChecker) SetConfirm(val *ObjectConfirm) *ComponentChecker {
	r.Confirm = val
	return r
}

// SetBehaviors set ComponentChecker.Behaviors attribute
func (r *ComponentChecker) SetBehaviors(val ...any) *ComponentChecker {
	r.Behaviors = val
	return r
}

// SetHoverTips set ComponentChecker.HoverTips attribute
func (r *ComponentChecker) SetHoverTips(val *ObjectText) *ComponentChecker {
	r.HoverTips = val
	return r
}

// SetDisabled set ComponentChecker.Disabled attribute
func (r *ComponentChecker) SetDisabled(val bool) *ComponentChecker {
	r.Disabled = val
	return r
}

// SetDisabledTips set ComponentChecker.DisabledTips attribute
func (r *ComponentChecker) SetDisabledTips(val *ObjectText) *ComponentChecker {
	r.DisabledTips = val
	return r
}

// toMap conv ComponentChecker to map
func (r *ComponentChecker) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 13)
	if r.Name != "" {
		res["name"] = r.Name
	}
	if r.Checked != false {
		res["checked"] = r.Checked
	}
	if r.Text != nil {
		res["text"] = r.Text
	}
	if r.OverallCheckable != false {
		res["overall_checkable"] = r.OverallCheckable
	}
	if r.ButtonArea != nil {
		res["button_area"] = r.ButtonArea
	}
	if r.CheckedStyle != nil {
		res["checked_style"] = r.CheckedStyle
	}
	if r.Margin != "" {
		res["margin"] = r.Margin
	}
	if r.Padding != "" {
		res["padding"] = r.Padding
	}
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	if len(r.Behaviors) != 0 {
		res["behaviors"] = r.Behaviors
	}
	if r.HoverTips != nil {
		res["hover_tips"] = r.HoverTips
	}
	if r.Disabled != false {
		res["disabled"] = r.Disabled
	}
	if r.DisabledTips != nil {
		res["disabled_tips"] = r.DisabledTips
	}
	return res
}
