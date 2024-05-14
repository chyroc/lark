package card

func Text(text string) *ObjectText {
	return &ObjectText{
		Tag:     "plain_text",
		Content: text,
	}
}

func TextI18n(i18n map[Language]string) *ObjectText {
	return &ObjectText{
		Tag:  "plain_text",
		I18n: i18n,
	}
}

func Md(md string) *ObjectText {
	return &ObjectText{
		Tag:     "lark_md",
		Content: md,
	}
}

//go:generate generate_set_attrs -type=ObjectText
type ObjectText struct {
	componentBase

	// 文本类型的标签。可取值：
	//
	// plain_text：普通文本内容
	// lark_md：支持部分 Markdown 语法的文本内容。详情参考下文 lark_md 支持的 Markdown 语法
	// 注意：飞书卡片搭建工具中仅支持使用 plain_text 类型的普通文本组件。你可使用富文本组件添加 Markdown 格式的文本。
	Tag string `json:"tag,omitempty"`

	// 文本内容。
	//
	// 当 tag 为 lark_md 时，支持部分 Markdown 语法的文本内容。详情参考下文 lark_md 支持的 Markdown 语法。
	Content string `json:"content"`

	// 多语言标题内容
	//
	// 必须配置 content 或 i18n 两个属性的其中一个。如果同时配置仅生效 i18n。
	// 标题内容最多 2 行, 超出 2 行的内容用 ... 省略。
	I18n map[Language]string `json:"i18n,omitempty"`

	// 文本大小。可取值：
	//
	// normal：正文（14px）
	// heading：标题（16px）
	// notation：辅助信息（12px）
	//
	// 支持本参数的组件:
	// - 勾选器: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/checker#:~:text=%E2%94%94-,text_size,-%E5%90%A6
	// - 普通文本: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/plain-text#:~:text=%E2%94%94-,text_size,-%E5%90%A6
	TextSize TextSize `json:"text_size,omitempty"`

	// 文本的颜色。仅在 tag 为 plain_text 时生效。可取值：
	//
	// default：客户端浅色主题模式下为黑色; 客户端深色主题模式下为白色
	// 颜色的枚举值。详情参考颜色枚举值
	//
	// 支持本参数的组件:
	// - 勾选器: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/checker#:~:text=%E2%94%94-,text_color,-%E5%90%A6
	// - 普通文本: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/plain-text#:~:text=%E2%94%94-,text_color,-%E5%90%A6
	TextColor Color `json:"text_color,omitempty"`

	// 文本对齐方式。可取值：
	//
	// left：左对齐
	// center：居中对齐
	// right：右对齐
	//
	// 支持本参数的组件:
	// - 勾选器: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/checker#:~:text=%E2%94%94-,text_align,-%E5%90%A6
	// - 普通文本: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/plain-text#:~:text=%E2%94%94-,text_align,-%E5%90%A6
	TextAlign TextAlign `json:"text_align,omitempty"`

	// 内容最大显示行数，超出设置行的内容用 ... 省略
	//
	// 支持本参数的组件:
	// - 普通文本: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/plain-text#:~:text=%E2%94%94-,lines,-%E5%90%A6
	Lines int64 `json:"lines,omitempty"`
}

// SetTag set ObjectText.Tag attribute
func (r *ObjectText) SetTag(val string) *ObjectText {
	r.Tag = val
	return r
}

// SetContent set ObjectText.Content attribute
func (r *ObjectText) SetContent(val string) *ObjectText {
	r.Content = val
	return r
}

// SetI18n set ObjectText.I18n attribute
func (r *ObjectText) SetI18n(val map[Language]string) *ObjectText {
	r.I18n = val
	return r
}

// SetTextSize set ObjectText.TextSize attribute
func (r *ObjectText) SetTextSize(val TextSize) *ObjectText {
	r.TextSize = val
	return r
}

// SetTextSizeHeading0 set ObjectText.TextSize attribute to TextSizeHeading0
func (r *ObjectText) SetTextSizeHeading0() *ObjectText {
	r.TextSize = TextSizeHeading0
	return r
}

// SetTextSizeHeading1 set ObjectText.TextSize attribute to TextSizeHeading1
func (r *ObjectText) SetTextSizeHeading1() *ObjectText {
	r.TextSize = TextSizeHeading1
	return r
}

// SetTextSizeHeading2 set ObjectText.TextSize attribute to TextSizeHeading2
func (r *ObjectText) SetTextSizeHeading2() *ObjectText {
	r.TextSize = TextSizeHeading2
	return r
}

// SetTextSizeHeading3 set ObjectText.TextSize attribute to TextSizeHeading3
func (r *ObjectText) SetTextSizeHeading3() *ObjectText {
	r.TextSize = TextSizeHeading3
	return r
}

// SetTextSizeHeading4 set ObjectText.TextSize attribute to TextSizeHeading4
func (r *ObjectText) SetTextSizeHeading4() *ObjectText {
	r.TextSize = TextSizeHeading4
	return r
}

// SetTextSizeHeading set ObjectText.TextSize attribute to TextSizeHeading
func (r *ObjectText) SetTextSizeHeading() *ObjectText {
	r.TextSize = TextSizeHeading
	return r
}

// SetTextSizeNormal set ObjectText.TextSize attribute to TextSizeNormal
func (r *ObjectText) SetTextSizeNormal() *ObjectText {
	r.TextSize = TextSizeNormal
	return r
}

// SetTextSizeNotation set ObjectText.TextSize attribute to TextSizeNotation
func (r *ObjectText) SetTextSizeNotation() *ObjectText {
	r.TextSize = TextSizeNotation
	return r
}

// SetTextSizeXxxxLarge set ObjectText.TextSize attribute to TextSizeXxxxLarge
func (r *ObjectText) SetTextSizeXxxxLarge() *ObjectText {
	r.TextSize = TextSizeXxxxLarge
	return r
}

// SetTextSizeXxxLarge set ObjectText.TextSize attribute to TextSizeXxxLarge
func (r *ObjectText) SetTextSizeXxxLarge() *ObjectText {
	r.TextSize = TextSizeXxxLarge
	return r
}

// SetTextSizeXxLarge set ObjectText.TextSize attribute to TextSizeXxLarge
func (r *ObjectText) SetTextSizeXxLarge() *ObjectText {
	r.TextSize = TextSizeXxLarge
	return r
}

// SetTextSizeXLarge set ObjectText.TextSize attribute to TextSizeXLarge
func (r *ObjectText) SetTextSizeXLarge() *ObjectText {
	r.TextSize = TextSizeXLarge
	return r
}

// SetTextSizeLarge set ObjectText.TextSize attribute to TextSizeLarge
func (r *ObjectText) SetTextSizeLarge() *ObjectText {
	r.TextSize = TextSizeLarge
	return r
}

// SetTextSizeMedium set ObjectText.TextSize attribute to TextSizeMedium
func (r *ObjectText) SetTextSizeMedium() *ObjectText {
	r.TextSize = TextSizeMedium
	return r
}

// SetTextSizeSmall set ObjectText.TextSize attribute to TextSizeSmall
func (r *ObjectText) SetTextSizeSmall() *ObjectText {
	r.TextSize = TextSizeSmall
	return r
}

// SetTextSizeXSmall set ObjectText.TextSize attribute to TextSizeXSmall
func (r *ObjectText) SetTextSizeXSmall() *ObjectText {
	r.TextSize = TextSizeXSmall
	return r
}

// SetTextColor set ObjectText.TextColor attribute
func (r *ObjectText) SetTextColor(val Color) *ObjectText {
	r.TextColor = val
	return r
}

// SetTextColorDefault set ObjectText.TextColor attribute to ColorDefault
func (r *ObjectText) SetTextColorDefault() *ObjectText {
	r.TextColor = ColorDefault
	return r
}

// SetTextColorBgWhite set ObjectText.TextColor attribute to ColorBgWhite
func (r *ObjectText) SetTextColorBgWhite() *ObjectText {
	r.TextColor = ColorBgWhite
	return r
}

// SetTextColorWhite set ObjectText.TextColor attribute to ColorWhite
func (r *ObjectText) SetTextColorWhite() *ObjectText {
	r.TextColor = ColorWhite
	return r
}

// SetTextColorBlue set ObjectText.TextColor attribute to ColorBlue
func (r *ObjectText) SetTextColorBlue() *ObjectText {
	r.TextColor = ColorBlue
	return r
}

// SetTextColorBlue50 set ObjectText.TextColor attribute to ColorBlue50
func (r *ObjectText) SetTextColorBlue50() *ObjectText {
	r.TextColor = ColorBlue50
	return r
}

// SetTextColorBlue100 set ObjectText.TextColor attribute to ColorBlue100
func (r *ObjectText) SetTextColorBlue100() *ObjectText {
	r.TextColor = ColorBlue100
	return r
}

// SetTextColorBlue200 set ObjectText.TextColor attribute to ColorBlue200
func (r *ObjectText) SetTextColorBlue200() *ObjectText {
	r.TextColor = ColorBlue200
	return r
}

// SetTextColorBlue300 set ObjectText.TextColor attribute to ColorBlue300
func (r *ObjectText) SetTextColorBlue300() *ObjectText {
	r.TextColor = ColorBlue300
	return r
}

// SetTextColorBlue350 set ObjectText.TextColor attribute to ColorBlue350
func (r *ObjectText) SetTextColorBlue350() *ObjectText {
	r.TextColor = ColorBlue350
	return r
}

// SetTextColorBlue400 set ObjectText.TextColor attribute to ColorBlue400
func (r *ObjectText) SetTextColorBlue400() *ObjectText {
	r.TextColor = ColorBlue400
	return r
}

// SetTextColorBlue500 set ObjectText.TextColor attribute to ColorBlue500
func (r *ObjectText) SetTextColorBlue500() *ObjectText {
	r.TextColor = ColorBlue500
	return r
}

// SetTextColorBlue600 set ObjectText.TextColor attribute to ColorBlue600
func (r *ObjectText) SetTextColorBlue600() *ObjectText {
	r.TextColor = ColorBlue600
	return r
}

// SetTextColorBlue700 set ObjectText.TextColor attribute to ColorBlue700
func (r *ObjectText) SetTextColorBlue700() *ObjectText {
	r.TextColor = ColorBlue700
	return r
}

// SetTextColorBlue800 set ObjectText.TextColor attribute to ColorBlue800
func (r *ObjectText) SetTextColorBlue800() *ObjectText {
	r.TextColor = ColorBlue800
	return r
}

// SetTextColorBlue900 set ObjectText.TextColor attribute to ColorBlue900
func (r *ObjectText) SetTextColorBlue900() *ObjectText {
	r.TextColor = ColorBlue900
	return r
}

// SetTextColorCarmine set ObjectText.TextColor attribute to ColorCarmine
func (r *ObjectText) SetTextColorCarmine() *ObjectText {
	r.TextColor = ColorCarmine
	return r
}

// SetTextColorCarmine50 set ObjectText.TextColor attribute to ColorCarmine50
func (r *ObjectText) SetTextColorCarmine50() *ObjectText {
	r.TextColor = ColorCarmine50
	return r
}

// SetTextColorCarmine100 set ObjectText.TextColor attribute to ColorCarmine100
func (r *ObjectText) SetTextColorCarmine100() *ObjectText {
	r.TextColor = ColorCarmine100
	return r
}

// SetTextColorCarmine200 set ObjectText.TextColor attribute to ColorCarmine200
func (r *ObjectText) SetTextColorCarmine200() *ObjectText {
	r.TextColor = ColorCarmine200
	return r
}

// SetTextColorCarmine300 set ObjectText.TextColor attribute to ColorCarmine300
func (r *ObjectText) SetTextColorCarmine300() *ObjectText {
	r.TextColor = ColorCarmine300
	return r
}

// SetTextColorCarmine350 set ObjectText.TextColor attribute to ColorCarmine350
func (r *ObjectText) SetTextColorCarmine350() *ObjectText {
	r.TextColor = ColorCarmine350
	return r
}

// SetTextColorCarmine400 set ObjectText.TextColor attribute to ColorCarmine400
func (r *ObjectText) SetTextColorCarmine400() *ObjectText {
	r.TextColor = ColorCarmine400
	return r
}

// SetTextColorCarmine500 set ObjectText.TextColor attribute to ColorCarmine500
func (r *ObjectText) SetTextColorCarmine500() *ObjectText {
	r.TextColor = ColorCarmine500
	return r
}

// SetTextColorCarmine600 set ObjectText.TextColor attribute to ColorCarmine600
func (r *ObjectText) SetTextColorCarmine600() *ObjectText {
	r.TextColor = ColorCarmine600
	return r
}

// SetTextColorCarmine700 set ObjectText.TextColor attribute to ColorCarmine700
func (r *ObjectText) SetTextColorCarmine700() *ObjectText {
	r.TextColor = ColorCarmine700
	return r
}

// SetTextColorCarmine800 set ObjectText.TextColor attribute to ColorCarmine800
func (r *ObjectText) SetTextColorCarmine800() *ObjectText {
	r.TextColor = ColorCarmine800
	return r
}

// SetTextColorCarmine900 set ObjectText.TextColor attribute to ColorCarmine900
func (r *ObjectText) SetTextColorCarmine900() *ObjectText {
	r.TextColor = ColorCarmine900
	return r
}

// SetTextColorGreen set ObjectText.TextColor attribute to ColorGreen
func (r *ObjectText) SetTextColorGreen() *ObjectText {
	r.TextColor = ColorGreen
	return r
}

// SetTextColorGreen50 set ObjectText.TextColor attribute to ColorGreen50
func (r *ObjectText) SetTextColorGreen50() *ObjectText {
	r.TextColor = ColorGreen50
	return r
}

// SetTextColorGreen100 set ObjectText.TextColor attribute to ColorGreen100
func (r *ObjectText) SetTextColorGreen100() *ObjectText {
	r.TextColor = ColorGreen100
	return r
}

// SetTextColorGreen200 set ObjectText.TextColor attribute to ColorGreen200
func (r *ObjectText) SetTextColorGreen200() *ObjectText {
	r.TextColor = ColorGreen200
	return r
}

// SetTextColorGreen300 set ObjectText.TextColor attribute to ColorGreen300
func (r *ObjectText) SetTextColorGreen300() *ObjectText {
	r.TextColor = ColorGreen300
	return r
}

// SetTextColorGreen350 set ObjectText.TextColor attribute to ColorGreen350
func (r *ObjectText) SetTextColorGreen350() *ObjectText {
	r.TextColor = ColorGreen350
	return r
}

// SetTextColorGreen400 set ObjectText.TextColor attribute to ColorGreen400
func (r *ObjectText) SetTextColorGreen400() *ObjectText {
	r.TextColor = ColorGreen400
	return r
}

// SetTextColorGreen500 set ObjectText.TextColor attribute to ColorGreen500
func (r *ObjectText) SetTextColorGreen500() *ObjectText {
	r.TextColor = ColorGreen500
	return r
}

// SetTextColorGreen600 set ObjectText.TextColor attribute to ColorGreen600
func (r *ObjectText) SetTextColorGreen600() *ObjectText {
	r.TextColor = ColorGreen600
	return r
}

// SetTextColorGreen700 set ObjectText.TextColor attribute to ColorGreen700
func (r *ObjectText) SetTextColorGreen700() *ObjectText {
	r.TextColor = ColorGreen700
	return r
}

// SetTextColorGreen800 set ObjectText.TextColor attribute to ColorGreen800
func (r *ObjectText) SetTextColorGreen800() *ObjectText {
	r.TextColor = ColorGreen800
	return r
}

// SetTextColorGreen900 set ObjectText.TextColor attribute to ColorGreen900
func (r *ObjectText) SetTextColorGreen900() *ObjectText {
	r.TextColor = ColorGreen900
	return r
}

// SetTextColorIndigo set ObjectText.TextColor attribute to ColorIndigo
func (r *ObjectText) SetTextColorIndigo() *ObjectText {
	r.TextColor = ColorIndigo
	return r
}

// SetTextColorIndigo50 set ObjectText.TextColor attribute to ColorIndigo50
func (r *ObjectText) SetTextColorIndigo50() *ObjectText {
	r.TextColor = ColorIndigo50
	return r
}

// SetTextColorIndigo100 set ObjectText.TextColor attribute to ColorIndigo100
func (r *ObjectText) SetTextColorIndigo100() *ObjectText {
	r.TextColor = ColorIndigo100
	return r
}

// SetTextColorIndigo200 set ObjectText.TextColor attribute to ColorIndigo200
func (r *ObjectText) SetTextColorIndigo200() *ObjectText {
	r.TextColor = ColorIndigo200
	return r
}

// SetTextColorIndigo300 set ObjectText.TextColor attribute to ColorIndigo300
func (r *ObjectText) SetTextColorIndigo300() *ObjectText {
	r.TextColor = ColorIndigo300
	return r
}

// SetTextColorIndigo350 set ObjectText.TextColor attribute to ColorIndigo350
func (r *ObjectText) SetTextColorIndigo350() *ObjectText {
	r.TextColor = ColorIndigo350
	return r
}

// SetTextColorIndigo400 set ObjectText.TextColor attribute to ColorIndigo400
func (r *ObjectText) SetTextColorIndigo400() *ObjectText {
	r.TextColor = ColorIndigo400
	return r
}

// SetTextColorIndigo500 set ObjectText.TextColor attribute to ColorIndigo500
func (r *ObjectText) SetTextColorIndigo500() *ObjectText {
	r.TextColor = ColorIndigo500
	return r
}

// SetTextColorIndigo600 set ObjectText.TextColor attribute to ColorIndigo600
func (r *ObjectText) SetTextColorIndigo600() *ObjectText {
	r.TextColor = ColorIndigo600
	return r
}

// SetTextColorIndigo700 set ObjectText.TextColor attribute to ColorIndigo700
func (r *ObjectText) SetTextColorIndigo700() *ObjectText {
	r.TextColor = ColorIndigo700
	return r
}

// SetTextColorIndigo800 set ObjectText.TextColor attribute to ColorIndigo800
func (r *ObjectText) SetTextColorIndigo800() *ObjectText {
	r.TextColor = ColorIndigo800
	return r
}

// SetTextColorIndigo900 set ObjectText.TextColor attribute to ColorIndigo900
func (r *ObjectText) SetTextColorIndigo900() *ObjectText {
	r.TextColor = ColorIndigo900
	return r
}

// SetTextColorLime set ObjectText.TextColor attribute to ColorLime
func (r *ObjectText) SetTextColorLime() *ObjectText {
	r.TextColor = ColorLime
	return r
}

// SetTextColorLime50 set ObjectText.TextColor attribute to ColorLime50
func (r *ObjectText) SetTextColorLime50() *ObjectText {
	r.TextColor = ColorLime50
	return r
}

// SetTextColorLime100 set ObjectText.TextColor attribute to ColorLime100
func (r *ObjectText) SetTextColorLime100() *ObjectText {
	r.TextColor = ColorLime100
	return r
}

// SetTextColorLime200 set ObjectText.TextColor attribute to ColorLime200
func (r *ObjectText) SetTextColorLime200() *ObjectText {
	r.TextColor = ColorLime200
	return r
}

// SetTextColorLime300 set ObjectText.TextColor attribute to ColorLime300
func (r *ObjectText) SetTextColorLime300() *ObjectText {
	r.TextColor = ColorLime300
	return r
}

// SetTextColorLime350 set ObjectText.TextColor attribute to ColorLime350
func (r *ObjectText) SetTextColorLime350() *ObjectText {
	r.TextColor = ColorLime350
	return r
}

// SetTextColorLime400 set ObjectText.TextColor attribute to ColorLime400
func (r *ObjectText) SetTextColorLime400() *ObjectText {
	r.TextColor = ColorLime400
	return r
}

// SetTextColorLime500 set ObjectText.TextColor attribute to ColorLime500
func (r *ObjectText) SetTextColorLime500() *ObjectText {
	r.TextColor = ColorLime500
	return r
}

// SetTextColorLime600 set ObjectText.TextColor attribute to ColorLime600
func (r *ObjectText) SetTextColorLime600() *ObjectText {
	r.TextColor = ColorLime600
	return r
}

// SetTextColorLime700 set ObjectText.TextColor attribute to ColorLime700
func (r *ObjectText) SetTextColorLime700() *ObjectText {
	r.TextColor = ColorLime700
	return r
}

// SetTextColorLime800 set ObjectText.TextColor attribute to ColorLime800
func (r *ObjectText) SetTextColorLime800() *ObjectText {
	r.TextColor = ColorLime800
	return r
}

// SetTextColorLime900 set ObjectText.TextColor attribute to ColorLime900
func (r *ObjectText) SetTextColorLime900() *ObjectText {
	r.TextColor = ColorLime900
	return r
}

// SetTextColorGrey set ObjectText.TextColor attribute to ColorGrey
func (r *ObjectText) SetTextColorGrey() *ObjectText {
	r.TextColor = ColorGrey
	return r
}

// SetTextColorGrey00 set ObjectText.TextColor attribute to ColorGrey00
func (r *ObjectText) SetTextColorGrey00() *ObjectText {
	r.TextColor = ColorGrey00
	return r
}

// SetTextColorGrey50 set ObjectText.TextColor attribute to ColorGrey50
func (r *ObjectText) SetTextColorGrey50() *ObjectText {
	r.TextColor = ColorGrey50
	return r
}

// SetTextColorGrey100 set ObjectText.TextColor attribute to ColorGrey100
func (r *ObjectText) SetTextColorGrey100() *ObjectText {
	r.TextColor = ColorGrey100
	return r
}

// SetTextColorGrey200 set ObjectText.TextColor attribute to ColorGrey200
func (r *ObjectText) SetTextColorGrey200() *ObjectText {
	r.TextColor = ColorGrey200
	return r
}

// SetTextColorGrey300 set ObjectText.TextColor attribute to ColorGrey300
func (r *ObjectText) SetTextColorGrey300() *ObjectText {
	r.TextColor = ColorGrey300
	return r
}

// SetTextColorGrey350 set ObjectText.TextColor attribute to ColorGrey350
func (r *ObjectText) SetTextColorGrey350() *ObjectText {
	r.TextColor = ColorGrey350
	return r
}

// SetTextColorGrey400 set ObjectText.TextColor attribute to ColorGrey400
func (r *ObjectText) SetTextColorGrey400() *ObjectText {
	r.TextColor = ColorGrey400
	return r
}

// SetTextColorGrey500 set ObjectText.TextColor attribute to ColorGrey500
func (r *ObjectText) SetTextColorGrey500() *ObjectText {
	r.TextColor = ColorGrey500
	return r
}

// SetTextColorGrey600 set ObjectText.TextColor attribute to ColorGrey600
func (r *ObjectText) SetTextColorGrey600() *ObjectText {
	r.TextColor = ColorGrey600
	return r
}

// SetTextColorGrey650 set ObjectText.TextColor attribute to ColorGrey650
func (r *ObjectText) SetTextColorGrey650() *ObjectText {
	r.TextColor = ColorGrey650
	return r
}

// SetTextColorGrey700 set ObjectText.TextColor attribute to ColorGrey700
func (r *ObjectText) SetTextColorGrey700() *ObjectText {
	r.TextColor = ColorGrey700
	return r
}

// SetTextColorGrey800 set ObjectText.TextColor attribute to ColorGrey800
func (r *ObjectText) SetTextColorGrey800() *ObjectText {
	r.TextColor = ColorGrey800
	return r
}

// SetTextColorGrey900 set ObjectText.TextColor attribute to ColorGrey900
func (r *ObjectText) SetTextColorGrey900() *ObjectText {
	r.TextColor = ColorGrey900
	return r
}

// SetTextColorGrey950 set ObjectText.TextColor attribute to ColorGrey950
func (r *ObjectText) SetTextColorGrey950() *ObjectText {
	r.TextColor = ColorGrey950
	return r
}

// SetTextColorGrey1000 set ObjectText.TextColor attribute to ColorGrey1000
func (r *ObjectText) SetTextColorGrey1000() *ObjectText {
	r.TextColor = ColorGrey1000
	return r
}

// SetTextColorOrange set ObjectText.TextColor attribute to ColorOrange
func (r *ObjectText) SetTextColorOrange() *ObjectText {
	r.TextColor = ColorOrange
	return r
}

// SetTextColorOrange50 set ObjectText.TextColor attribute to ColorOrange50
func (r *ObjectText) SetTextColorOrange50() *ObjectText {
	r.TextColor = ColorOrange50
	return r
}

// SetTextColorOrange100 set ObjectText.TextColor attribute to ColorOrange100
func (r *ObjectText) SetTextColorOrange100() *ObjectText {
	r.TextColor = ColorOrange100
	return r
}

// SetTextColorOrange200 set ObjectText.TextColor attribute to ColorOrange200
func (r *ObjectText) SetTextColorOrange200() *ObjectText {
	r.TextColor = ColorOrange200
	return r
}

// SetTextColorOrange300 set ObjectText.TextColor attribute to ColorOrange300
func (r *ObjectText) SetTextColorOrange300() *ObjectText {
	r.TextColor = ColorOrange300
	return r
}

// SetTextColorOrange350 set ObjectText.TextColor attribute to ColorOrange350
func (r *ObjectText) SetTextColorOrange350() *ObjectText {
	r.TextColor = ColorOrange350
	return r
}

// SetTextColorOrange400 set ObjectText.TextColor attribute to ColorOrange400
func (r *ObjectText) SetTextColorOrange400() *ObjectText {
	r.TextColor = ColorOrange400
	return r
}

// SetTextColorOrange500 set ObjectText.TextColor attribute to ColorOrange500
func (r *ObjectText) SetTextColorOrange500() *ObjectText {
	r.TextColor = ColorOrange500
	return r
}

// SetTextColorOrange600 set ObjectText.TextColor attribute to ColorOrange600
func (r *ObjectText) SetTextColorOrange600() *ObjectText {
	r.TextColor = ColorOrange600
	return r
}

// SetTextColorOrange700 set ObjectText.TextColor attribute to ColorOrange700
func (r *ObjectText) SetTextColorOrange700() *ObjectText {
	r.TextColor = ColorOrange700
	return r
}

// SetTextColorOrange800 set ObjectText.TextColor attribute to ColorOrange800
func (r *ObjectText) SetTextColorOrange800() *ObjectText {
	r.TextColor = ColorOrange800
	return r
}

// SetTextColorOrange900 set ObjectText.TextColor attribute to ColorOrange900
func (r *ObjectText) SetTextColorOrange900() *ObjectText {
	r.TextColor = ColorOrange900
	return r
}

// SetTextColorPurple set ObjectText.TextColor attribute to ColorPurple
func (r *ObjectText) SetTextColorPurple() *ObjectText {
	r.TextColor = ColorPurple
	return r
}

// SetTextColorPurple50 set ObjectText.TextColor attribute to ColorPurple50
func (r *ObjectText) SetTextColorPurple50() *ObjectText {
	r.TextColor = ColorPurple50
	return r
}

// SetTextColorPurple100 set ObjectText.TextColor attribute to ColorPurple100
func (r *ObjectText) SetTextColorPurple100() *ObjectText {
	r.TextColor = ColorPurple100
	return r
}

// SetTextColorPurple200 set ObjectText.TextColor attribute to ColorPurple200
func (r *ObjectText) SetTextColorPurple200() *ObjectText {
	r.TextColor = ColorPurple200
	return r
}

// SetTextColorPurple300 set ObjectText.TextColor attribute to ColorPurple300
func (r *ObjectText) SetTextColorPurple300() *ObjectText {
	r.TextColor = ColorPurple300
	return r
}

// SetTextColorPurple350 set ObjectText.TextColor attribute to ColorPurple350
func (r *ObjectText) SetTextColorPurple350() *ObjectText {
	r.TextColor = ColorPurple350
	return r
}

// SetTextColorPurple400 set ObjectText.TextColor attribute to ColorPurple400
func (r *ObjectText) SetTextColorPurple400() *ObjectText {
	r.TextColor = ColorPurple400
	return r
}

// SetTextColorPurple500 set ObjectText.TextColor attribute to ColorPurple500
func (r *ObjectText) SetTextColorPurple500() *ObjectText {
	r.TextColor = ColorPurple500
	return r
}

// SetTextColorPurple600 set ObjectText.TextColor attribute to ColorPurple600
func (r *ObjectText) SetTextColorPurple600() *ObjectText {
	r.TextColor = ColorPurple600
	return r
}

// SetTextColorPurple700 set ObjectText.TextColor attribute to ColorPurple700
func (r *ObjectText) SetTextColorPurple700() *ObjectText {
	r.TextColor = ColorPurple700
	return r
}

// SetTextColorPurple800 set ObjectText.TextColor attribute to ColorPurple800
func (r *ObjectText) SetTextColorPurple800() *ObjectText {
	r.TextColor = ColorPurple800
	return r
}

// SetTextColorPurple900 set ObjectText.TextColor attribute to ColorPurple900
func (r *ObjectText) SetTextColorPurple900() *ObjectText {
	r.TextColor = ColorPurple900
	return r
}

// SetTextColorRed set ObjectText.TextColor attribute to ColorRed
func (r *ObjectText) SetTextColorRed() *ObjectText {
	r.TextColor = ColorRed
	return r
}

// SetTextColorRed50 set ObjectText.TextColor attribute to ColorRed50
func (r *ObjectText) SetTextColorRed50() *ObjectText {
	r.TextColor = ColorRed50
	return r
}

// SetTextColorRed100 set ObjectText.TextColor attribute to ColorRed100
func (r *ObjectText) SetTextColorRed100() *ObjectText {
	r.TextColor = ColorRed100
	return r
}

// SetTextColorRed200 set ObjectText.TextColor attribute to ColorRed200
func (r *ObjectText) SetTextColorRed200() *ObjectText {
	r.TextColor = ColorRed200
	return r
}

// SetTextColorRed300 set ObjectText.TextColor attribute to ColorRed300
func (r *ObjectText) SetTextColorRed300() *ObjectText {
	r.TextColor = ColorRed300
	return r
}

// SetTextColorRed350 set ObjectText.TextColor attribute to ColorRed350
func (r *ObjectText) SetTextColorRed350() *ObjectText {
	r.TextColor = ColorRed350
	return r
}

// SetTextColorRed400 set ObjectText.TextColor attribute to ColorRed400
func (r *ObjectText) SetTextColorRed400() *ObjectText {
	r.TextColor = ColorRed400
	return r
}

// SetTextColorRed500 set ObjectText.TextColor attribute to ColorRed500
func (r *ObjectText) SetTextColorRed500() *ObjectText {
	r.TextColor = ColorRed500
	return r
}

// SetTextColorRed600 set ObjectText.TextColor attribute to ColorRed600
func (r *ObjectText) SetTextColorRed600() *ObjectText {
	r.TextColor = ColorRed600
	return r
}

// SetTextColorRed700 set ObjectText.TextColor attribute to ColorRed700
func (r *ObjectText) SetTextColorRed700() *ObjectText {
	r.TextColor = ColorRed700
	return r
}

// SetTextColorRed800 set ObjectText.TextColor attribute to ColorRed800
func (r *ObjectText) SetTextColorRed800() *ObjectText {
	r.TextColor = ColorRed800
	return r
}

// SetTextColorRed900 set ObjectText.TextColor attribute to ColorRed900
func (r *ObjectText) SetTextColorRed900() *ObjectText {
	r.TextColor = ColorRed900
	return r
}

// SetTextColorSunflower set ObjectText.TextColor attribute to ColorSunflower
func (r *ObjectText) SetTextColorSunflower() *ObjectText {
	r.TextColor = ColorSunflower
	return r
}

// SetTextColorSunflower50 set ObjectText.TextColor attribute to ColorSunflower50
func (r *ObjectText) SetTextColorSunflower50() *ObjectText {
	r.TextColor = ColorSunflower50
	return r
}

// SetTextColorSunflower100 set ObjectText.TextColor attribute to ColorSunflower100
func (r *ObjectText) SetTextColorSunflower100() *ObjectText {
	r.TextColor = ColorSunflower100
	return r
}

// SetTextColorSunflower200 set ObjectText.TextColor attribute to ColorSunflower200
func (r *ObjectText) SetTextColorSunflower200() *ObjectText {
	r.TextColor = ColorSunflower200
	return r
}

// SetTextColorSunflower300 set ObjectText.TextColor attribute to ColorSunflower300
func (r *ObjectText) SetTextColorSunflower300() *ObjectText {
	r.TextColor = ColorSunflower300
	return r
}

// SetTextColorSunflower350 set ObjectText.TextColor attribute to ColorSunflower350
func (r *ObjectText) SetTextColorSunflower350() *ObjectText {
	r.TextColor = ColorSunflower350
	return r
}

// SetTextColorSunflower400 set ObjectText.TextColor attribute to ColorSunflower400
func (r *ObjectText) SetTextColorSunflower400() *ObjectText {
	r.TextColor = ColorSunflower400
	return r
}

// SetTextColorSunflower500 set ObjectText.TextColor attribute to ColorSunflower500
func (r *ObjectText) SetTextColorSunflower500() *ObjectText {
	r.TextColor = ColorSunflower500
	return r
}

// SetTextColorSunflower600 set ObjectText.TextColor attribute to ColorSunflower600
func (r *ObjectText) SetTextColorSunflower600() *ObjectText {
	r.TextColor = ColorSunflower600
	return r
}

// SetTextColorSunflower700 set ObjectText.TextColor attribute to ColorSunflower700
func (r *ObjectText) SetTextColorSunflower700() *ObjectText {
	r.TextColor = ColorSunflower700
	return r
}

// SetTextColorSunflower800 set ObjectText.TextColor attribute to ColorSunflower800
func (r *ObjectText) SetTextColorSunflower800() *ObjectText {
	r.TextColor = ColorSunflower800
	return r
}

// SetTextColorSunflower900 set ObjectText.TextColor attribute to ColorSunflower900
func (r *ObjectText) SetTextColorSunflower900() *ObjectText {
	r.TextColor = ColorSunflower900
	return r
}

// SetTextColorTurquoise set ObjectText.TextColor attribute to ColorTurquoise
func (r *ObjectText) SetTextColorTurquoise() *ObjectText {
	r.TextColor = ColorTurquoise
	return r
}

// SetTextColorTurquoise50 set ObjectText.TextColor attribute to ColorTurquoise50
func (r *ObjectText) SetTextColorTurquoise50() *ObjectText {
	r.TextColor = ColorTurquoise50
	return r
}

// SetTextColorTurquoise100 set ObjectText.TextColor attribute to ColorTurquoise100
func (r *ObjectText) SetTextColorTurquoise100() *ObjectText {
	r.TextColor = ColorTurquoise100
	return r
}

// SetTextColorTurquoise200 set ObjectText.TextColor attribute to ColorTurquoise200
func (r *ObjectText) SetTextColorTurquoise200() *ObjectText {
	r.TextColor = ColorTurquoise200
	return r
}

// SetTextColorTurquoise300 set ObjectText.TextColor attribute to ColorTurquoise300
func (r *ObjectText) SetTextColorTurquoise300() *ObjectText {
	r.TextColor = ColorTurquoise300
	return r
}

// SetTextColorTurquoise350 set ObjectText.TextColor attribute to ColorTurquoise350
func (r *ObjectText) SetTextColorTurquoise350() *ObjectText {
	r.TextColor = ColorTurquoise350
	return r
}

// SetTextColorTurquoise400 set ObjectText.TextColor attribute to ColorTurquoise400
func (r *ObjectText) SetTextColorTurquoise400() *ObjectText {
	r.TextColor = ColorTurquoise400
	return r
}

// SetTextColorTurquoise500 set ObjectText.TextColor attribute to ColorTurquoise500
func (r *ObjectText) SetTextColorTurquoise500() *ObjectText {
	r.TextColor = ColorTurquoise500
	return r
}

// SetTextColorTurquoise600 set ObjectText.TextColor attribute to ColorTurquoise600
func (r *ObjectText) SetTextColorTurquoise600() *ObjectText {
	r.TextColor = ColorTurquoise600
	return r
}

// SetTextColorTurquoise700 set ObjectText.TextColor attribute to ColorTurquoise700
func (r *ObjectText) SetTextColorTurquoise700() *ObjectText {
	r.TextColor = ColorTurquoise700
	return r
}

// SetTextColorTurquoise800 set ObjectText.TextColor attribute to ColorTurquoise800
func (r *ObjectText) SetTextColorTurquoise800() *ObjectText {
	r.TextColor = ColorTurquoise800
	return r
}

// SetTextColorTurquoise900 set ObjectText.TextColor attribute to ColorTurquoise900
func (r *ObjectText) SetTextColorTurquoise900() *ObjectText {
	r.TextColor = ColorTurquoise900
	return r
}

// SetTextColorViolet set ObjectText.TextColor attribute to ColorViolet
func (r *ObjectText) SetTextColorViolet() *ObjectText {
	r.TextColor = ColorViolet
	return r
}

// SetTextColorViolet50 set ObjectText.TextColor attribute to ColorViolet50
func (r *ObjectText) SetTextColorViolet50() *ObjectText {
	r.TextColor = ColorViolet50
	return r
}

// SetTextColorViolet100 set ObjectText.TextColor attribute to ColorViolet100
func (r *ObjectText) SetTextColorViolet100() *ObjectText {
	r.TextColor = ColorViolet100
	return r
}

// SetTextColorViolet200 set ObjectText.TextColor attribute to ColorViolet200
func (r *ObjectText) SetTextColorViolet200() *ObjectText {
	r.TextColor = ColorViolet200
	return r
}

// SetTextColorViolet300 set ObjectText.TextColor attribute to ColorViolet300
func (r *ObjectText) SetTextColorViolet300() *ObjectText {
	r.TextColor = ColorViolet300
	return r
}

// SetTextColorViolet350 set ObjectText.TextColor attribute to ColorViolet350
func (r *ObjectText) SetTextColorViolet350() *ObjectText {
	r.TextColor = ColorViolet350
	return r
}

// SetTextColorViolet400 set ObjectText.TextColor attribute to ColorViolet400
func (r *ObjectText) SetTextColorViolet400() *ObjectText {
	r.TextColor = ColorViolet400
	return r
}

// SetTextColorViolet500 set ObjectText.TextColor attribute to ColorViolet500
func (r *ObjectText) SetTextColorViolet500() *ObjectText {
	r.TextColor = ColorViolet500
	return r
}

// SetTextColorViolet600 set ObjectText.TextColor attribute to ColorViolet600
func (r *ObjectText) SetTextColorViolet600() *ObjectText {
	r.TextColor = ColorViolet600
	return r
}

// SetTextColorViolet700 set ObjectText.TextColor attribute to ColorViolet700
func (r *ObjectText) SetTextColorViolet700() *ObjectText {
	r.TextColor = ColorViolet700
	return r
}

// SetTextColorViolet800 set ObjectText.TextColor attribute to ColorViolet800
func (r *ObjectText) SetTextColorViolet800() *ObjectText {
	r.TextColor = ColorViolet800
	return r
}

// SetTextColorViolet900 set ObjectText.TextColor attribute to ColorViolet900
func (r *ObjectText) SetTextColorViolet900() *ObjectText {
	r.TextColor = ColorViolet900
	return r
}

// SetTextColorWathet set ObjectText.TextColor attribute to ColorWathet
func (r *ObjectText) SetTextColorWathet() *ObjectText {
	r.TextColor = ColorWathet
	return r
}

// SetTextColorWathet50 set ObjectText.TextColor attribute to ColorWathet50
func (r *ObjectText) SetTextColorWathet50() *ObjectText {
	r.TextColor = ColorWathet50
	return r
}

// SetTextColorWathet100 set ObjectText.TextColor attribute to ColorWathet100
func (r *ObjectText) SetTextColorWathet100() *ObjectText {
	r.TextColor = ColorWathet100
	return r
}

// SetTextColorWathet200 set ObjectText.TextColor attribute to ColorWathet200
func (r *ObjectText) SetTextColorWathet200() *ObjectText {
	r.TextColor = ColorWathet200
	return r
}

// SetTextColorWathet300 set ObjectText.TextColor attribute to ColorWathet300
func (r *ObjectText) SetTextColorWathet300() *ObjectText {
	r.TextColor = ColorWathet300
	return r
}

// SetTextColorWathet350 set ObjectText.TextColor attribute to ColorWathet350
func (r *ObjectText) SetTextColorWathet350() *ObjectText {
	r.TextColor = ColorWathet350
	return r
}

// SetTextColorWathet400 set ObjectText.TextColor attribute to ColorWathet400
func (r *ObjectText) SetTextColorWathet400() *ObjectText {
	r.TextColor = ColorWathet400
	return r
}

// SetTextColorWathet500 set ObjectText.TextColor attribute to ColorWathet500
func (r *ObjectText) SetTextColorWathet500() *ObjectText {
	r.TextColor = ColorWathet500
	return r
}

// SetTextColorWathet600 set ObjectText.TextColor attribute to ColorWathet600
func (r *ObjectText) SetTextColorWathet600() *ObjectText {
	r.TextColor = ColorWathet600
	return r
}

// SetTextColorWathet700 set ObjectText.TextColor attribute to ColorWathet700
func (r *ObjectText) SetTextColorWathet700() *ObjectText {
	r.TextColor = ColorWathet700
	return r
}

// SetTextColorWathet800 set ObjectText.TextColor attribute to ColorWathet800
func (r *ObjectText) SetTextColorWathet800() *ObjectText {
	r.TextColor = ColorWathet800
	return r
}

// SetTextColorWathet900 set ObjectText.TextColor attribute to ColorWathet900
func (r *ObjectText) SetTextColorWathet900() *ObjectText {
	r.TextColor = ColorWathet900
	return r
}

// SetTextColorYellow set ObjectText.TextColor attribute to ColorYellow
func (r *ObjectText) SetTextColorYellow() *ObjectText {
	r.TextColor = ColorYellow
	return r
}

// SetTextColorYellow50 set ObjectText.TextColor attribute to ColorYellow50
func (r *ObjectText) SetTextColorYellow50() *ObjectText {
	r.TextColor = ColorYellow50
	return r
}

// SetTextColorYellow100 set ObjectText.TextColor attribute to ColorYellow100
func (r *ObjectText) SetTextColorYellow100() *ObjectText {
	r.TextColor = ColorYellow100
	return r
}

// SetTextColorYellow200 set ObjectText.TextColor attribute to ColorYellow200
func (r *ObjectText) SetTextColorYellow200() *ObjectText {
	r.TextColor = ColorYellow200
	return r
}

// SetTextColorYellow300 set ObjectText.TextColor attribute to ColorYellow300
func (r *ObjectText) SetTextColorYellow300() *ObjectText {
	r.TextColor = ColorYellow300
	return r
}

// SetTextColorYellow350 set ObjectText.TextColor attribute to ColorYellow350
func (r *ObjectText) SetTextColorYellow350() *ObjectText {
	r.TextColor = ColorYellow350
	return r
}

// SetTextColorYellow400 set ObjectText.TextColor attribute to ColorYellow400
func (r *ObjectText) SetTextColorYellow400() *ObjectText {
	r.TextColor = ColorYellow400
	return r
}

// SetTextColorYellow500 set ObjectText.TextColor attribute to ColorYellow500
func (r *ObjectText) SetTextColorYellow500() *ObjectText {
	r.TextColor = ColorYellow500
	return r
}

// SetTextColorYellow600 set ObjectText.TextColor attribute to ColorYellow600
func (r *ObjectText) SetTextColorYellow600() *ObjectText {
	r.TextColor = ColorYellow600
	return r
}

// SetTextColorYellow700 set ObjectText.TextColor attribute to ColorYellow700
func (r *ObjectText) SetTextColorYellow700() *ObjectText {
	r.TextColor = ColorYellow700
	return r
}

// SetTextColorYellow800 set ObjectText.TextColor attribute to ColorYellow800
func (r *ObjectText) SetTextColorYellow800() *ObjectText {
	r.TextColor = ColorYellow800
	return r
}

// SetTextColorYellow900 set ObjectText.TextColor attribute to ColorYellow900
func (r *ObjectText) SetTextColorYellow900() *ObjectText {
	r.TextColor = ColorYellow900
	return r
}

// SetTextAlign set ObjectText.TextAlign attribute
func (r *ObjectText) SetTextAlign(val TextAlign) *ObjectText {
	r.TextAlign = val
	return r
}

// SetTextAlignLeft set ObjectText.TextAlign attribute to TextAlignLeft
func (r *ObjectText) SetTextAlignLeft() *ObjectText {
	r.TextAlign = TextAlignLeft
	return r
}

// SetTextAlignCenter set ObjectText.TextAlign attribute to TextAlignCenter
func (r *ObjectText) SetTextAlignCenter() *ObjectText {
	r.TextAlign = TextAlignCenter
	return r
}

// SetTextAlignRight set ObjectText.TextAlign attribute to TextAlignRight
func (r *ObjectText) SetTextAlignRight() *ObjectText {
	r.TextAlign = TextAlignRight
	return r
}

// SetLines set ObjectText.Lines attribute
func (r *ObjectText) SetLines(val int64) *ObjectText {
	r.Lines = val
	return r
}
