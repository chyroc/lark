package card

import (
	"encoding/json"
)

func Column(elements ...Component) *ComponentColumn {
	return &ComponentColumn{
		BackgroundStyle: ColorDefault,
		VerticalSpacing: VerticalSpacingDefault,
		Elements:        elements,
	}
}

// ComponentColumn
// /go:generate generate_iface_unmarshal -type=ComponentColumn
// /go:generate generate_to_map -type=ComponentColumn
//
//go:generate generate_set_attrs -type=ComponentColumn
type ComponentColumn struct {
	componentBase

	// 列的背景色样式。可取值：
	//
	// default: 默认的白底样式, 客户端深色主题下为黑底样式
	// 卡片支持的颜色枚举值和 RGBA 语法自定义颜色。参考颜色枚举值
	BackgroundStyle Color `json:"background_style,omitempty"`

	// 列宽度。仅 flex_mode 为 none 时, 生效此属性。取值：
	//
	// auto: 列宽度与列内元素宽度一致
	// weighted: 列宽度按 weight 参数定义的权重分布
	// 具体数值, 如 100px。取值范围为 [16,600]px。V7.4 及以上版本支持该枚举
	Width ColumnWidth `json:"width,omitempty"`

	// 当 width 字段取值为 weighted 时生效, 表示当前列的宽度占比。取值范围为 1 ~ 5 之间的整数。
	Weight int64 `json:"weight,omitempty"`

	// 列垂直居中的方式。可取值：
	//
	// top: 上对齐
	// center: 居中对齐
	// bottom: 下对齐
	VerticalAlign VerticalAlign `json:"vertical_align,omitempty"`

	// 列内组件的纵向间距。取值：
	//
	// default: 默认间距, 8px
	// medium: 中等间距
	// large: 大间距
	// 具体数值, 如 8px。取值范围为 [0,28]px
	VerticalSpacing VerticalSpacing `json:"vertical_spacing,omitempty"`

	// 列的内边距。值的取值范围为 [0,28]px。可选值：
	//
	// 单值, 如 "10px", 表示列的四个外边距都为 10 px。
	// 多值, 如 "4px 12px 4px 12px", 表示列的上、右、下、左的外边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔。
	Padding Padding `json:"padding,omitempty"`

	// 列容器中内嵌的组件。可内嵌组件参考上文嵌套关系。
	Elements []Component `json:"elements,omitempty" unmarshal:"unmarshalComponent"`

	// 设置点击列时的交互配置。当前仅支持跳转交互。如果布局容器内有交互组件, 则优先响应交互组件定义的交互。
	Action *ComponentAction `json:"action,omitempty"`
}

type ColumnWidth = string

const (
	ColumnWidthAuto     ColumnWidth = "auto"     // 列宽度与列内元素宽度一致
	ColumnWidthWeighted ColumnWidth = "weighted" // 列宽度按 weight 参数定义的权重分布
)

type VerticalSpacing = string

const (
	VerticalSpacingDefault VerticalSpacing = "default" // 默认间距, 8px
	VerticalSpacingMedium  VerticalSpacing = "medium"  // 中等间距
	VerticalSpacingLarge   VerticalSpacing = "large"   // 大间距
	// 具体数值, 如 8px。取值范围为 [0,28]px
)

// MarshalJSON ...
func (r ComponentColumn) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "column"
	return json.Marshal(m)
}

// toMap conv ComponentColumn to map
func (r *ComponentColumn) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 8)
	if r.BackgroundStyle != "" {
		res["background_style"] = r.BackgroundStyle
	}
	if r.Width != "" {
		res["width"] = r.Width
	}
	if r.Weight != 0 {
		res["weight"] = r.Weight
	}
	if r.VerticalAlign != "" {
		res["vertical_align"] = r.VerticalAlign
	}
	if r.VerticalSpacing != "" {
		res["vertical_spacing"] = r.VerticalSpacing
	}
	if r.Padding != "" {
		res["padding"] = r.Padding
	}
	if len(r.Elements) != 0 {
		res["elements"] = r.Elements
	}
	if r.Action != nil {
		res["action"] = r.Action
	}
	return res
}

// unmarshalComponentColumn generated to unmarshal ComponentColumn iface
type unmarshalComponentColumn struct {
	BackgroundStyle Color             `json:"background_style,omitempty"`
	Width           string            `json:"width,omitempty"`
	Weight          int64             `json:"weight,omitempty"`
	VerticalAlign   VerticalAlign     `json:"vertical_align,omitempty"`
	VerticalSpacing string            `json:"vertical_spacing,omitempty"`
	Padding         Padding           `json:"padding,omitempty"`
	Elements        []json.RawMessage `json:"elements,omitempty"`
	Action          *ComponentAction  `json:"action,omitempty"`
}

// UnmarshalJSON generated to unmarshal ComponentColumn iface
func (r *ComponentColumn) UnmarshalJSON(bs []byte) error {
	obj := new(unmarshalComponentColumn)
	err := json.Unmarshal(bs, obj)
	if err != nil {
		return err
	}
	r.BackgroundStyle = obj.BackgroundStyle
	r.Width = obj.Width
	r.Weight = obj.Weight
	r.VerticalAlign = obj.VerticalAlign
	r.VerticalSpacing = obj.VerticalSpacing
	r.Padding = obj.Padding
	r.Elements = make([]Component, len(obj.Elements))
	for i, v := range obj.Elements {
		r.Elements[i], err = unmarshalComponent(v)
		if err != nil {
			return err
		}
	}
	r.Action = obj.Action
	return nil
}

// SetBackgroundStyle set ComponentColumn.BackgroundStyle attribute
func (r *ComponentColumn) SetBackgroundStyle(val Color) *ComponentColumn {
	r.BackgroundStyle = val
	return r
}

// SetBackgroundStyleDefault set ComponentColumn.BackgroundStyle attribute to ColorDefault
func (r *ComponentColumn) SetBackgroundStyleDefault() *ComponentColumn {
	r.BackgroundStyle = ColorDefault
	return r
}

// SetBackgroundStyleBgWhite set ComponentColumn.BackgroundStyle attribute to ColorBgWhite
func (r *ComponentColumn) SetBackgroundStyleBgWhite() *ComponentColumn {
	r.BackgroundStyle = ColorBgWhite
	return r
}

// SetBackgroundStyleWhite set ComponentColumn.BackgroundStyle attribute to ColorWhite
func (r *ComponentColumn) SetBackgroundStyleWhite() *ComponentColumn {
	r.BackgroundStyle = ColorWhite
	return r
}

// SetBackgroundStyleBlue set ComponentColumn.BackgroundStyle attribute to ColorBlue
func (r *ComponentColumn) SetBackgroundStyleBlue() *ComponentColumn {
	r.BackgroundStyle = ColorBlue
	return r
}

// SetBackgroundStyleBlue50 set ComponentColumn.BackgroundStyle attribute to ColorBlue50
func (r *ComponentColumn) SetBackgroundStyleBlue50() *ComponentColumn {
	r.BackgroundStyle = ColorBlue50
	return r
}

// SetBackgroundStyleBlue100 set ComponentColumn.BackgroundStyle attribute to ColorBlue100
func (r *ComponentColumn) SetBackgroundStyleBlue100() *ComponentColumn {
	r.BackgroundStyle = ColorBlue100
	return r
}

// SetBackgroundStyleBlue200 set ComponentColumn.BackgroundStyle attribute to ColorBlue200
func (r *ComponentColumn) SetBackgroundStyleBlue200() *ComponentColumn {
	r.BackgroundStyle = ColorBlue200
	return r
}

// SetBackgroundStyleBlue300 set ComponentColumn.BackgroundStyle attribute to ColorBlue300
func (r *ComponentColumn) SetBackgroundStyleBlue300() *ComponentColumn {
	r.BackgroundStyle = ColorBlue300
	return r
}

// SetBackgroundStyleBlue350 set ComponentColumn.BackgroundStyle attribute to ColorBlue350
func (r *ComponentColumn) SetBackgroundStyleBlue350() *ComponentColumn {
	r.BackgroundStyle = ColorBlue350
	return r
}

// SetBackgroundStyleBlue400 set ComponentColumn.BackgroundStyle attribute to ColorBlue400
func (r *ComponentColumn) SetBackgroundStyleBlue400() *ComponentColumn {
	r.BackgroundStyle = ColorBlue400
	return r
}

// SetBackgroundStyleBlue500 set ComponentColumn.BackgroundStyle attribute to ColorBlue500
func (r *ComponentColumn) SetBackgroundStyleBlue500() *ComponentColumn {
	r.BackgroundStyle = ColorBlue500
	return r
}

// SetBackgroundStyleBlue600 set ComponentColumn.BackgroundStyle attribute to ColorBlue600
func (r *ComponentColumn) SetBackgroundStyleBlue600() *ComponentColumn {
	r.BackgroundStyle = ColorBlue600
	return r
}

// SetBackgroundStyleBlue700 set ComponentColumn.BackgroundStyle attribute to ColorBlue700
func (r *ComponentColumn) SetBackgroundStyleBlue700() *ComponentColumn {
	r.BackgroundStyle = ColorBlue700
	return r
}

// SetBackgroundStyleBlue800 set ComponentColumn.BackgroundStyle attribute to ColorBlue800
func (r *ComponentColumn) SetBackgroundStyleBlue800() *ComponentColumn {
	r.BackgroundStyle = ColorBlue800
	return r
}

// SetBackgroundStyleBlue900 set ComponentColumn.BackgroundStyle attribute to ColorBlue900
func (r *ComponentColumn) SetBackgroundStyleBlue900() *ComponentColumn {
	r.BackgroundStyle = ColorBlue900
	return r
}

// SetBackgroundStyleCarmine set ComponentColumn.BackgroundStyle attribute to ColorCarmine
func (r *ComponentColumn) SetBackgroundStyleCarmine() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine
	return r
}

// SetBackgroundStyleCarmine50 set ComponentColumn.BackgroundStyle attribute to ColorCarmine50
func (r *ComponentColumn) SetBackgroundStyleCarmine50() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine50
	return r
}

// SetBackgroundStyleCarmine100 set ComponentColumn.BackgroundStyle attribute to ColorCarmine100
func (r *ComponentColumn) SetBackgroundStyleCarmine100() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine100
	return r
}

// SetBackgroundStyleCarmine200 set ComponentColumn.BackgroundStyle attribute to ColorCarmine200
func (r *ComponentColumn) SetBackgroundStyleCarmine200() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine200
	return r
}

// SetBackgroundStyleCarmine300 set ComponentColumn.BackgroundStyle attribute to ColorCarmine300
func (r *ComponentColumn) SetBackgroundStyleCarmine300() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine300
	return r
}

// SetBackgroundStyleCarmine350 set ComponentColumn.BackgroundStyle attribute to ColorCarmine350
func (r *ComponentColumn) SetBackgroundStyleCarmine350() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine350
	return r
}

// SetBackgroundStyleCarmine400 set ComponentColumn.BackgroundStyle attribute to ColorCarmine400
func (r *ComponentColumn) SetBackgroundStyleCarmine400() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine400
	return r
}

// SetBackgroundStyleCarmine500 set ComponentColumn.BackgroundStyle attribute to ColorCarmine500
func (r *ComponentColumn) SetBackgroundStyleCarmine500() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine500
	return r
}

// SetBackgroundStyleCarmine600 set ComponentColumn.BackgroundStyle attribute to ColorCarmine600
func (r *ComponentColumn) SetBackgroundStyleCarmine600() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine600
	return r
}

// SetBackgroundStyleCarmine700 set ComponentColumn.BackgroundStyle attribute to ColorCarmine700
func (r *ComponentColumn) SetBackgroundStyleCarmine700() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine700
	return r
}

// SetBackgroundStyleCarmine800 set ComponentColumn.BackgroundStyle attribute to ColorCarmine800
func (r *ComponentColumn) SetBackgroundStyleCarmine800() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine800
	return r
}

// SetBackgroundStyleCarmine900 set ComponentColumn.BackgroundStyle attribute to ColorCarmine900
func (r *ComponentColumn) SetBackgroundStyleCarmine900() *ComponentColumn {
	r.BackgroundStyle = ColorCarmine900
	return r
}

// SetBackgroundStyleGreen set ComponentColumn.BackgroundStyle attribute to ColorGreen
func (r *ComponentColumn) SetBackgroundStyleGreen() *ComponentColumn {
	r.BackgroundStyle = ColorGreen
	return r
}

// SetBackgroundStyleGreen50 set ComponentColumn.BackgroundStyle attribute to ColorGreen50
func (r *ComponentColumn) SetBackgroundStyleGreen50() *ComponentColumn {
	r.BackgroundStyle = ColorGreen50
	return r
}

// SetBackgroundStyleGreen100 set ComponentColumn.BackgroundStyle attribute to ColorGreen100
func (r *ComponentColumn) SetBackgroundStyleGreen100() *ComponentColumn {
	r.BackgroundStyle = ColorGreen100
	return r
}

// SetBackgroundStyleGreen200 set ComponentColumn.BackgroundStyle attribute to ColorGreen200
func (r *ComponentColumn) SetBackgroundStyleGreen200() *ComponentColumn {
	r.BackgroundStyle = ColorGreen200
	return r
}

// SetBackgroundStyleGreen300 set ComponentColumn.BackgroundStyle attribute to ColorGreen300
func (r *ComponentColumn) SetBackgroundStyleGreen300() *ComponentColumn {
	r.BackgroundStyle = ColorGreen300
	return r
}

// SetBackgroundStyleGreen350 set ComponentColumn.BackgroundStyle attribute to ColorGreen350
func (r *ComponentColumn) SetBackgroundStyleGreen350() *ComponentColumn {
	r.BackgroundStyle = ColorGreen350
	return r
}

// SetBackgroundStyleGreen400 set ComponentColumn.BackgroundStyle attribute to ColorGreen400
func (r *ComponentColumn) SetBackgroundStyleGreen400() *ComponentColumn {
	r.BackgroundStyle = ColorGreen400
	return r
}

// SetBackgroundStyleGreen500 set ComponentColumn.BackgroundStyle attribute to ColorGreen500
func (r *ComponentColumn) SetBackgroundStyleGreen500() *ComponentColumn {
	r.BackgroundStyle = ColorGreen500
	return r
}

// SetBackgroundStyleGreen600 set ComponentColumn.BackgroundStyle attribute to ColorGreen600
func (r *ComponentColumn) SetBackgroundStyleGreen600() *ComponentColumn {
	r.BackgroundStyle = ColorGreen600
	return r
}

// SetBackgroundStyleGreen700 set ComponentColumn.BackgroundStyle attribute to ColorGreen700
func (r *ComponentColumn) SetBackgroundStyleGreen700() *ComponentColumn {
	r.BackgroundStyle = ColorGreen700
	return r
}

// SetBackgroundStyleGreen800 set ComponentColumn.BackgroundStyle attribute to ColorGreen800
func (r *ComponentColumn) SetBackgroundStyleGreen800() *ComponentColumn {
	r.BackgroundStyle = ColorGreen800
	return r
}

// SetBackgroundStyleGreen900 set ComponentColumn.BackgroundStyle attribute to ColorGreen900
func (r *ComponentColumn) SetBackgroundStyleGreen900() *ComponentColumn {
	r.BackgroundStyle = ColorGreen900
	return r
}

// SetBackgroundStyleIndigo set ComponentColumn.BackgroundStyle attribute to ColorIndigo
func (r *ComponentColumn) SetBackgroundStyleIndigo() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo
	return r
}

// SetBackgroundStyleIndigo50 set ComponentColumn.BackgroundStyle attribute to ColorIndigo50
func (r *ComponentColumn) SetBackgroundStyleIndigo50() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo50
	return r
}

// SetBackgroundStyleIndigo100 set ComponentColumn.BackgroundStyle attribute to ColorIndigo100
func (r *ComponentColumn) SetBackgroundStyleIndigo100() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo100
	return r
}

// SetBackgroundStyleIndigo200 set ComponentColumn.BackgroundStyle attribute to ColorIndigo200
func (r *ComponentColumn) SetBackgroundStyleIndigo200() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo200
	return r
}

// SetBackgroundStyleIndigo300 set ComponentColumn.BackgroundStyle attribute to ColorIndigo300
func (r *ComponentColumn) SetBackgroundStyleIndigo300() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo300
	return r
}

// SetBackgroundStyleIndigo350 set ComponentColumn.BackgroundStyle attribute to ColorIndigo350
func (r *ComponentColumn) SetBackgroundStyleIndigo350() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo350
	return r
}

// SetBackgroundStyleIndigo400 set ComponentColumn.BackgroundStyle attribute to ColorIndigo400
func (r *ComponentColumn) SetBackgroundStyleIndigo400() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo400
	return r
}

// SetBackgroundStyleIndigo500 set ComponentColumn.BackgroundStyle attribute to ColorIndigo500
func (r *ComponentColumn) SetBackgroundStyleIndigo500() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo500
	return r
}

// SetBackgroundStyleIndigo600 set ComponentColumn.BackgroundStyle attribute to ColorIndigo600
func (r *ComponentColumn) SetBackgroundStyleIndigo600() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo600
	return r
}

// SetBackgroundStyleIndigo700 set ComponentColumn.BackgroundStyle attribute to ColorIndigo700
func (r *ComponentColumn) SetBackgroundStyleIndigo700() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo700
	return r
}

// SetBackgroundStyleIndigo800 set ComponentColumn.BackgroundStyle attribute to ColorIndigo800
func (r *ComponentColumn) SetBackgroundStyleIndigo800() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo800
	return r
}

// SetBackgroundStyleIndigo900 set ComponentColumn.BackgroundStyle attribute to ColorIndigo900
func (r *ComponentColumn) SetBackgroundStyleIndigo900() *ComponentColumn {
	r.BackgroundStyle = ColorIndigo900
	return r
}

// SetBackgroundStyleLime set ComponentColumn.BackgroundStyle attribute to ColorLime
func (r *ComponentColumn) SetBackgroundStyleLime() *ComponentColumn {
	r.BackgroundStyle = ColorLime
	return r
}

// SetBackgroundStyleLime50 set ComponentColumn.BackgroundStyle attribute to ColorLime50
func (r *ComponentColumn) SetBackgroundStyleLime50() *ComponentColumn {
	r.BackgroundStyle = ColorLime50
	return r
}

// SetBackgroundStyleLime100 set ComponentColumn.BackgroundStyle attribute to ColorLime100
func (r *ComponentColumn) SetBackgroundStyleLime100() *ComponentColumn {
	r.BackgroundStyle = ColorLime100
	return r
}

// SetBackgroundStyleLime200 set ComponentColumn.BackgroundStyle attribute to ColorLime200
func (r *ComponentColumn) SetBackgroundStyleLime200() *ComponentColumn {
	r.BackgroundStyle = ColorLime200
	return r
}

// SetBackgroundStyleLime300 set ComponentColumn.BackgroundStyle attribute to ColorLime300
func (r *ComponentColumn) SetBackgroundStyleLime300() *ComponentColumn {
	r.BackgroundStyle = ColorLime300
	return r
}

// SetBackgroundStyleLime350 set ComponentColumn.BackgroundStyle attribute to ColorLime350
func (r *ComponentColumn) SetBackgroundStyleLime350() *ComponentColumn {
	r.BackgroundStyle = ColorLime350
	return r
}

// SetBackgroundStyleLime400 set ComponentColumn.BackgroundStyle attribute to ColorLime400
func (r *ComponentColumn) SetBackgroundStyleLime400() *ComponentColumn {
	r.BackgroundStyle = ColorLime400
	return r
}

// SetBackgroundStyleLime500 set ComponentColumn.BackgroundStyle attribute to ColorLime500
func (r *ComponentColumn) SetBackgroundStyleLime500() *ComponentColumn {
	r.BackgroundStyle = ColorLime500
	return r
}

// SetBackgroundStyleLime600 set ComponentColumn.BackgroundStyle attribute to ColorLime600
func (r *ComponentColumn) SetBackgroundStyleLime600() *ComponentColumn {
	r.BackgroundStyle = ColorLime600
	return r
}

// SetBackgroundStyleLime700 set ComponentColumn.BackgroundStyle attribute to ColorLime700
func (r *ComponentColumn) SetBackgroundStyleLime700() *ComponentColumn {
	r.BackgroundStyle = ColorLime700
	return r
}

// SetBackgroundStyleLime800 set ComponentColumn.BackgroundStyle attribute to ColorLime800
func (r *ComponentColumn) SetBackgroundStyleLime800() *ComponentColumn {
	r.BackgroundStyle = ColorLime800
	return r
}

// SetBackgroundStyleLime900 set ComponentColumn.BackgroundStyle attribute to ColorLime900
func (r *ComponentColumn) SetBackgroundStyleLime900() *ComponentColumn {
	r.BackgroundStyle = ColorLime900
	return r
}

// SetBackgroundStyleGrey set ComponentColumn.BackgroundStyle attribute to ColorGrey
func (r *ComponentColumn) SetBackgroundStyleGrey() *ComponentColumn {
	r.BackgroundStyle = ColorGrey
	return r
}

// SetBackgroundStyleGrey00 set ComponentColumn.BackgroundStyle attribute to ColorGrey00
func (r *ComponentColumn) SetBackgroundStyleGrey00() *ComponentColumn {
	r.BackgroundStyle = ColorGrey00
	return r
}

// SetBackgroundStyleGrey50 set ComponentColumn.BackgroundStyle attribute to ColorGrey50
func (r *ComponentColumn) SetBackgroundStyleGrey50() *ComponentColumn {
	r.BackgroundStyle = ColorGrey50
	return r
}

// SetBackgroundStyleGrey100 set ComponentColumn.BackgroundStyle attribute to ColorGrey100
func (r *ComponentColumn) SetBackgroundStyleGrey100() *ComponentColumn {
	r.BackgroundStyle = ColorGrey100
	return r
}

// SetBackgroundStyleGrey200 set ComponentColumn.BackgroundStyle attribute to ColorGrey200
func (r *ComponentColumn) SetBackgroundStyleGrey200() *ComponentColumn {
	r.BackgroundStyle = ColorGrey200
	return r
}

// SetBackgroundStyleGrey300 set ComponentColumn.BackgroundStyle attribute to ColorGrey300
func (r *ComponentColumn) SetBackgroundStyleGrey300() *ComponentColumn {
	r.BackgroundStyle = ColorGrey300
	return r
}

// SetBackgroundStyleGrey350 set ComponentColumn.BackgroundStyle attribute to ColorGrey350
func (r *ComponentColumn) SetBackgroundStyleGrey350() *ComponentColumn {
	r.BackgroundStyle = ColorGrey350
	return r
}

// SetBackgroundStyleGrey400 set ComponentColumn.BackgroundStyle attribute to ColorGrey400
func (r *ComponentColumn) SetBackgroundStyleGrey400() *ComponentColumn {
	r.BackgroundStyle = ColorGrey400
	return r
}

// SetBackgroundStyleGrey500 set ComponentColumn.BackgroundStyle attribute to ColorGrey500
func (r *ComponentColumn) SetBackgroundStyleGrey500() *ComponentColumn {
	r.BackgroundStyle = ColorGrey500
	return r
}

// SetBackgroundStyleGrey600 set ComponentColumn.BackgroundStyle attribute to ColorGrey600
func (r *ComponentColumn) SetBackgroundStyleGrey600() *ComponentColumn {
	r.BackgroundStyle = ColorGrey600
	return r
}

// SetBackgroundStyleGrey650 set ComponentColumn.BackgroundStyle attribute to ColorGrey650
func (r *ComponentColumn) SetBackgroundStyleGrey650() *ComponentColumn {
	r.BackgroundStyle = ColorGrey650
	return r
}

// SetBackgroundStyleGrey700 set ComponentColumn.BackgroundStyle attribute to ColorGrey700
func (r *ComponentColumn) SetBackgroundStyleGrey700() *ComponentColumn {
	r.BackgroundStyle = ColorGrey700
	return r
}

// SetBackgroundStyleGrey800 set ComponentColumn.BackgroundStyle attribute to ColorGrey800
func (r *ComponentColumn) SetBackgroundStyleGrey800() *ComponentColumn {
	r.BackgroundStyle = ColorGrey800
	return r
}

// SetBackgroundStyleGrey900 set ComponentColumn.BackgroundStyle attribute to ColorGrey900
func (r *ComponentColumn) SetBackgroundStyleGrey900() *ComponentColumn {
	r.BackgroundStyle = ColorGrey900
	return r
}

// SetBackgroundStyleGrey950 set ComponentColumn.BackgroundStyle attribute to ColorGrey950
func (r *ComponentColumn) SetBackgroundStyleGrey950() *ComponentColumn {
	r.BackgroundStyle = ColorGrey950
	return r
}

// SetBackgroundStyleGrey1000 set ComponentColumn.BackgroundStyle attribute to ColorGrey1000
func (r *ComponentColumn) SetBackgroundStyleGrey1000() *ComponentColumn {
	r.BackgroundStyle = ColorGrey1000
	return r
}

// SetBackgroundStyleOrange set ComponentColumn.BackgroundStyle attribute to ColorOrange
func (r *ComponentColumn) SetBackgroundStyleOrange() *ComponentColumn {
	r.BackgroundStyle = ColorOrange
	return r
}

// SetBackgroundStyleOrange50 set ComponentColumn.BackgroundStyle attribute to ColorOrange50
func (r *ComponentColumn) SetBackgroundStyleOrange50() *ComponentColumn {
	r.BackgroundStyle = ColorOrange50
	return r
}

// SetBackgroundStyleOrange100 set ComponentColumn.BackgroundStyle attribute to ColorOrange100
func (r *ComponentColumn) SetBackgroundStyleOrange100() *ComponentColumn {
	r.BackgroundStyle = ColorOrange100
	return r
}

// SetBackgroundStyleOrange200 set ComponentColumn.BackgroundStyle attribute to ColorOrange200
func (r *ComponentColumn) SetBackgroundStyleOrange200() *ComponentColumn {
	r.BackgroundStyle = ColorOrange200
	return r
}

// SetBackgroundStyleOrange300 set ComponentColumn.BackgroundStyle attribute to ColorOrange300
func (r *ComponentColumn) SetBackgroundStyleOrange300() *ComponentColumn {
	r.BackgroundStyle = ColorOrange300
	return r
}

// SetBackgroundStyleOrange350 set ComponentColumn.BackgroundStyle attribute to ColorOrange350
func (r *ComponentColumn) SetBackgroundStyleOrange350() *ComponentColumn {
	r.BackgroundStyle = ColorOrange350
	return r
}

// SetBackgroundStyleOrange400 set ComponentColumn.BackgroundStyle attribute to ColorOrange400
func (r *ComponentColumn) SetBackgroundStyleOrange400() *ComponentColumn {
	r.BackgroundStyle = ColorOrange400
	return r
}

// SetBackgroundStyleOrange500 set ComponentColumn.BackgroundStyle attribute to ColorOrange500
func (r *ComponentColumn) SetBackgroundStyleOrange500() *ComponentColumn {
	r.BackgroundStyle = ColorOrange500
	return r
}

// SetBackgroundStyleOrange600 set ComponentColumn.BackgroundStyle attribute to ColorOrange600
func (r *ComponentColumn) SetBackgroundStyleOrange600() *ComponentColumn {
	r.BackgroundStyle = ColorOrange600
	return r
}

// SetBackgroundStyleOrange700 set ComponentColumn.BackgroundStyle attribute to ColorOrange700
func (r *ComponentColumn) SetBackgroundStyleOrange700() *ComponentColumn {
	r.BackgroundStyle = ColorOrange700
	return r
}

// SetBackgroundStyleOrange800 set ComponentColumn.BackgroundStyle attribute to ColorOrange800
func (r *ComponentColumn) SetBackgroundStyleOrange800() *ComponentColumn {
	r.BackgroundStyle = ColorOrange800
	return r
}

// SetBackgroundStyleOrange900 set ComponentColumn.BackgroundStyle attribute to ColorOrange900
func (r *ComponentColumn) SetBackgroundStyleOrange900() *ComponentColumn {
	r.BackgroundStyle = ColorOrange900
	return r
}

// SetBackgroundStylePurple set ComponentColumn.BackgroundStyle attribute to ColorPurple
func (r *ComponentColumn) SetBackgroundStylePurple() *ComponentColumn {
	r.BackgroundStyle = ColorPurple
	return r
}

// SetBackgroundStylePurple50 set ComponentColumn.BackgroundStyle attribute to ColorPurple50
func (r *ComponentColumn) SetBackgroundStylePurple50() *ComponentColumn {
	r.BackgroundStyle = ColorPurple50
	return r
}

// SetBackgroundStylePurple100 set ComponentColumn.BackgroundStyle attribute to ColorPurple100
func (r *ComponentColumn) SetBackgroundStylePurple100() *ComponentColumn {
	r.BackgroundStyle = ColorPurple100
	return r
}

// SetBackgroundStylePurple200 set ComponentColumn.BackgroundStyle attribute to ColorPurple200
func (r *ComponentColumn) SetBackgroundStylePurple200() *ComponentColumn {
	r.BackgroundStyle = ColorPurple200
	return r
}

// SetBackgroundStylePurple300 set ComponentColumn.BackgroundStyle attribute to ColorPurple300
func (r *ComponentColumn) SetBackgroundStylePurple300() *ComponentColumn {
	r.BackgroundStyle = ColorPurple300
	return r
}

// SetBackgroundStylePurple350 set ComponentColumn.BackgroundStyle attribute to ColorPurple350
func (r *ComponentColumn) SetBackgroundStylePurple350() *ComponentColumn {
	r.BackgroundStyle = ColorPurple350
	return r
}

// SetBackgroundStylePurple400 set ComponentColumn.BackgroundStyle attribute to ColorPurple400
func (r *ComponentColumn) SetBackgroundStylePurple400() *ComponentColumn {
	r.BackgroundStyle = ColorPurple400
	return r
}

// SetBackgroundStylePurple500 set ComponentColumn.BackgroundStyle attribute to ColorPurple500
func (r *ComponentColumn) SetBackgroundStylePurple500() *ComponentColumn {
	r.BackgroundStyle = ColorPurple500
	return r
}

// SetBackgroundStylePurple600 set ComponentColumn.BackgroundStyle attribute to ColorPurple600
func (r *ComponentColumn) SetBackgroundStylePurple600() *ComponentColumn {
	r.BackgroundStyle = ColorPurple600
	return r
}

// SetBackgroundStylePurple700 set ComponentColumn.BackgroundStyle attribute to ColorPurple700
func (r *ComponentColumn) SetBackgroundStylePurple700() *ComponentColumn {
	r.BackgroundStyle = ColorPurple700
	return r
}

// SetBackgroundStylePurple800 set ComponentColumn.BackgroundStyle attribute to ColorPurple800
func (r *ComponentColumn) SetBackgroundStylePurple800() *ComponentColumn {
	r.BackgroundStyle = ColorPurple800
	return r
}

// SetBackgroundStylePurple900 set ComponentColumn.BackgroundStyle attribute to ColorPurple900
func (r *ComponentColumn) SetBackgroundStylePurple900() *ComponentColumn {
	r.BackgroundStyle = ColorPurple900
	return r
}

// SetBackgroundStyleRed set ComponentColumn.BackgroundStyle attribute to ColorRed
func (r *ComponentColumn) SetBackgroundStyleRed() *ComponentColumn {
	r.BackgroundStyle = ColorRed
	return r
}

// SetBackgroundStyleRed50 set ComponentColumn.BackgroundStyle attribute to ColorRed50
func (r *ComponentColumn) SetBackgroundStyleRed50() *ComponentColumn {
	r.BackgroundStyle = ColorRed50
	return r
}

// SetBackgroundStyleRed100 set ComponentColumn.BackgroundStyle attribute to ColorRed100
func (r *ComponentColumn) SetBackgroundStyleRed100() *ComponentColumn {
	r.BackgroundStyle = ColorRed100
	return r
}

// SetBackgroundStyleRed200 set ComponentColumn.BackgroundStyle attribute to ColorRed200
func (r *ComponentColumn) SetBackgroundStyleRed200() *ComponentColumn {
	r.BackgroundStyle = ColorRed200
	return r
}

// SetBackgroundStyleRed300 set ComponentColumn.BackgroundStyle attribute to ColorRed300
func (r *ComponentColumn) SetBackgroundStyleRed300() *ComponentColumn {
	r.BackgroundStyle = ColorRed300
	return r
}

// SetBackgroundStyleRed350 set ComponentColumn.BackgroundStyle attribute to ColorRed350
func (r *ComponentColumn) SetBackgroundStyleRed350() *ComponentColumn {
	r.BackgroundStyle = ColorRed350
	return r
}

// SetBackgroundStyleRed400 set ComponentColumn.BackgroundStyle attribute to ColorRed400
func (r *ComponentColumn) SetBackgroundStyleRed400() *ComponentColumn {
	r.BackgroundStyle = ColorRed400
	return r
}

// SetBackgroundStyleRed500 set ComponentColumn.BackgroundStyle attribute to ColorRed500
func (r *ComponentColumn) SetBackgroundStyleRed500() *ComponentColumn {
	r.BackgroundStyle = ColorRed500
	return r
}

// SetBackgroundStyleRed600 set ComponentColumn.BackgroundStyle attribute to ColorRed600
func (r *ComponentColumn) SetBackgroundStyleRed600() *ComponentColumn {
	r.BackgroundStyle = ColorRed600
	return r
}

// SetBackgroundStyleRed700 set ComponentColumn.BackgroundStyle attribute to ColorRed700
func (r *ComponentColumn) SetBackgroundStyleRed700() *ComponentColumn {
	r.BackgroundStyle = ColorRed700
	return r
}

// SetBackgroundStyleRed800 set ComponentColumn.BackgroundStyle attribute to ColorRed800
func (r *ComponentColumn) SetBackgroundStyleRed800() *ComponentColumn {
	r.BackgroundStyle = ColorRed800
	return r
}

// SetBackgroundStyleRed900 set ComponentColumn.BackgroundStyle attribute to ColorRed900
func (r *ComponentColumn) SetBackgroundStyleRed900() *ComponentColumn {
	r.BackgroundStyle = ColorRed900
	return r
}

// SetBackgroundStyleSunflower set ComponentColumn.BackgroundStyle attribute to ColorSunflower
func (r *ComponentColumn) SetBackgroundStyleSunflower() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower
	return r
}

// SetBackgroundStyleSunflower50 set ComponentColumn.BackgroundStyle attribute to ColorSunflower50
func (r *ComponentColumn) SetBackgroundStyleSunflower50() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower50
	return r
}

// SetBackgroundStyleSunflower100 set ComponentColumn.BackgroundStyle attribute to ColorSunflower100
func (r *ComponentColumn) SetBackgroundStyleSunflower100() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower100
	return r
}

// SetBackgroundStyleSunflower200 set ComponentColumn.BackgroundStyle attribute to ColorSunflower200
func (r *ComponentColumn) SetBackgroundStyleSunflower200() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower200
	return r
}

// SetBackgroundStyleSunflower300 set ComponentColumn.BackgroundStyle attribute to ColorSunflower300
func (r *ComponentColumn) SetBackgroundStyleSunflower300() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower300
	return r
}

// SetBackgroundStyleSunflower350 set ComponentColumn.BackgroundStyle attribute to ColorSunflower350
func (r *ComponentColumn) SetBackgroundStyleSunflower350() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower350
	return r
}

// SetBackgroundStyleSunflower400 set ComponentColumn.BackgroundStyle attribute to ColorSunflower400
func (r *ComponentColumn) SetBackgroundStyleSunflower400() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower400
	return r
}

// SetBackgroundStyleSunflower500 set ComponentColumn.BackgroundStyle attribute to ColorSunflower500
func (r *ComponentColumn) SetBackgroundStyleSunflower500() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower500
	return r
}

// SetBackgroundStyleSunflower600 set ComponentColumn.BackgroundStyle attribute to ColorSunflower600
func (r *ComponentColumn) SetBackgroundStyleSunflower600() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower600
	return r
}

// SetBackgroundStyleSunflower700 set ComponentColumn.BackgroundStyle attribute to ColorSunflower700
func (r *ComponentColumn) SetBackgroundStyleSunflower700() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower700
	return r
}

// SetBackgroundStyleSunflower800 set ComponentColumn.BackgroundStyle attribute to ColorSunflower800
func (r *ComponentColumn) SetBackgroundStyleSunflower800() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower800
	return r
}

// SetBackgroundStyleSunflower900 set ComponentColumn.BackgroundStyle attribute to ColorSunflower900
func (r *ComponentColumn) SetBackgroundStyleSunflower900() *ComponentColumn {
	r.BackgroundStyle = ColorSunflower900
	return r
}

// SetBackgroundStyleTurquoise set ComponentColumn.BackgroundStyle attribute to ColorTurquoise
func (r *ComponentColumn) SetBackgroundStyleTurquoise() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise
	return r
}

// SetBackgroundStyleTurquoise50 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise50
func (r *ComponentColumn) SetBackgroundStyleTurquoise50() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise50
	return r
}

// SetBackgroundStyleTurquoise100 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise100
func (r *ComponentColumn) SetBackgroundStyleTurquoise100() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise100
	return r
}

// SetBackgroundStyleTurquoise200 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise200
func (r *ComponentColumn) SetBackgroundStyleTurquoise200() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise200
	return r
}

// SetBackgroundStyleTurquoise300 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise300
func (r *ComponentColumn) SetBackgroundStyleTurquoise300() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise300
	return r
}

// SetBackgroundStyleTurquoise350 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise350
func (r *ComponentColumn) SetBackgroundStyleTurquoise350() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise350
	return r
}

// SetBackgroundStyleTurquoise400 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise400
func (r *ComponentColumn) SetBackgroundStyleTurquoise400() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise400
	return r
}

// SetBackgroundStyleTurquoise500 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise500
func (r *ComponentColumn) SetBackgroundStyleTurquoise500() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise500
	return r
}

// SetBackgroundStyleTurquoise600 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise600
func (r *ComponentColumn) SetBackgroundStyleTurquoise600() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise600
	return r
}

// SetBackgroundStyleTurquoise700 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise700
func (r *ComponentColumn) SetBackgroundStyleTurquoise700() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise700
	return r
}

// SetBackgroundStyleTurquoise800 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise800
func (r *ComponentColumn) SetBackgroundStyleTurquoise800() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise800
	return r
}

// SetBackgroundStyleTurquoise900 set ComponentColumn.BackgroundStyle attribute to ColorTurquoise900
func (r *ComponentColumn) SetBackgroundStyleTurquoise900() *ComponentColumn {
	r.BackgroundStyle = ColorTurquoise900
	return r
}

// SetBackgroundStyleViolet set ComponentColumn.BackgroundStyle attribute to ColorViolet
func (r *ComponentColumn) SetBackgroundStyleViolet() *ComponentColumn {
	r.BackgroundStyle = ColorViolet
	return r
}

// SetBackgroundStyleViolet50 set ComponentColumn.BackgroundStyle attribute to ColorViolet50
func (r *ComponentColumn) SetBackgroundStyleViolet50() *ComponentColumn {
	r.BackgroundStyle = ColorViolet50
	return r
}

// SetBackgroundStyleViolet100 set ComponentColumn.BackgroundStyle attribute to ColorViolet100
func (r *ComponentColumn) SetBackgroundStyleViolet100() *ComponentColumn {
	r.BackgroundStyle = ColorViolet100
	return r
}

// SetBackgroundStyleViolet200 set ComponentColumn.BackgroundStyle attribute to ColorViolet200
func (r *ComponentColumn) SetBackgroundStyleViolet200() *ComponentColumn {
	r.BackgroundStyle = ColorViolet200
	return r
}

// SetBackgroundStyleViolet300 set ComponentColumn.BackgroundStyle attribute to ColorViolet300
func (r *ComponentColumn) SetBackgroundStyleViolet300() *ComponentColumn {
	r.BackgroundStyle = ColorViolet300
	return r
}

// SetBackgroundStyleViolet350 set ComponentColumn.BackgroundStyle attribute to ColorViolet350
func (r *ComponentColumn) SetBackgroundStyleViolet350() *ComponentColumn {
	r.BackgroundStyle = ColorViolet350
	return r
}

// SetBackgroundStyleViolet400 set ComponentColumn.BackgroundStyle attribute to ColorViolet400
func (r *ComponentColumn) SetBackgroundStyleViolet400() *ComponentColumn {
	r.BackgroundStyle = ColorViolet400
	return r
}

// SetBackgroundStyleViolet500 set ComponentColumn.BackgroundStyle attribute to ColorViolet500
func (r *ComponentColumn) SetBackgroundStyleViolet500() *ComponentColumn {
	r.BackgroundStyle = ColorViolet500
	return r
}

// SetBackgroundStyleViolet600 set ComponentColumn.BackgroundStyle attribute to ColorViolet600
func (r *ComponentColumn) SetBackgroundStyleViolet600() *ComponentColumn {
	r.BackgroundStyle = ColorViolet600
	return r
}

// SetBackgroundStyleViolet700 set ComponentColumn.BackgroundStyle attribute to ColorViolet700
func (r *ComponentColumn) SetBackgroundStyleViolet700() *ComponentColumn {
	r.BackgroundStyle = ColorViolet700
	return r
}

// SetBackgroundStyleViolet800 set ComponentColumn.BackgroundStyle attribute to ColorViolet800
func (r *ComponentColumn) SetBackgroundStyleViolet800() *ComponentColumn {
	r.BackgroundStyle = ColorViolet800
	return r
}

// SetBackgroundStyleViolet900 set ComponentColumn.BackgroundStyle attribute to ColorViolet900
func (r *ComponentColumn) SetBackgroundStyleViolet900() *ComponentColumn {
	r.BackgroundStyle = ColorViolet900
	return r
}

// SetBackgroundStyleWathet set ComponentColumn.BackgroundStyle attribute to ColorWathet
func (r *ComponentColumn) SetBackgroundStyleWathet() *ComponentColumn {
	r.BackgroundStyle = ColorWathet
	return r
}

// SetBackgroundStyleWathet50 set ComponentColumn.BackgroundStyle attribute to ColorWathet50
func (r *ComponentColumn) SetBackgroundStyleWathet50() *ComponentColumn {
	r.BackgroundStyle = ColorWathet50
	return r
}

// SetBackgroundStyleWathet100 set ComponentColumn.BackgroundStyle attribute to ColorWathet100
func (r *ComponentColumn) SetBackgroundStyleWathet100() *ComponentColumn {
	r.BackgroundStyle = ColorWathet100
	return r
}

// SetBackgroundStyleWathet200 set ComponentColumn.BackgroundStyle attribute to ColorWathet200
func (r *ComponentColumn) SetBackgroundStyleWathet200() *ComponentColumn {
	r.BackgroundStyle = ColorWathet200
	return r
}

// SetBackgroundStyleWathet300 set ComponentColumn.BackgroundStyle attribute to ColorWathet300
func (r *ComponentColumn) SetBackgroundStyleWathet300() *ComponentColumn {
	r.BackgroundStyle = ColorWathet300
	return r
}

// SetBackgroundStyleWathet350 set ComponentColumn.BackgroundStyle attribute to ColorWathet350
func (r *ComponentColumn) SetBackgroundStyleWathet350() *ComponentColumn {
	r.BackgroundStyle = ColorWathet350
	return r
}

// SetBackgroundStyleWathet400 set ComponentColumn.BackgroundStyle attribute to ColorWathet400
func (r *ComponentColumn) SetBackgroundStyleWathet400() *ComponentColumn {
	r.BackgroundStyle = ColorWathet400
	return r
}

// SetBackgroundStyleWathet500 set ComponentColumn.BackgroundStyle attribute to ColorWathet500
func (r *ComponentColumn) SetBackgroundStyleWathet500() *ComponentColumn {
	r.BackgroundStyle = ColorWathet500
	return r
}

// SetBackgroundStyleWathet600 set ComponentColumn.BackgroundStyle attribute to ColorWathet600
func (r *ComponentColumn) SetBackgroundStyleWathet600() *ComponentColumn {
	r.BackgroundStyle = ColorWathet600
	return r
}

// SetBackgroundStyleWathet700 set ComponentColumn.BackgroundStyle attribute to ColorWathet700
func (r *ComponentColumn) SetBackgroundStyleWathet700() *ComponentColumn {
	r.BackgroundStyle = ColorWathet700
	return r
}

// SetBackgroundStyleWathet800 set ComponentColumn.BackgroundStyle attribute to ColorWathet800
func (r *ComponentColumn) SetBackgroundStyleWathet800() *ComponentColumn {
	r.BackgroundStyle = ColorWathet800
	return r
}

// SetBackgroundStyleWathet900 set ComponentColumn.BackgroundStyle attribute to ColorWathet900
func (r *ComponentColumn) SetBackgroundStyleWathet900() *ComponentColumn {
	r.BackgroundStyle = ColorWathet900
	return r
}

// SetBackgroundStyleYellow set ComponentColumn.BackgroundStyle attribute to ColorYellow
func (r *ComponentColumn) SetBackgroundStyleYellow() *ComponentColumn {
	r.BackgroundStyle = ColorYellow
	return r
}

// SetBackgroundStyleYellow50 set ComponentColumn.BackgroundStyle attribute to ColorYellow50
func (r *ComponentColumn) SetBackgroundStyleYellow50() *ComponentColumn {
	r.BackgroundStyle = ColorYellow50
	return r
}

// SetBackgroundStyleYellow100 set ComponentColumn.BackgroundStyle attribute to ColorYellow100
func (r *ComponentColumn) SetBackgroundStyleYellow100() *ComponentColumn {
	r.BackgroundStyle = ColorYellow100
	return r
}

// SetBackgroundStyleYellow200 set ComponentColumn.BackgroundStyle attribute to ColorYellow200
func (r *ComponentColumn) SetBackgroundStyleYellow200() *ComponentColumn {
	r.BackgroundStyle = ColorYellow200
	return r
}

// SetBackgroundStyleYellow300 set ComponentColumn.BackgroundStyle attribute to ColorYellow300
func (r *ComponentColumn) SetBackgroundStyleYellow300() *ComponentColumn {
	r.BackgroundStyle = ColorYellow300
	return r
}

// SetBackgroundStyleYellow350 set ComponentColumn.BackgroundStyle attribute to ColorYellow350
func (r *ComponentColumn) SetBackgroundStyleYellow350() *ComponentColumn {
	r.BackgroundStyle = ColorYellow350
	return r
}

// SetBackgroundStyleYellow400 set ComponentColumn.BackgroundStyle attribute to ColorYellow400
func (r *ComponentColumn) SetBackgroundStyleYellow400() *ComponentColumn {
	r.BackgroundStyle = ColorYellow400
	return r
}

// SetBackgroundStyleYellow500 set ComponentColumn.BackgroundStyle attribute to ColorYellow500
func (r *ComponentColumn) SetBackgroundStyleYellow500() *ComponentColumn {
	r.BackgroundStyle = ColorYellow500
	return r
}

// SetBackgroundStyleYellow600 set ComponentColumn.BackgroundStyle attribute to ColorYellow600
func (r *ComponentColumn) SetBackgroundStyleYellow600() *ComponentColumn {
	r.BackgroundStyle = ColorYellow600
	return r
}

// SetBackgroundStyleYellow700 set ComponentColumn.BackgroundStyle attribute to ColorYellow700
func (r *ComponentColumn) SetBackgroundStyleYellow700() *ComponentColumn {
	r.BackgroundStyle = ColorYellow700
	return r
}

// SetBackgroundStyleYellow800 set ComponentColumn.BackgroundStyle attribute to ColorYellow800
func (r *ComponentColumn) SetBackgroundStyleYellow800() *ComponentColumn {
	r.BackgroundStyle = ColorYellow800
	return r
}

// SetBackgroundStyleYellow900 set ComponentColumn.BackgroundStyle attribute to ColorYellow900
func (r *ComponentColumn) SetBackgroundStyleYellow900() *ComponentColumn {
	r.BackgroundStyle = ColorYellow900
	return r
}

// SetWidth set ComponentColumn.Width attribute
func (r *ComponentColumn) SetWidth(val ColumnWidth) *ComponentColumn {
	r.Width = val
	return r
}

// SetWidthAuto set ComponentColumn.Width attribute to ColumnWidthAuto
func (r *ComponentColumn) SetWidthAuto() *ComponentColumn {
	r.Width = ColumnWidthAuto
	return r
}

// SetWidthWeighted set ComponentColumn.Width attribute to ColumnWidthWeighted
func (r *ComponentColumn) SetWidthWeighted() *ComponentColumn {
	r.Width = ColumnWidthWeighted
	return r
}

// SetWeight set ComponentColumn.Weight attribute
func (r *ComponentColumn) SetWeight(val int64) *ComponentColumn {
	r.Weight = val
	return r
}

// SetVerticalAlign set ComponentColumn.VerticalAlign attribute
func (r *ComponentColumn) SetVerticalAlign(val VerticalAlign) *ComponentColumn {
	r.VerticalAlign = val
	return r
}

// SetVerticalAlignTop set ComponentColumn.VerticalAlign attribute to VerticalAlignTop
func (r *ComponentColumn) SetVerticalAlignTop() *ComponentColumn {
	r.VerticalAlign = VerticalAlignTop
	return r
}

// SetVerticalAlignCenter set ComponentColumn.VerticalAlign attribute to VerticalAlignCenter
func (r *ComponentColumn) SetVerticalAlignCenter() *ComponentColumn {
	r.VerticalAlign = VerticalAlignCenter
	return r
}

// SetVerticalAlignBottom set ComponentColumn.VerticalAlign attribute to VerticalAlignBottom
func (r *ComponentColumn) SetVerticalAlignBottom() *ComponentColumn {
	r.VerticalAlign = VerticalAlignBottom
	return r
}

// SetVerticalSpacing set ComponentColumn.VerticalSpacing attribute
func (r *ComponentColumn) SetVerticalSpacing(val VerticalSpacing) *ComponentColumn {
	r.VerticalSpacing = val
	return r
}

// SetVerticalSpacingDefault set ComponentColumn.VerticalSpacing attribute to VerticalSpacingDefault
func (r *ComponentColumn) SetVerticalSpacingDefault() *ComponentColumn {
	r.VerticalSpacing = VerticalSpacingDefault
	return r
}

// SetVerticalSpacingMedium set ComponentColumn.VerticalSpacing attribute to VerticalSpacingMedium
func (r *ComponentColumn) SetVerticalSpacingMedium() *ComponentColumn {
	r.VerticalSpacing = VerticalSpacingMedium
	return r
}

// SetVerticalSpacingLarge set ComponentColumn.VerticalSpacing attribute to VerticalSpacingLarge
func (r *ComponentColumn) SetVerticalSpacingLarge() *ComponentColumn {
	r.VerticalSpacing = VerticalSpacingLarge
	return r
}

// SetPadding set ComponentColumn.Padding attribute
func (r *ComponentColumn) SetPadding(val Padding) *ComponentColumn {
	r.Padding = val
	return r
}

// SetElements set ComponentColumn.Elements attribute
func (r *ComponentColumn) SetElements(val ...Component) *ComponentColumn {
	r.Elements = val
	return r
}

// SetAction set ComponentColumn.Action attribute
func (r *ComponentColumn) SetAction(val *ComponentAction) *ComponentColumn {
	r.Action = val
	return r
}
