package card

import "encoding/json"

func ColumnSet(columns ...*ComponentColumn) *ComponentColumnSet {
	return &ComponentColumnSet{
		Columns:           columns,
		HorizontalSpacing: HorizontalSpacingDefault,
		BackgroundStyle:   ColorDefault,
	}
}

// ComponentColumnSet 分栏
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/containers/column-set
//
// 分栏组件提供卡片内布局的能力, 并提供对齐方式、容器宽度、交互方式等属性。你可以使用分栏组件横向排布多个列容器, 在列容器内自由组合图文内容, 搭建出如数据表、商品或文章列表、差旅信息等图文并茂、交互友好的卡片。
//
// 注意事项
// 分栏组件最多支持嵌套五层组件。建议你避免在分栏中嵌套多层组件。多层嵌套会压缩内容的展示空间, 影响卡片的展示效果。
//
// 应用场景
// 数据报表推送场景：使用分栏可以快速构建整齐、自适应屏幕的多列数据表, 解决了传统报表构建时繁琐的排版过程, 以及无法自适应各类屏幕的样式问题。
// 图文混排场景：分栏灵活的横纵列排版能力, 使你可以快速构建理想的图文卡片。有效降低手动调整图文排版的耗时。
// 嵌套规则
// 分栏组件由分栏本身的属性（column_set）和列容器（column）组成。一个分栏组件中内可以添加多个列容器, 每个列容器中可内嵌多个组件, 支持内嵌的组件有：
//
// 容器类：分栏和循环容器组件
// 展示类：除标题组件、表格组件以外的所有展示组件
// 交互类：除折叠按钮组以外的所有交互类组件
//
//go:generate generate_set_attrs -type=ComponentColumnSet
//go:generate generate_to_map -type=ComponentColumnSet
type ComponentColumnSet struct {
	componentBase

	// 各列之间的水平分栏间距。取值：
	//
	// default: 默认间距, 8px
	// small: 窄间距, 4px
	// large: 大间距, 12px
	// [0,28px]: 自定义间距
	HorizontalSpacing HorizontalSpacing `json:"horizontal_spacing,omitempty"`

	// 列容器水平对齐的方式。可取值：
	//
	// left: 左对齐
	// center: 居中对齐
	// right: 右对齐
	HorizontalAlign HorizontalAlign `json:"horizontal_align,omitempty"`

	// 列的外边距。值的取值范围为 [0,28]px。可选值：
	//
	// 单值, 如 "10px", 表示列的四个外边距都为 10 px。
	// 多值, 如 "4px 12px 4px 12px", 表示列的上、右、下、左的外边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔。
	// 注意：首行列的上外边距强制为 0, 末行列的下外边距强制为 0。
	Margin Margin `json:"margin,omitempty"`

	// 移动端和 PC 端的窄屏幕下, 各列的自适应方式。取值：
	//
	// none: 不做布局上的自适应, 在窄屏幕下按比例压缩列宽度
	// stretch: 列布局变为行布局, 且每列（行）宽度强制拉伸为 100%, 所有列自适应为上下堆叠排布
	// flow: 列流式排布（自动换行）, 当一行展示不下一列时, 自动换至下一行展示
	// bisect: 两列等分布局
	// trisect: 三列等分布局
	FlexMode FlexMode `json:"flex_mode,omitempty"`

	// 分栏的背景色样式。可取值：
	//
	// default：默认的白底样式, 客户端深色主题下为黑底样式
	// 卡片支持的颜色枚举值和 RGBA 语法自定义颜色。参考颜色枚举值。
	// 注意：当存在分栏的嵌套时, 上层分栏的颜色覆盖下层分栏的颜色。
	BackgroundStyle Color `json:"background_style,omitempty"`

	// 分栏中列的配置。
	Columns []*ComponentColumn `json:"columns,omitempty"`

	// 设置点击分栏时的交互配置。当前仅支持跳转交互。如果布局容器内有交互组件, 则优先响应交互组件定义的交互。
}

type HorizontalSpacing = string

const (
	HorizontalSpacingDefault HorizontalSpacing = "default" // 默认间距, 8px
	HorizontalSpacingSmall   HorizontalSpacing = "small"   // 窄间距, 4px
	HorizontalSpacingLarge   HorizontalSpacing = "large"   // 大间距, 12px
	// [0,28px]: 自定义间距
)

type HorizontalAlign = string

const (
	HorizontalAlignLeft   HorizontalAlign = "left"   // 左对齐
	HorizontalAlignCenter HorizontalAlign = "center" // 居中对齐
	HorizontalAlignRight  HorizontalAlign = "right"  // 右对齐
)

type FlexMode = string

const (
	FlexModeNone    FlexMode = "none"    // 不做布局上的自适应, 在窄屏幕下按比例压缩列宽度
	FlexModeStretch FlexMode = "stretch" // 列布局变为行布局, 且每列（行）宽度强制拉伸为 100%, 所有列自适应为上下堆叠排布
	FlexModeFlow    FlexMode = "flow"    // 列流式排布（自动换行）, 当一行展示不下一列时, 自动换至下一行展示
	FlexModeBisect  FlexMode = "bisect"  // 两列等分布局
	FlexModeTrisect FlexMode = "trisect" // 三列等分布局
)

// MarshalJSON ...
func (r ComponentColumnSet) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "column_set"
	return json.Marshal(m)
}

// SetHorizontalSpacing set ComponentColumnSet.HorizontalSpacing attribute
func (r *ComponentColumnSet) SetHorizontalSpacing(val HorizontalSpacing) *ComponentColumnSet {
	r.HorizontalSpacing = val
	return r
}

// SetHorizontalSpacingDefault set ComponentColumnSet.HorizontalSpacing attribute to HorizontalSpacingDefault
func (r *ComponentColumnSet) SetHorizontalSpacingDefault() *ComponentColumnSet {
	r.HorizontalSpacing = HorizontalSpacingDefault
	return r
}

// SetHorizontalSpacingSmall set ComponentColumnSet.HorizontalSpacing attribute to HorizontalSpacingSmall
func (r *ComponentColumnSet) SetHorizontalSpacingSmall() *ComponentColumnSet {
	r.HorizontalSpacing = HorizontalSpacingSmall
	return r
}

// SetHorizontalSpacingLarge set ComponentColumnSet.HorizontalSpacing attribute to HorizontalSpacingLarge
func (r *ComponentColumnSet) SetHorizontalSpacingLarge() *ComponentColumnSet {
	r.HorizontalSpacing = HorizontalSpacingLarge
	return r
}

// SetHorizontalAlign set ComponentColumnSet.HorizontalAlign attribute
func (r *ComponentColumnSet) SetHorizontalAlign(val HorizontalAlign) *ComponentColumnSet {
	r.HorizontalAlign = val
	return r
}

// SetHorizontalAlignLeft set ComponentColumnSet.HorizontalAlign attribute to HorizontalAlignLeft
func (r *ComponentColumnSet) SetHorizontalAlignLeft() *ComponentColumnSet {
	r.HorizontalAlign = HorizontalAlignLeft
	return r
}

// SetHorizontalAlignCenter set ComponentColumnSet.HorizontalAlign attribute to HorizontalAlignCenter
func (r *ComponentColumnSet) SetHorizontalAlignCenter() *ComponentColumnSet {
	r.HorizontalAlign = HorizontalAlignCenter
	return r
}

// SetHorizontalAlignRight set ComponentColumnSet.HorizontalAlign attribute to HorizontalAlignRight
func (r *ComponentColumnSet) SetHorizontalAlignRight() *ComponentColumnSet {
	r.HorizontalAlign = HorizontalAlignRight
	return r
}

// SetMargin set ComponentColumnSet.Margin attribute
func (r *ComponentColumnSet) SetMargin(val Margin) *ComponentColumnSet {
	r.Margin = val
	return r
}

// SetFlexMode set ComponentColumnSet.FlexMode attribute
func (r *ComponentColumnSet) SetFlexMode(val FlexMode) *ComponentColumnSet {
	r.FlexMode = val
	return r
}

// SetFlexModeNone set ComponentColumnSet.FlexMode attribute to FlexModeNone
func (r *ComponentColumnSet) SetFlexModeNone() *ComponentColumnSet {
	r.FlexMode = FlexModeNone
	return r
}

// SetFlexModeStretch set ComponentColumnSet.FlexMode attribute to FlexModeStretch
func (r *ComponentColumnSet) SetFlexModeStretch() *ComponentColumnSet {
	r.FlexMode = FlexModeStretch
	return r
}

// SetFlexModeFlow set ComponentColumnSet.FlexMode attribute to FlexModeFlow
func (r *ComponentColumnSet) SetFlexModeFlow() *ComponentColumnSet {
	r.FlexMode = FlexModeFlow
	return r
}

// SetFlexModeBisect set ComponentColumnSet.FlexMode attribute to FlexModeBisect
func (r *ComponentColumnSet) SetFlexModeBisect() *ComponentColumnSet {
	r.FlexMode = FlexModeBisect
	return r
}

// SetFlexModeTrisect set ComponentColumnSet.FlexMode attribute to FlexModeTrisect
func (r *ComponentColumnSet) SetFlexModeTrisect() *ComponentColumnSet {
	r.FlexMode = FlexModeTrisect
	return r
}

// SetBackgroundStyle set ComponentColumnSet.BackgroundStyle attribute
func (r *ComponentColumnSet) SetBackgroundStyle(val Color) *ComponentColumnSet {
	r.BackgroundStyle = val
	return r
}

// SetBackgroundStyleDefault set ComponentColumnSet.BackgroundStyle attribute to ColorDefault
func (r *ComponentColumnSet) SetBackgroundStyleDefault() *ComponentColumnSet {
	r.BackgroundStyle = ColorDefault
	return r
}

// SetBackgroundStyleBgWhite set ComponentColumnSet.BackgroundStyle attribute to ColorBgWhite
func (r *ComponentColumnSet) SetBackgroundStyleBgWhite() *ComponentColumnSet {
	r.BackgroundStyle = ColorBgWhite
	return r
}

// SetBackgroundStyleWhite set ComponentColumnSet.BackgroundStyle attribute to ColorWhite
func (r *ComponentColumnSet) SetBackgroundStyleWhite() *ComponentColumnSet {
	r.BackgroundStyle = ColorWhite
	return r
}

// SetBackgroundStyleBlue set ComponentColumnSet.BackgroundStyle attribute to ColorBlue
func (r *ComponentColumnSet) SetBackgroundStyleBlue() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue
	return r
}

// SetBackgroundStyleBlue50 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue50
func (r *ComponentColumnSet) SetBackgroundStyleBlue50() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue50
	return r
}

// SetBackgroundStyleBlue100 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue100
func (r *ComponentColumnSet) SetBackgroundStyleBlue100() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue100
	return r
}

// SetBackgroundStyleBlue200 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue200
func (r *ComponentColumnSet) SetBackgroundStyleBlue200() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue200
	return r
}

// SetBackgroundStyleBlue300 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue300
func (r *ComponentColumnSet) SetBackgroundStyleBlue300() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue300
	return r
}

// SetBackgroundStyleBlue350 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue350
func (r *ComponentColumnSet) SetBackgroundStyleBlue350() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue350
	return r
}

// SetBackgroundStyleBlue400 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue400
func (r *ComponentColumnSet) SetBackgroundStyleBlue400() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue400
	return r
}

// SetBackgroundStyleBlue500 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue500
func (r *ComponentColumnSet) SetBackgroundStyleBlue500() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue500
	return r
}

// SetBackgroundStyleBlue600 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue600
func (r *ComponentColumnSet) SetBackgroundStyleBlue600() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue600
	return r
}

// SetBackgroundStyleBlue700 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue700
func (r *ComponentColumnSet) SetBackgroundStyleBlue700() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue700
	return r
}

// SetBackgroundStyleBlue800 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue800
func (r *ComponentColumnSet) SetBackgroundStyleBlue800() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue800
	return r
}

// SetBackgroundStyleBlue900 set ComponentColumnSet.BackgroundStyle attribute to ColorBlue900
func (r *ComponentColumnSet) SetBackgroundStyleBlue900() *ComponentColumnSet {
	r.BackgroundStyle = ColorBlue900
	return r
}

// SetBackgroundStyleCarmine set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine
func (r *ComponentColumnSet) SetBackgroundStyleCarmine() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine
	return r
}

// SetBackgroundStyleCarmine50 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine50
func (r *ComponentColumnSet) SetBackgroundStyleCarmine50() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine50
	return r
}

// SetBackgroundStyleCarmine100 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine100
func (r *ComponentColumnSet) SetBackgroundStyleCarmine100() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine100
	return r
}

// SetBackgroundStyleCarmine200 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine200
func (r *ComponentColumnSet) SetBackgroundStyleCarmine200() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine200
	return r
}

// SetBackgroundStyleCarmine300 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine300
func (r *ComponentColumnSet) SetBackgroundStyleCarmine300() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine300
	return r
}

// SetBackgroundStyleCarmine350 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine350
func (r *ComponentColumnSet) SetBackgroundStyleCarmine350() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine350
	return r
}

// SetBackgroundStyleCarmine400 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine400
func (r *ComponentColumnSet) SetBackgroundStyleCarmine400() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine400
	return r
}

// SetBackgroundStyleCarmine500 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine500
func (r *ComponentColumnSet) SetBackgroundStyleCarmine500() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine500
	return r
}

// SetBackgroundStyleCarmine600 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine600
func (r *ComponentColumnSet) SetBackgroundStyleCarmine600() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine600
	return r
}

// SetBackgroundStyleCarmine700 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine700
func (r *ComponentColumnSet) SetBackgroundStyleCarmine700() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine700
	return r
}

// SetBackgroundStyleCarmine800 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine800
func (r *ComponentColumnSet) SetBackgroundStyleCarmine800() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine800
	return r
}

// SetBackgroundStyleCarmine900 set ComponentColumnSet.BackgroundStyle attribute to ColorCarmine900
func (r *ComponentColumnSet) SetBackgroundStyleCarmine900() *ComponentColumnSet {
	r.BackgroundStyle = ColorCarmine900
	return r
}

// SetBackgroundStyleGreen set ComponentColumnSet.BackgroundStyle attribute to ColorGreen
func (r *ComponentColumnSet) SetBackgroundStyleGreen() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen
	return r
}

// SetBackgroundStyleGreen50 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen50
func (r *ComponentColumnSet) SetBackgroundStyleGreen50() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen50
	return r
}

// SetBackgroundStyleGreen100 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen100
func (r *ComponentColumnSet) SetBackgroundStyleGreen100() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen100
	return r
}

// SetBackgroundStyleGreen200 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen200
func (r *ComponentColumnSet) SetBackgroundStyleGreen200() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen200
	return r
}

// SetBackgroundStyleGreen300 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen300
func (r *ComponentColumnSet) SetBackgroundStyleGreen300() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen300
	return r
}

// SetBackgroundStyleGreen350 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen350
func (r *ComponentColumnSet) SetBackgroundStyleGreen350() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen350
	return r
}

// SetBackgroundStyleGreen400 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen400
func (r *ComponentColumnSet) SetBackgroundStyleGreen400() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen400
	return r
}

// SetBackgroundStyleGreen500 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen500
func (r *ComponentColumnSet) SetBackgroundStyleGreen500() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen500
	return r
}

// SetBackgroundStyleGreen600 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen600
func (r *ComponentColumnSet) SetBackgroundStyleGreen600() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen600
	return r
}

// SetBackgroundStyleGreen700 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen700
func (r *ComponentColumnSet) SetBackgroundStyleGreen700() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen700
	return r
}

// SetBackgroundStyleGreen800 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen800
func (r *ComponentColumnSet) SetBackgroundStyleGreen800() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen800
	return r
}

// SetBackgroundStyleGreen900 set ComponentColumnSet.BackgroundStyle attribute to ColorGreen900
func (r *ComponentColumnSet) SetBackgroundStyleGreen900() *ComponentColumnSet {
	r.BackgroundStyle = ColorGreen900
	return r
}

// SetBackgroundStyleIndigo set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo
func (r *ComponentColumnSet) SetBackgroundStyleIndigo() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo
	return r
}

// SetBackgroundStyleIndigo50 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo50
func (r *ComponentColumnSet) SetBackgroundStyleIndigo50() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo50
	return r
}

// SetBackgroundStyleIndigo100 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo100
func (r *ComponentColumnSet) SetBackgroundStyleIndigo100() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo100
	return r
}

// SetBackgroundStyleIndigo200 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo200
func (r *ComponentColumnSet) SetBackgroundStyleIndigo200() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo200
	return r
}

// SetBackgroundStyleIndigo300 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo300
func (r *ComponentColumnSet) SetBackgroundStyleIndigo300() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo300
	return r
}

// SetBackgroundStyleIndigo350 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo350
func (r *ComponentColumnSet) SetBackgroundStyleIndigo350() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo350
	return r
}

// SetBackgroundStyleIndigo400 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo400
func (r *ComponentColumnSet) SetBackgroundStyleIndigo400() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo400
	return r
}

// SetBackgroundStyleIndigo500 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo500
func (r *ComponentColumnSet) SetBackgroundStyleIndigo500() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo500
	return r
}

// SetBackgroundStyleIndigo600 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo600
func (r *ComponentColumnSet) SetBackgroundStyleIndigo600() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo600
	return r
}

// SetBackgroundStyleIndigo700 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo700
func (r *ComponentColumnSet) SetBackgroundStyleIndigo700() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo700
	return r
}

// SetBackgroundStyleIndigo800 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo800
func (r *ComponentColumnSet) SetBackgroundStyleIndigo800() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo800
	return r
}

// SetBackgroundStyleIndigo900 set ComponentColumnSet.BackgroundStyle attribute to ColorIndigo900
func (r *ComponentColumnSet) SetBackgroundStyleIndigo900() *ComponentColumnSet {
	r.BackgroundStyle = ColorIndigo900
	return r
}

// SetBackgroundStyleLime set ComponentColumnSet.BackgroundStyle attribute to ColorLime
func (r *ComponentColumnSet) SetBackgroundStyleLime() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime
	return r
}

// SetBackgroundStyleLime50 set ComponentColumnSet.BackgroundStyle attribute to ColorLime50
func (r *ComponentColumnSet) SetBackgroundStyleLime50() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime50
	return r
}

// SetBackgroundStyleLime100 set ComponentColumnSet.BackgroundStyle attribute to ColorLime100
func (r *ComponentColumnSet) SetBackgroundStyleLime100() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime100
	return r
}

// SetBackgroundStyleLime200 set ComponentColumnSet.BackgroundStyle attribute to ColorLime200
func (r *ComponentColumnSet) SetBackgroundStyleLime200() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime200
	return r
}

// SetBackgroundStyleLime300 set ComponentColumnSet.BackgroundStyle attribute to ColorLime300
func (r *ComponentColumnSet) SetBackgroundStyleLime300() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime300
	return r
}

// SetBackgroundStyleLime350 set ComponentColumnSet.BackgroundStyle attribute to ColorLime350
func (r *ComponentColumnSet) SetBackgroundStyleLime350() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime350
	return r
}

// SetBackgroundStyleLime400 set ComponentColumnSet.BackgroundStyle attribute to ColorLime400
func (r *ComponentColumnSet) SetBackgroundStyleLime400() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime400
	return r
}

// SetBackgroundStyleLime500 set ComponentColumnSet.BackgroundStyle attribute to ColorLime500
func (r *ComponentColumnSet) SetBackgroundStyleLime500() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime500
	return r
}

// SetBackgroundStyleLime600 set ComponentColumnSet.BackgroundStyle attribute to ColorLime600
func (r *ComponentColumnSet) SetBackgroundStyleLime600() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime600
	return r
}

// SetBackgroundStyleLime700 set ComponentColumnSet.BackgroundStyle attribute to ColorLime700
func (r *ComponentColumnSet) SetBackgroundStyleLime700() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime700
	return r
}

// SetBackgroundStyleLime800 set ComponentColumnSet.BackgroundStyle attribute to ColorLime800
func (r *ComponentColumnSet) SetBackgroundStyleLime800() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime800
	return r
}

// SetBackgroundStyleLime900 set ComponentColumnSet.BackgroundStyle attribute to ColorLime900
func (r *ComponentColumnSet) SetBackgroundStyleLime900() *ComponentColumnSet {
	r.BackgroundStyle = ColorLime900
	return r
}

// SetBackgroundStyleGrey set ComponentColumnSet.BackgroundStyle attribute to ColorGrey
func (r *ComponentColumnSet) SetBackgroundStyleGrey() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey
	return r
}

// SetBackgroundStyleGrey00 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey00
func (r *ComponentColumnSet) SetBackgroundStyleGrey00() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey00
	return r
}

// SetBackgroundStyleGrey50 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey50
func (r *ComponentColumnSet) SetBackgroundStyleGrey50() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey50
	return r
}

// SetBackgroundStyleGrey100 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey100
func (r *ComponentColumnSet) SetBackgroundStyleGrey100() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey100
	return r
}

// SetBackgroundStyleGrey200 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey200
func (r *ComponentColumnSet) SetBackgroundStyleGrey200() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey200
	return r
}

// SetBackgroundStyleGrey300 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey300
func (r *ComponentColumnSet) SetBackgroundStyleGrey300() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey300
	return r
}

// SetBackgroundStyleGrey350 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey350
func (r *ComponentColumnSet) SetBackgroundStyleGrey350() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey350
	return r
}

// SetBackgroundStyleGrey400 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey400
func (r *ComponentColumnSet) SetBackgroundStyleGrey400() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey400
	return r
}

// SetBackgroundStyleGrey500 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey500
func (r *ComponentColumnSet) SetBackgroundStyleGrey500() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey500
	return r
}

// SetBackgroundStyleGrey600 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey600
func (r *ComponentColumnSet) SetBackgroundStyleGrey600() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey600
	return r
}

// SetBackgroundStyleGrey650 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey650
func (r *ComponentColumnSet) SetBackgroundStyleGrey650() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey650
	return r
}

// SetBackgroundStyleGrey700 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey700
func (r *ComponentColumnSet) SetBackgroundStyleGrey700() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey700
	return r
}

// SetBackgroundStyleGrey800 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey800
func (r *ComponentColumnSet) SetBackgroundStyleGrey800() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey800
	return r
}

// SetBackgroundStyleGrey900 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey900
func (r *ComponentColumnSet) SetBackgroundStyleGrey900() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey900
	return r
}

// SetBackgroundStyleGrey950 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey950
func (r *ComponentColumnSet) SetBackgroundStyleGrey950() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey950
	return r
}

// SetBackgroundStyleGrey1000 set ComponentColumnSet.BackgroundStyle attribute to ColorGrey1000
func (r *ComponentColumnSet) SetBackgroundStyleGrey1000() *ComponentColumnSet {
	r.BackgroundStyle = ColorGrey1000
	return r
}

// SetBackgroundStyleOrange set ComponentColumnSet.BackgroundStyle attribute to ColorOrange
func (r *ComponentColumnSet) SetBackgroundStyleOrange() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange
	return r
}

// SetBackgroundStyleOrange50 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange50
func (r *ComponentColumnSet) SetBackgroundStyleOrange50() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange50
	return r
}

// SetBackgroundStyleOrange100 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange100
func (r *ComponentColumnSet) SetBackgroundStyleOrange100() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange100
	return r
}

// SetBackgroundStyleOrange200 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange200
func (r *ComponentColumnSet) SetBackgroundStyleOrange200() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange200
	return r
}

// SetBackgroundStyleOrange300 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange300
func (r *ComponentColumnSet) SetBackgroundStyleOrange300() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange300
	return r
}

// SetBackgroundStyleOrange350 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange350
func (r *ComponentColumnSet) SetBackgroundStyleOrange350() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange350
	return r
}

// SetBackgroundStyleOrange400 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange400
func (r *ComponentColumnSet) SetBackgroundStyleOrange400() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange400
	return r
}

// SetBackgroundStyleOrange500 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange500
func (r *ComponentColumnSet) SetBackgroundStyleOrange500() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange500
	return r
}

// SetBackgroundStyleOrange600 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange600
func (r *ComponentColumnSet) SetBackgroundStyleOrange600() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange600
	return r
}

// SetBackgroundStyleOrange700 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange700
func (r *ComponentColumnSet) SetBackgroundStyleOrange700() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange700
	return r
}

// SetBackgroundStyleOrange800 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange800
func (r *ComponentColumnSet) SetBackgroundStyleOrange800() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange800
	return r
}

// SetBackgroundStyleOrange900 set ComponentColumnSet.BackgroundStyle attribute to ColorOrange900
func (r *ComponentColumnSet) SetBackgroundStyleOrange900() *ComponentColumnSet {
	r.BackgroundStyle = ColorOrange900
	return r
}

// SetBackgroundStylePurple set ComponentColumnSet.BackgroundStyle attribute to ColorPurple
func (r *ComponentColumnSet) SetBackgroundStylePurple() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple
	return r
}

// SetBackgroundStylePurple50 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple50
func (r *ComponentColumnSet) SetBackgroundStylePurple50() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple50
	return r
}

// SetBackgroundStylePurple100 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple100
func (r *ComponentColumnSet) SetBackgroundStylePurple100() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple100
	return r
}

// SetBackgroundStylePurple200 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple200
func (r *ComponentColumnSet) SetBackgroundStylePurple200() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple200
	return r
}

// SetBackgroundStylePurple300 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple300
func (r *ComponentColumnSet) SetBackgroundStylePurple300() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple300
	return r
}

// SetBackgroundStylePurple350 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple350
func (r *ComponentColumnSet) SetBackgroundStylePurple350() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple350
	return r
}

// SetBackgroundStylePurple400 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple400
func (r *ComponentColumnSet) SetBackgroundStylePurple400() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple400
	return r
}

// SetBackgroundStylePurple500 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple500
func (r *ComponentColumnSet) SetBackgroundStylePurple500() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple500
	return r
}

// SetBackgroundStylePurple600 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple600
func (r *ComponentColumnSet) SetBackgroundStylePurple600() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple600
	return r
}

// SetBackgroundStylePurple700 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple700
func (r *ComponentColumnSet) SetBackgroundStylePurple700() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple700
	return r
}

// SetBackgroundStylePurple800 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple800
func (r *ComponentColumnSet) SetBackgroundStylePurple800() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple800
	return r
}

// SetBackgroundStylePurple900 set ComponentColumnSet.BackgroundStyle attribute to ColorPurple900
func (r *ComponentColumnSet) SetBackgroundStylePurple900() *ComponentColumnSet {
	r.BackgroundStyle = ColorPurple900
	return r
}

// SetBackgroundStyleRed set ComponentColumnSet.BackgroundStyle attribute to ColorRed
func (r *ComponentColumnSet) SetBackgroundStyleRed() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed
	return r
}

// SetBackgroundStyleRed50 set ComponentColumnSet.BackgroundStyle attribute to ColorRed50
func (r *ComponentColumnSet) SetBackgroundStyleRed50() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed50
	return r
}

// SetBackgroundStyleRed100 set ComponentColumnSet.BackgroundStyle attribute to ColorRed100
func (r *ComponentColumnSet) SetBackgroundStyleRed100() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed100
	return r
}

// SetBackgroundStyleRed200 set ComponentColumnSet.BackgroundStyle attribute to ColorRed200
func (r *ComponentColumnSet) SetBackgroundStyleRed200() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed200
	return r
}

// SetBackgroundStyleRed300 set ComponentColumnSet.BackgroundStyle attribute to ColorRed300
func (r *ComponentColumnSet) SetBackgroundStyleRed300() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed300
	return r
}

// SetBackgroundStyleRed350 set ComponentColumnSet.BackgroundStyle attribute to ColorRed350
func (r *ComponentColumnSet) SetBackgroundStyleRed350() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed350
	return r
}

// SetBackgroundStyleRed400 set ComponentColumnSet.BackgroundStyle attribute to ColorRed400
func (r *ComponentColumnSet) SetBackgroundStyleRed400() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed400
	return r
}

// SetBackgroundStyleRed500 set ComponentColumnSet.BackgroundStyle attribute to ColorRed500
func (r *ComponentColumnSet) SetBackgroundStyleRed500() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed500
	return r
}

// SetBackgroundStyleRed600 set ComponentColumnSet.BackgroundStyle attribute to ColorRed600
func (r *ComponentColumnSet) SetBackgroundStyleRed600() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed600
	return r
}

// SetBackgroundStyleRed700 set ComponentColumnSet.BackgroundStyle attribute to ColorRed700
func (r *ComponentColumnSet) SetBackgroundStyleRed700() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed700
	return r
}

// SetBackgroundStyleRed800 set ComponentColumnSet.BackgroundStyle attribute to ColorRed800
func (r *ComponentColumnSet) SetBackgroundStyleRed800() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed800
	return r
}

// SetBackgroundStyleRed900 set ComponentColumnSet.BackgroundStyle attribute to ColorRed900
func (r *ComponentColumnSet) SetBackgroundStyleRed900() *ComponentColumnSet {
	r.BackgroundStyle = ColorRed900
	return r
}

// SetBackgroundStyleSunflower set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower
func (r *ComponentColumnSet) SetBackgroundStyleSunflower() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower
	return r
}

// SetBackgroundStyleSunflower50 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower50
func (r *ComponentColumnSet) SetBackgroundStyleSunflower50() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower50
	return r
}

// SetBackgroundStyleSunflower100 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower100
func (r *ComponentColumnSet) SetBackgroundStyleSunflower100() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower100
	return r
}

// SetBackgroundStyleSunflower200 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower200
func (r *ComponentColumnSet) SetBackgroundStyleSunflower200() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower200
	return r
}

// SetBackgroundStyleSunflower300 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower300
func (r *ComponentColumnSet) SetBackgroundStyleSunflower300() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower300
	return r
}

// SetBackgroundStyleSunflower350 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower350
func (r *ComponentColumnSet) SetBackgroundStyleSunflower350() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower350
	return r
}

// SetBackgroundStyleSunflower400 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower400
func (r *ComponentColumnSet) SetBackgroundStyleSunflower400() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower400
	return r
}

// SetBackgroundStyleSunflower500 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower500
func (r *ComponentColumnSet) SetBackgroundStyleSunflower500() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower500
	return r
}

// SetBackgroundStyleSunflower600 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower600
func (r *ComponentColumnSet) SetBackgroundStyleSunflower600() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower600
	return r
}

// SetBackgroundStyleSunflower700 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower700
func (r *ComponentColumnSet) SetBackgroundStyleSunflower700() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower700
	return r
}

// SetBackgroundStyleSunflower800 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower800
func (r *ComponentColumnSet) SetBackgroundStyleSunflower800() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower800
	return r
}

// SetBackgroundStyleSunflower900 set ComponentColumnSet.BackgroundStyle attribute to ColorSunflower900
func (r *ComponentColumnSet) SetBackgroundStyleSunflower900() *ComponentColumnSet {
	r.BackgroundStyle = ColorSunflower900
	return r
}

// SetBackgroundStyleTurquoise set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise
	return r
}

// SetBackgroundStyleTurquoise50 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise50
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise50() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise50
	return r
}

// SetBackgroundStyleTurquoise100 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise100
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise100() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise100
	return r
}

// SetBackgroundStyleTurquoise200 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise200
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise200() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise200
	return r
}

// SetBackgroundStyleTurquoise300 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise300
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise300() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise300
	return r
}

// SetBackgroundStyleTurquoise350 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise350
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise350() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise350
	return r
}

// SetBackgroundStyleTurquoise400 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise400
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise400() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise400
	return r
}

// SetBackgroundStyleTurquoise500 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise500
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise500() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise500
	return r
}

// SetBackgroundStyleTurquoise600 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise600
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise600() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise600
	return r
}

// SetBackgroundStyleTurquoise700 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise700
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise700() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise700
	return r
}

// SetBackgroundStyleTurquoise800 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise800
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise800() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise800
	return r
}

// SetBackgroundStyleTurquoise900 set ComponentColumnSet.BackgroundStyle attribute to ColorTurquoise900
func (r *ComponentColumnSet) SetBackgroundStyleTurquoise900() *ComponentColumnSet {
	r.BackgroundStyle = ColorTurquoise900
	return r
}

// SetBackgroundStyleViolet set ComponentColumnSet.BackgroundStyle attribute to ColorViolet
func (r *ComponentColumnSet) SetBackgroundStyleViolet() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet
	return r
}

// SetBackgroundStyleViolet50 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet50
func (r *ComponentColumnSet) SetBackgroundStyleViolet50() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet50
	return r
}

// SetBackgroundStyleViolet100 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet100
func (r *ComponentColumnSet) SetBackgroundStyleViolet100() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet100
	return r
}

// SetBackgroundStyleViolet200 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet200
func (r *ComponentColumnSet) SetBackgroundStyleViolet200() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet200
	return r
}

// SetBackgroundStyleViolet300 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet300
func (r *ComponentColumnSet) SetBackgroundStyleViolet300() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet300
	return r
}

// SetBackgroundStyleViolet350 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet350
func (r *ComponentColumnSet) SetBackgroundStyleViolet350() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet350
	return r
}

// SetBackgroundStyleViolet400 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet400
func (r *ComponentColumnSet) SetBackgroundStyleViolet400() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet400
	return r
}

// SetBackgroundStyleViolet500 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet500
func (r *ComponentColumnSet) SetBackgroundStyleViolet500() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet500
	return r
}

// SetBackgroundStyleViolet600 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet600
func (r *ComponentColumnSet) SetBackgroundStyleViolet600() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet600
	return r
}

// SetBackgroundStyleViolet700 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet700
func (r *ComponentColumnSet) SetBackgroundStyleViolet700() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet700
	return r
}

// SetBackgroundStyleViolet800 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet800
func (r *ComponentColumnSet) SetBackgroundStyleViolet800() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet800
	return r
}

// SetBackgroundStyleViolet900 set ComponentColumnSet.BackgroundStyle attribute to ColorViolet900
func (r *ComponentColumnSet) SetBackgroundStyleViolet900() *ComponentColumnSet {
	r.BackgroundStyle = ColorViolet900
	return r
}

// SetBackgroundStyleWathet set ComponentColumnSet.BackgroundStyle attribute to ColorWathet
func (r *ComponentColumnSet) SetBackgroundStyleWathet() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet
	return r
}

// SetBackgroundStyleWathet50 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet50
func (r *ComponentColumnSet) SetBackgroundStyleWathet50() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet50
	return r
}

// SetBackgroundStyleWathet100 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet100
func (r *ComponentColumnSet) SetBackgroundStyleWathet100() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet100
	return r
}

// SetBackgroundStyleWathet200 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet200
func (r *ComponentColumnSet) SetBackgroundStyleWathet200() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet200
	return r
}

// SetBackgroundStyleWathet300 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet300
func (r *ComponentColumnSet) SetBackgroundStyleWathet300() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet300
	return r
}

// SetBackgroundStyleWathet350 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet350
func (r *ComponentColumnSet) SetBackgroundStyleWathet350() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet350
	return r
}

// SetBackgroundStyleWathet400 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet400
func (r *ComponentColumnSet) SetBackgroundStyleWathet400() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet400
	return r
}

// SetBackgroundStyleWathet500 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet500
func (r *ComponentColumnSet) SetBackgroundStyleWathet500() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet500
	return r
}

// SetBackgroundStyleWathet600 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet600
func (r *ComponentColumnSet) SetBackgroundStyleWathet600() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet600
	return r
}

// SetBackgroundStyleWathet700 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet700
func (r *ComponentColumnSet) SetBackgroundStyleWathet700() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet700
	return r
}

// SetBackgroundStyleWathet800 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet800
func (r *ComponentColumnSet) SetBackgroundStyleWathet800() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet800
	return r
}

// SetBackgroundStyleWathet900 set ComponentColumnSet.BackgroundStyle attribute to ColorWathet900
func (r *ComponentColumnSet) SetBackgroundStyleWathet900() *ComponentColumnSet {
	r.BackgroundStyle = ColorWathet900
	return r
}

// SetBackgroundStyleYellow set ComponentColumnSet.BackgroundStyle attribute to ColorYellow
func (r *ComponentColumnSet) SetBackgroundStyleYellow() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow
	return r
}

// SetBackgroundStyleYellow50 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow50
func (r *ComponentColumnSet) SetBackgroundStyleYellow50() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow50
	return r
}

// SetBackgroundStyleYellow100 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow100
func (r *ComponentColumnSet) SetBackgroundStyleYellow100() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow100
	return r
}

// SetBackgroundStyleYellow200 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow200
func (r *ComponentColumnSet) SetBackgroundStyleYellow200() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow200
	return r
}

// SetBackgroundStyleYellow300 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow300
func (r *ComponentColumnSet) SetBackgroundStyleYellow300() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow300
	return r
}

// SetBackgroundStyleYellow350 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow350
func (r *ComponentColumnSet) SetBackgroundStyleYellow350() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow350
	return r
}

// SetBackgroundStyleYellow400 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow400
func (r *ComponentColumnSet) SetBackgroundStyleYellow400() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow400
	return r
}

// SetBackgroundStyleYellow500 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow500
func (r *ComponentColumnSet) SetBackgroundStyleYellow500() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow500
	return r
}

// SetBackgroundStyleYellow600 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow600
func (r *ComponentColumnSet) SetBackgroundStyleYellow600() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow600
	return r
}

// SetBackgroundStyleYellow700 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow700
func (r *ComponentColumnSet) SetBackgroundStyleYellow700() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow700
	return r
}

// SetBackgroundStyleYellow800 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow800
func (r *ComponentColumnSet) SetBackgroundStyleYellow800() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow800
	return r
}

// SetBackgroundStyleYellow900 set ComponentColumnSet.BackgroundStyle attribute to ColorYellow900
func (r *ComponentColumnSet) SetBackgroundStyleYellow900() *ComponentColumnSet {
	r.BackgroundStyle = ColorYellow900
	return r
}

// SetColumns set ComponentColumnSet.Columns attribute
func (r *ComponentColumnSet) SetColumns(val ...*ComponentColumn) *ComponentColumnSet {
	r.Columns = val
	return r
}

// toMap conv ComponentColumnSet to map
func (r *ComponentColumnSet) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 6)
	if r.HorizontalSpacing != "" {
		res["horizontal_spacing"] = r.HorizontalSpacing
	}
	if r.HorizontalAlign != "" {
		res["horizontal_align"] = r.HorizontalAlign
	}
	if r.Margin != "" {
		res["margin"] = r.Margin
	}
	if r.FlexMode != "" {
		res["flex_mode"] = r.FlexMode
	}
	if r.BackgroundStyle != "" {
		res["background_style"] = r.BackgroundStyle
	}
	if len(r.Columns) != 0 {
		res["columns"] = r.Columns
	}
	return res
}
