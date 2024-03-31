package card

import "encoding/json"

func TextTag(text string) *ObjectTextTag {
	return &ObjectTextTag{
		Text: Text(text),
	}
}

//go:generate generate_set_attrs -type=ObjectTextTag
//go:generate generate_to_map -type=ObjectTextTag
type ObjectTextTag struct {
	// 后缀标签的内容。基于文本组件的 plain_text 模式定义内容。示例值：
	Text *ObjectText `json:"text,omitempty"`

	// 后缀标签的颜色, 默认为蓝色（blue）。可选值与示例效果参见下文的后缀标签颜色枚举。
	Color TextTagColor `json:"color,omitempty"`
}

func (r ObjectTextTag) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "text_tag"
	return json.Marshal(m)
}

type TextTagColor = string

const (
	TextTagColorNeutral   TextTagColor = "neutral"   // 中性
	TextTagColorBlue      TextTagColor = "blue"      // 蓝
	TextTagColorTurquoise TextTagColor = "turquoise" // 青
	TextTagColorLime      TextTagColor = "lime"      // 绿
	TextTagColorOrange    TextTagColor = "orange"    // 橙
	TextTagColorViolet    TextTagColor = "violet"    // 紫红
	TextTagColorIndigo    TextTagColor = "indigo"    // 靛蓝
	TextTagColorWathet    TextTagColor = "wathet"    // 天蓝
	TextTagColorGreen     TextTagColor = "green"     // 绿
	TextTagColorYellow    TextTagColor = "yellow"    // 黄
	TextTagColorRed       TextTagColor = "red"       // 红
	TextTagColorPurple    TextTagColor = "purple"    // 紫
	TextTagColorCarmine   TextTagColor = "carmine"   // 粉
)

// SetText set ObjectTextTag.Text attribute
func (r *ObjectTextTag) SetText(val *ObjectText) *ObjectTextTag {
	r.Text = val
	return r
}

// SetColor set ObjectTextTag.Color attribute
func (r *ObjectTextTag) SetColor(val TextTagColor) *ObjectTextTag {
	r.Color = val
	return r
}

// SetColorNeutral set ObjectTextTag.Color attribute to TextTagColorNeutral
func (r *ObjectTextTag) SetColorNeutral() *ObjectTextTag {
	r.Color = TextTagColorNeutral
	return r
}

// SetColorBlue set ObjectTextTag.Color attribute to TextTagColorBlue
func (r *ObjectTextTag) SetColorBlue() *ObjectTextTag {
	r.Color = TextTagColorBlue
	return r
}

// SetColorTurquoise set ObjectTextTag.Color attribute to TextTagColorTurquoise
func (r *ObjectTextTag) SetColorTurquoise() *ObjectTextTag {
	r.Color = TextTagColorTurquoise
	return r
}

// SetColorLime set ObjectTextTag.Color attribute to TextTagColorLime
func (r *ObjectTextTag) SetColorLime() *ObjectTextTag {
	r.Color = TextTagColorLime
	return r
}

// SetColorOrange set ObjectTextTag.Color attribute to TextTagColorOrange
func (r *ObjectTextTag) SetColorOrange() *ObjectTextTag {
	r.Color = TextTagColorOrange
	return r
}

// SetColorViolet set ObjectTextTag.Color attribute to TextTagColorViolet
func (r *ObjectTextTag) SetColorViolet() *ObjectTextTag {
	r.Color = TextTagColorViolet
	return r
}

// SetColorIndigo set ObjectTextTag.Color attribute to TextTagColorIndigo
func (r *ObjectTextTag) SetColorIndigo() *ObjectTextTag {
	r.Color = TextTagColorIndigo
	return r
}

// SetColorWathet set ObjectTextTag.Color attribute to TextTagColorWathet
func (r *ObjectTextTag) SetColorWathet() *ObjectTextTag {
	r.Color = TextTagColorWathet
	return r
}

// SetColorGreen set ObjectTextTag.Color attribute to TextTagColorGreen
func (r *ObjectTextTag) SetColorGreen() *ObjectTextTag {
	r.Color = TextTagColorGreen
	return r
}

// SetColorYellow set ObjectTextTag.Color attribute to TextTagColorYellow
func (r *ObjectTextTag) SetColorYellow() *ObjectTextTag {
	r.Color = TextTagColorYellow
	return r
}

// SetColorRed set ObjectTextTag.Color attribute to TextTagColorRed
func (r *ObjectTextTag) SetColorRed() *ObjectTextTag {
	r.Color = TextTagColorRed
	return r
}

// SetColorPurple set ObjectTextTag.Color attribute to TextTagColorPurple
func (r *ObjectTextTag) SetColorPurple() *ObjectTextTag {
	r.Color = TextTagColorPurple
	return r
}

// SetColorCarmine set ObjectTextTag.Color attribute to TextTagColorCarmine
func (r *ObjectTextTag) SetColorCarmine() *ObjectTextTag {
	r.Color = TextTagColorCarmine
	return r
}

// toMap conv ObjectTextTag to map
func (r *ObjectTextTag) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 2)
	if r.Text != nil {
		res["text"] = r.Text
	}
	if r.Color != "" {
		res["color"] = r.Color
	}
	return res
}
