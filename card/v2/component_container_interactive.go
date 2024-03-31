package card

import "encoding/json"

func Interactive(elements ...Component) *ComponentInteractive {
	return &ComponentInteractive{
		Elements: elements,
	}
}

// ComponentInteractive 交互容器
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/containers/interactive-container
//
// 你可基于业务需求在交互容器中内嵌组件, 并灵活组合多个交互容器, 并统一定义多个交互容器的样式、交互能力等, 实现多种组合效果和丰富的卡片交互。
//
// 注意事项
// 交互容器仅支持通过撰写卡片 JSON 代码的方式使用, 暂不支持在卡片搭建工具上构建使用。
// 交互容器支持飞书 V7.4 及以上版本的客户端。在低于该版本的飞书客户端上, 交互容器的内容将展示为“请升级至最新版本客户端, 以查看内容”的占位图。
// 容器类组件最多支持嵌套五层组件。建议你避免在交互容器中嵌套多层组件。多层嵌套会压缩内容的展示空间, 影响卡片的展示效果。
//
// 嵌套规则
// 交互容器仅支持内嵌普通文本、富文本、图片、备注、分栏、勾选器、交互容器组件。
// 交互容器支持内嵌在卡片根节点、循环容器、分栏、表单容器、交互容器组件中。
//
// 回调结构
// 为组件成功配置交互后, 用户基于组件进行交互时, 你在开发者后台配置的请求地址将会收到回调数据。
//
// 如果你添加的是新版卡片回传交互回调(card.action.trigger), 可参考卡片回传交互了解回调结构。
// 如果你添加的是旧版卡片回传交互回调(card.action.trigger_v1), 可参考消息卡片回传交互（旧）了解回调结构。
//
//go:generate generate_set_attrs -type=ComponentInteractive
//go:generate generate_to_map -type=ComponentInteractive
type ComponentInteractive struct {
	componentBase

	// 交互容器的宽度。可取值:
	//
	// fill: 卡片最大支持宽度
	// auto: 自适应宽度
	// [16,999]px: 自定义宽度, 如 "20px"。最小宽度为 16px
	Width Width `json:"width,omitempty"`

	// 交互容器的高度。可取值:
	//
	// auto: 自适应高度
	// [10,999]px: 自定义高度, 如 "20px"
	Height Height `json:"height,omitempty"`

	// 交互容器的背景色样式。可取值:
	//
	// default: 默认的白底样式, 客户端深色主题下为黑底
	// laser: 镭射渐变彩色样式
	// 卡片支持的颜色枚举值和 RGBA 语法自定义颜色。参考颜色枚举值
	BackgroundStyle Color `json:"background_style,omitempty"`

	// 是否展示边框, 粗细固定为 1px。
	HasBorder bool `json:"has_border,omitempty"`

	// 边框的颜色, 仅 has_border 为 true 时, 此字段生效。
	// 枚举值为卡片支持的颜色枚举值和 RGBA 语法自定义颜色, 参考颜色枚举值 Color
	BorderColor Color `json:"border_color,omitempty"`

	// 交互容器的圆角半径, 单位是像素（px）或百分比（%）。取值遵循以下格式:
	//
	// [0,∞]px, 如 "10px"
	// [0,100]%, 如 "30%"
	CornerRadius CornerRadius `json:"corner_radius,omitempty"`

	// 交互容器的内边距。值的取值范围为 [0,28]px。支持填写单值或多值：
	//
	// 单值：如 "10px", 表示容器内四个内边距都为 10px
	// 多值：如 "4px 12px 4px 12px", 表示容器内上、右、下、左的内边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔
	Padding Padding `json:"padding,omitempty"`

	// 设置点击交互容器时的交互配置。如果交互容器内有交互组件, 则优先响应交互组件定义的交互。详情参考配置卡片交互。
	Behaviors []*ObjectBehavior `json:"behaviors,omitempty"`

	// 用户在 PC 端将光标悬浮在交互容器上方时的文案提醒。默认为空。
	HoverTips *ObjectText `json:"hover_tips,omitempty"`

	// 是否禁用交互容器。可选值：
	//
	// true：禁用整个容器
	// false：容器组件保持可用状态
	Disabled bool `json:"disabled,omitempty"`

	// 禁用交互容器后, 用户触发交互时的弹窗文案提醒。默认为空, 即不弹窗。
	DisabledTips *ObjectText `json:"disabled_tips,omitempty"`

	// 二次确认弹窗配置。指在用户提交时弹出二次确认弹窗提示; 只有用户点击确认后, 才提交输入的内容。该字段默认提供了确认和取消按钮, 你只需要配置弹窗的标题与内容即可。
	//
	// 注意：confirm 字段仅在用户点击包含提交属性的按钮时才会触发二次确认弹窗。
	Confirm *ObjectConfirm `json:"confirm,omitempty"`

	// 交互容器内嵌的组件。仅支持内嵌普通文本、富文本、图片、备注、分栏、勾选器、交互容器组件。
	Elements []Component `json:"elements,omitempty"`
}

// MarshalJSON ...
func (r ComponentInteractive) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "interactive_container"
	return json.Marshal(m)
}

// SetWidth set ComponentInteractive.Width attribute
func (r *ComponentInteractive) SetWidth(val Width) *ComponentInteractive {
	r.Width = val
	return r
}

// SetWidthDefault set ComponentInteractive.Width attribute to WidthDefault
func (r *ComponentInteractive) SetWidthDefault() *ComponentInteractive {
	r.Width = WidthDefault
	return r
}

// SetWidthFill set ComponentInteractive.Width attribute to WidthFill
func (r *ComponentInteractive) SetWidthFill() *ComponentInteractive {
	r.Width = WidthFill
	return r
}

// SetWidthAuto set ComponentInteractive.Width attribute to WidthAuto
func (r *ComponentInteractive) SetWidthAuto() *ComponentInteractive {
	r.Width = WidthAuto
	return r
}

// SetHeight set ComponentInteractive.Height attribute
func (r *ComponentInteractive) SetHeight(val Height) *ComponentInteractive {
	r.Height = val
	return r
}

// SetHeightAuto set ComponentInteractive.Height attribute to HeightAuto
func (r *ComponentInteractive) SetHeightAuto() *ComponentInteractive {
	r.Height = HeightAuto
	return r
}

// SetBackgroundStyle set ComponentInteractive.BackgroundStyle attribute
func (r *ComponentInteractive) SetBackgroundStyle(val Color) *ComponentInteractive {
	r.BackgroundStyle = val
	return r
}

// SetBackgroundStyleDefault set ComponentInteractive.BackgroundStyle attribute to ColorDefault
func (r *ComponentInteractive) SetBackgroundStyleDefault() *ComponentInteractive {
	r.BackgroundStyle = ColorDefault
	return r
}

// SetBackgroundStyleBgWhite set ComponentInteractive.BackgroundStyle attribute to ColorBgWhite
func (r *ComponentInteractive) SetBackgroundStyleBgWhite() *ComponentInteractive {
	r.BackgroundStyle = ColorBgWhite
	return r
}

// SetBackgroundStyleWhite set ComponentInteractive.BackgroundStyle attribute to ColorWhite
func (r *ComponentInteractive) SetBackgroundStyleWhite() *ComponentInteractive {
	r.BackgroundStyle = ColorWhite
	return r
}

// SetBackgroundStyleBlue set ComponentInteractive.BackgroundStyle attribute to ColorBlue
func (r *ComponentInteractive) SetBackgroundStyleBlue() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue
	return r
}

// SetBackgroundStyleBlue50 set ComponentInteractive.BackgroundStyle attribute to ColorBlue50
func (r *ComponentInteractive) SetBackgroundStyleBlue50() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue50
	return r
}

// SetBackgroundStyleBlue100 set ComponentInteractive.BackgroundStyle attribute to ColorBlue100
func (r *ComponentInteractive) SetBackgroundStyleBlue100() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue100
	return r
}

// SetBackgroundStyleBlue200 set ComponentInteractive.BackgroundStyle attribute to ColorBlue200
func (r *ComponentInteractive) SetBackgroundStyleBlue200() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue200
	return r
}

// SetBackgroundStyleBlue300 set ComponentInteractive.BackgroundStyle attribute to ColorBlue300
func (r *ComponentInteractive) SetBackgroundStyleBlue300() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue300
	return r
}

// SetBackgroundStyleBlue350 set ComponentInteractive.BackgroundStyle attribute to ColorBlue350
func (r *ComponentInteractive) SetBackgroundStyleBlue350() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue350
	return r
}

// SetBackgroundStyleBlue400 set ComponentInteractive.BackgroundStyle attribute to ColorBlue400
func (r *ComponentInteractive) SetBackgroundStyleBlue400() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue400
	return r
}

// SetBackgroundStyleBlue500 set ComponentInteractive.BackgroundStyle attribute to ColorBlue500
func (r *ComponentInteractive) SetBackgroundStyleBlue500() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue500
	return r
}

// SetBackgroundStyleBlue600 set ComponentInteractive.BackgroundStyle attribute to ColorBlue600
func (r *ComponentInteractive) SetBackgroundStyleBlue600() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue600
	return r
}

// SetBackgroundStyleBlue700 set ComponentInteractive.BackgroundStyle attribute to ColorBlue700
func (r *ComponentInteractive) SetBackgroundStyleBlue700() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue700
	return r
}

// SetBackgroundStyleBlue800 set ComponentInteractive.BackgroundStyle attribute to ColorBlue800
func (r *ComponentInteractive) SetBackgroundStyleBlue800() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue800
	return r
}

// SetBackgroundStyleBlue900 set ComponentInteractive.BackgroundStyle attribute to ColorBlue900
func (r *ComponentInteractive) SetBackgroundStyleBlue900() *ComponentInteractive {
	r.BackgroundStyle = ColorBlue900
	return r
}

// SetBackgroundStyleCarmine set ComponentInteractive.BackgroundStyle attribute to ColorCarmine
func (r *ComponentInteractive) SetBackgroundStyleCarmine() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine
	return r
}

// SetBackgroundStyleCarmine50 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine50
func (r *ComponentInteractive) SetBackgroundStyleCarmine50() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine50
	return r
}

// SetBackgroundStyleCarmine100 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine100
func (r *ComponentInteractive) SetBackgroundStyleCarmine100() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine100
	return r
}

// SetBackgroundStyleCarmine200 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine200
func (r *ComponentInteractive) SetBackgroundStyleCarmine200() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine200
	return r
}

// SetBackgroundStyleCarmine300 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine300
func (r *ComponentInteractive) SetBackgroundStyleCarmine300() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine300
	return r
}

// SetBackgroundStyleCarmine350 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine350
func (r *ComponentInteractive) SetBackgroundStyleCarmine350() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine350
	return r
}

// SetBackgroundStyleCarmine400 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine400
func (r *ComponentInteractive) SetBackgroundStyleCarmine400() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine400
	return r
}

// SetBackgroundStyleCarmine500 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine500
func (r *ComponentInteractive) SetBackgroundStyleCarmine500() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine500
	return r
}

// SetBackgroundStyleCarmine600 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine600
func (r *ComponentInteractive) SetBackgroundStyleCarmine600() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine600
	return r
}

// SetBackgroundStyleCarmine700 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine700
func (r *ComponentInteractive) SetBackgroundStyleCarmine700() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine700
	return r
}

// SetBackgroundStyleCarmine800 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine800
func (r *ComponentInteractive) SetBackgroundStyleCarmine800() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine800
	return r
}

// SetBackgroundStyleCarmine900 set ComponentInteractive.BackgroundStyle attribute to ColorCarmine900
func (r *ComponentInteractive) SetBackgroundStyleCarmine900() *ComponentInteractive {
	r.BackgroundStyle = ColorCarmine900
	return r
}

// SetBackgroundStyleGreen set ComponentInteractive.BackgroundStyle attribute to ColorGreen
func (r *ComponentInteractive) SetBackgroundStyleGreen() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen
	return r
}

// SetBackgroundStyleGreen50 set ComponentInteractive.BackgroundStyle attribute to ColorGreen50
func (r *ComponentInteractive) SetBackgroundStyleGreen50() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen50
	return r
}

// SetBackgroundStyleGreen100 set ComponentInteractive.BackgroundStyle attribute to ColorGreen100
func (r *ComponentInteractive) SetBackgroundStyleGreen100() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen100
	return r
}

// SetBackgroundStyleGreen200 set ComponentInteractive.BackgroundStyle attribute to ColorGreen200
func (r *ComponentInteractive) SetBackgroundStyleGreen200() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen200
	return r
}

// SetBackgroundStyleGreen300 set ComponentInteractive.BackgroundStyle attribute to ColorGreen300
func (r *ComponentInteractive) SetBackgroundStyleGreen300() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen300
	return r
}

// SetBackgroundStyleGreen350 set ComponentInteractive.BackgroundStyle attribute to ColorGreen350
func (r *ComponentInteractive) SetBackgroundStyleGreen350() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen350
	return r
}

// SetBackgroundStyleGreen400 set ComponentInteractive.BackgroundStyle attribute to ColorGreen400
func (r *ComponentInteractive) SetBackgroundStyleGreen400() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen400
	return r
}

// SetBackgroundStyleGreen500 set ComponentInteractive.BackgroundStyle attribute to ColorGreen500
func (r *ComponentInteractive) SetBackgroundStyleGreen500() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen500
	return r
}

// SetBackgroundStyleGreen600 set ComponentInteractive.BackgroundStyle attribute to ColorGreen600
func (r *ComponentInteractive) SetBackgroundStyleGreen600() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen600
	return r
}

// SetBackgroundStyleGreen700 set ComponentInteractive.BackgroundStyle attribute to ColorGreen700
func (r *ComponentInteractive) SetBackgroundStyleGreen700() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen700
	return r
}

// SetBackgroundStyleGreen800 set ComponentInteractive.BackgroundStyle attribute to ColorGreen800
func (r *ComponentInteractive) SetBackgroundStyleGreen800() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen800
	return r
}

// SetBackgroundStyleGreen900 set ComponentInteractive.BackgroundStyle attribute to ColorGreen900
func (r *ComponentInteractive) SetBackgroundStyleGreen900() *ComponentInteractive {
	r.BackgroundStyle = ColorGreen900
	return r
}

// SetBackgroundStyleIndigo set ComponentInteractive.BackgroundStyle attribute to ColorIndigo
func (r *ComponentInteractive) SetBackgroundStyleIndigo() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo
	return r
}

// SetBackgroundStyleIndigo50 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo50
func (r *ComponentInteractive) SetBackgroundStyleIndigo50() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo50
	return r
}

// SetBackgroundStyleIndigo100 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo100
func (r *ComponentInteractive) SetBackgroundStyleIndigo100() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo100
	return r
}

// SetBackgroundStyleIndigo200 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo200
func (r *ComponentInteractive) SetBackgroundStyleIndigo200() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo200
	return r
}

// SetBackgroundStyleIndigo300 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo300
func (r *ComponentInteractive) SetBackgroundStyleIndigo300() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo300
	return r
}

// SetBackgroundStyleIndigo350 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo350
func (r *ComponentInteractive) SetBackgroundStyleIndigo350() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo350
	return r
}

// SetBackgroundStyleIndigo400 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo400
func (r *ComponentInteractive) SetBackgroundStyleIndigo400() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo400
	return r
}

// SetBackgroundStyleIndigo500 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo500
func (r *ComponentInteractive) SetBackgroundStyleIndigo500() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo500
	return r
}

// SetBackgroundStyleIndigo600 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo600
func (r *ComponentInteractive) SetBackgroundStyleIndigo600() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo600
	return r
}

// SetBackgroundStyleIndigo700 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo700
func (r *ComponentInteractive) SetBackgroundStyleIndigo700() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo700
	return r
}

// SetBackgroundStyleIndigo800 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo800
func (r *ComponentInteractive) SetBackgroundStyleIndigo800() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo800
	return r
}

// SetBackgroundStyleIndigo900 set ComponentInteractive.BackgroundStyle attribute to ColorIndigo900
func (r *ComponentInteractive) SetBackgroundStyleIndigo900() *ComponentInteractive {
	r.BackgroundStyle = ColorIndigo900
	return r
}

// SetBackgroundStyleLime set ComponentInteractive.BackgroundStyle attribute to ColorLime
func (r *ComponentInteractive) SetBackgroundStyleLime() *ComponentInteractive {
	r.BackgroundStyle = ColorLime
	return r
}

// SetBackgroundStyleLime50 set ComponentInteractive.BackgroundStyle attribute to ColorLime50
func (r *ComponentInteractive) SetBackgroundStyleLime50() *ComponentInteractive {
	r.BackgroundStyle = ColorLime50
	return r
}

// SetBackgroundStyleLime100 set ComponentInteractive.BackgroundStyle attribute to ColorLime100
func (r *ComponentInteractive) SetBackgroundStyleLime100() *ComponentInteractive {
	r.BackgroundStyle = ColorLime100
	return r
}

// SetBackgroundStyleLime200 set ComponentInteractive.BackgroundStyle attribute to ColorLime200
func (r *ComponentInteractive) SetBackgroundStyleLime200() *ComponentInteractive {
	r.BackgroundStyle = ColorLime200
	return r
}

// SetBackgroundStyleLime300 set ComponentInteractive.BackgroundStyle attribute to ColorLime300
func (r *ComponentInteractive) SetBackgroundStyleLime300() *ComponentInteractive {
	r.BackgroundStyle = ColorLime300
	return r
}

// SetBackgroundStyleLime350 set ComponentInteractive.BackgroundStyle attribute to ColorLime350
func (r *ComponentInteractive) SetBackgroundStyleLime350() *ComponentInteractive {
	r.BackgroundStyle = ColorLime350
	return r
}

// SetBackgroundStyleLime400 set ComponentInteractive.BackgroundStyle attribute to ColorLime400
func (r *ComponentInteractive) SetBackgroundStyleLime400() *ComponentInteractive {
	r.BackgroundStyle = ColorLime400
	return r
}

// SetBackgroundStyleLime500 set ComponentInteractive.BackgroundStyle attribute to ColorLime500
func (r *ComponentInteractive) SetBackgroundStyleLime500() *ComponentInteractive {
	r.BackgroundStyle = ColorLime500
	return r
}

// SetBackgroundStyleLime600 set ComponentInteractive.BackgroundStyle attribute to ColorLime600
func (r *ComponentInteractive) SetBackgroundStyleLime600() *ComponentInteractive {
	r.BackgroundStyle = ColorLime600
	return r
}

// SetBackgroundStyleLime700 set ComponentInteractive.BackgroundStyle attribute to ColorLime700
func (r *ComponentInteractive) SetBackgroundStyleLime700() *ComponentInteractive {
	r.BackgroundStyle = ColorLime700
	return r
}

// SetBackgroundStyleLime800 set ComponentInteractive.BackgroundStyle attribute to ColorLime800
func (r *ComponentInteractive) SetBackgroundStyleLime800() *ComponentInteractive {
	r.BackgroundStyle = ColorLime800
	return r
}

// SetBackgroundStyleLime900 set ComponentInteractive.BackgroundStyle attribute to ColorLime900
func (r *ComponentInteractive) SetBackgroundStyleLime900() *ComponentInteractive {
	r.BackgroundStyle = ColorLime900
	return r
}

// SetBackgroundStyleGrey set ComponentInteractive.BackgroundStyle attribute to ColorGrey
func (r *ComponentInteractive) SetBackgroundStyleGrey() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey
	return r
}

// SetBackgroundStyleGrey00 set ComponentInteractive.BackgroundStyle attribute to ColorGrey00
func (r *ComponentInteractive) SetBackgroundStyleGrey00() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey00
	return r
}

// SetBackgroundStyleGrey50 set ComponentInteractive.BackgroundStyle attribute to ColorGrey50
func (r *ComponentInteractive) SetBackgroundStyleGrey50() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey50
	return r
}

// SetBackgroundStyleGrey100 set ComponentInteractive.BackgroundStyle attribute to ColorGrey100
func (r *ComponentInteractive) SetBackgroundStyleGrey100() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey100
	return r
}

// SetBackgroundStyleGrey200 set ComponentInteractive.BackgroundStyle attribute to ColorGrey200
func (r *ComponentInteractive) SetBackgroundStyleGrey200() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey200
	return r
}

// SetBackgroundStyleGrey300 set ComponentInteractive.BackgroundStyle attribute to ColorGrey300
func (r *ComponentInteractive) SetBackgroundStyleGrey300() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey300
	return r
}

// SetBackgroundStyleGrey350 set ComponentInteractive.BackgroundStyle attribute to ColorGrey350
func (r *ComponentInteractive) SetBackgroundStyleGrey350() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey350
	return r
}

// SetBackgroundStyleGrey400 set ComponentInteractive.BackgroundStyle attribute to ColorGrey400
func (r *ComponentInteractive) SetBackgroundStyleGrey400() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey400
	return r
}

// SetBackgroundStyleGrey500 set ComponentInteractive.BackgroundStyle attribute to ColorGrey500
func (r *ComponentInteractive) SetBackgroundStyleGrey500() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey500
	return r
}

// SetBackgroundStyleGrey600 set ComponentInteractive.BackgroundStyle attribute to ColorGrey600
func (r *ComponentInteractive) SetBackgroundStyleGrey600() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey600
	return r
}

// SetBackgroundStyleGrey650 set ComponentInteractive.BackgroundStyle attribute to ColorGrey650
func (r *ComponentInteractive) SetBackgroundStyleGrey650() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey650
	return r
}

// SetBackgroundStyleGrey700 set ComponentInteractive.BackgroundStyle attribute to ColorGrey700
func (r *ComponentInteractive) SetBackgroundStyleGrey700() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey700
	return r
}

// SetBackgroundStyleGrey800 set ComponentInteractive.BackgroundStyle attribute to ColorGrey800
func (r *ComponentInteractive) SetBackgroundStyleGrey800() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey800
	return r
}

// SetBackgroundStyleGrey900 set ComponentInteractive.BackgroundStyle attribute to ColorGrey900
func (r *ComponentInteractive) SetBackgroundStyleGrey900() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey900
	return r
}

// SetBackgroundStyleGrey950 set ComponentInteractive.BackgroundStyle attribute to ColorGrey950
func (r *ComponentInteractive) SetBackgroundStyleGrey950() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey950
	return r
}

// SetBackgroundStyleGrey1000 set ComponentInteractive.BackgroundStyle attribute to ColorGrey1000
func (r *ComponentInteractive) SetBackgroundStyleGrey1000() *ComponentInteractive {
	r.BackgroundStyle = ColorGrey1000
	return r
}

// SetBackgroundStyleOrange set ComponentInteractive.BackgroundStyle attribute to ColorOrange
func (r *ComponentInteractive) SetBackgroundStyleOrange() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange
	return r
}

// SetBackgroundStyleOrange50 set ComponentInteractive.BackgroundStyle attribute to ColorOrange50
func (r *ComponentInteractive) SetBackgroundStyleOrange50() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange50
	return r
}

// SetBackgroundStyleOrange100 set ComponentInteractive.BackgroundStyle attribute to ColorOrange100
func (r *ComponentInteractive) SetBackgroundStyleOrange100() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange100
	return r
}

// SetBackgroundStyleOrange200 set ComponentInteractive.BackgroundStyle attribute to ColorOrange200
func (r *ComponentInteractive) SetBackgroundStyleOrange200() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange200
	return r
}

// SetBackgroundStyleOrange300 set ComponentInteractive.BackgroundStyle attribute to ColorOrange300
func (r *ComponentInteractive) SetBackgroundStyleOrange300() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange300
	return r
}

// SetBackgroundStyleOrange350 set ComponentInteractive.BackgroundStyle attribute to ColorOrange350
func (r *ComponentInteractive) SetBackgroundStyleOrange350() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange350
	return r
}

// SetBackgroundStyleOrange400 set ComponentInteractive.BackgroundStyle attribute to ColorOrange400
func (r *ComponentInteractive) SetBackgroundStyleOrange400() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange400
	return r
}

// SetBackgroundStyleOrange500 set ComponentInteractive.BackgroundStyle attribute to ColorOrange500
func (r *ComponentInteractive) SetBackgroundStyleOrange500() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange500
	return r
}

// SetBackgroundStyleOrange600 set ComponentInteractive.BackgroundStyle attribute to ColorOrange600
func (r *ComponentInteractive) SetBackgroundStyleOrange600() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange600
	return r
}

// SetBackgroundStyleOrange700 set ComponentInteractive.BackgroundStyle attribute to ColorOrange700
func (r *ComponentInteractive) SetBackgroundStyleOrange700() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange700
	return r
}

// SetBackgroundStyleOrange800 set ComponentInteractive.BackgroundStyle attribute to ColorOrange800
func (r *ComponentInteractive) SetBackgroundStyleOrange800() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange800
	return r
}

// SetBackgroundStyleOrange900 set ComponentInteractive.BackgroundStyle attribute to ColorOrange900
func (r *ComponentInteractive) SetBackgroundStyleOrange900() *ComponentInteractive {
	r.BackgroundStyle = ColorOrange900
	return r
}

// SetBackgroundStylePurple set ComponentInteractive.BackgroundStyle attribute to ColorPurple
func (r *ComponentInteractive) SetBackgroundStylePurple() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple
	return r
}

// SetBackgroundStylePurple50 set ComponentInteractive.BackgroundStyle attribute to ColorPurple50
func (r *ComponentInteractive) SetBackgroundStylePurple50() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple50
	return r
}

// SetBackgroundStylePurple100 set ComponentInteractive.BackgroundStyle attribute to ColorPurple100
func (r *ComponentInteractive) SetBackgroundStylePurple100() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple100
	return r
}

// SetBackgroundStylePurple200 set ComponentInteractive.BackgroundStyle attribute to ColorPurple200
func (r *ComponentInteractive) SetBackgroundStylePurple200() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple200
	return r
}

// SetBackgroundStylePurple300 set ComponentInteractive.BackgroundStyle attribute to ColorPurple300
func (r *ComponentInteractive) SetBackgroundStylePurple300() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple300
	return r
}

// SetBackgroundStylePurple350 set ComponentInteractive.BackgroundStyle attribute to ColorPurple350
func (r *ComponentInteractive) SetBackgroundStylePurple350() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple350
	return r
}

// SetBackgroundStylePurple400 set ComponentInteractive.BackgroundStyle attribute to ColorPurple400
func (r *ComponentInteractive) SetBackgroundStylePurple400() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple400
	return r
}

// SetBackgroundStylePurple500 set ComponentInteractive.BackgroundStyle attribute to ColorPurple500
func (r *ComponentInteractive) SetBackgroundStylePurple500() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple500
	return r
}

// SetBackgroundStylePurple600 set ComponentInteractive.BackgroundStyle attribute to ColorPurple600
func (r *ComponentInteractive) SetBackgroundStylePurple600() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple600
	return r
}

// SetBackgroundStylePurple700 set ComponentInteractive.BackgroundStyle attribute to ColorPurple700
func (r *ComponentInteractive) SetBackgroundStylePurple700() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple700
	return r
}

// SetBackgroundStylePurple800 set ComponentInteractive.BackgroundStyle attribute to ColorPurple800
func (r *ComponentInteractive) SetBackgroundStylePurple800() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple800
	return r
}

// SetBackgroundStylePurple900 set ComponentInteractive.BackgroundStyle attribute to ColorPurple900
func (r *ComponentInteractive) SetBackgroundStylePurple900() *ComponentInteractive {
	r.BackgroundStyle = ColorPurple900
	return r
}

// SetBackgroundStyleRed set ComponentInteractive.BackgroundStyle attribute to ColorRed
func (r *ComponentInteractive) SetBackgroundStyleRed() *ComponentInteractive {
	r.BackgroundStyle = ColorRed
	return r
}

// SetBackgroundStyleRed50 set ComponentInteractive.BackgroundStyle attribute to ColorRed50
func (r *ComponentInteractive) SetBackgroundStyleRed50() *ComponentInteractive {
	r.BackgroundStyle = ColorRed50
	return r
}

// SetBackgroundStyleRed100 set ComponentInteractive.BackgroundStyle attribute to ColorRed100
func (r *ComponentInteractive) SetBackgroundStyleRed100() *ComponentInteractive {
	r.BackgroundStyle = ColorRed100
	return r
}

// SetBackgroundStyleRed200 set ComponentInteractive.BackgroundStyle attribute to ColorRed200
func (r *ComponentInteractive) SetBackgroundStyleRed200() *ComponentInteractive {
	r.BackgroundStyle = ColorRed200
	return r
}

// SetBackgroundStyleRed300 set ComponentInteractive.BackgroundStyle attribute to ColorRed300
func (r *ComponentInteractive) SetBackgroundStyleRed300() *ComponentInteractive {
	r.BackgroundStyle = ColorRed300
	return r
}

// SetBackgroundStyleRed350 set ComponentInteractive.BackgroundStyle attribute to ColorRed350
func (r *ComponentInteractive) SetBackgroundStyleRed350() *ComponentInteractive {
	r.BackgroundStyle = ColorRed350
	return r
}

// SetBackgroundStyleRed400 set ComponentInteractive.BackgroundStyle attribute to ColorRed400
func (r *ComponentInteractive) SetBackgroundStyleRed400() *ComponentInteractive {
	r.BackgroundStyle = ColorRed400
	return r
}

// SetBackgroundStyleRed500 set ComponentInteractive.BackgroundStyle attribute to ColorRed500
func (r *ComponentInteractive) SetBackgroundStyleRed500() *ComponentInteractive {
	r.BackgroundStyle = ColorRed500
	return r
}

// SetBackgroundStyleRed600 set ComponentInteractive.BackgroundStyle attribute to ColorRed600
func (r *ComponentInteractive) SetBackgroundStyleRed600() *ComponentInteractive {
	r.BackgroundStyle = ColorRed600
	return r
}

// SetBackgroundStyleRed700 set ComponentInteractive.BackgroundStyle attribute to ColorRed700
func (r *ComponentInteractive) SetBackgroundStyleRed700() *ComponentInteractive {
	r.BackgroundStyle = ColorRed700
	return r
}

// SetBackgroundStyleRed800 set ComponentInteractive.BackgroundStyle attribute to ColorRed800
func (r *ComponentInteractive) SetBackgroundStyleRed800() *ComponentInteractive {
	r.BackgroundStyle = ColorRed800
	return r
}

// SetBackgroundStyleRed900 set ComponentInteractive.BackgroundStyle attribute to ColorRed900
func (r *ComponentInteractive) SetBackgroundStyleRed900() *ComponentInteractive {
	r.BackgroundStyle = ColorRed900
	return r
}

// SetBackgroundStyleSunflower set ComponentInteractive.BackgroundStyle attribute to ColorSunflower
func (r *ComponentInteractive) SetBackgroundStyleSunflower() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower
	return r
}

// SetBackgroundStyleSunflower50 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower50
func (r *ComponentInteractive) SetBackgroundStyleSunflower50() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower50
	return r
}

// SetBackgroundStyleSunflower100 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower100
func (r *ComponentInteractive) SetBackgroundStyleSunflower100() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower100
	return r
}

// SetBackgroundStyleSunflower200 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower200
func (r *ComponentInteractive) SetBackgroundStyleSunflower200() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower200
	return r
}

// SetBackgroundStyleSunflower300 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower300
func (r *ComponentInteractive) SetBackgroundStyleSunflower300() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower300
	return r
}

// SetBackgroundStyleSunflower350 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower350
func (r *ComponentInteractive) SetBackgroundStyleSunflower350() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower350
	return r
}

// SetBackgroundStyleSunflower400 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower400
func (r *ComponentInteractive) SetBackgroundStyleSunflower400() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower400
	return r
}

// SetBackgroundStyleSunflower500 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower500
func (r *ComponentInteractive) SetBackgroundStyleSunflower500() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower500
	return r
}

// SetBackgroundStyleSunflower600 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower600
func (r *ComponentInteractive) SetBackgroundStyleSunflower600() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower600
	return r
}

// SetBackgroundStyleSunflower700 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower700
func (r *ComponentInteractive) SetBackgroundStyleSunflower700() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower700
	return r
}

// SetBackgroundStyleSunflower800 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower800
func (r *ComponentInteractive) SetBackgroundStyleSunflower800() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower800
	return r
}

// SetBackgroundStyleSunflower900 set ComponentInteractive.BackgroundStyle attribute to ColorSunflower900
func (r *ComponentInteractive) SetBackgroundStyleSunflower900() *ComponentInteractive {
	r.BackgroundStyle = ColorSunflower900
	return r
}

// SetBackgroundStyleTurquoise set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise
func (r *ComponentInteractive) SetBackgroundStyleTurquoise() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise
	return r
}

// SetBackgroundStyleTurquoise50 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise50
func (r *ComponentInteractive) SetBackgroundStyleTurquoise50() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise50
	return r
}

// SetBackgroundStyleTurquoise100 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise100
func (r *ComponentInteractive) SetBackgroundStyleTurquoise100() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise100
	return r
}

// SetBackgroundStyleTurquoise200 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise200
func (r *ComponentInteractive) SetBackgroundStyleTurquoise200() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise200
	return r
}

// SetBackgroundStyleTurquoise300 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise300
func (r *ComponentInteractive) SetBackgroundStyleTurquoise300() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise300
	return r
}

// SetBackgroundStyleTurquoise350 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise350
func (r *ComponentInteractive) SetBackgroundStyleTurquoise350() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise350
	return r
}

// SetBackgroundStyleTurquoise400 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise400
func (r *ComponentInteractive) SetBackgroundStyleTurquoise400() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise400
	return r
}

// SetBackgroundStyleTurquoise500 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise500
func (r *ComponentInteractive) SetBackgroundStyleTurquoise500() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise500
	return r
}

// SetBackgroundStyleTurquoise600 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise600
func (r *ComponentInteractive) SetBackgroundStyleTurquoise600() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise600
	return r
}

// SetBackgroundStyleTurquoise700 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise700
func (r *ComponentInteractive) SetBackgroundStyleTurquoise700() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise700
	return r
}

// SetBackgroundStyleTurquoise800 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise800
func (r *ComponentInteractive) SetBackgroundStyleTurquoise800() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise800
	return r
}

// SetBackgroundStyleTurquoise900 set ComponentInteractive.BackgroundStyle attribute to ColorTurquoise900
func (r *ComponentInteractive) SetBackgroundStyleTurquoise900() *ComponentInteractive {
	r.BackgroundStyle = ColorTurquoise900
	return r
}

// SetBackgroundStyleViolet set ComponentInteractive.BackgroundStyle attribute to ColorViolet
func (r *ComponentInteractive) SetBackgroundStyleViolet() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet
	return r
}

// SetBackgroundStyleViolet50 set ComponentInteractive.BackgroundStyle attribute to ColorViolet50
func (r *ComponentInteractive) SetBackgroundStyleViolet50() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet50
	return r
}

// SetBackgroundStyleViolet100 set ComponentInteractive.BackgroundStyle attribute to ColorViolet100
func (r *ComponentInteractive) SetBackgroundStyleViolet100() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet100
	return r
}

// SetBackgroundStyleViolet200 set ComponentInteractive.BackgroundStyle attribute to ColorViolet200
func (r *ComponentInteractive) SetBackgroundStyleViolet200() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet200
	return r
}

// SetBackgroundStyleViolet300 set ComponentInteractive.BackgroundStyle attribute to ColorViolet300
func (r *ComponentInteractive) SetBackgroundStyleViolet300() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet300
	return r
}

// SetBackgroundStyleViolet350 set ComponentInteractive.BackgroundStyle attribute to ColorViolet350
func (r *ComponentInteractive) SetBackgroundStyleViolet350() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet350
	return r
}

// SetBackgroundStyleViolet400 set ComponentInteractive.BackgroundStyle attribute to ColorViolet400
func (r *ComponentInteractive) SetBackgroundStyleViolet400() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet400
	return r
}

// SetBackgroundStyleViolet500 set ComponentInteractive.BackgroundStyle attribute to ColorViolet500
func (r *ComponentInteractive) SetBackgroundStyleViolet500() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet500
	return r
}

// SetBackgroundStyleViolet600 set ComponentInteractive.BackgroundStyle attribute to ColorViolet600
func (r *ComponentInteractive) SetBackgroundStyleViolet600() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet600
	return r
}

// SetBackgroundStyleViolet700 set ComponentInteractive.BackgroundStyle attribute to ColorViolet700
func (r *ComponentInteractive) SetBackgroundStyleViolet700() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet700
	return r
}

// SetBackgroundStyleViolet800 set ComponentInteractive.BackgroundStyle attribute to ColorViolet800
func (r *ComponentInteractive) SetBackgroundStyleViolet800() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet800
	return r
}

// SetBackgroundStyleViolet900 set ComponentInteractive.BackgroundStyle attribute to ColorViolet900
func (r *ComponentInteractive) SetBackgroundStyleViolet900() *ComponentInteractive {
	r.BackgroundStyle = ColorViolet900
	return r
}

// SetBackgroundStyleWathet set ComponentInteractive.BackgroundStyle attribute to ColorWathet
func (r *ComponentInteractive) SetBackgroundStyleWathet() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet
	return r
}

// SetBackgroundStyleWathet50 set ComponentInteractive.BackgroundStyle attribute to ColorWathet50
func (r *ComponentInteractive) SetBackgroundStyleWathet50() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet50
	return r
}

// SetBackgroundStyleWathet100 set ComponentInteractive.BackgroundStyle attribute to ColorWathet100
func (r *ComponentInteractive) SetBackgroundStyleWathet100() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet100
	return r
}

// SetBackgroundStyleWathet200 set ComponentInteractive.BackgroundStyle attribute to ColorWathet200
func (r *ComponentInteractive) SetBackgroundStyleWathet200() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet200
	return r
}

// SetBackgroundStyleWathet300 set ComponentInteractive.BackgroundStyle attribute to ColorWathet300
func (r *ComponentInteractive) SetBackgroundStyleWathet300() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet300
	return r
}

// SetBackgroundStyleWathet350 set ComponentInteractive.BackgroundStyle attribute to ColorWathet350
func (r *ComponentInteractive) SetBackgroundStyleWathet350() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet350
	return r
}

// SetBackgroundStyleWathet400 set ComponentInteractive.BackgroundStyle attribute to ColorWathet400
func (r *ComponentInteractive) SetBackgroundStyleWathet400() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet400
	return r
}

// SetBackgroundStyleWathet500 set ComponentInteractive.BackgroundStyle attribute to ColorWathet500
func (r *ComponentInteractive) SetBackgroundStyleWathet500() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet500
	return r
}

// SetBackgroundStyleWathet600 set ComponentInteractive.BackgroundStyle attribute to ColorWathet600
func (r *ComponentInteractive) SetBackgroundStyleWathet600() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet600
	return r
}

// SetBackgroundStyleWathet700 set ComponentInteractive.BackgroundStyle attribute to ColorWathet700
func (r *ComponentInteractive) SetBackgroundStyleWathet700() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet700
	return r
}

// SetBackgroundStyleWathet800 set ComponentInteractive.BackgroundStyle attribute to ColorWathet800
func (r *ComponentInteractive) SetBackgroundStyleWathet800() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet800
	return r
}

// SetBackgroundStyleWathet900 set ComponentInteractive.BackgroundStyle attribute to ColorWathet900
func (r *ComponentInteractive) SetBackgroundStyleWathet900() *ComponentInteractive {
	r.BackgroundStyle = ColorWathet900
	return r
}

// SetBackgroundStyleYellow set ComponentInteractive.BackgroundStyle attribute to ColorYellow
func (r *ComponentInteractive) SetBackgroundStyleYellow() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow
	return r
}

// SetBackgroundStyleYellow50 set ComponentInteractive.BackgroundStyle attribute to ColorYellow50
func (r *ComponentInteractive) SetBackgroundStyleYellow50() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow50
	return r
}

// SetBackgroundStyleYellow100 set ComponentInteractive.BackgroundStyle attribute to ColorYellow100
func (r *ComponentInteractive) SetBackgroundStyleYellow100() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow100
	return r
}

// SetBackgroundStyleYellow200 set ComponentInteractive.BackgroundStyle attribute to ColorYellow200
func (r *ComponentInteractive) SetBackgroundStyleYellow200() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow200
	return r
}

// SetBackgroundStyleYellow300 set ComponentInteractive.BackgroundStyle attribute to ColorYellow300
func (r *ComponentInteractive) SetBackgroundStyleYellow300() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow300
	return r
}

// SetBackgroundStyleYellow350 set ComponentInteractive.BackgroundStyle attribute to ColorYellow350
func (r *ComponentInteractive) SetBackgroundStyleYellow350() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow350
	return r
}

// SetBackgroundStyleYellow400 set ComponentInteractive.BackgroundStyle attribute to ColorYellow400
func (r *ComponentInteractive) SetBackgroundStyleYellow400() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow400
	return r
}

// SetBackgroundStyleYellow500 set ComponentInteractive.BackgroundStyle attribute to ColorYellow500
func (r *ComponentInteractive) SetBackgroundStyleYellow500() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow500
	return r
}

// SetBackgroundStyleYellow600 set ComponentInteractive.BackgroundStyle attribute to ColorYellow600
func (r *ComponentInteractive) SetBackgroundStyleYellow600() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow600
	return r
}

// SetBackgroundStyleYellow700 set ComponentInteractive.BackgroundStyle attribute to ColorYellow700
func (r *ComponentInteractive) SetBackgroundStyleYellow700() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow700
	return r
}

// SetBackgroundStyleYellow800 set ComponentInteractive.BackgroundStyle attribute to ColorYellow800
func (r *ComponentInteractive) SetBackgroundStyleYellow800() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow800
	return r
}

// SetBackgroundStyleYellow900 set ComponentInteractive.BackgroundStyle attribute to ColorYellow900
func (r *ComponentInteractive) SetBackgroundStyleYellow900() *ComponentInteractive {
	r.BackgroundStyle = ColorYellow900
	return r
}

// SetHasBorder set ComponentInteractive.HasBorder attribute
func (r *ComponentInteractive) SetHasBorder(val bool) *ComponentInteractive {
	r.HasBorder = val
	return r
}

// SetBorderColor set ComponentInteractive.BorderColor attribute
func (r *ComponentInteractive) SetBorderColor(val Color) *ComponentInteractive {
	r.BorderColor = val
	return r
}

// SetBorderColorDefault set ComponentInteractive.BorderColor attribute to ColorDefault
func (r *ComponentInteractive) SetBorderColorDefault() *ComponentInteractive {
	r.BorderColor = ColorDefault
	return r
}

// SetBorderColorBgWhite set ComponentInteractive.BorderColor attribute to ColorBgWhite
func (r *ComponentInteractive) SetBorderColorBgWhite() *ComponentInteractive {
	r.BorderColor = ColorBgWhite
	return r
}

// SetBorderColorWhite set ComponentInteractive.BorderColor attribute to ColorWhite
func (r *ComponentInteractive) SetBorderColorWhite() *ComponentInteractive {
	r.BorderColor = ColorWhite
	return r
}

// SetBorderColorBlue set ComponentInteractive.BorderColor attribute to ColorBlue
func (r *ComponentInteractive) SetBorderColorBlue() *ComponentInteractive {
	r.BorderColor = ColorBlue
	return r
}

// SetBorderColorBlue50 set ComponentInteractive.BorderColor attribute to ColorBlue50
func (r *ComponentInteractive) SetBorderColorBlue50() *ComponentInteractive {
	r.BorderColor = ColorBlue50
	return r
}

// SetBorderColorBlue100 set ComponentInteractive.BorderColor attribute to ColorBlue100
func (r *ComponentInteractive) SetBorderColorBlue100() *ComponentInteractive {
	r.BorderColor = ColorBlue100
	return r
}

// SetBorderColorBlue200 set ComponentInteractive.BorderColor attribute to ColorBlue200
func (r *ComponentInteractive) SetBorderColorBlue200() *ComponentInteractive {
	r.BorderColor = ColorBlue200
	return r
}

// SetBorderColorBlue300 set ComponentInteractive.BorderColor attribute to ColorBlue300
func (r *ComponentInteractive) SetBorderColorBlue300() *ComponentInteractive {
	r.BorderColor = ColorBlue300
	return r
}

// SetBorderColorBlue350 set ComponentInteractive.BorderColor attribute to ColorBlue350
func (r *ComponentInteractive) SetBorderColorBlue350() *ComponentInteractive {
	r.BorderColor = ColorBlue350
	return r
}

// SetBorderColorBlue400 set ComponentInteractive.BorderColor attribute to ColorBlue400
func (r *ComponentInteractive) SetBorderColorBlue400() *ComponentInteractive {
	r.BorderColor = ColorBlue400
	return r
}

// SetBorderColorBlue500 set ComponentInteractive.BorderColor attribute to ColorBlue500
func (r *ComponentInteractive) SetBorderColorBlue500() *ComponentInteractive {
	r.BorderColor = ColorBlue500
	return r
}

// SetBorderColorBlue600 set ComponentInteractive.BorderColor attribute to ColorBlue600
func (r *ComponentInteractive) SetBorderColorBlue600() *ComponentInteractive {
	r.BorderColor = ColorBlue600
	return r
}

// SetBorderColorBlue700 set ComponentInteractive.BorderColor attribute to ColorBlue700
func (r *ComponentInteractive) SetBorderColorBlue700() *ComponentInteractive {
	r.BorderColor = ColorBlue700
	return r
}

// SetBorderColorBlue800 set ComponentInteractive.BorderColor attribute to ColorBlue800
func (r *ComponentInteractive) SetBorderColorBlue800() *ComponentInteractive {
	r.BorderColor = ColorBlue800
	return r
}

// SetBorderColorBlue900 set ComponentInteractive.BorderColor attribute to ColorBlue900
func (r *ComponentInteractive) SetBorderColorBlue900() *ComponentInteractive {
	r.BorderColor = ColorBlue900
	return r
}

// SetBorderColorCarmine set ComponentInteractive.BorderColor attribute to ColorCarmine
func (r *ComponentInteractive) SetBorderColorCarmine() *ComponentInteractive {
	r.BorderColor = ColorCarmine
	return r
}

// SetBorderColorCarmine50 set ComponentInteractive.BorderColor attribute to ColorCarmine50
func (r *ComponentInteractive) SetBorderColorCarmine50() *ComponentInteractive {
	r.BorderColor = ColorCarmine50
	return r
}

// SetBorderColorCarmine100 set ComponentInteractive.BorderColor attribute to ColorCarmine100
func (r *ComponentInteractive) SetBorderColorCarmine100() *ComponentInteractive {
	r.BorderColor = ColorCarmine100
	return r
}

// SetBorderColorCarmine200 set ComponentInteractive.BorderColor attribute to ColorCarmine200
func (r *ComponentInteractive) SetBorderColorCarmine200() *ComponentInteractive {
	r.BorderColor = ColorCarmine200
	return r
}

// SetBorderColorCarmine300 set ComponentInteractive.BorderColor attribute to ColorCarmine300
func (r *ComponentInteractive) SetBorderColorCarmine300() *ComponentInteractive {
	r.BorderColor = ColorCarmine300
	return r
}

// SetBorderColorCarmine350 set ComponentInteractive.BorderColor attribute to ColorCarmine350
func (r *ComponentInteractive) SetBorderColorCarmine350() *ComponentInteractive {
	r.BorderColor = ColorCarmine350
	return r
}

// SetBorderColorCarmine400 set ComponentInteractive.BorderColor attribute to ColorCarmine400
func (r *ComponentInteractive) SetBorderColorCarmine400() *ComponentInteractive {
	r.BorderColor = ColorCarmine400
	return r
}

// SetBorderColorCarmine500 set ComponentInteractive.BorderColor attribute to ColorCarmine500
func (r *ComponentInteractive) SetBorderColorCarmine500() *ComponentInteractive {
	r.BorderColor = ColorCarmine500
	return r
}

// SetBorderColorCarmine600 set ComponentInteractive.BorderColor attribute to ColorCarmine600
func (r *ComponentInteractive) SetBorderColorCarmine600() *ComponentInteractive {
	r.BorderColor = ColorCarmine600
	return r
}

// SetBorderColorCarmine700 set ComponentInteractive.BorderColor attribute to ColorCarmine700
func (r *ComponentInteractive) SetBorderColorCarmine700() *ComponentInteractive {
	r.BorderColor = ColorCarmine700
	return r
}

// SetBorderColorCarmine800 set ComponentInteractive.BorderColor attribute to ColorCarmine800
func (r *ComponentInteractive) SetBorderColorCarmine800() *ComponentInteractive {
	r.BorderColor = ColorCarmine800
	return r
}

// SetBorderColorCarmine900 set ComponentInteractive.BorderColor attribute to ColorCarmine900
func (r *ComponentInteractive) SetBorderColorCarmine900() *ComponentInteractive {
	r.BorderColor = ColorCarmine900
	return r
}

// SetBorderColorGreen set ComponentInteractive.BorderColor attribute to ColorGreen
func (r *ComponentInteractive) SetBorderColorGreen() *ComponentInteractive {
	r.BorderColor = ColorGreen
	return r
}

// SetBorderColorGreen50 set ComponentInteractive.BorderColor attribute to ColorGreen50
func (r *ComponentInteractive) SetBorderColorGreen50() *ComponentInteractive {
	r.BorderColor = ColorGreen50
	return r
}

// SetBorderColorGreen100 set ComponentInteractive.BorderColor attribute to ColorGreen100
func (r *ComponentInteractive) SetBorderColorGreen100() *ComponentInteractive {
	r.BorderColor = ColorGreen100
	return r
}

// SetBorderColorGreen200 set ComponentInteractive.BorderColor attribute to ColorGreen200
func (r *ComponentInteractive) SetBorderColorGreen200() *ComponentInteractive {
	r.BorderColor = ColorGreen200
	return r
}

// SetBorderColorGreen300 set ComponentInteractive.BorderColor attribute to ColorGreen300
func (r *ComponentInteractive) SetBorderColorGreen300() *ComponentInteractive {
	r.BorderColor = ColorGreen300
	return r
}

// SetBorderColorGreen350 set ComponentInteractive.BorderColor attribute to ColorGreen350
func (r *ComponentInteractive) SetBorderColorGreen350() *ComponentInteractive {
	r.BorderColor = ColorGreen350
	return r
}

// SetBorderColorGreen400 set ComponentInteractive.BorderColor attribute to ColorGreen400
func (r *ComponentInteractive) SetBorderColorGreen400() *ComponentInteractive {
	r.BorderColor = ColorGreen400
	return r
}

// SetBorderColorGreen500 set ComponentInteractive.BorderColor attribute to ColorGreen500
func (r *ComponentInteractive) SetBorderColorGreen500() *ComponentInteractive {
	r.BorderColor = ColorGreen500
	return r
}

// SetBorderColorGreen600 set ComponentInteractive.BorderColor attribute to ColorGreen600
func (r *ComponentInteractive) SetBorderColorGreen600() *ComponentInteractive {
	r.BorderColor = ColorGreen600
	return r
}

// SetBorderColorGreen700 set ComponentInteractive.BorderColor attribute to ColorGreen700
func (r *ComponentInteractive) SetBorderColorGreen700() *ComponentInteractive {
	r.BorderColor = ColorGreen700
	return r
}

// SetBorderColorGreen800 set ComponentInteractive.BorderColor attribute to ColorGreen800
func (r *ComponentInteractive) SetBorderColorGreen800() *ComponentInteractive {
	r.BorderColor = ColorGreen800
	return r
}

// SetBorderColorGreen900 set ComponentInteractive.BorderColor attribute to ColorGreen900
func (r *ComponentInteractive) SetBorderColorGreen900() *ComponentInteractive {
	r.BorderColor = ColorGreen900
	return r
}

// SetBorderColorIndigo set ComponentInteractive.BorderColor attribute to ColorIndigo
func (r *ComponentInteractive) SetBorderColorIndigo() *ComponentInteractive {
	r.BorderColor = ColorIndigo
	return r
}

// SetBorderColorIndigo50 set ComponentInteractive.BorderColor attribute to ColorIndigo50
func (r *ComponentInteractive) SetBorderColorIndigo50() *ComponentInteractive {
	r.BorderColor = ColorIndigo50
	return r
}

// SetBorderColorIndigo100 set ComponentInteractive.BorderColor attribute to ColorIndigo100
func (r *ComponentInteractive) SetBorderColorIndigo100() *ComponentInteractive {
	r.BorderColor = ColorIndigo100
	return r
}

// SetBorderColorIndigo200 set ComponentInteractive.BorderColor attribute to ColorIndigo200
func (r *ComponentInteractive) SetBorderColorIndigo200() *ComponentInteractive {
	r.BorderColor = ColorIndigo200
	return r
}

// SetBorderColorIndigo300 set ComponentInteractive.BorderColor attribute to ColorIndigo300
func (r *ComponentInteractive) SetBorderColorIndigo300() *ComponentInteractive {
	r.BorderColor = ColorIndigo300
	return r
}

// SetBorderColorIndigo350 set ComponentInteractive.BorderColor attribute to ColorIndigo350
func (r *ComponentInteractive) SetBorderColorIndigo350() *ComponentInteractive {
	r.BorderColor = ColorIndigo350
	return r
}

// SetBorderColorIndigo400 set ComponentInteractive.BorderColor attribute to ColorIndigo400
func (r *ComponentInteractive) SetBorderColorIndigo400() *ComponentInteractive {
	r.BorderColor = ColorIndigo400
	return r
}

// SetBorderColorIndigo500 set ComponentInteractive.BorderColor attribute to ColorIndigo500
func (r *ComponentInteractive) SetBorderColorIndigo500() *ComponentInteractive {
	r.BorderColor = ColorIndigo500
	return r
}

// SetBorderColorIndigo600 set ComponentInteractive.BorderColor attribute to ColorIndigo600
func (r *ComponentInteractive) SetBorderColorIndigo600() *ComponentInteractive {
	r.BorderColor = ColorIndigo600
	return r
}

// SetBorderColorIndigo700 set ComponentInteractive.BorderColor attribute to ColorIndigo700
func (r *ComponentInteractive) SetBorderColorIndigo700() *ComponentInteractive {
	r.BorderColor = ColorIndigo700
	return r
}

// SetBorderColorIndigo800 set ComponentInteractive.BorderColor attribute to ColorIndigo800
func (r *ComponentInteractive) SetBorderColorIndigo800() *ComponentInteractive {
	r.BorderColor = ColorIndigo800
	return r
}

// SetBorderColorIndigo900 set ComponentInteractive.BorderColor attribute to ColorIndigo900
func (r *ComponentInteractive) SetBorderColorIndigo900() *ComponentInteractive {
	r.BorderColor = ColorIndigo900
	return r
}

// SetBorderColorLime set ComponentInteractive.BorderColor attribute to ColorLime
func (r *ComponentInteractive) SetBorderColorLime() *ComponentInteractive {
	r.BorderColor = ColorLime
	return r
}

// SetBorderColorLime50 set ComponentInteractive.BorderColor attribute to ColorLime50
func (r *ComponentInteractive) SetBorderColorLime50() *ComponentInteractive {
	r.BorderColor = ColorLime50
	return r
}

// SetBorderColorLime100 set ComponentInteractive.BorderColor attribute to ColorLime100
func (r *ComponentInteractive) SetBorderColorLime100() *ComponentInteractive {
	r.BorderColor = ColorLime100
	return r
}

// SetBorderColorLime200 set ComponentInteractive.BorderColor attribute to ColorLime200
func (r *ComponentInteractive) SetBorderColorLime200() *ComponentInteractive {
	r.BorderColor = ColorLime200
	return r
}

// SetBorderColorLime300 set ComponentInteractive.BorderColor attribute to ColorLime300
func (r *ComponentInteractive) SetBorderColorLime300() *ComponentInteractive {
	r.BorderColor = ColorLime300
	return r
}

// SetBorderColorLime350 set ComponentInteractive.BorderColor attribute to ColorLime350
func (r *ComponentInteractive) SetBorderColorLime350() *ComponentInteractive {
	r.BorderColor = ColorLime350
	return r
}

// SetBorderColorLime400 set ComponentInteractive.BorderColor attribute to ColorLime400
func (r *ComponentInteractive) SetBorderColorLime400() *ComponentInteractive {
	r.BorderColor = ColorLime400
	return r
}

// SetBorderColorLime500 set ComponentInteractive.BorderColor attribute to ColorLime500
func (r *ComponentInteractive) SetBorderColorLime500() *ComponentInteractive {
	r.BorderColor = ColorLime500
	return r
}

// SetBorderColorLime600 set ComponentInteractive.BorderColor attribute to ColorLime600
func (r *ComponentInteractive) SetBorderColorLime600() *ComponentInteractive {
	r.BorderColor = ColorLime600
	return r
}

// SetBorderColorLime700 set ComponentInteractive.BorderColor attribute to ColorLime700
func (r *ComponentInteractive) SetBorderColorLime700() *ComponentInteractive {
	r.BorderColor = ColorLime700
	return r
}

// SetBorderColorLime800 set ComponentInteractive.BorderColor attribute to ColorLime800
func (r *ComponentInteractive) SetBorderColorLime800() *ComponentInteractive {
	r.BorderColor = ColorLime800
	return r
}

// SetBorderColorLime900 set ComponentInteractive.BorderColor attribute to ColorLime900
func (r *ComponentInteractive) SetBorderColorLime900() *ComponentInteractive {
	r.BorderColor = ColorLime900
	return r
}

// SetBorderColorGrey set ComponentInteractive.BorderColor attribute to ColorGrey
func (r *ComponentInteractive) SetBorderColorGrey() *ComponentInteractive {
	r.BorderColor = ColorGrey
	return r
}

// SetBorderColorGrey00 set ComponentInteractive.BorderColor attribute to ColorGrey00
func (r *ComponentInteractive) SetBorderColorGrey00() *ComponentInteractive {
	r.BorderColor = ColorGrey00
	return r
}

// SetBorderColorGrey50 set ComponentInteractive.BorderColor attribute to ColorGrey50
func (r *ComponentInteractive) SetBorderColorGrey50() *ComponentInteractive {
	r.BorderColor = ColorGrey50
	return r
}

// SetBorderColorGrey100 set ComponentInteractive.BorderColor attribute to ColorGrey100
func (r *ComponentInteractive) SetBorderColorGrey100() *ComponentInteractive {
	r.BorderColor = ColorGrey100
	return r
}

// SetBorderColorGrey200 set ComponentInteractive.BorderColor attribute to ColorGrey200
func (r *ComponentInteractive) SetBorderColorGrey200() *ComponentInteractive {
	r.BorderColor = ColorGrey200
	return r
}

// SetBorderColorGrey300 set ComponentInteractive.BorderColor attribute to ColorGrey300
func (r *ComponentInteractive) SetBorderColorGrey300() *ComponentInteractive {
	r.BorderColor = ColorGrey300
	return r
}

// SetBorderColorGrey350 set ComponentInteractive.BorderColor attribute to ColorGrey350
func (r *ComponentInteractive) SetBorderColorGrey350() *ComponentInteractive {
	r.BorderColor = ColorGrey350
	return r
}

// SetBorderColorGrey400 set ComponentInteractive.BorderColor attribute to ColorGrey400
func (r *ComponentInteractive) SetBorderColorGrey400() *ComponentInteractive {
	r.BorderColor = ColorGrey400
	return r
}

// SetBorderColorGrey500 set ComponentInteractive.BorderColor attribute to ColorGrey500
func (r *ComponentInteractive) SetBorderColorGrey500() *ComponentInteractive {
	r.BorderColor = ColorGrey500
	return r
}

// SetBorderColorGrey600 set ComponentInteractive.BorderColor attribute to ColorGrey600
func (r *ComponentInteractive) SetBorderColorGrey600() *ComponentInteractive {
	r.BorderColor = ColorGrey600
	return r
}

// SetBorderColorGrey650 set ComponentInteractive.BorderColor attribute to ColorGrey650
func (r *ComponentInteractive) SetBorderColorGrey650() *ComponentInteractive {
	r.BorderColor = ColorGrey650
	return r
}

// SetBorderColorGrey700 set ComponentInteractive.BorderColor attribute to ColorGrey700
func (r *ComponentInteractive) SetBorderColorGrey700() *ComponentInteractive {
	r.BorderColor = ColorGrey700
	return r
}

// SetBorderColorGrey800 set ComponentInteractive.BorderColor attribute to ColorGrey800
func (r *ComponentInteractive) SetBorderColorGrey800() *ComponentInteractive {
	r.BorderColor = ColorGrey800
	return r
}

// SetBorderColorGrey900 set ComponentInteractive.BorderColor attribute to ColorGrey900
func (r *ComponentInteractive) SetBorderColorGrey900() *ComponentInteractive {
	r.BorderColor = ColorGrey900
	return r
}

// SetBorderColorGrey950 set ComponentInteractive.BorderColor attribute to ColorGrey950
func (r *ComponentInteractive) SetBorderColorGrey950() *ComponentInteractive {
	r.BorderColor = ColorGrey950
	return r
}

// SetBorderColorGrey1000 set ComponentInteractive.BorderColor attribute to ColorGrey1000
func (r *ComponentInteractive) SetBorderColorGrey1000() *ComponentInteractive {
	r.BorderColor = ColorGrey1000
	return r
}

// SetBorderColorOrange set ComponentInteractive.BorderColor attribute to ColorOrange
func (r *ComponentInteractive) SetBorderColorOrange() *ComponentInteractive {
	r.BorderColor = ColorOrange
	return r
}

// SetBorderColorOrange50 set ComponentInteractive.BorderColor attribute to ColorOrange50
func (r *ComponentInteractive) SetBorderColorOrange50() *ComponentInteractive {
	r.BorderColor = ColorOrange50
	return r
}

// SetBorderColorOrange100 set ComponentInteractive.BorderColor attribute to ColorOrange100
func (r *ComponentInteractive) SetBorderColorOrange100() *ComponentInteractive {
	r.BorderColor = ColorOrange100
	return r
}

// SetBorderColorOrange200 set ComponentInteractive.BorderColor attribute to ColorOrange200
func (r *ComponentInteractive) SetBorderColorOrange200() *ComponentInteractive {
	r.BorderColor = ColorOrange200
	return r
}

// SetBorderColorOrange300 set ComponentInteractive.BorderColor attribute to ColorOrange300
func (r *ComponentInteractive) SetBorderColorOrange300() *ComponentInteractive {
	r.BorderColor = ColorOrange300
	return r
}

// SetBorderColorOrange350 set ComponentInteractive.BorderColor attribute to ColorOrange350
func (r *ComponentInteractive) SetBorderColorOrange350() *ComponentInteractive {
	r.BorderColor = ColorOrange350
	return r
}

// SetBorderColorOrange400 set ComponentInteractive.BorderColor attribute to ColorOrange400
func (r *ComponentInteractive) SetBorderColorOrange400() *ComponentInteractive {
	r.BorderColor = ColorOrange400
	return r
}

// SetBorderColorOrange500 set ComponentInteractive.BorderColor attribute to ColorOrange500
func (r *ComponentInteractive) SetBorderColorOrange500() *ComponentInteractive {
	r.BorderColor = ColorOrange500
	return r
}

// SetBorderColorOrange600 set ComponentInteractive.BorderColor attribute to ColorOrange600
func (r *ComponentInteractive) SetBorderColorOrange600() *ComponentInteractive {
	r.BorderColor = ColorOrange600
	return r
}

// SetBorderColorOrange700 set ComponentInteractive.BorderColor attribute to ColorOrange700
func (r *ComponentInteractive) SetBorderColorOrange700() *ComponentInteractive {
	r.BorderColor = ColorOrange700
	return r
}

// SetBorderColorOrange800 set ComponentInteractive.BorderColor attribute to ColorOrange800
func (r *ComponentInteractive) SetBorderColorOrange800() *ComponentInteractive {
	r.BorderColor = ColorOrange800
	return r
}

// SetBorderColorOrange900 set ComponentInteractive.BorderColor attribute to ColorOrange900
func (r *ComponentInteractive) SetBorderColorOrange900() *ComponentInteractive {
	r.BorderColor = ColorOrange900
	return r
}

// SetBorderColorPurple set ComponentInteractive.BorderColor attribute to ColorPurple
func (r *ComponentInteractive) SetBorderColorPurple() *ComponentInteractive {
	r.BorderColor = ColorPurple
	return r
}

// SetBorderColorPurple50 set ComponentInteractive.BorderColor attribute to ColorPurple50
func (r *ComponentInteractive) SetBorderColorPurple50() *ComponentInteractive {
	r.BorderColor = ColorPurple50
	return r
}

// SetBorderColorPurple100 set ComponentInteractive.BorderColor attribute to ColorPurple100
func (r *ComponentInteractive) SetBorderColorPurple100() *ComponentInteractive {
	r.BorderColor = ColorPurple100
	return r
}

// SetBorderColorPurple200 set ComponentInteractive.BorderColor attribute to ColorPurple200
func (r *ComponentInteractive) SetBorderColorPurple200() *ComponentInteractive {
	r.BorderColor = ColorPurple200
	return r
}

// SetBorderColorPurple300 set ComponentInteractive.BorderColor attribute to ColorPurple300
func (r *ComponentInteractive) SetBorderColorPurple300() *ComponentInteractive {
	r.BorderColor = ColorPurple300
	return r
}

// SetBorderColorPurple350 set ComponentInteractive.BorderColor attribute to ColorPurple350
func (r *ComponentInteractive) SetBorderColorPurple350() *ComponentInteractive {
	r.BorderColor = ColorPurple350
	return r
}

// SetBorderColorPurple400 set ComponentInteractive.BorderColor attribute to ColorPurple400
func (r *ComponentInteractive) SetBorderColorPurple400() *ComponentInteractive {
	r.BorderColor = ColorPurple400
	return r
}

// SetBorderColorPurple500 set ComponentInteractive.BorderColor attribute to ColorPurple500
func (r *ComponentInteractive) SetBorderColorPurple500() *ComponentInteractive {
	r.BorderColor = ColorPurple500
	return r
}

// SetBorderColorPurple600 set ComponentInteractive.BorderColor attribute to ColorPurple600
func (r *ComponentInteractive) SetBorderColorPurple600() *ComponentInteractive {
	r.BorderColor = ColorPurple600
	return r
}

// SetBorderColorPurple700 set ComponentInteractive.BorderColor attribute to ColorPurple700
func (r *ComponentInteractive) SetBorderColorPurple700() *ComponentInteractive {
	r.BorderColor = ColorPurple700
	return r
}

// SetBorderColorPurple800 set ComponentInteractive.BorderColor attribute to ColorPurple800
func (r *ComponentInteractive) SetBorderColorPurple800() *ComponentInteractive {
	r.BorderColor = ColorPurple800
	return r
}

// SetBorderColorPurple900 set ComponentInteractive.BorderColor attribute to ColorPurple900
func (r *ComponentInteractive) SetBorderColorPurple900() *ComponentInteractive {
	r.BorderColor = ColorPurple900
	return r
}

// SetBorderColorRed set ComponentInteractive.BorderColor attribute to ColorRed
func (r *ComponentInteractive) SetBorderColorRed() *ComponentInteractive {
	r.BorderColor = ColorRed
	return r
}

// SetBorderColorRed50 set ComponentInteractive.BorderColor attribute to ColorRed50
func (r *ComponentInteractive) SetBorderColorRed50() *ComponentInteractive {
	r.BorderColor = ColorRed50
	return r
}

// SetBorderColorRed100 set ComponentInteractive.BorderColor attribute to ColorRed100
func (r *ComponentInteractive) SetBorderColorRed100() *ComponentInteractive {
	r.BorderColor = ColorRed100
	return r
}

// SetBorderColorRed200 set ComponentInteractive.BorderColor attribute to ColorRed200
func (r *ComponentInteractive) SetBorderColorRed200() *ComponentInteractive {
	r.BorderColor = ColorRed200
	return r
}

// SetBorderColorRed300 set ComponentInteractive.BorderColor attribute to ColorRed300
func (r *ComponentInteractive) SetBorderColorRed300() *ComponentInteractive {
	r.BorderColor = ColorRed300
	return r
}

// SetBorderColorRed350 set ComponentInteractive.BorderColor attribute to ColorRed350
func (r *ComponentInteractive) SetBorderColorRed350() *ComponentInteractive {
	r.BorderColor = ColorRed350
	return r
}

// SetBorderColorRed400 set ComponentInteractive.BorderColor attribute to ColorRed400
func (r *ComponentInteractive) SetBorderColorRed400() *ComponentInteractive {
	r.BorderColor = ColorRed400
	return r
}

// SetBorderColorRed500 set ComponentInteractive.BorderColor attribute to ColorRed500
func (r *ComponentInteractive) SetBorderColorRed500() *ComponentInteractive {
	r.BorderColor = ColorRed500
	return r
}

// SetBorderColorRed600 set ComponentInteractive.BorderColor attribute to ColorRed600
func (r *ComponentInteractive) SetBorderColorRed600() *ComponentInteractive {
	r.BorderColor = ColorRed600
	return r
}

// SetBorderColorRed700 set ComponentInteractive.BorderColor attribute to ColorRed700
func (r *ComponentInteractive) SetBorderColorRed700() *ComponentInteractive {
	r.BorderColor = ColorRed700
	return r
}

// SetBorderColorRed800 set ComponentInteractive.BorderColor attribute to ColorRed800
func (r *ComponentInteractive) SetBorderColorRed800() *ComponentInteractive {
	r.BorderColor = ColorRed800
	return r
}

// SetBorderColorRed900 set ComponentInteractive.BorderColor attribute to ColorRed900
func (r *ComponentInteractive) SetBorderColorRed900() *ComponentInteractive {
	r.BorderColor = ColorRed900
	return r
}

// SetBorderColorSunflower set ComponentInteractive.BorderColor attribute to ColorSunflower
func (r *ComponentInteractive) SetBorderColorSunflower() *ComponentInteractive {
	r.BorderColor = ColorSunflower
	return r
}

// SetBorderColorSunflower50 set ComponentInteractive.BorderColor attribute to ColorSunflower50
func (r *ComponentInteractive) SetBorderColorSunflower50() *ComponentInteractive {
	r.BorderColor = ColorSunflower50
	return r
}

// SetBorderColorSunflower100 set ComponentInteractive.BorderColor attribute to ColorSunflower100
func (r *ComponentInteractive) SetBorderColorSunflower100() *ComponentInteractive {
	r.BorderColor = ColorSunflower100
	return r
}

// SetBorderColorSunflower200 set ComponentInteractive.BorderColor attribute to ColorSunflower200
func (r *ComponentInteractive) SetBorderColorSunflower200() *ComponentInteractive {
	r.BorderColor = ColorSunflower200
	return r
}

// SetBorderColorSunflower300 set ComponentInteractive.BorderColor attribute to ColorSunflower300
func (r *ComponentInteractive) SetBorderColorSunflower300() *ComponentInteractive {
	r.BorderColor = ColorSunflower300
	return r
}

// SetBorderColorSunflower350 set ComponentInteractive.BorderColor attribute to ColorSunflower350
func (r *ComponentInteractive) SetBorderColorSunflower350() *ComponentInteractive {
	r.BorderColor = ColorSunflower350
	return r
}

// SetBorderColorSunflower400 set ComponentInteractive.BorderColor attribute to ColorSunflower400
func (r *ComponentInteractive) SetBorderColorSunflower400() *ComponentInteractive {
	r.BorderColor = ColorSunflower400
	return r
}

// SetBorderColorSunflower500 set ComponentInteractive.BorderColor attribute to ColorSunflower500
func (r *ComponentInteractive) SetBorderColorSunflower500() *ComponentInteractive {
	r.BorderColor = ColorSunflower500
	return r
}

// SetBorderColorSunflower600 set ComponentInteractive.BorderColor attribute to ColorSunflower600
func (r *ComponentInteractive) SetBorderColorSunflower600() *ComponentInteractive {
	r.BorderColor = ColorSunflower600
	return r
}

// SetBorderColorSunflower700 set ComponentInteractive.BorderColor attribute to ColorSunflower700
func (r *ComponentInteractive) SetBorderColorSunflower700() *ComponentInteractive {
	r.BorderColor = ColorSunflower700
	return r
}

// SetBorderColorSunflower800 set ComponentInteractive.BorderColor attribute to ColorSunflower800
func (r *ComponentInteractive) SetBorderColorSunflower800() *ComponentInteractive {
	r.BorderColor = ColorSunflower800
	return r
}

// SetBorderColorSunflower900 set ComponentInteractive.BorderColor attribute to ColorSunflower900
func (r *ComponentInteractive) SetBorderColorSunflower900() *ComponentInteractive {
	r.BorderColor = ColorSunflower900
	return r
}

// SetBorderColorTurquoise set ComponentInteractive.BorderColor attribute to ColorTurquoise
func (r *ComponentInteractive) SetBorderColorTurquoise() *ComponentInteractive {
	r.BorderColor = ColorTurquoise
	return r
}

// SetBorderColorTurquoise50 set ComponentInteractive.BorderColor attribute to ColorTurquoise50
func (r *ComponentInteractive) SetBorderColorTurquoise50() *ComponentInteractive {
	r.BorderColor = ColorTurquoise50
	return r
}

// SetBorderColorTurquoise100 set ComponentInteractive.BorderColor attribute to ColorTurquoise100
func (r *ComponentInteractive) SetBorderColorTurquoise100() *ComponentInteractive {
	r.BorderColor = ColorTurquoise100
	return r
}

// SetBorderColorTurquoise200 set ComponentInteractive.BorderColor attribute to ColorTurquoise200
func (r *ComponentInteractive) SetBorderColorTurquoise200() *ComponentInteractive {
	r.BorderColor = ColorTurquoise200
	return r
}

// SetBorderColorTurquoise300 set ComponentInteractive.BorderColor attribute to ColorTurquoise300
func (r *ComponentInteractive) SetBorderColorTurquoise300() *ComponentInteractive {
	r.BorderColor = ColorTurquoise300
	return r
}

// SetBorderColorTurquoise350 set ComponentInteractive.BorderColor attribute to ColorTurquoise350
func (r *ComponentInteractive) SetBorderColorTurquoise350() *ComponentInteractive {
	r.BorderColor = ColorTurquoise350
	return r
}

// SetBorderColorTurquoise400 set ComponentInteractive.BorderColor attribute to ColorTurquoise400
func (r *ComponentInteractive) SetBorderColorTurquoise400() *ComponentInteractive {
	r.BorderColor = ColorTurquoise400
	return r
}

// SetBorderColorTurquoise500 set ComponentInteractive.BorderColor attribute to ColorTurquoise500
func (r *ComponentInteractive) SetBorderColorTurquoise500() *ComponentInteractive {
	r.BorderColor = ColorTurquoise500
	return r
}

// SetBorderColorTurquoise600 set ComponentInteractive.BorderColor attribute to ColorTurquoise600
func (r *ComponentInteractive) SetBorderColorTurquoise600() *ComponentInteractive {
	r.BorderColor = ColorTurquoise600
	return r
}

// SetBorderColorTurquoise700 set ComponentInteractive.BorderColor attribute to ColorTurquoise700
func (r *ComponentInteractive) SetBorderColorTurquoise700() *ComponentInteractive {
	r.BorderColor = ColorTurquoise700
	return r
}

// SetBorderColorTurquoise800 set ComponentInteractive.BorderColor attribute to ColorTurquoise800
func (r *ComponentInteractive) SetBorderColorTurquoise800() *ComponentInteractive {
	r.BorderColor = ColorTurquoise800
	return r
}

// SetBorderColorTurquoise900 set ComponentInteractive.BorderColor attribute to ColorTurquoise900
func (r *ComponentInteractive) SetBorderColorTurquoise900() *ComponentInteractive {
	r.BorderColor = ColorTurquoise900
	return r
}

// SetBorderColorViolet set ComponentInteractive.BorderColor attribute to ColorViolet
func (r *ComponentInteractive) SetBorderColorViolet() *ComponentInteractive {
	r.BorderColor = ColorViolet
	return r
}

// SetBorderColorViolet50 set ComponentInteractive.BorderColor attribute to ColorViolet50
func (r *ComponentInteractive) SetBorderColorViolet50() *ComponentInteractive {
	r.BorderColor = ColorViolet50
	return r
}

// SetBorderColorViolet100 set ComponentInteractive.BorderColor attribute to ColorViolet100
func (r *ComponentInteractive) SetBorderColorViolet100() *ComponentInteractive {
	r.BorderColor = ColorViolet100
	return r
}

// SetBorderColorViolet200 set ComponentInteractive.BorderColor attribute to ColorViolet200
func (r *ComponentInteractive) SetBorderColorViolet200() *ComponentInteractive {
	r.BorderColor = ColorViolet200
	return r
}

// SetBorderColorViolet300 set ComponentInteractive.BorderColor attribute to ColorViolet300
func (r *ComponentInteractive) SetBorderColorViolet300() *ComponentInteractive {
	r.BorderColor = ColorViolet300
	return r
}

// SetBorderColorViolet350 set ComponentInteractive.BorderColor attribute to ColorViolet350
func (r *ComponentInteractive) SetBorderColorViolet350() *ComponentInteractive {
	r.BorderColor = ColorViolet350
	return r
}

// SetBorderColorViolet400 set ComponentInteractive.BorderColor attribute to ColorViolet400
func (r *ComponentInteractive) SetBorderColorViolet400() *ComponentInteractive {
	r.BorderColor = ColorViolet400
	return r
}

// SetBorderColorViolet500 set ComponentInteractive.BorderColor attribute to ColorViolet500
func (r *ComponentInteractive) SetBorderColorViolet500() *ComponentInteractive {
	r.BorderColor = ColorViolet500
	return r
}

// SetBorderColorViolet600 set ComponentInteractive.BorderColor attribute to ColorViolet600
func (r *ComponentInteractive) SetBorderColorViolet600() *ComponentInteractive {
	r.BorderColor = ColorViolet600
	return r
}

// SetBorderColorViolet700 set ComponentInteractive.BorderColor attribute to ColorViolet700
func (r *ComponentInteractive) SetBorderColorViolet700() *ComponentInteractive {
	r.BorderColor = ColorViolet700
	return r
}

// SetBorderColorViolet800 set ComponentInteractive.BorderColor attribute to ColorViolet800
func (r *ComponentInteractive) SetBorderColorViolet800() *ComponentInteractive {
	r.BorderColor = ColorViolet800
	return r
}

// SetBorderColorViolet900 set ComponentInteractive.BorderColor attribute to ColorViolet900
func (r *ComponentInteractive) SetBorderColorViolet900() *ComponentInteractive {
	r.BorderColor = ColorViolet900
	return r
}

// SetBorderColorWathet set ComponentInteractive.BorderColor attribute to ColorWathet
func (r *ComponentInteractive) SetBorderColorWathet() *ComponentInteractive {
	r.BorderColor = ColorWathet
	return r
}

// SetBorderColorWathet50 set ComponentInteractive.BorderColor attribute to ColorWathet50
func (r *ComponentInteractive) SetBorderColorWathet50() *ComponentInteractive {
	r.BorderColor = ColorWathet50
	return r
}

// SetBorderColorWathet100 set ComponentInteractive.BorderColor attribute to ColorWathet100
func (r *ComponentInteractive) SetBorderColorWathet100() *ComponentInteractive {
	r.BorderColor = ColorWathet100
	return r
}

// SetBorderColorWathet200 set ComponentInteractive.BorderColor attribute to ColorWathet200
func (r *ComponentInteractive) SetBorderColorWathet200() *ComponentInteractive {
	r.BorderColor = ColorWathet200
	return r
}

// SetBorderColorWathet300 set ComponentInteractive.BorderColor attribute to ColorWathet300
func (r *ComponentInteractive) SetBorderColorWathet300() *ComponentInteractive {
	r.BorderColor = ColorWathet300
	return r
}

// SetBorderColorWathet350 set ComponentInteractive.BorderColor attribute to ColorWathet350
func (r *ComponentInteractive) SetBorderColorWathet350() *ComponentInteractive {
	r.BorderColor = ColorWathet350
	return r
}

// SetBorderColorWathet400 set ComponentInteractive.BorderColor attribute to ColorWathet400
func (r *ComponentInteractive) SetBorderColorWathet400() *ComponentInteractive {
	r.BorderColor = ColorWathet400
	return r
}

// SetBorderColorWathet500 set ComponentInteractive.BorderColor attribute to ColorWathet500
func (r *ComponentInteractive) SetBorderColorWathet500() *ComponentInteractive {
	r.BorderColor = ColorWathet500
	return r
}

// SetBorderColorWathet600 set ComponentInteractive.BorderColor attribute to ColorWathet600
func (r *ComponentInteractive) SetBorderColorWathet600() *ComponentInteractive {
	r.BorderColor = ColorWathet600
	return r
}

// SetBorderColorWathet700 set ComponentInteractive.BorderColor attribute to ColorWathet700
func (r *ComponentInteractive) SetBorderColorWathet700() *ComponentInteractive {
	r.BorderColor = ColorWathet700
	return r
}

// SetBorderColorWathet800 set ComponentInteractive.BorderColor attribute to ColorWathet800
func (r *ComponentInteractive) SetBorderColorWathet800() *ComponentInteractive {
	r.BorderColor = ColorWathet800
	return r
}

// SetBorderColorWathet900 set ComponentInteractive.BorderColor attribute to ColorWathet900
func (r *ComponentInteractive) SetBorderColorWathet900() *ComponentInteractive {
	r.BorderColor = ColorWathet900
	return r
}

// SetBorderColorYellow set ComponentInteractive.BorderColor attribute to ColorYellow
func (r *ComponentInteractive) SetBorderColorYellow() *ComponentInteractive {
	r.BorderColor = ColorYellow
	return r
}

// SetBorderColorYellow50 set ComponentInteractive.BorderColor attribute to ColorYellow50
func (r *ComponentInteractive) SetBorderColorYellow50() *ComponentInteractive {
	r.BorderColor = ColorYellow50
	return r
}

// SetBorderColorYellow100 set ComponentInteractive.BorderColor attribute to ColorYellow100
func (r *ComponentInteractive) SetBorderColorYellow100() *ComponentInteractive {
	r.BorderColor = ColorYellow100
	return r
}

// SetBorderColorYellow200 set ComponentInteractive.BorderColor attribute to ColorYellow200
func (r *ComponentInteractive) SetBorderColorYellow200() *ComponentInteractive {
	r.BorderColor = ColorYellow200
	return r
}

// SetBorderColorYellow300 set ComponentInteractive.BorderColor attribute to ColorYellow300
func (r *ComponentInteractive) SetBorderColorYellow300() *ComponentInteractive {
	r.BorderColor = ColorYellow300
	return r
}

// SetBorderColorYellow350 set ComponentInteractive.BorderColor attribute to ColorYellow350
func (r *ComponentInteractive) SetBorderColorYellow350() *ComponentInteractive {
	r.BorderColor = ColorYellow350
	return r
}

// SetBorderColorYellow400 set ComponentInteractive.BorderColor attribute to ColorYellow400
func (r *ComponentInteractive) SetBorderColorYellow400() *ComponentInteractive {
	r.BorderColor = ColorYellow400
	return r
}

// SetBorderColorYellow500 set ComponentInteractive.BorderColor attribute to ColorYellow500
func (r *ComponentInteractive) SetBorderColorYellow500() *ComponentInteractive {
	r.BorderColor = ColorYellow500
	return r
}

// SetBorderColorYellow600 set ComponentInteractive.BorderColor attribute to ColorYellow600
func (r *ComponentInteractive) SetBorderColorYellow600() *ComponentInteractive {
	r.BorderColor = ColorYellow600
	return r
}

// SetBorderColorYellow700 set ComponentInteractive.BorderColor attribute to ColorYellow700
func (r *ComponentInteractive) SetBorderColorYellow700() *ComponentInteractive {
	r.BorderColor = ColorYellow700
	return r
}

// SetBorderColorYellow800 set ComponentInteractive.BorderColor attribute to ColorYellow800
func (r *ComponentInteractive) SetBorderColorYellow800() *ComponentInteractive {
	r.BorderColor = ColorYellow800
	return r
}

// SetBorderColorYellow900 set ComponentInteractive.BorderColor attribute to ColorYellow900
func (r *ComponentInteractive) SetBorderColorYellow900() *ComponentInteractive {
	r.BorderColor = ColorYellow900
	return r
}

// SetCornerRadius set ComponentInteractive.CornerRadius attribute
func (r *ComponentInteractive) SetCornerRadius(val CornerRadius) *ComponentInteractive {
	r.CornerRadius = val
	return r
}

// SetPadding set ComponentInteractive.Padding attribute
func (r *ComponentInteractive) SetPadding(val Padding) *ComponentInteractive {
	r.Padding = val
	return r
}

// SetBehaviors set ComponentInteractive.Behaviors attribute
func (r *ComponentInteractive) SetBehaviors(val ...*ObjectBehavior) *ComponentInteractive {
	r.Behaviors = val
	return r
}

// SetHoverTips set ComponentInteractive.HoverTips attribute
func (r *ComponentInteractive) SetHoverTips(val *ObjectText) *ComponentInteractive {
	r.HoverTips = val
	return r
}

// SetDisabled set ComponentInteractive.Disabled attribute
func (r *ComponentInteractive) SetDisabled(val bool) *ComponentInteractive {
	r.Disabled = val
	return r
}

// SetDisabledTips set ComponentInteractive.DisabledTips attribute
func (r *ComponentInteractive) SetDisabledTips(val *ObjectText) *ComponentInteractive {
	r.DisabledTips = val
	return r
}

// SetConfirm set ComponentInteractive.Confirm attribute
func (r *ComponentInteractive) SetConfirm(val *ObjectConfirm) *ComponentInteractive {
	r.Confirm = val
	return r
}

// SetElements set ComponentInteractive.Elements attribute
func (r *ComponentInteractive) SetElements(val ...Component) *ComponentInteractive {
	r.Elements = val
	return r
}

// toMap conv ComponentInteractive to map
func (r *ComponentInteractive) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 13)
	if r.Width != "" {
		res["width"] = r.Width
	}
	if r.Height != "" {
		res["height"] = r.Height
	}
	if r.BackgroundStyle != "" {
		res["background_style"] = r.BackgroundStyle
	}
	if r.HasBorder != false {
		res["has_border"] = r.HasBorder
	}
	if r.BorderColor != "" {
		res["border_color"] = r.BorderColor
	}
	if r.CornerRadius != "" {
		res["corner_radius"] = r.CornerRadius
	}
	if r.Padding != "" {
		res["padding"] = r.Padding
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
	if r.Confirm != nil {
		res["confirm"] = r.Confirm
	}
	if len(r.Elements) != 0 {
		res["elements"] = r.Elements
	}
	return res
}
