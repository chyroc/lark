package card

func OverflowOption(text string, value any) *ObjectOverflowOption {
	return &ObjectOverflowOption{
		Text:  Text(text),
		Value: value,
	}
}

// ObjectOverflowOption ...
//
//go:generate generate_set_attrs -type=ObjectOverflowOption
type ObjectOverflowOption struct {
	// 按钮上的文本。
	Text *ObjectText `json:"text,omitempty"`

	// 为按钮添加多端的跳转链接。
	MultiURL *ObjectMultiURL `json:"multi_url,omitempty"`

	// 该按钮的回传参数值。当用户点击选项后, 应用会将该值返回至卡片请求地址。
	Value any `json:"value,omitempty"`
}

// SetText set ObjectOverflowOption.Text attribute
func (r *ObjectOverflowOption) SetText(val *ObjectText) *ObjectOverflowOption {
	r.Text = val
	return r
}

// SetMultiURL set ObjectOverflowOption.MultiURL attribute
func (r *ObjectOverflowOption) SetMultiURL(val *ObjectMultiURL) *ObjectOverflowOption {
	r.MultiURL = val
	return r
}

// SetValue set ObjectOverflowOption.Value attribute
func (r *ObjectOverflowOption) SetValue(val any) *ObjectOverflowOption {
	r.Value = val
	return r
}
