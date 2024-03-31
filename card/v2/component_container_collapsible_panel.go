package card

import (
	"encoding/json"
)

// CollapsiblePanel 折叠面板
func CollapsiblePanel(elements ...Component) *ComponentCollapsiblePanel {
	return &ComponentCollapsiblePanel{
		Elements: elements,
	}
}

// ComponentCollapsiblePanel 折叠面板
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/containers/collapsible-panel
//
// 折叠面板允许在卡片中折叠次要信息, 如备注、较长文本等, 以突出主要信息。
//
// 注意事项
// 折叠面板仅支持通过撰写卡片 JSON 代码的方式使用, 暂不支持在卡片搭建工具上构建使用。
// 折叠面板支持飞书 V7.9 及以上版本的客户端。在低于该版本的飞书客户端上, 折叠面板的内容将展示为“请升级至最新版本客户端, 以查看内容”的占位图。
// 容器类组件最多支持嵌套五层组件。建议你避免在折叠面板中嵌套多层组件。多层嵌套会压缩内容的展示空间, 影响卡片的展示效果。
//
// 嵌套规则
// 折叠面板不支持内嵌表单容器（form）组件。
//
//go:generate generate_iface_unmarshal -type=ComponentCollapsiblePanel
//go:generate generate_set_attrs -type=ComponentCollapsiblePanel
//go:generate generate_to_map -type=ComponentCollapsiblePanel
type ComponentCollapsiblePanel struct {
	componentBase

	// 面板是否展开。可选值:
	//
	// true: 面板为展开状态
	// false: 面板为折叠状态。默认为折叠状态
	Expanded bool `json:"expanded,omitempty"`

	// 折叠面板的背景色, 默认为透明。枚举值参见颜色枚举值。
	BackgroundColor Color `json:"background_color,omitempty"`

	// 折叠面板的标题设置。
	Header *CollapsiblePanelHeader `json:"header,omitempty"`

	// 边框设置。默认不显示边框。
	Border *Border `json:"border,omitempty"`

	// 面板内元素垂直边距设置。
	VerticalSpacing VerticalSpacing `json:"vertical_spacing,omitempty"`

	// 内容区的内边距。值的取值范围为 [0,28]px。支持填写单值或多值：
	//
	// 单值：如 "4px", 表示组件内四个内边距都为 4px
	// 多值：如 "4px 12px 4px 12px", 表示容器内上、右、下、左的内边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔
	Padding Padding `json:"padding,omitempty"`

	// 各个组件的 JSON 结构。暂不支持表单（form）组件。
	Elements []Component `json:"elements,omitempty" unmarshal:"unmarshalComponent"`
}

// MarshalJSON ...
func (r ComponentCollapsiblePanel) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "collapsible_panel"
	return json.Marshal(m)
}

//go:generate generate_to_map -type=CollapsiblePanelHeader
//go:generate generate_set_attrs -type=CollapsiblePanelHeader
type CollapsiblePanelHeader struct {
	// 标题文本设置。
	Title *ObjectText `json:"title,omitempty"`

	// 折叠面板标题区域的背景颜色设置, 默认为透明色。枚举值参见颜色枚举值。
	//
	// 注意: 如果你未设置此字段, 则折叠面板的标题区域的背景色由 background_color 字段决定。
	BackgroundColor Color `json:"background_color,omitempty"`

	// 标题区域的垂直居中方式。可取值:
	//
	// top: 标题区域垂直居中于面板区域的顶部
	// center: 标题区域垂直居中于面板区域的中间
	// bottom: 标题区域垂直居中于面板区域的底部
	VerticalAlign VerticalAlign `json:"vertical_align,omitempty"`

	// 标题区域的内边距。值的取值范围为 [0,28]px。支持填写单值或多值:
	//
	// 单值: 如 "4px", 表示组件内四个内边距都为 4px
	// 多值: 如 "4px 12px 4px 12px", 表示容器内上、右、下、左的内边距分别为 4px, 12px, 4px, 12px。四个值必填, 使用空格间隔
	Padding Padding `json:"padding,omitempty"`

	// 添加图标作为标题前缀或后缀图标。支持自定义或使用图标库中的图标。示例代码如下:
	Icon *ObjectIcon `json:"icon,omitempty"`

	// 图标的位置。可选值:
	//
	// left: 图标在标题区域最左侧
	// right: 图标在标题区域最右侧
	// follow_text: 图标在文本右侧
	IconPosition string `json:"icon_position,omitempty"`

	// 折叠面板展开时图标旋转的角度, 正值为顺时针, 负值为逆时针。可选值:
	//
	// -180: 逆时针旋转 180 度
	// -90: 逆时针旋转 90 度
	// 90: 顺时针旋转 90 度
	// 180: 顺时针旋转 180 度
	IconExpandedAngle int `json:"icon_expanded_angle,omitempty"`
}

// Border 边框
type Border struct {
	// 边框颜色设置。枚举值参见颜色枚举值。
	Color Color `json:"color,omitempty"`

	// 圆角设置。
	CornerRadius CornerRadius `json:"corner_radius,omitempty"`
}

// unmarshalComponentCollapsiblePanel generated to unmarshal ComponentCollapsiblePanel iface
type unmarshalComponentCollapsiblePanel struct {
	Expanded        bool                    `json:"expanded,omitempty"`
	BackgroundColor Color                   `json:"background_color,omitempty"`
	Header          *CollapsiblePanelHeader `json:"header,omitempty"`
	Border          *Border                 `json:"border,omitempty"`
	VerticalSpacing VerticalSpacing         `json:"vertical_spacing,omitempty"`
	Padding         Padding                 `json:"padding,omitempty"`
	Elements        []json.RawMessage       `json:"elements,omitempty"`
}

// UnmarshalJSON generated to unmarshal ComponentCollapsiblePanel iface
func (r *ComponentCollapsiblePanel) UnmarshalJSON(bs []byte) error {
	obj := new(unmarshalComponentCollapsiblePanel)
	err := json.Unmarshal(bs, obj)
	if err != nil {
		return err
	}
	r.Expanded = obj.Expanded
	r.BackgroundColor = obj.BackgroundColor
	r.Header = obj.Header
	r.Border = obj.Border
	r.VerticalSpacing = obj.VerticalSpacing
	r.Padding = obj.Padding
	r.Elements = make([]Component, len(obj.Elements))
	for i, v := range obj.Elements {
		r.Elements[i], err = unmarshalComponent(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetExpanded set ComponentCollapsiblePanel.Expanded attribute
func (r *ComponentCollapsiblePanel) SetExpanded(val bool) *ComponentCollapsiblePanel {
	r.Expanded = val
	return r
}

// SetBackgroundColor set ComponentCollapsiblePanel.BackgroundColor attribute
func (r *ComponentCollapsiblePanel) SetBackgroundColor(val Color) *ComponentCollapsiblePanel {
	r.BackgroundColor = val
	return r
}

// SetBackgroundColorDefault set ComponentCollapsiblePanel.BackgroundColor attribute to ColorDefault
func (r *ComponentCollapsiblePanel) SetBackgroundColorDefault() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorDefault
	return r
}

// SetBackgroundColorBgWhite set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBgWhite
func (r *ComponentCollapsiblePanel) SetBackgroundColorBgWhite() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBgWhite
	return r
}

// SetBackgroundColorWhite set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWhite
func (r *ComponentCollapsiblePanel) SetBackgroundColorWhite() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWhite
	return r
}

// SetBackgroundColorBlue set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue
	return r
}

// SetBackgroundColorBlue50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue50
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue50
	return r
}

// SetBackgroundColorBlue100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue100
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue100
	return r
}

// SetBackgroundColorBlue200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue200
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue200
	return r
}

// SetBackgroundColorBlue300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue300
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue300
	return r
}

// SetBackgroundColorBlue350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue350
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue350
	return r
}

// SetBackgroundColorBlue400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue400
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue400
	return r
}

// SetBackgroundColorBlue500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue500
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue500
	return r
}

// SetBackgroundColorBlue600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue600
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue600
	return r
}

// SetBackgroundColorBlue700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue700
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue700
	return r
}

// SetBackgroundColorBlue800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue800
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue800
	return r
}

// SetBackgroundColorBlue900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorBlue900
func (r *ComponentCollapsiblePanel) SetBackgroundColorBlue900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorBlue900
	return r
}

// SetBackgroundColorCarmine set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine
	return r
}

// SetBackgroundColorCarmine50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine50
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine50
	return r
}

// SetBackgroundColorCarmine100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine100
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine100
	return r
}

// SetBackgroundColorCarmine200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine200
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine200
	return r
}

// SetBackgroundColorCarmine300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine300
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine300
	return r
}

// SetBackgroundColorCarmine350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine350
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine350
	return r
}

// SetBackgroundColorCarmine400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine400
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine400
	return r
}

// SetBackgroundColorCarmine500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine500
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine500
	return r
}

// SetBackgroundColorCarmine600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine600
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine600
	return r
}

// SetBackgroundColorCarmine700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine700
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine700
	return r
}

// SetBackgroundColorCarmine800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine800
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine800
	return r
}

// SetBackgroundColorCarmine900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorCarmine900
func (r *ComponentCollapsiblePanel) SetBackgroundColorCarmine900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorCarmine900
	return r
}

// SetBackgroundColorGreen set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen
	return r
}

// SetBackgroundColorGreen50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen50
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen50
	return r
}

// SetBackgroundColorGreen100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen100
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen100
	return r
}

// SetBackgroundColorGreen200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen200
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen200
	return r
}

// SetBackgroundColorGreen300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen300
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen300
	return r
}

// SetBackgroundColorGreen350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen350
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen350
	return r
}

// SetBackgroundColorGreen400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen400
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen400
	return r
}

// SetBackgroundColorGreen500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen500
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen500
	return r
}

// SetBackgroundColorGreen600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen600
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen600
	return r
}

// SetBackgroundColorGreen700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen700
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen700
	return r
}

// SetBackgroundColorGreen800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen800
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen800
	return r
}

// SetBackgroundColorGreen900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGreen900
func (r *ComponentCollapsiblePanel) SetBackgroundColorGreen900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGreen900
	return r
}

// SetBackgroundColorIndigo set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo
	return r
}

// SetBackgroundColorIndigo50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo50
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo50
	return r
}

// SetBackgroundColorIndigo100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo100
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo100
	return r
}

// SetBackgroundColorIndigo200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo200
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo200
	return r
}

// SetBackgroundColorIndigo300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo300
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo300
	return r
}

// SetBackgroundColorIndigo350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo350
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo350
	return r
}

// SetBackgroundColorIndigo400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo400
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo400
	return r
}

// SetBackgroundColorIndigo500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo500
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo500
	return r
}

// SetBackgroundColorIndigo600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo600
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo600
	return r
}

// SetBackgroundColorIndigo700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo700
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo700
	return r
}

// SetBackgroundColorIndigo800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo800
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo800
	return r
}

// SetBackgroundColorIndigo900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorIndigo900
func (r *ComponentCollapsiblePanel) SetBackgroundColorIndigo900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorIndigo900
	return r
}

// SetBackgroundColorLime set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime
	return r
}

// SetBackgroundColorLime50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime50
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime50
	return r
}

// SetBackgroundColorLime100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime100
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime100
	return r
}

// SetBackgroundColorLime200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime200
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime200
	return r
}

// SetBackgroundColorLime300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime300
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime300
	return r
}

// SetBackgroundColorLime350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime350
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime350
	return r
}

// SetBackgroundColorLime400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime400
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime400
	return r
}

// SetBackgroundColorLime500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime500
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime500
	return r
}

// SetBackgroundColorLime600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime600
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime600
	return r
}

// SetBackgroundColorLime700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime700
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime700
	return r
}

// SetBackgroundColorLime800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime800
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime800
	return r
}

// SetBackgroundColorLime900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorLime900
func (r *ComponentCollapsiblePanel) SetBackgroundColorLime900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorLime900
	return r
}

// SetBackgroundColorGrey set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey
	return r
}

// SetBackgroundColorGrey00 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey00
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey00() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey00
	return r
}

// SetBackgroundColorGrey50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey50
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey50
	return r
}

// SetBackgroundColorGrey100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey100
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey100
	return r
}

// SetBackgroundColorGrey200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey200
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey200
	return r
}

// SetBackgroundColorGrey300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey300
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey300
	return r
}

// SetBackgroundColorGrey350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey350
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey350
	return r
}

// SetBackgroundColorGrey400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey400
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey400
	return r
}

// SetBackgroundColorGrey500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey500
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey500
	return r
}

// SetBackgroundColorGrey600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey600
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey600
	return r
}

// SetBackgroundColorGrey650 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey650
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey650() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey650
	return r
}

// SetBackgroundColorGrey700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey700
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey700
	return r
}

// SetBackgroundColorGrey800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey800
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey800
	return r
}

// SetBackgroundColorGrey900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey900
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey900
	return r
}

// SetBackgroundColorGrey950 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey950
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey950() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey950
	return r
}

// SetBackgroundColorGrey1000 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorGrey1000
func (r *ComponentCollapsiblePanel) SetBackgroundColorGrey1000() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorGrey1000
	return r
}

// SetBackgroundColorOrange set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange
	return r
}

// SetBackgroundColorOrange50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange50
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange50
	return r
}

// SetBackgroundColorOrange100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange100
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange100
	return r
}

// SetBackgroundColorOrange200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange200
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange200
	return r
}

// SetBackgroundColorOrange300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange300
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange300
	return r
}

// SetBackgroundColorOrange350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange350
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange350
	return r
}

// SetBackgroundColorOrange400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange400
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange400
	return r
}

// SetBackgroundColorOrange500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange500
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange500
	return r
}

// SetBackgroundColorOrange600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange600
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange600
	return r
}

// SetBackgroundColorOrange700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange700
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange700
	return r
}

// SetBackgroundColorOrange800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange800
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange800
	return r
}

// SetBackgroundColorOrange900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorOrange900
func (r *ComponentCollapsiblePanel) SetBackgroundColorOrange900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorOrange900
	return r
}

// SetBackgroundColorPurple set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple
	return r
}

// SetBackgroundColorPurple50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple50
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple50
	return r
}

// SetBackgroundColorPurple100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple100
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple100
	return r
}

// SetBackgroundColorPurple200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple200
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple200
	return r
}

// SetBackgroundColorPurple300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple300
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple300
	return r
}

// SetBackgroundColorPurple350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple350
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple350
	return r
}

// SetBackgroundColorPurple400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple400
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple400
	return r
}

// SetBackgroundColorPurple500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple500
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple500
	return r
}

// SetBackgroundColorPurple600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple600
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple600
	return r
}

// SetBackgroundColorPurple700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple700
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple700
	return r
}

// SetBackgroundColorPurple800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple800
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple800
	return r
}

// SetBackgroundColorPurple900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorPurple900
func (r *ComponentCollapsiblePanel) SetBackgroundColorPurple900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorPurple900
	return r
}

// SetBackgroundColorRed set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed
	return r
}

// SetBackgroundColorRed50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed50
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed50
	return r
}

// SetBackgroundColorRed100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed100
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed100
	return r
}

// SetBackgroundColorRed200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed200
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed200
	return r
}

// SetBackgroundColorRed300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed300
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed300
	return r
}

// SetBackgroundColorRed350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed350
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed350
	return r
}

// SetBackgroundColorRed400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed400
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed400
	return r
}

// SetBackgroundColorRed500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed500
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed500
	return r
}

// SetBackgroundColorRed600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed600
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed600
	return r
}

// SetBackgroundColorRed700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed700
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed700
	return r
}

// SetBackgroundColorRed800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed800
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed800
	return r
}

// SetBackgroundColorRed900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorRed900
func (r *ComponentCollapsiblePanel) SetBackgroundColorRed900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorRed900
	return r
}

// SetBackgroundColorSunflower set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower
	return r
}

// SetBackgroundColorSunflower50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower50
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower50
	return r
}

// SetBackgroundColorSunflower100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower100
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower100
	return r
}

// SetBackgroundColorSunflower200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower200
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower200
	return r
}

// SetBackgroundColorSunflower300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower300
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower300
	return r
}

// SetBackgroundColorSunflower350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower350
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower350
	return r
}

// SetBackgroundColorSunflower400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower400
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower400
	return r
}

// SetBackgroundColorSunflower500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower500
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower500
	return r
}

// SetBackgroundColorSunflower600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower600
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower600
	return r
}

// SetBackgroundColorSunflower700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower700
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower700
	return r
}

// SetBackgroundColorSunflower800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower800
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower800
	return r
}

// SetBackgroundColorSunflower900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorSunflower900
func (r *ComponentCollapsiblePanel) SetBackgroundColorSunflower900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorSunflower900
	return r
}

// SetBackgroundColorTurquoise set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise
	return r
}

// SetBackgroundColorTurquoise50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise50
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise50
	return r
}

// SetBackgroundColorTurquoise100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise100
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise100
	return r
}

// SetBackgroundColorTurquoise200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise200
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise200
	return r
}

// SetBackgroundColorTurquoise300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise300
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise300
	return r
}

// SetBackgroundColorTurquoise350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise350
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise350
	return r
}

// SetBackgroundColorTurquoise400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise400
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise400
	return r
}

// SetBackgroundColorTurquoise500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise500
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise500
	return r
}

// SetBackgroundColorTurquoise600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise600
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise600
	return r
}

// SetBackgroundColorTurquoise700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise700
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise700
	return r
}

// SetBackgroundColorTurquoise800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise800
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise800
	return r
}

// SetBackgroundColorTurquoise900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorTurquoise900
func (r *ComponentCollapsiblePanel) SetBackgroundColorTurquoise900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorTurquoise900
	return r
}

// SetBackgroundColorViolet set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet
	return r
}

// SetBackgroundColorViolet50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet50
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet50
	return r
}

// SetBackgroundColorViolet100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet100
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet100
	return r
}

// SetBackgroundColorViolet200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet200
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet200
	return r
}

// SetBackgroundColorViolet300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet300
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet300
	return r
}

// SetBackgroundColorViolet350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet350
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet350
	return r
}

// SetBackgroundColorViolet400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet400
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet400
	return r
}

// SetBackgroundColorViolet500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet500
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet500
	return r
}

// SetBackgroundColorViolet600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet600
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet600
	return r
}

// SetBackgroundColorViolet700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet700
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet700
	return r
}

// SetBackgroundColorViolet800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet800
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet800
	return r
}

// SetBackgroundColorViolet900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorViolet900
func (r *ComponentCollapsiblePanel) SetBackgroundColorViolet900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorViolet900
	return r
}

// SetBackgroundColorWathet set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet
	return r
}

// SetBackgroundColorWathet50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet50
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet50
	return r
}

// SetBackgroundColorWathet100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet100
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet100
	return r
}

// SetBackgroundColorWathet200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet200
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet200
	return r
}

// SetBackgroundColorWathet300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet300
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet300
	return r
}

// SetBackgroundColorWathet350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet350
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet350
	return r
}

// SetBackgroundColorWathet400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet400
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet400
	return r
}

// SetBackgroundColorWathet500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet500
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet500
	return r
}

// SetBackgroundColorWathet600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet600
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet600
	return r
}

// SetBackgroundColorWathet700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet700
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet700
	return r
}

// SetBackgroundColorWathet800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet800
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet800
	return r
}

// SetBackgroundColorWathet900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorWathet900
func (r *ComponentCollapsiblePanel) SetBackgroundColorWathet900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorWathet900
	return r
}

// SetBackgroundColorYellow set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow
	return r
}

// SetBackgroundColorYellow50 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow50
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow50() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow50
	return r
}

// SetBackgroundColorYellow100 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow100
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow100() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow100
	return r
}

// SetBackgroundColorYellow200 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow200
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow200() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow200
	return r
}

// SetBackgroundColorYellow300 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow300
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow300() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow300
	return r
}

// SetBackgroundColorYellow350 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow350
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow350() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow350
	return r
}

// SetBackgroundColorYellow400 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow400
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow400() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow400
	return r
}

// SetBackgroundColorYellow500 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow500
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow500() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow500
	return r
}

// SetBackgroundColorYellow600 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow600
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow600() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow600
	return r
}

// SetBackgroundColorYellow700 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow700
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow700() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow700
	return r
}

// SetBackgroundColorYellow800 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow800
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow800() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow800
	return r
}

// SetBackgroundColorYellow900 set ComponentCollapsiblePanel.BackgroundColor attribute to ColorYellow900
func (r *ComponentCollapsiblePanel) SetBackgroundColorYellow900() *ComponentCollapsiblePanel {
	r.BackgroundColor = ColorYellow900
	return r
}

// SetHeader set ComponentCollapsiblePanel.Header attribute
func (r *ComponentCollapsiblePanel) SetHeader(val *CollapsiblePanelHeader) *ComponentCollapsiblePanel {
	r.Header = val
	return r
}

// SetBorder set ComponentCollapsiblePanel.Border attribute
func (r *ComponentCollapsiblePanel) SetBorder(val *Border) *ComponentCollapsiblePanel {
	r.Border = val
	return r
}

// SetVerticalSpacing set ComponentCollapsiblePanel.VerticalSpacing attribute
func (r *ComponentCollapsiblePanel) SetVerticalSpacing(val VerticalSpacing) *ComponentCollapsiblePanel {
	r.VerticalSpacing = val
	return r
}

// SetVerticalSpacingDefault set ComponentCollapsiblePanel.VerticalSpacing attribute to VerticalSpacingDefault
func (r *ComponentCollapsiblePanel) SetVerticalSpacingDefault() *ComponentCollapsiblePanel {
	r.VerticalSpacing = VerticalSpacingDefault
	return r
}

// SetVerticalSpacingMedium set ComponentCollapsiblePanel.VerticalSpacing attribute to VerticalSpacingMedium
func (r *ComponentCollapsiblePanel) SetVerticalSpacingMedium() *ComponentCollapsiblePanel {
	r.VerticalSpacing = VerticalSpacingMedium
	return r
}

// SetVerticalSpacingLarge set ComponentCollapsiblePanel.VerticalSpacing attribute to VerticalSpacingLarge
func (r *ComponentCollapsiblePanel) SetVerticalSpacingLarge() *ComponentCollapsiblePanel {
	r.VerticalSpacing = VerticalSpacingLarge
	return r
}

// SetPadding set ComponentCollapsiblePanel.Padding attribute
func (r *ComponentCollapsiblePanel) SetPadding(val Padding) *ComponentCollapsiblePanel {
	r.Padding = val
	return r
}

// SetElements set ComponentCollapsiblePanel.Elements attribute
func (r *ComponentCollapsiblePanel) SetElements(val ...Component) *ComponentCollapsiblePanel {
	r.Elements = val
	return r
}

// toMap conv ComponentCollapsiblePanel to map
func (r *ComponentCollapsiblePanel) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 7)
	if r.Expanded != false {
		res["expanded"] = r.Expanded
	}
	if r.BackgroundColor != "" {
		res["background_color"] = r.BackgroundColor
	}
	if r.Header != nil {
		res["header"] = r.Header
	}
	if r.Border != nil {
		res["border"] = r.Border
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
	return res
}

// toMap conv CollapsiblePanelHeader to map
func (r *CollapsiblePanelHeader) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 7)
	if r.Title != nil {
		res["title"] = r.Title
	}
	if r.BackgroundColor != "" {
		res["background_color"] = r.BackgroundColor
	}
	if r.VerticalAlign != "" {
		res["vertical_align"] = r.VerticalAlign
	}
	if r.Padding != "" {
		res["padding"] = r.Padding
	}
	if r.Icon != nil {
		res["icon"] = r.Icon
	}
	if r.IconPosition != "" {
		res["icon_position"] = r.IconPosition
	}
	if r.IconExpandedAngle != 0 {
		res["icon_expanded_angle"] = r.IconExpandedAngle
	}
	return res
}

// SetTitle set CollapsiblePanelHeader.Title attribute
func (r *CollapsiblePanelHeader) SetTitle(val *ObjectText) *CollapsiblePanelHeader {
	r.Title = val
	return r
}

// SetBackgroundColor set CollapsiblePanelHeader.BackgroundColor attribute
func (r *CollapsiblePanelHeader) SetBackgroundColor(val Color) *CollapsiblePanelHeader {
	r.BackgroundColor = val
	return r
}

// SetBackgroundColorDefault set CollapsiblePanelHeader.BackgroundColor attribute to ColorDefault
func (r *CollapsiblePanelHeader) SetBackgroundColorDefault() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorDefault
	return r
}

// SetBackgroundColorBgWhite set CollapsiblePanelHeader.BackgroundColor attribute to ColorBgWhite
func (r *CollapsiblePanelHeader) SetBackgroundColorBgWhite() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBgWhite
	return r
}

// SetBackgroundColorWhite set CollapsiblePanelHeader.BackgroundColor attribute to ColorWhite
func (r *CollapsiblePanelHeader) SetBackgroundColorWhite() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWhite
	return r
}

// SetBackgroundColorBlue set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue
	return r
}

// SetBackgroundColorBlue50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue50
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue50
	return r
}

// SetBackgroundColorBlue100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue100
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue100
	return r
}

// SetBackgroundColorBlue200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue200
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue200
	return r
}

// SetBackgroundColorBlue300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue300
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue300
	return r
}

// SetBackgroundColorBlue350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue350
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue350
	return r
}

// SetBackgroundColorBlue400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue400
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue400
	return r
}

// SetBackgroundColorBlue500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue500
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue500
	return r
}

// SetBackgroundColorBlue600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue600
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue600
	return r
}

// SetBackgroundColorBlue700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue700
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue700
	return r
}

// SetBackgroundColorBlue800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue800
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue800
	return r
}

// SetBackgroundColorBlue900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorBlue900
func (r *CollapsiblePanelHeader) SetBackgroundColorBlue900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorBlue900
	return r
}

// SetBackgroundColorCarmine set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine
	return r
}

// SetBackgroundColorCarmine50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine50
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine50
	return r
}

// SetBackgroundColorCarmine100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine100
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine100
	return r
}

// SetBackgroundColorCarmine200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine200
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine200
	return r
}

// SetBackgroundColorCarmine300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine300
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine300
	return r
}

// SetBackgroundColorCarmine350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine350
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine350
	return r
}

// SetBackgroundColorCarmine400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine400
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine400
	return r
}

// SetBackgroundColorCarmine500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine500
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine500
	return r
}

// SetBackgroundColorCarmine600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine600
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine600
	return r
}

// SetBackgroundColorCarmine700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine700
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine700
	return r
}

// SetBackgroundColorCarmine800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine800
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine800
	return r
}

// SetBackgroundColorCarmine900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorCarmine900
func (r *CollapsiblePanelHeader) SetBackgroundColorCarmine900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorCarmine900
	return r
}

// SetBackgroundColorGreen set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen
	return r
}

// SetBackgroundColorGreen50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen50
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen50
	return r
}

// SetBackgroundColorGreen100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen100
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen100
	return r
}

// SetBackgroundColorGreen200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen200
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen200
	return r
}

// SetBackgroundColorGreen300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen300
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen300
	return r
}

// SetBackgroundColorGreen350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen350
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen350
	return r
}

// SetBackgroundColorGreen400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen400
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen400
	return r
}

// SetBackgroundColorGreen500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen500
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen500
	return r
}

// SetBackgroundColorGreen600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen600
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen600
	return r
}

// SetBackgroundColorGreen700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen700
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen700
	return r
}

// SetBackgroundColorGreen800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen800
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen800
	return r
}

// SetBackgroundColorGreen900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGreen900
func (r *CollapsiblePanelHeader) SetBackgroundColorGreen900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGreen900
	return r
}

// SetBackgroundColorIndigo set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo
	return r
}

// SetBackgroundColorIndigo50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo50
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo50
	return r
}

// SetBackgroundColorIndigo100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo100
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo100
	return r
}

// SetBackgroundColorIndigo200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo200
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo200
	return r
}

// SetBackgroundColorIndigo300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo300
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo300
	return r
}

// SetBackgroundColorIndigo350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo350
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo350
	return r
}

// SetBackgroundColorIndigo400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo400
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo400
	return r
}

// SetBackgroundColorIndigo500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo500
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo500
	return r
}

// SetBackgroundColorIndigo600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo600
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo600
	return r
}

// SetBackgroundColorIndigo700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo700
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo700
	return r
}

// SetBackgroundColorIndigo800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo800
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo800
	return r
}

// SetBackgroundColorIndigo900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorIndigo900
func (r *CollapsiblePanelHeader) SetBackgroundColorIndigo900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorIndigo900
	return r
}

// SetBackgroundColorLime set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime
func (r *CollapsiblePanelHeader) SetBackgroundColorLime() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime
	return r
}

// SetBackgroundColorLime50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime50
func (r *CollapsiblePanelHeader) SetBackgroundColorLime50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime50
	return r
}

// SetBackgroundColorLime100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime100
func (r *CollapsiblePanelHeader) SetBackgroundColorLime100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime100
	return r
}

// SetBackgroundColorLime200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime200
func (r *CollapsiblePanelHeader) SetBackgroundColorLime200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime200
	return r
}

// SetBackgroundColorLime300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime300
func (r *CollapsiblePanelHeader) SetBackgroundColorLime300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime300
	return r
}

// SetBackgroundColorLime350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime350
func (r *CollapsiblePanelHeader) SetBackgroundColorLime350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime350
	return r
}

// SetBackgroundColorLime400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime400
func (r *CollapsiblePanelHeader) SetBackgroundColorLime400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime400
	return r
}

// SetBackgroundColorLime500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime500
func (r *CollapsiblePanelHeader) SetBackgroundColorLime500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime500
	return r
}

// SetBackgroundColorLime600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime600
func (r *CollapsiblePanelHeader) SetBackgroundColorLime600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime600
	return r
}

// SetBackgroundColorLime700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime700
func (r *CollapsiblePanelHeader) SetBackgroundColorLime700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime700
	return r
}

// SetBackgroundColorLime800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime800
func (r *CollapsiblePanelHeader) SetBackgroundColorLime800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime800
	return r
}

// SetBackgroundColorLime900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorLime900
func (r *CollapsiblePanelHeader) SetBackgroundColorLime900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorLime900
	return r
}

// SetBackgroundColorGrey set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey
	return r
}

// SetBackgroundColorGrey00 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey00
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey00() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey00
	return r
}

// SetBackgroundColorGrey50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey50
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey50
	return r
}

// SetBackgroundColorGrey100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey100
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey100
	return r
}

// SetBackgroundColorGrey200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey200
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey200
	return r
}

// SetBackgroundColorGrey300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey300
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey300
	return r
}

// SetBackgroundColorGrey350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey350
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey350
	return r
}

// SetBackgroundColorGrey400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey400
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey400
	return r
}

// SetBackgroundColorGrey500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey500
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey500
	return r
}

// SetBackgroundColorGrey600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey600
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey600
	return r
}

// SetBackgroundColorGrey650 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey650
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey650() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey650
	return r
}

// SetBackgroundColorGrey700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey700
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey700
	return r
}

// SetBackgroundColorGrey800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey800
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey800
	return r
}

// SetBackgroundColorGrey900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey900
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey900
	return r
}

// SetBackgroundColorGrey950 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey950
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey950() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey950
	return r
}

// SetBackgroundColorGrey1000 set CollapsiblePanelHeader.BackgroundColor attribute to ColorGrey1000
func (r *CollapsiblePanelHeader) SetBackgroundColorGrey1000() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorGrey1000
	return r
}

// SetBackgroundColorOrange set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange
	return r
}

// SetBackgroundColorOrange50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange50
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange50
	return r
}

// SetBackgroundColorOrange100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange100
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange100
	return r
}

// SetBackgroundColorOrange200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange200
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange200
	return r
}

// SetBackgroundColorOrange300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange300
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange300
	return r
}

// SetBackgroundColorOrange350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange350
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange350
	return r
}

// SetBackgroundColorOrange400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange400
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange400
	return r
}

// SetBackgroundColorOrange500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange500
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange500
	return r
}

// SetBackgroundColorOrange600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange600
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange600
	return r
}

// SetBackgroundColorOrange700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange700
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange700
	return r
}

// SetBackgroundColorOrange800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange800
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange800
	return r
}

// SetBackgroundColorOrange900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorOrange900
func (r *CollapsiblePanelHeader) SetBackgroundColorOrange900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorOrange900
	return r
}

// SetBackgroundColorPurple set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple
	return r
}

// SetBackgroundColorPurple50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple50
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple50
	return r
}

// SetBackgroundColorPurple100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple100
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple100
	return r
}

// SetBackgroundColorPurple200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple200
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple200
	return r
}

// SetBackgroundColorPurple300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple300
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple300
	return r
}

// SetBackgroundColorPurple350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple350
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple350
	return r
}

// SetBackgroundColorPurple400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple400
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple400
	return r
}

// SetBackgroundColorPurple500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple500
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple500
	return r
}

// SetBackgroundColorPurple600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple600
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple600
	return r
}

// SetBackgroundColorPurple700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple700
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple700
	return r
}

// SetBackgroundColorPurple800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple800
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple800
	return r
}

// SetBackgroundColorPurple900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorPurple900
func (r *CollapsiblePanelHeader) SetBackgroundColorPurple900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorPurple900
	return r
}

// SetBackgroundColorRed set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed
func (r *CollapsiblePanelHeader) SetBackgroundColorRed() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed
	return r
}

// SetBackgroundColorRed50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed50
func (r *CollapsiblePanelHeader) SetBackgroundColorRed50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed50
	return r
}

// SetBackgroundColorRed100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed100
func (r *CollapsiblePanelHeader) SetBackgroundColorRed100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed100
	return r
}

// SetBackgroundColorRed200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed200
func (r *CollapsiblePanelHeader) SetBackgroundColorRed200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed200
	return r
}

// SetBackgroundColorRed300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed300
func (r *CollapsiblePanelHeader) SetBackgroundColorRed300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed300
	return r
}

// SetBackgroundColorRed350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed350
func (r *CollapsiblePanelHeader) SetBackgroundColorRed350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed350
	return r
}

// SetBackgroundColorRed400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed400
func (r *CollapsiblePanelHeader) SetBackgroundColorRed400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed400
	return r
}

// SetBackgroundColorRed500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed500
func (r *CollapsiblePanelHeader) SetBackgroundColorRed500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed500
	return r
}

// SetBackgroundColorRed600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed600
func (r *CollapsiblePanelHeader) SetBackgroundColorRed600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed600
	return r
}

// SetBackgroundColorRed700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed700
func (r *CollapsiblePanelHeader) SetBackgroundColorRed700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed700
	return r
}

// SetBackgroundColorRed800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed800
func (r *CollapsiblePanelHeader) SetBackgroundColorRed800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed800
	return r
}

// SetBackgroundColorRed900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorRed900
func (r *CollapsiblePanelHeader) SetBackgroundColorRed900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorRed900
	return r
}

// SetBackgroundColorSunflower set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower
	return r
}

// SetBackgroundColorSunflower50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower50
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower50
	return r
}

// SetBackgroundColorSunflower100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower100
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower100
	return r
}

// SetBackgroundColorSunflower200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower200
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower200
	return r
}

// SetBackgroundColorSunflower300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower300
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower300
	return r
}

// SetBackgroundColorSunflower350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower350
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower350
	return r
}

// SetBackgroundColorSunflower400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower400
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower400
	return r
}

// SetBackgroundColorSunflower500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower500
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower500
	return r
}

// SetBackgroundColorSunflower600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower600
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower600
	return r
}

// SetBackgroundColorSunflower700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower700
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower700
	return r
}

// SetBackgroundColorSunflower800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower800
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower800
	return r
}

// SetBackgroundColorSunflower900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorSunflower900
func (r *CollapsiblePanelHeader) SetBackgroundColorSunflower900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorSunflower900
	return r
}

// SetBackgroundColorTurquoise set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise
	return r
}

// SetBackgroundColorTurquoise50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise50
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise50
	return r
}

// SetBackgroundColorTurquoise100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise100
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise100
	return r
}

// SetBackgroundColorTurquoise200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise200
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise200
	return r
}

// SetBackgroundColorTurquoise300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise300
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise300
	return r
}

// SetBackgroundColorTurquoise350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise350
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise350
	return r
}

// SetBackgroundColorTurquoise400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise400
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise400
	return r
}

// SetBackgroundColorTurquoise500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise500
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise500
	return r
}

// SetBackgroundColorTurquoise600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise600
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise600
	return r
}

// SetBackgroundColorTurquoise700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise700
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise700
	return r
}

// SetBackgroundColorTurquoise800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise800
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise800
	return r
}

// SetBackgroundColorTurquoise900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorTurquoise900
func (r *CollapsiblePanelHeader) SetBackgroundColorTurquoise900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorTurquoise900
	return r
}

// SetBackgroundColorViolet set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet
	return r
}

// SetBackgroundColorViolet50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet50
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet50
	return r
}

// SetBackgroundColorViolet100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet100
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet100
	return r
}

// SetBackgroundColorViolet200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet200
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet200
	return r
}

// SetBackgroundColorViolet300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet300
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet300
	return r
}

// SetBackgroundColorViolet350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet350
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet350
	return r
}

// SetBackgroundColorViolet400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet400
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet400
	return r
}

// SetBackgroundColorViolet500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet500
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet500
	return r
}

// SetBackgroundColorViolet600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet600
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet600
	return r
}

// SetBackgroundColorViolet700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet700
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet700
	return r
}

// SetBackgroundColorViolet800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet800
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet800
	return r
}

// SetBackgroundColorViolet900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorViolet900
func (r *CollapsiblePanelHeader) SetBackgroundColorViolet900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorViolet900
	return r
}

// SetBackgroundColorWathet set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet
	return r
}

// SetBackgroundColorWathet50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet50
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet50
	return r
}

// SetBackgroundColorWathet100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet100
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet100
	return r
}

// SetBackgroundColorWathet200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet200
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet200
	return r
}

// SetBackgroundColorWathet300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet300
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet300
	return r
}

// SetBackgroundColorWathet350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet350
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet350
	return r
}

// SetBackgroundColorWathet400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet400
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet400
	return r
}

// SetBackgroundColorWathet500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet500
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet500
	return r
}

// SetBackgroundColorWathet600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet600
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet600
	return r
}

// SetBackgroundColorWathet700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet700
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet700
	return r
}

// SetBackgroundColorWathet800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet800
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet800
	return r
}

// SetBackgroundColorWathet900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorWathet900
func (r *CollapsiblePanelHeader) SetBackgroundColorWathet900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorWathet900
	return r
}

// SetBackgroundColorYellow set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow
	return r
}

// SetBackgroundColorYellow50 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow50
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow50() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow50
	return r
}

// SetBackgroundColorYellow100 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow100
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow100() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow100
	return r
}

// SetBackgroundColorYellow200 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow200
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow200() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow200
	return r
}

// SetBackgroundColorYellow300 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow300
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow300() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow300
	return r
}

// SetBackgroundColorYellow350 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow350
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow350() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow350
	return r
}

// SetBackgroundColorYellow400 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow400
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow400() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow400
	return r
}

// SetBackgroundColorYellow500 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow500
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow500() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow500
	return r
}

// SetBackgroundColorYellow600 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow600
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow600() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow600
	return r
}

// SetBackgroundColorYellow700 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow700
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow700() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow700
	return r
}

// SetBackgroundColorYellow800 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow800
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow800() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow800
	return r
}

// SetBackgroundColorYellow900 set CollapsiblePanelHeader.BackgroundColor attribute to ColorYellow900
func (r *CollapsiblePanelHeader) SetBackgroundColorYellow900() *CollapsiblePanelHeader {
	r.BackgroundColor = ColorYellow900
	return r
}

// SetVerticalAlign set CollapsiblePanelHeader.VerticalAlign attribute
func (r *CollapsiblePanelHeader) SetVerticalAlign(val VerticalAlign) *CollapsiblePanelHeader {
	r.VerticalAlign = val
	return r
}

// SetVerticalAlignTop set CollapsiblePanelHeader.VerticalAlign attribute to VerticalAlignTop
func (r *CollapsiblePanelHeader) SetVerticalAlignTop() *CollapsiblePanelHeader {
	r.VerticalAlign = VerticalAlignTop
	return r
}

// SetVerticalAlignCenter set CollapsiblePanelHeader.VerticalAlign attribute to VerticalAlignCenter
func (r *CollapsiblePanelHeader) SetVerticalAlignCenter() *CollapsiblePanelHeader {
	r.VerticalAlign = VerticalAlignCenter
	return r
}

// SetVerticalAlignBottom set CollapsiblePanelHeader.VerticalAlign attribute to VerticalAlignBottom
func (r *CollapsiblePanelHeader) SetVerticalAlignBottom() *CollapsiblePanelHeader {
	r.VerticalAlign = VerticalAlignBottom
	return r
}

// SetPadding set CollapsiblePanelHeader.Padding attribute
func (r *CollapsiblePanelHeader) SetPadding(val Padding) *CollapsiblePanelHeader {
	r.Padding = val
	return r
}

// SetIcon set CollapsiblePanelHeader.Icon attribute
func (r *CollapsiblePanelHeader) SetIcon(val *ObjectIcon) *CollapsiblePanelHeader {
	r.Icon = val
	return r
}

// SetIconPosition set CollapsiblePanelHeader.IconPosition attribute
func (r *CollapsiblePanelHeader) SetIconPosition(val string) *CollapsiblePanelHeader {
	r.IconPosition = val
	return r
}

// SetIconExpandedAngle set CollapsiblePanelHeader.IconExpandedAngle attribute
func (r *CollapsiblePanelHeader) SetIconExpandedAngle(val int) *CollapsiblePanelHeader {
	r.IconExpandedAngle = val
	return r
}
