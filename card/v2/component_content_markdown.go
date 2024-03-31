package card

import (
	"encoding/json"
)

func Markdown(content string) *ComponentMarkdown {
	return &ComponentMarkdown{
		Content: content,
	}
}

// ComponentMarkdown 富文本（Markdown）
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/rich-text
//
// 卡片的富文本（Markdown）组件支持渲染文本、图片、分割线等元素。本文档介绍富文本组件以及对应的 JSON 参数说明。
//
// 注意事项
// 飞书卡片搭建工具暂不支持直接预览富文本组件中的有序列表、无序列表、代码块、和部分彩色文本样式语法。你可通过点击工具右上角的“向我发送预览”在飞书客户端中预览效果。
//
//go:generate generate_set_attrs -type=ComponentMarkdown
//go:generate generate_to_map -type=ComponentMarkdown
type ComponentMarkdown struct {
	componentBase

	// 设置文本内容的对齐方式。可取值有：
	//
	// left：左对齐
	// center：居中对齐
	// right：右对齐
	TextAlign TextAlign `json:"text_align,omitempty"`

	// 文本大小。可取值如下所示。如果你填写了其它值，卡片将展示为 normal 字段对应的字号。
	//
	// heading-0：特大标题（30px）
	// heading-1：一级标题（24px）
	// heading-2：二级标题（20 px）
	// heading-3：三级标题（18px）
	// heading-4：四级标题（16px）
	// heading：标题（16px）
	// normal：正文（14px）
	// notation：辅助信息（12px）
	// xxxx-large：30px
	// xxx-large：24px
	// xx-large：20px
	// x-large：18px
	// large：16px
	// medium：14px
	// small：12px
	// x-small：10px
	TextSize TextSize `json:"text_size,omitempty"`

	// 添加图标作为文本前缀图标。支持自定义或使用图标库中的图标。
	Icon *ObjectIcon `json:"icon,omitempty"`

	// 配置差异化跳转链接，实现“不同设备跳转链接不同”的效果。
	Href map[string]*ObjectMultiURL `json:"href,omitempty"`

	// Markdown 文本内容。了解支持的语法，参考下文。
	Content string `json:"content,omitempty"`
}

func (r ComponentMarkdown) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "markdown"
	return json.Marshal(m)
}

// SetTextAlign set ComponentMarkdown.TextAlign attribute
func (r *ComponentMarkdown) SetTextAlign(val TextAlign) *ComponentMarkdown {
	r.TextAlign = val
	return r
}

// SetTextAlignLeft set ComponentMarkdown.TextAlign attribute to TextAlignLeft
func (r *ComponentMarkdown) SetTextAlignLeft() *ComponentMarkdown {
	r.TextAlign = TextAlignLeft
	return r
}

// SetTextAlignCenter set ComponentMarkdown.TextAlign attribute to TextAlignCenter
func (r *ComponentMarkdown) SetTextAlignCenter() *ComponentMarkdown {
	r.TextAlign = TextAlignCenter
	return r
}

// SetTextAlignRight set ComponentMarkdown.TextAlign attribute to TextAlignRight
func (r *ComponentMarkdown) SetTextAlignRight() *ComponentMarkdown {
	r.TextAlign = TextAlignRight
	return r
}

// SetTextSize set ComponentMarkdown.TextSize attribute
func (r *ComponentMarkdown) SetTextSize(val TextSize) *ComponentMarkdown {
	r.TextSize = val
	return r
}

// SetTextSizeHeading0 set ComponentMarkdown.TextSize attribute to TextSizeHeading0
func (r *ComponentMarkdown) SetTextSizeHeading0() *ComponentMarkdown {
	r.TextSize = TextSizeHeading0
	return r
}

// SetTextSizeHeading1 set ComponentMarkdown.TextSize attribute to TextSizeHeading1
func (r *ComponentMarkdown) SetTextSizeHeading1() *ComponentMarkdown {
	r.TextSize = TextSizeHeading1
	return r
}

// SetTextSizeHeading2 set ComponentMarkdown.TextSize attribute to TextSizeHeading2
func (r *ComponentMarkdown) SetTextSizeHeading2() *ComponentMarkdown {
	r.TextSize = TextSizeHeading2
	return r
}

// SetTextSizeHeading3 set ComponentMarkdown.TextSize attribute to TextSizeHeading3
func (r *ComponentMarkdown) SetTextSizeHeading3() *ComponentMarkdown {
	r.TextSize = TextSizeHeading3
	return r
}

// SetTextSizeHeading4 set ComponentMarkdown.TextSize attribute to TextSizeHeading4
func (r *ComponentMarkdown) SetTextSizeHeading4() *ComponentMarkdown {
	r.TextSize = TextSizeHeading4
	return r
}

// SetTextSizeHeading set ComponentMarkdown.TextSize attribute to TextSizeHeading
func (r *ComponentMarkdown) SetTextSizeHeading() *ComponentMarkdown {
	r.TextSize = TextSizeHeading
	return r
}

// SetTextSizeNormal set ComponentMarkdown.TextSize attribute to TextSizeNormal
func (r *ComponentMarkdown) SetTextSizeNormal() *ComponentMarkdown {
	r.TextSize = TextSizeNormal
	return r
}

// SetTextSizeNotation set ComponentMarkdown.TextSize attribute to TextSizeNotation
func (r *ComponentMarkdown) SetTextSizeNotation() *ComponentMarkdown {
	r.TextSize = TextSizeNotation
	return r
}

// SetTextSizeXxxxLarge set ComponentMarkdown.TextSize attribute to TextSizeXxxxLarge
func (r *ComponentMarkdown) SetTextSizeXxxxLarge() *ComponentMarkdown {
	r.TextSize = TextSizeXxxxLarge
	return r
}

// SetTextSizeXxxLarge set ComponentMarkdown.TextSize attribute to TextSizeXxxLarge
func (r *ComponentMarkdown) SetTextSizeXxxLarge() *ComponentMarkdown {
	r.TextSize = TextSizeXxxLarge
	return r
}

// SetTextSizeXxLarge set ComponentMarkdown.TextSize attribute to TextSizeXxLarge
func (r *ComponentMarkdown) SetTextSizeXxLarge() *ComponentMarkdown {
	r.TextSize = TextSizeXxLarge
	return r
}

// SetTextSizeXLarge set ComponentMarkdown.TextSize attribute to TextSizeXLarge
func (r *ComponentMarkdown) SetTextSizeXLarge() *ComponentMarkdown {
	r.TextSize = TextSizeXLarge
	return r
}

// SetTextSizeLarge set ComponentMarkdown.TextSize attribute to TextSizeLarge
func (r *ComponentMarkdown) SetTextSizeLarge() *ComponentMarkdown {
	r.TextSize = TextSizeLarge
	return r
}

// SetTextSizeMedium set ComponentMarkdown.TextSize attribute to TextSizeMedium
func (r *ComponentMarkdown) SetTextSizeMedium() *ComponentMarkdown {
	r.TextSize = TextSizeMedium
	return r
}

// SetTextSizeSmall set ComponentMarkdown.TextSize attribute to TextSizeSmall
func (r *ComponentMarkdown) SetTextSizeSmall() *ComponentMarkdown {
	r.TextSize = TextSizeSmall
	return r
}

// SetTextSizeXSmall set ComponentMarkdown.TextSize attribute to TextSizeXSmall
func (r *ComponentMarkdown) SetTextSizeXSmall() *ComponentMarkdown {
	r.TextSize = TextSizeXSmall
	return r
}

// SetIcon set ComponentMarkdown.Icon attribute
func (r *ComponentMarkdown) SetIcon(val *ObjectIcon) *ComponentMarkdown {
	r.Icon = val
	return r
}

// SetHref set ComponentMarkdown.Href attribute
func (r *ComponentMarkdown) SetHref(val map[string]*ObjectMultiURL) *ComponentMarkdown {
	r.Href = val
	return r
}

// SetContent set ComponentMarkdown.Content attribute
func (r *ComponentMarkdown) SetContent(val string) *ComponentMarkdown {
	r.Content = val
	return r
}

// toMap conv ComponentMarkdown to map
func (r *ComponentMarkdown) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 5)
	if r.TextAlign != "" {
		res["text_align"] = r.TextAlign
	}
	if r.TextSize != "" {
		res["text_size"] = r.TextSize
	}
	if r.Icon != nil {
		res["icon"] = r.Icon
	}
	if len(r.Href) != 0 {
		res["href"] = r.Href
	}
	if r.Content != "" {
		res["content"] = r.Content
	}
	return res
}
