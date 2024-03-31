package card

import "encoding/json"

func PlainText(text string) *ComponentPlainText {
	return &ComponentPlainText{
		Text: Text(text),
	}
}

// ComponentPlainText 普通文本
//
// docs: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-components/content-components/plain-text
//
// 卡片的普通文本组件支持添加普通文本和前缀图标, 并设置文本大小、颜色、对齐方式等展示样式。
//
//go:generate generate_set_attrs -type=ComponentPlainText
//go:generate generate_to_map -type=ComponentPlainText
type ComponentPlainText struct {
	componentBase

	// 配置卡片的普通文本信息
	Text *ObjectText `json:"text,omitempty"`
	Icon *ObjectIcon `json:"icon,omitempty"`
}

func (r ComponentPlainText) MarshalJSON() ([]byte, error) {
	m := r.toMap()
	m["tag"] = "div"
	return json.Marshal(m)
}

// SetText set ComponentPlainText.Text attribute
func (r *ComponentPlainText) SetText(val *ObjectText) *ComponentPlainText {
	r.Text = val
	return r
}

// SetIcon set ComponentPlainText.Icon attribute
func (r *ComponentPlainText) SetIcon(val *ObjectIcon) *ComponentPlainText {
	r.Icon = val
	return r
}

// toMap conv ComponentPlainText to map
func (r *ComponentPlainText) toMap() map[string]interface{} {
	res := make(map[string]interface{}, 2)
	if r.Text != nil {
		res["text"] = r.Text
	}
	if r.Icon != nil {
		res["icon"] = r.Icon
	}
	return res
}
