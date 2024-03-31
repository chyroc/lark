package card

import (
	"encoding/json"
)

func StandardIcon(token string) *ObjectIcon {
	return &ObjectIcon{
		Token: token,
	}
}

func CustomIcon(imgKey string) *ObjectIcon {
	return &ObjectIcon{
		ImgKey: imgKey,
	}
}

// ObjectIcon 图标
//
// 可用于标题: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/title
// 可用于富文本: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/rich-text
//
//go:generate generate_set_attrs -type=ObjectIcon
type ObjectIcon struct {
	componentBase

	// 图标库中图标的 token。当 tag 为 standard_icon 时生效。枚举值参见图标库。
	Token string `json:"token,omitempty"`

	// 图标的颜色。支持设置线性和面性图标（即 token 末尾为 outlined 或 filled 的图标）的颜色。当 tag 为 standard_icon 时生效。枚举值参见颜色枚举值。
	Color Color `json:"color,omitempty"`

	// 自定义前缀图标的图片 key。当 tag 为 custom_icon 时生效。图标 key 的获取方式：调用上传图片接口, 上传用于发送消息的图片, 并在返回值中获取图片的 image_key。
	ImgKey string `json:"img_key,omitempty"`

	// 图标的尺寸。支持 "[1,999] [1,999]px"。
	// size 参数仅支持在折叠面板中使用: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/containers/collapsible-panel#:~:text=%E5%9B%BE%E6%A0%87%E7%9A%84%E5%B0%BA%E5%AF%B8%E3%80%82%E6%94%AF%E6%8C%81%20%22%5B1%2C999%5D%20%5B1%2C999%5Dpx%22%E3%80%82
	Size string `json:"size,omitempty"`
}

func (r ObjectIcon) MarshalJSON() ([]byte, error) {
	// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/plain-text#:~:text=%E5%9B%BE%E6%A0%87%E7%B1%BB%E5%9E%8B%E7%9A%84,%E5%9B%BE%E7%89%87%E4%BD%9C%E4%B8%BA%E5%9B%BE%E6%A0%87%E3%80%82
	// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/interactive-components/button#:~:text=%E6%B7%BB%E5%8A%A0%E5%9B%BE%E6%A0%87%E4%BD%9C%E4%B8%BA%E6%96%87%E6%9C%AC%E5%89%8D%E7%BC%80%E5%9B%BE%E6%A0%87%E3%80%82%E6%94%AF%E6%8C%81%E8%87%AA%E5%AE%9A%E4%B9%89%E6%88%96%E4%BD%BF%E7%94%A8%E5%9B%BE%E6%A0%87%E5%BA%93%E4%B8%AD%E7%9A%84%E5%9B%BE%E6%A0%87%E3%80%82
	if r.Token != "" {
		if r.Color == "" {
			return json.Marshal(map[string]any{
				"token": r.Token,
				"tag":   "standard_icon",
			})
		} else {
			return json.Marshal(map[string]any{
				"token": r.Token,
				"color": r.Color,
				"tag":   "standard_icon",
			})
		}
	} else if r.ImgKey != "" {
		return json.Marshal(map[string]any{
			"img_key": r.ImgKey,
			"tag":     "custom_icon",
		})
	} else {
		return []byte("null"), nil
	}
}

// SetToken set ObjectIcon.Token attribute
func (r *ObjectIcon) SetToken(val string) *ObjectIcon {
	r.Token = val
	return r
}

// SetColor set ObjectIcon.Color attribute
func (r *ObjectIcon) SetColor(val Color) *ObjectIcon {
	r.Color = val
	return r
}

// SetColorDefault set ObjectIcon.Color attribute to ColorDefault
func (r *ObjectIcon) SetColorDefault() *ObjectIcon {
	r.Color = ColorDefault
	return r
}

// SetColorBgWhite set ObjectIcon.Color attribute to ColorBgWhite
func (r *ObjectIcon) SetColorBgWhite() *ObjectIcon {
	r.Color = ColorBgWhite
	return r
}

// SetColorWhite set ObjectIcon.Color attribute to ColorWhite
func (r *ObjectIcon) SetColorWhite() *ObjectIcon {
	r.Color = ColorWhite
	return r
}

// SetColorBlue set ObjectIcon.Color attribute to ColorBlue
func (r *ObjectIcon) SetColorBlue() *ObjectIcon {
	r.Color = ColorBlue
	return r
}

// SetColorBlue50 set ObjectIcon.Color attribute to ColorBlue50
func (r *ObjectIcon) SetColorBlue50() *ObjectIcon {
	r.Color = ColorBlue50
	return r
}

// SetColorBlue100 set ObjectIcon.Color attribute to ColorBlue100
func (r *ObjectIcon) SetColorBlue100() *ObjectIcon {
	r.Color = ColorBlue100
	return r
}

// SetColorBlue200 set ObjectIcon.Color attribute to ColorBlue200
func (r *ObjectIcon) SetColorBlue200() *ObjectIcon {
	r.Color = ColorBlue200
	return r
}

// SetColorBlue300 set ObjectIcon.Color attribute to ColorBlue300
func (r *ObjectIcon) SetColorBlue300() *ObjectIcon {
	r.Color = ColorBlue300
	return r
}

// SetColorBlue350 set ObjectIcon.Color attribute to ColorBlue350
func (r *ObjectIcon) SetColorBlue350() *ObjectIcon {
	r.Color = ColorBlue350
	return r
}

// SetColorBlue400 set ObjectIcon.Color attribute to ColorBlue400
func (r *ObjectIcon) SetColorBlue400() *ObjectIcon {
	r.Color = ColorBlue400
	return r
}

// SetColorBlue500 set ObjectIcon.Color attribute to ColorBlue500
func (r *ObjectIcon) SetColorBlue500() *ObjectIcon {
	r.Color = ColorBlue500
	return r
}

// SetColorBlue600 set ObjectIcon.Color attribute to ColorBlue600
func (r *ObjectIcon) SetColorBlue600() *ObjectIcon {
	r.Color = ColorBlue600
	return r
}

// SetColorBlue700 set ObjectIcon.Color attribute to ColorBlue700
func (r *ObjectIcon) SetColorBlue700() *ObjectIcon {
	r.Color = ColorBlue700
	return r
}

// SetColorBlue800 set ObjectIcon.Color attribute to ColorBlue800
func (r *ObjectIcon) SetColorBlue800() *ObjectIcon {
	r.Color = ColorBlue800
	return r
}

// SetColorBlue900 set ObjectIcon.Color attribute to ColorBlue900
func (r *ObjectIcon) SetColorBlue900() *ObjectIcon {
	r.Color = ColorBlue900
	return r
}

// SetColorCarmine set ObjectIcon.Color attribute to ColorCarmine
func (r *ObjectIcon) SetColorCarmine() *ObjectIcon {
	r.Color = ColorCarmine
	return r
}

// SetColorCarmine50 set ObjectIcon.Color attribute to ColorCarmine50
func (r *ObjectIcon) SetColorCarmine50() *ObjectIcon {
	r.Color = ColorCarmine50
	return r
}

// SetColorCarmine100 set ObjectIcon.Color attribute to ColorCarmine100
func (r *ObjectIcon) SetColorCarmine100() *ObjectIcon {
	r.Color = ColorCarmine100
	return r
}

// SetColorCarmine200 set ObjectIcon.Color attribute to ColorCarmine200
func (r *ObjectIcon) SetColorCarmine200() *ObjectIcon {
	r.Color = ColorCarmine200
	return r
}

// SetColorCarmine300 set ObjectIcon.Color attribute to ColorCarmine300
func (r *ObjectIcon) SetColorCarmine300() *ObjectIcon {
	r.Color = ColorCarmine300
	return r
}

// SetColorCarmine350 set ObjectIcon.Color attribute to ColorCarmine350
func (r *ObjectIcon) SetColorCarmine350() *ObjectIcon {
	r.Color = ColorCarmine350
	return r
}

// SetColorCarmine400 set ObjectIcon.Color attribute to ColorCarmine400
func (r *ObjectIcon) SetColorCarmine400() *ObjectIcon {
	r.Color = ColorCarmine400
	return r
}

// SetColorCarmine500 set ObjectIcon.Color attribute to ColorCarmine500
func (r *ObjectIcon) SetColorCarmine500() *ObjectIcon {
	r.Color = ColorCarmine500
	return r
}

// SetColorCarmine600 set ObjectIcon.Color attribute to ColorCarmine600
func (r *ObjectIcon) SetColorCarmine600() *ObjectIcon {
	r.Color = ColorCarmine600
	return r
}

// SetColorCarmine700 set ObjectIcon.Color attribute to ColorCarmine700
func (r *ObjectIcon) SetColorCarmine700() *ObjectIcon {
	r.Color = ColorCarmine700
	return r
}

// SetColorCarmine800 set ObjectIcon.Color attribute to ColorCarmine800
func (r *ObjectIcon) SetColorCarmine800() *ObjectIcon {
	r.Color = ColorCarmine800
	return r
}

// SetColorCarmine900 set ObjectIcon.Color attribute to ColorCarmine900
func (r *ObjectIcon) SetColorCarmine900() *ObjectIcon {
	r.Color = ColorCarmine900
	return r
}

// SetColorGreen set ObjectIcon.Color attribute to ColorGreen
func (r *ObjectIcon) SetColorGreen() *ObjectIcon {
	r.Color = ColorGreen
	return r
}

// SetColorGreen50 set ObjectIcon.Color attribute to ColorGreen50
func (r *ObjectIcon) SetColorGreen50() *ObjectIcon {
	r.Color = ColorGreen50
	return r
}

// SetColorGreen100 set ObjectIcon.Color attribute to ColorGreen100
func (r *ObjectIcon) SetColorGreen100() *ObjectIcon {
	r.Color = ColorGreen100
	return r
}

// SetColorGreen200 set ObjectIcon.Color attribute to ColorGreen200
func (r *ObjectIcon) SetColorGreen200() *ObjectIcon {
	r.Color = ColorGreen200
	return r
}

// SetColorGreen300 set ObjectIcon.Color attribute to ColorGreen300
func (r *ObjectIcon) SetColorGreen300() *ObjectIcon {
	r.Color = ColorGreen300
	return r
}

// SetColorGreen350 set ObjectIcon.Color attribute to ColorGreen350
func (r *ObjectIcon) SetColorGreen350() *ObjectIcon {
	r.Color = ColorGreen350
	return r
}

// SetColorGreen400 set ObjectIcon.Color attribute to ColorGreen400
func (r *ObjectIcon) SetColorGreen400() *ObjectIcon {
	r.Color = ColorGreen400
	return r
}

// SetColorGreen500 set ObjectIcon.Color attribute to ColorGreen500
func (r *ObjectIcon) SetColorGreen500() *ObjectIcon {
	r.Color = ColorGreen500
	return r
}

// SetColorGreen600 set ObjectIcon.Color attribute to ColorGreen600
func (r *ObjectIcon) SetColorGreen600() *ObjectIcon {
	r.Color = ColorGreen600
	return r
}

// SetColorGreen700 set ObjectIcon.Color attribute to ColorGreen700
func (r *ObjectIcon) SetColorGreen700() *ObjectIcon {
	r.Color = ColorGreen700
	return r
}

// SetColorGreen800 set ObjectIcon.Color attribute to ColorGreen800
func (r *ObjectIcon) SetColorGreen800() *ObjectIcon {
	r.Color = ColorGreen800
	return r
}

// SetColorGreen900 set ObjectIcon.Color attribute to ColorGreen900
func (r *ObjectIcon) SetColorGreen900() *ObjectIcon {
	r.Color = ColorGreen900
	return r
}

// SetColorIndigo set ObjectIcon.Color attribute to ColorIndigo
func (r *ObjectIcon) SetColorIndigo() *ObjectIcon {
	r.Color = ColorIndigo
	return r
}

// SetColorIndigo50 set ObjectIcon.Color attribute to ColorIndigo50
func (r *ObjectIcon) SetColorIndigo50() *ObjectIcon {
	r.Color = ColorIndigo50
	return r
}

// SetColorIndigo100 set ObjectIcon.Color attribute to ColorIndigo100
func (r *ObjectIcon) SetColorIndigo100() *ObjectIcon {
	r.Color = ColorIndigo100
	return r
}

// SetColorIndigo200 set ObjectIcon.Color attribute to ColorIndigo200
func (r *ObjectIcon) SetColorIndigo200() *ObjectIcon {
	r.Color = ColorIndigo200
	return r
}

// SetColorIndigo300 set ObjectIcon.Color attribute to ColorIndigo300
func (r *ObjectIcon) SetColorIndigo300() *ObjectIcon {
	r.Color = ColorIndigo300
	return r
}

// SetColorIndigo350 set ObjectIcon.Color attribute to ColorIndigo350
func (r *ObjectIcon) SetColorIndigo350() *ObjectIcon {
	r.Color = ColorIndigo350
	return r
}

// SetColorIndigo400 set ObjectIcon.Color attribute to ColorIndigo400
func (r *ObjectIcon) SetColorIndigo400() *ObjectIcon {
	r.Color = ColorIndigo400
	return r
}

// SetColorIndigo500 set ObjectIcon.Color attribute to ColorIndigo500
func (r *ObjectIcon) SetColorIndigo500() *ObjectIcon {
	r.Color = ColorIndigo500
	return r
}

// SetColorIndigo600 set ObjectIcon.Color attribute to ColorIndigo600
func (r *ObjectIcon) SetColorIndigo600() *ObjectIcon {
	r.Color = ColorIndigo600
	return r
}

// SetColorIndigo700 set ObjectIcon.Color attribute to ColorIndigo700
func (r *ObjectIcon) SetColorIndigo700() *ObjectIcon {
	r.Color = ColorIndigo700
	return r
}

// SetColorIndigo800 set ObjectIcon.Color attribute to ColorIndigo800
func (r *ObjectIcon) SetColorIndigo800() *ObjectIcon {
	r.Color = ColorIndigo800
	return r
}

// SetColorIndigo900 set ObjectIcon.Color attribute to ColorIndigo900
func (r *ObjectIcon) SetColorIndigo900() *ObjectIcon {
	r.Color = ColorIndigo900
	return r
}

// SetColorLime set ObjectIcon.Color attribute to ColorLime
func (r *ObjectIcon) SetColorLime() *ObjectIcon {
	r.Color = ColorLime
	return r
}

// SetColorLime50 set ObjectIcon.Color attribute to ColorLime50
func (r *ObjectIcon) SetColorLime50() *ObjectIcon {
	r.Color = ColorLime50
	return r
}

// SetColorLime100 set ObjectIcon.Color attribute to ColorLime100
func (r *ObjectIcon) SetColorLime100() *ObjectIcon {
	r.Color = ColorLime100
	return r
}

// SetColorLime200 set ObjectIcon.Color attribute to ColorLime200
func (r *ObjectIcon) SetColorLime200() *ObjectIcon {
	r.Color = ColorLime200
	return r
}

// SetColorLime300 set ObjectIcon.Color attribute to ColorLime300
func (r *ObjectIcon) SetColorLime300() *ObjectIcon {
	r.Color = ColorLime300
	return r
}

// SetColorLime350 set ObjectIcon.Color attribute to ColorLime350
func (r *ObjectIcon) SetColorLime350() *ObjectIcon {
	r.Color = ColorLime350
	return r
}

// SetColorLime400 set ObjectIcon.Color attribute to ColorLime400
func (r *ObjectIcon) SetColorLime400() *ObjectIcon {
	r.Color = ColorLime400
	return r
}

// SetColorLime500 set ObjectIcon.Color attribute to ColorLime500
func (r *ObjectIcon) SetColorLime500() *ObjectIcon {
	r.Color = ColorLime500
	return r
}

// SetColorLime600 set ObjectIcon.Color attribute to ColorLime600
func (r *ObjectIcon) SetColorLime600() *ObjectIcon {
	r.Color = ColorLime600
	return r
}

// SetColorLime700 set ObjectIcon.Color attribute to ColorLime700
func (r *ObjectIcon) SetColorLime700() *ObjectIcon {
	r.Color = ColorLime700
	return r
}

// SetColorLime800 set ObjectIcon.Color attribute to ColorLime800
func (r *ObjectIcon) SetColorLime800() *ObjectIcon {
	r.Color = ColorLime800
	return r
}

// SetColorLime900 set ObjectIcon.Color attribute to ColorLime900
func (r *ObjectIcon) SetColorLime900() *ObjectIcon {
	r.Color = ColorLime900
	return r
}

// SetColorGrey set ObjectIcon.Color attribute to ColorGrey
func (r *ObjectIcon) SetColorGrey() *ObjectIcon {
	r.Color = ColorGrey
	return r
}

// SetColorGrey00 set ObjectIcon.Color attribute to ColorGrey00
func (r *ObjectIcon) SetColorGrey00() *ObjectIcon {
	r.Color = ColorGrey00
	return r
}

// SetColorGrey50 set ObjectIcon.Color attribute to ColorGrey50
func (r *ObjectIcon) SetColorGrey50() *ObjectIcon {
	r.Color = ColorGrey50
	return r
}

// SetColorGrey100 set ObjectIcon.Color attribute to ColorGrey100
func (r *ObjectIcon) SetColorGrey100() *ObjectIcon {
	r.Color = ColorGrey100
	return r
}

// SetColorGrey200 set ObjectIcon.Color attribute to ColorGrey200
func (r *ObjectIcon) SetColorGrey200() *ObjectIcon {
	r.Color = ColorGrey200
	return r
}

// SetColorGrey300 set ObjectIcon.Color attribute to ColorGrey300
func (r *ObjectIcon) SetColorGrey300() *ObjectIcon {
	r.Color = ColorGrey300
	return r
}

// SetColorGrey350 set ObjectIcon.Color attribute to ColorGrey350
func (r *ObjectIcon) SetColorGrey350() *ObjectIcon {
	r.Color = ColorGrey350
	return r
}

// SetColorGrey400 set ObjectIcon.Color attribute to ColorGrey400
func (r *ObjectIcon) SetColorGrey400() *ObjectIcon {
	r.Color = ColorGrey400
	return r
}

// SetColorGrey500 set ObjectIcon.Color attribute to ColorGrey500
func (r *ObjectIcon) SetColorGrey500() *ObjectIcon {
	r.Color = ColorGrey500
	return r
}

// SetColorGrey600 set ObjectIcon.Color attribute to ColorGrey600
func (r *ObjectIcon) SetColorGrey600() *ObjectIcon {
	r.Color = ColorGrey600
	return r
}

// SetColorGrey650 set ObjectIcon.Color attribute to ColorGrey650
func (r *ObjectIcon) SetColorGrey650() *ObjectIcon {
	r.Color = ColorGrey650
	return r
}

// SetColorGrey700 set ObjectIcon.Color attribute to ColorGrey700
func (r *ObjectIcon) SetColorGrey700() *ObjectIcon {
	r.Color = ColorGrey700
	return r
}

// SetColorGrey800 set ObjectIcon.Color attribute to ColorGrey800
func (r *ObjectIcon) SetColorGrey800() *ObjectIcon {
	r.Color = ColorGrey800
	return r
}

// SetColorGrey900 set ObjectIcon.Color attribute to ColorGrey900
func (r *ObjectIcon) SetColorGrey900() *ObjectIcon {
	r.Color = ColorGrey900
	return r
}

// SetColorGrey950 set ObjectIcon.Color attribute to ColorGrey950
func (r *ObjectIcon) SetColorGrey950() *ObjectIcon {
	r.Color = ColorGrey950
	return r
}

// SetColorGrey1000 set ObjectIcon.Color attribute to ColorGrey1000
func (r *ObjectIcon) SetColorGrey1000() *ObjectIcon {
	r.Color = ColorGrey1000
	return r
}

// SetColorOrange set ObjectIcon.Color attribute to ColorOrange
func (r *ObjectIcon) SetColorOrange() *ObjectIcon {
	r.Color = ColorOrange
	return r
}

// SetColorOrange50 set ObjectIcon.Color attribute to ColorOrange50
func (r *ObjectIcon) SetColorOrange50() *ObjectIcon {
	r.Color = ColorOrange50
	return r
}

// SetColorOrange100 set ObjectIcon.Color attribute to ColorOrange100
func (r *ObjectIcon) SetColorOrange100() *ObjectIcon {
	r.Color = ColorOrange100
	return r
}

// SetColorOrange200 set ObjectIcon.Color attribute to ColorOrange200
func (r *ObjectIcon) SetColorOrange200() *ObjectIcon {
	r.Color = ColorOrange200
	return r
}

// SetColorOrange300 set ObjectIcon.Color attribute to ColorOrange300
func (r *ObjectIcon) SetColorOrange300() *ObjectIcon {
	r.Color = ColorOrange300
	return r
}

// SetColorOrange350 set ObjectIcon.Color attribute to ColorOrange350
func (r *ObjectIcon) SetColorOrange350() *ObjectIcon {
	r.Color = ColorOrange350
	return r
}

// SetColorOrange400 set ObjectIcon.Color attribute to ColorOrange400
func (r *ObjectIcon) SetColorOrange400() *ObjectIcon {
	r.Color = ColorOrange400
	return r
}

// SetColorOrange500 set ObjectIcon.Color attribute to ColorOrange500
func (r *ObjectIcon) SetColorOrange500() *ObjectIcon {
	r.Color = ColorOrange500
	return r
}

// SetColorOrange600 set ObjectIcon.Color attribute to ColorOrange600
func (r *ObjectIcon) SetColorOrange600() *ObjectIcon {
	r.Color = ColorOrange600
	return r
}

// SetColorOrange700 set ObjectIcon.Color attribute to ColorOrange700
func (r *ObjectIcon) SetColorOrange700() *ObjectIcon {
	r.Color = ColorOrange700
	return r
}

// SetColorOrange800 set ObjectIcon.Color attribute to ColorOrange800
func (r *ObjectIcon) SetColorOrange800() *ObjectIcon {
	r.Color = ColorOrange800
	return r
}

// SetColorOrange900 set ObjectIcon.Color attribute to ColorOrange900
func (r *ObjectIcon) SetColorOrange900() *ObjectIcon {
	r.Color = ColorOrange900
	return r
}

// SetColorPurple set ObjectIcon.Color attribute to ColorPurple
func (r *ObjectIcon) SetColorPurple() *ObjectIcon {
	r.Color = ColorPurple
	return r
}

// SetColorPurple50 set ObjectIcon.Color attribute to ColorPurple50
func (r *ObjectIcon) SetColorPurple50() *ObjectIcon {
	r.Color = ColorPurple50
	return r
}

// SetColorPurple100 set ObjectIcon.Color attribute to ColorPurple100
func (r *ObjectIcon) SetColorPurple100() *ObjectIcon {
	r.Color = ColorPurple100
	return r
}

// SetColorPurple200 set ObjectIcon.Color attribute to ColorPurple200
func (r *ObjectIcon) SetColorPurple200() *ObjectIcon {
	r.Color = ColorPurple200
	return r
}

// SetColorPurple300 set ObjectIcon.Color attribute to ColorPurple300
func (r *ObjectIcon) SetColorPurple300() *ObjectIcon {
	r.Color = ColorPurple300
	return r
}

// SetColorPurple350 set ObjectIcon.Color attribute to ColorPurple350
func (r *ObjectIcon) SetColorPurple350() *ObjectIcon {
	r.Color = ColorPurple350
	return r
}

// SetColorPurple400 set ObjectIcon.Color attribute to ColorPurple400
func (r *ObjectIcon) SetColorPurple400() *ObjectIcon {
	r.Color = ColorPurple400
	return r
}

// SetColorPurple500 set ObjectIcon.Color attribute to ColorPurple500
func (r *ObjectIcon) SetColorPurple500() *ObjectIcon {
	r.Color = ColorPurple500
	return r
}

// SetColorPurple600 set ObjectIcon.Color attribute to ColorPurple600
func (r *ObjectIcon) SetColorPurple600() *ObjectIcon {
	r.Color = ColorPurple600
	return r
}

// SetColorPurple700 set ObjectIcon.Color attribute to ColorPurple700
func (r *ObjectIcon) SetColorPurple700() *ObjectIcon {
	r.Color = ColorPurple700
	return r
}

// SetColorPurple800 set ObjectIcon.Color attribute to ColorPurple800
func (r *ObjectIcon) SetColorPurple800() *ObjectIcon {
	r.Color = ColorPurple800
	return r
}

// SetColorPurple900 set ObjectIcon.Color attribute to ColorPurple900
func (r *ObjectIcon) SetColorPurple900() *ObjectIcon {
	r.Color = ColorPurple900
	return r
}

// SetColorRed set ObjectIcon.Color attribute to ColorRed
func (r *ObjectIcon) SetColorRed() *ObjectIcon {
	r.Color = ColorRed
	return r
}

// SetColorRed50 set ObjectIcon.Color attribute to ColorRed50
func (r *ObjectIcon) SetColorRed50() *ObjectIcon {
	r.Color = ColorRed50
	return r
}

// SetColorRed100 set ObjectIcon.Color attribute to ColorRed100
func (r *ObjectIcon) SetColorRed100() *ObjectIcon {
	r.Color = ColorRed100
	return r
}

// SetColorRed200 set ObjectIcon.Color attribute to ColorRed200
func (r *ObjectIcon) SetColorRed200() *ObjectIcon {
	r.Color = ColorRed200
	return r
}

// SetColorRed300 set ObjectIcon.Color attribute to ColorRed300
func (r *ObjectIcon) SetColorRed300() *ObjectIcon {
	r.Color = ColorRed300
	return r
}

// SetColorRed350 set ObjectIcon.Color attribute to ColorRed350
func (r *ObjectIcon) SetColorRed350() *ObjectIcon {
	r.Color = ColorRed350
	return r
}

// SetColorRed400 set ObjectIcon.Color attribute to ColorRed400
func (r *ObjectIcon) SetColorRed400() *ObjectIcon {
	r.Color = ColorRed400
	return r
}

// SetColorRed500 set ObjectIcon.Color attribute to ColorRed500
func (r *ObjectIcon) SetColorRed500() *ObjectIcon {
	r.Color = ColorRed500
	return r
}

// SetColorRed600 set ObjectIcon.Color attribute to ColorRed600
func (r *ObjectIcon) SetColorRed600() *ObjectIcon {
	r.Color = ColorRed600
	return r
}

// SetColorRed700 set ObjectIcon.Color attribute to ColorRed700
func (r *ObjectIcon) SetColorRed700() *ObjectIcon {
	r.Color = ColorRed700
	return r
}

// SetColorRed800 set ObjectIcon.Color attribute to ColorRed800
func (r *ObjectIcon) SetColorRed800() *ObjectIcon {
	r.Color = ColorRed800
	return r
}

// SetColorRed900 set ObjectIcon.Color attribute to ColorRed900
func (r *ObjectIcon) SetColorRed900() *ObjectIcon {
	r.Color = ColorRed900
	return r
}

// SetColorSunflower set ObjectIcon.Color attribute to ColorSunflower
func (r *ObjectIcon) SetColorSunflower() *ObjectIcon {
	r.Color = ColorSunflower
	return r
}

// SetColorSunflower50 set ObjectIcon.Color attribute to ColorSunflower50
func (r *ObjectIcon) SetColorSunflower50() *ObjectIcon {
	r.Color = ColorSunflower50
	return r
}

// SetColorSunflower100 set ObjectIcon.Color attribute to ColorSunflower100
func (r *ObjectIcon) SetColorSunflower100() *ObjectIcon {
	r.Color = ColorSunflower100
	return r
}

// SetColorSunflower200 set ObjectIcon.Color attribute to ColorSunflower200
func (r *ObjectIcon) SetColorSunflower200() *ObjectIcon {
	r.Color = ColorSunflower200
	return r
}

// SetColorSunflower300 set ObjectIcon.Color attribute to ColorSunflower300
func (r *ObjectIcon) SetColorSunflower300() *ObjectIcon {
	r.Color = ColorSunflower300
	return r
}

// SetColorSunflower350 set ObjectIcon.Color attribute to ColorSunflower350
func (r *ObjectIcon) SetColorSunflower350() *ObjectIcon {
	r.Color = ColorSunflower350
	return r
}

// SetColorSunflower400 set ObjectIcon.Color attribute to ColorSunflower400
func (r *ObjectIcon) SetColorSunflower400() *ObjectIcon {
	r.Color = ColorSunflower400
	return r
}

// SetColorSunflower500 set ObjectIcon.Color attribute to ColorSunflower500
func (r *ObjectIcon) SetColorSunflower500() *ObjectIcon {
	r.Color = ColorSunflower500
	return r
}

// SetColorSunflower600 set ObjectIcon.Color attribute to ColorSunflower600
func (r *ObjectIcon) SetColorSunflower600() *ObjectIcon {
	r.Color = ColorSunflower600
	return r
}

// SetColorSunflower700 set ObjectIcon.Color attribute to ColorSunflower700
func (r *ObjectIcon) SetColorSunflower700() *ObjectIcon {
	r.Color = ColorSunflower700
	return r
}

// SetColorSunflower800 set ObjectIcon.Color attribute to ColorSunflower800
func (r *ObjectIcon) SetColorSunflower800() *ObjectIcon {
	r.Color = ColorSunflower800
	return r
}

// SetColorSunflower900 set ObjectIcon.Color attribute to ColorSunflower900
func (r *ObjectIcon) SetColorSunflower900() *ObjectIcon {
	r.Color = ColorSunflower900
	return r
}

// SetColorTurquoise set ObjectIcon.Color attribute to ColorTurquoise
func (r *ObjectIcon) SetColorTurquoise() *ObjectIcon {
	r.Color = ColorTurquoise
	return r
}

// SetColorTurquoise50 set ObjectIcon.Color attribute to ColorTurquoise50
func (r *ObjectIcon) SetColorTurquoise50() *ObjectIcon {
	r.Color = ColorTurquoise50
	return r
}

// SetColorTurquoise100 set ObjectIcon.Color attribute to ColorTurquoise100
func (r *ObjectIcon) SetColorTurquoise100() *ObjectIcon {
	r.Color = ColorTurquoise100
	return r
}

// SetColorTurquoise200 set ObjectIcon.Color attribute to ColorTurquoise200
func (r *ObjectIcon) SetColorTurquoise200() *ObjectIcon {
	r.Color = ColorTurquoise200
	return r
}

// SetColorTurquoise300 set ObjectIcon.Color attribute to ColorTurquoise300
func (r *ObjectIcon) SetColorTurquoise300() *ObjectIcon {
	r.Color = ColorTurquoise300
	return r
}

// SetColorTurquoise350 set ObjectIcon.Color attribute to ColorTurquoise350
func (r *ObjectIcon) SetColorTurquoise350() *ObjectIcon {
	r.Color = ColorTurquoise350
	return r
}

// SetColorTurquoise400 set ObjectIcon.Color attribute to ColorTurquoise400
func (r *ObjectIcon) SetColorTurquoise400() *ObjectIcon {
	r.Color = ColorTurquoise400
	return r
}

// SetColorTurquoise500 set ObjectIcon.Color attribute to ColorTurquoise500
func (r *ObjectIcon) SetColorTurquoise500() *ObjectIcon {
	r.Color = ColorTurquoise500
	return r
}

// SetColorTurquoise600 set ObjectIcon.Color attribute to ColorTurquoise600
func (r *ObjectIcon) SetColorTurquoise600() *ObjectIcon {
	r.Color = ColorTurquoise600
	return r
}

// SetColorTurquoise700 set ObjectIcon.Color attribute to ColorTurquoise700
func (r *ObjectIcon) SetColorTurquoise700() *ObjectIcon {
	r.Color = ColorTurquoise700
	return r
}

// SetColorTurquoise800 set ObjectIcon.Color attribute to ColorTurquoise800
func (r *ObjectIcon) SetColorTurquoise800() *ObjectIcon {
	r.Color = ColorTurquoise800
	return r
}

// SetColorTurquoise900 set ObjectIcon.Color attribute to ColorTurquoise900
func (r *ObjectIcon) SetColorTurquoise900() *ObjectIcon {
	r.Color = ColorTurquoise900
	return r
}

// SetColorViolet set ObjectIcon.Color attribute to ColorViolet
func (r *ObjectIcon) SetColorViolet() *ObjectIcon {
	r.Color = ColorViolet
	return r
}

// SetColorViolet50 set ObjectIcon.Color attribute to ColorViolet50
func (r *ObjectIcon) SetColorViolet50() *ObjectIcon {
	r.Color = ColorViolet50
	return r
}

// SetColorViolet100 set ObjectIcon.Color attribute to ColorViolet100
func (r *ObjectIcon) SetColorViolet100() *ObjectIcon {
	r.Color = ColorViolet100
	return r
}

// SetColorViolet200 set ObjectIcon.Color attribute to ColorViolet200
func (r *ObjectIcon) SetColorViolet200() *ObjectIcon {
	r.Color = ColorViolet200
	return r
}

// SetColorViolet300 set ObjectIcon.Color attribute to ColorViolet300
func (r *ObjectIcon) SetColorViolet300() *ObjectIcon {
	r.Color = ColorViolet300
	return r
}

// SetColorViolet350 set ObjectIcon.Color attribute to ColorViolet350
func (r *ObjectIcon) SetColorViolet350() *ObjectIcon {
	r.Color = ColorViolet350
	return r
}

// SetColorViolet400 set ObjectIcon.Color attribute to ColorViolet400
func (r *ObjectIcon) SetColorViolet400() *ObjectIcon {
	r.Color = ColorViolet400
	return r
}

// SetColorViolet500 set ObjectIcon.Color attribute to ColorViolet500
func (r *ObjectIcon) SetColorViolet500() *ObjectIcon {
	r.Color = ColorViolet500
	return r
}

// SetColorViolet600 set ObjectIcon.Color attribute to ColorViolet600
func (r *ObjectIcon) SetColorViolet600() *ObjectIcon {
	r.Color = ColorViolet600
	return r
}

// SetColorViolet700 set ObjectIcon.Color attribute to ColorViolet700
func (r *ObjectIcon) SetColorViolet700() *ObjectIcon {
	r.Color = ColorViolet700
	return r
}

// SetColorViolet800 set ObjectIcon.Color attribute to ColorViolet800
func (r *ObjectIcon) SetColorViolet800() *ObjectIcon {
	r.Color = ColorViolet800
	return r
}

// SetColorViolet900 set ObjectIcon.Color attribute to ColorViolet900
func (r *ObjectIcon) SetColorViolet900() *ObjectIcon {
	r.Color = ColorViolet900
	return r
}

// SetColorWathet set ObjectIcon.Color attribute to ColorWathet
func (r *ObjectIcon) SetColorWathet() *ObjectIcon {
	r.Color = ColorWathet
	return r
}

// SetColorWathet50 set ObjectIcon.Color attribute to ColorWathet50
func (r *ObjectIcon) SetColorWathet50() *ObjectIcon {
	r.Color = ColorWathet50
	return r
}

// SetColorWathet100 set ObjectIcon.Color attribute to ColorWathet100
func (r *ObjectIcon) SetColorWathet100() *ObjectIcon {
	r.Color = ColorWathet100
	return r
}

// SetColorWathet200 set ObjectIcon.Color attribute to ColorWathet200
func (r *ObjectIcon) SetColorWathet200() *ObjectIcon {
	r.Color = ColorWathet200
	return r
}

// SetColorWathet300 set ObjectIcon.Color attribute to ColorWathet300
func (r *ObjectIcon) SetColorWathet300() *ObjectIcon {
	r.Color = ColorWathet300
	return r
}

// SetColorWathet350 set ObjectIcon.Color attribute to ColorWathet350
func (r *ObjectIcon) SetColorWathet350() *ObjectIcon {
	r.Color = ColorWathet350
	return r
}

// SetColorWathet400 set ObjectIcon.Color attribute to ColorWathet400
func (r *ObjectIcon) SetColorWathet400() *ObjectIcon {
	r.Color = ColorWathet400
	return r
}

// SetColorWathet500 set ObjectIcon.Color attribute to ColorWathet500
func (r *ObjectIcon) SetColorWathet500() *ObjectIcon {
	r.Color = ColorWathet500
	return r
}

// SetColorWathet600 set ObjectIcon.Color attribute to ColorWathet600
func (r *ObjectIcon) SetColorWathet600() *ObjectIcon {
	r.Color = ColorWathet600
	return r
}

// SetColorWathet700 set ObjectIcon.Color attribute to ColorWathet700
func (r *ObjectIcon) SetColorWathet700() *ObjectIcon {
	r.Color = ColorWathet700
	return r
}

// SetColorWathet800 set ObjectIcon.Color attribute to ColorWathet800
func (r *ObjectIcon) SetColorWathet800() *ObjectIcon {
	r.Color = ColorWathet800
	return r
}

// SetColorWathet900 set ObjectIcon.Color attribute to ColorWathet900
func (r *ObjectIcon) SetColorWathet900() *ObjectIcon {
	r.Color = ColorWathet900
	return r
}

// SetColorYellow set ObjectIcon.Color attribute to ColorYellow
func (r *ObjectIcon) SetColorYellow() *ObjectIcon {
	r.Color = ColorYellow
	return r
}

// SetColorYellow50 set ObjectIcon.Color attribute to ColorYellow50
func (r *ObjectIcon) SetColorYellow50() *ObjectIcon {
	r.Color = ColorYellow50
	return r
}

// SetColorYellow100 set ObjectIcon.Color attribute to ColorYellow100
func (r *ObjectIcon) SetColorYellow100() *ObjectIcon {
	r.Color = ColorYellow100
	return r
}

// SetColorYellow200 set ObjectIcon.Color attribute to ColorYellow200
func (r *ObjectIcon) SetColorYellow200() *ObjectIcon {
	r.Color = ColorYellow200
	return r
}

// SetColorYellow300 set ObjectIcon.Color attribute to ColorYellow300
func (r *ObjectIcon) SetColorYellow300() *ObjectIcon {
	r.Color = ColorYellow300
	return r
}

// SetColorYellow350 set ObjectIcon.Color attribute to ColorYellow350
func (r *ObjectIcon) SetColorYellow350() *ObjectIcon {
	r.Color = ColorYellow350
	return r
}

// SetColorYellow400 set ObjectIcon.Color attribute to ColorYellow400
func (r *ObjectIcon) SetColorYellow400() *ObjectIcon {
	r.Color = ColorYellow400
	return r
}

// SetColorYellow500 set ObjectIcon.Color attribute to ColorYellow500
func (r *ObjectIcon) SetColorYellow500() *ObjectIcon {
	r.Color = ColorYellow500
	return r
}

// SetColorYellow600 set ObjectIcon.Color attribute to ColorYellow600
func (r *ObjectIcon) SetColorYellow600() *ObjectIcon {
	r.Color = ColorYellow600
	return r
}

// SetColorYellow700 set ObjectIcon.Color attribute to ColorYellow700
func (r *ObjectIcon) SetColorYellow700() *ObjectIcon {
	r.Color = ColorYellow700
	return r
}

// SetColorYellow800 set ObjectIcon.Color attribute to ColorYellow800
func (r *ObjectIcon) SetColorYellow800() *ObjectIcon {
	r.Color = ColorYellow800
	return r
}

// SetColorYellow900 set ObjectIcon.Color attribute to ColorYellow900
func (r *ObjectIcon) SetColorYellow900() *ObjectIcon {
	r.Color = ColorYellow900
	return r
}

// SetImgKey set ObjectIcon.ImgKey attribute
func (r *ObjectIcon) SetImgKey(val string) *ObjectIcon {
	r.ImgKey = val
	return r
}

// SetSize set ObjectIcon.Size attribute
func (r *ObjectIcon) SetSize(val string) *ObjectIcon {
	r.Size = val
	return r
}
