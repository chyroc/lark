package card

import (
	"encoding/json"
)

func Header(title string) *ComponentHeader {
	return &ComponentHeader{
		Title:    Text(title),
		Template: TemplateDefault,
	}
}

// ComponentHeader 标题
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/title
//
//go:generate generate_set_attrs -type=ComponentHeader
//go:generate generate_to_map -type=ComponentHeader
type ComponentHeader struct {
	componentBase

	// 卡片主标题
	Title *ObjectText `json:"title,omitempty"`

	// 卡片副标题
	Subtitle *ObjectText `json:"subtitle,omitempty"`

	// 添加标题的后缀标签
	// 最多可添加 3 个标签内容, 如果配置的标签数量超过 3 个, 则取前 3 个标签进行展示。
	// 标签展示顺序与数组顺序一致。
	// 标题标签在飞书 V6.11 及以上版本开始生效。在旧版本客户端内将不会展示标题标签内容。
	// text_tag_list 和 i18n_text_tag_list 只能配置其中之一。如果同时配置仅生效 i18n_text_tag_list。
	TextTagList []*ObjectTextTag `json:"text_tag_list,omitempty"`

	// 配置后缀标签的多语言属性, 在所需语种字段下添加完整的后缀标签结构体即可。每个语言最多可配置 3 个标签内容, 如果配置的标签数量超过 3 个, 则取前 3 个标签进行展示。标签展示顺序与数组顺序一致。
	I18nTextTagList map[Language][]*ObjectTextTag `json:"i18n_text_tag_list,omitempty"`

	// 配置标题主题颜色
	Template Template `json:"template,omitempty"`

	// 图标
	Icon *ObjectIcon `json:"-,omitempty"`
}

// Template 参考 https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/title#39ee4e65
type Template = string

const (
	TemplateBlue      Template = "blue"      // 蓝色的
	TemplateWathet    Template = "wathet"    // 天蓝
	TemplateTurquoise Template = "turquoise" // 青
	TemplateGreen     Template = "green"     // 绿
	TemplateYellow    Template = "yellow"    // 黄
	TemplateOrange    Template = "orange"    // 橙
	TemplateRed       Template = "red"       // 红
	TemplateCarmine   Template = "carmine"   // 粉
	TemplateViolet    Template = "violet"    // 紫红
	TemplatePurple    Template = "purple"    // 紫
	TemplateIndigo    Template = "indigo"    // 靛蓝
	TemplateGrey      Template = "grey"      // 灰
	TemplateDefault   Template = "default"   // 默认
)

// MarshalJSON ...
func (r ComponentHeader) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	if r.Icon != nil {
		if r.Icon.Token != "" {
			if r.Icon.Color == "" {
				m["ud_icon"] = anyMap{"token": r.Icon.Token}
			} else {
				m["ud_icon"] = anyMap{
					"token": r.Icon.Token,
					"style": anyMap{"color": r.Icon.Color},
				}
			}
		} else {
			m["icon"] = anyMap{"img_key": r.Icon.ImgKey}
		}
	}
	return json.Marshal(m)
}

func (r *ComponentHeader) SetIconStandard(val string) *ComponentHeader {
	r.Icon = StandardIcon(val)
	return r
}

func (r *ComponentHeader) SetIconCustom(val string) *ComponentHeader {
	r.Icon = CustomIcon(val)
	return r
}

// SetTitle set ComponentHeader.Title attribute
func (r *ComponentHeader) SetTitle(val *ObjectText) *ComponentHeader {
	r.Title = val
	return r
}

// SetSubtitle set ComponentHeader.Subtitle attribute
func (r *ComponentHeader) SetSubtitle(val *ObjectText) *ComponentHeader {
	r.Subtitle = val
	return r
}

// SetTextTagList set ComponentHeader.TextTagList attribute
func (r *ComponentHeader) SetTextTagList(val ...*ObjectTextTag) *ComponentHeader {
	r.TextTagList = val
	return r
}

// SetI18nTextTagList set ComponentHeader.I18nTextTagList attribute
func (r *ComponentHeader) SetI18nTextTagList(val map[Language][]*ObjectTextTag) *ComponentHeader {
	r.I18nTextTagList = val
	return r
}

// SetTemplate set ComponentHeader.Template attribute
func (r *ComponentHeader) SetTemplate(val Template) *ComponentHeader {
	r.Template = val
	return r
}

// SetTemplateBlue set ComponentHeader.Template attribute to TemplateBlue
func (r *ComponentHeader) SetTemplateBlue() *ComponentHeader {
	r.Template = TemplateBlue
	return r
}

// SetTemplateWathet set ComponentHeader.Template attribute to TemplateWathet
func (r *ComponentHeader) SetTemplateWathet() *ComponentHeader {
	r.Template = TemplateWathet
	return r
}

// SetTemplateTurquoise set ComponentHeader.Template attribute to TemplateTurquoise
func (r *ComponentHeader) SetTemplateTurquoise() *ComponentHeader {
	r.Template = TemplateTurquoise
	return r
}

// SetTemplateGreen set ComponentHeader.Template attribute to TemplateGreen
func (r *ComponentHeader) SetTemplateGreen() *ComponentHeader {
	r.Template = TemplateGreen
	return r
}

// SetTemplateYellow set ComponentHeader.Template attribute to TemplateYellow
func (r *ComponentHeader) SetTemplateYellow() *ComponentHeader {
	r.Template = TemplateYellow
	return r
}

// SetTemplateOrange set ComponentHeader.Template attribute to TemplateOrange
func (r *ComponentHeader) SetTemplateOrange() *ComponentHeader {
	r.Template = TemplateOrange
	return r
}

// SetTemplateRed set ComponentHeader.Template attribute to TemplateRed
func (r *ComponentHeader) SetTemplateRed() *ComponentHeader {
	r.Template = TemplateRed
	return r
}

// SetTemplateCarmine set ComponentHeader.Template attribute to TemplateCarmine
func (r *ComponentHeader) SetTemplateCarmine() *ComponentHeader {
	r.Template = TemplateCarmine
	return r
}

// SetTemplateViolet set ComponentHeader.Template attribute to TemplateViolet
func (r *ComponentHeader) SetTemplateViolet() *ComponentHeader {
	r.Template = TemplateViolet
	return r
}

// SetTemplatePurple set ComponentHeader.Template attribute to TemplatePurple
func (r *ComponentHeader) SetTemplatePurple() *ComponentHeader {
	r.Template = TemplatePurple
	return r
}

// SetTemplateIndigo set ComponentHeader.Template attribute to TemplateIndigo
func (r *ComponentHeader) SetTemplateIndigo() *ComponentHeader {
	r.Template = TemplateIndigo
	return r
}

// SetTemplateGrey set ComponentHeader.Template attribute to TemplateGrey
func (r *ComponentHeader) SetTemplateGrey() *ComponentHeader {
	r.Template = TemplateGrey
	return r
}

// SetTemplateDefault set ComponentHeader.Template attribute to TemplateDefault
func (r *ComponentHeader) SetTemplateDefault() *ComponentHeader {
	r.Template = TemplateDefault
	return r
}

// SetIcon set ComponentHeader.Icon attribute
func (r *ComponentHeader) SetIcon(val *ObjectIcon) *ComponentHeader {
	r.Icon = val
	return r
}

// toMap conv ComponentHeader to map
func (r *ComponentHeader) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 5)
	if r.Title != nil {
		res["title"] = r.Title
	}
	if r.Subtitle != nil {
		res["subtitle"] = r.Subtitle
	}
	if len(r.TextTagList) != 0 {
		res["text_tag_list"] = r.TextTagList
	}
	if len(r.I18nTextTagList) != 0 {
		res["i18n_text_tag_list"] = r.I18nTextTagList
	}
	if r.Template != "" {
		res["template"] = r.Template
	}
	return res
}
