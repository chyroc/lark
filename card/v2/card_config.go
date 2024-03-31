package card

//go:generate generate_set_attrs -type=CardConfig
type CardConfig struct {
	// 是否允许转发卡片。取值：
	//
	// true：允许
	// false：不允许
	// 注意：该字段要求飞书客户端的版本为 V3.31.0 及以上。
	EnableForward bool `json:"enable_forward,omitempty"`

	// 是否为共享卡片。取值：
	//
	// true：是共享卡片, 更新卡片的内容对所有收到这张卡片的人员可见。
	// false：非共享卡片, 仅操作用户可见卡片的更新内容。
	UpdateMulti bool `json:"update_multi,omitempty"`

	// 卡片宽度模式。取值：
	//
	// default：默认宽度。PC 端宽版、iPad 端上的宽度上限为 600px。
	// fill：自适应屏幕宽度
	WidthMode Width `json:"width_mode,omitempty"`

	// 是否使用自定义翻译数据。取值：
	//
	// true：在用户点击消息翻译后, 使用 i18n 对应的目标语种作为翻译结果。若 i18n 取不到, 则使用当前内容请求飞书的机器翻译。
	// false：不使用自定义翻译数据, 直接请求飞书的机器翻译。
	UseCustomTranslation bool `json:"use_custom_translation,omitempty"`

	// 转发的卡片是否仍然支持回传交互。
	EnableForwardInteraction bool `json:"enable_forward_interaction,omitempty"`

	// 添加自定义字号和颜色。可应用于组件的 JSON 数据中, 设置字号和颜色属性。
	Style *CardConfigStyle `json:"style,omitempty"`
}

// SetEnableForward set CardConfig.EnableForward attribute
func (r *CardConfig) SetEnableForward(val bool) *CardConfig {
	r.EnableForward = val
	return r
}

// SetUpdateMulti set CardConfig.UpdateMulti attribute
func (r *CardConfig) SetUpdateMulti(val bool) *CardConfig {
	r.UpdateMulti = val
	return r
}

// SetWidthMode set CardConfig.WidthMode attribute
func (r *CardConfig) SetWidthMode(val Width) *CardConfig {
	r.WidthMode = val
	return r
}

// SetWidthModeDefault set CardConfig.WidthMode attribute to WidthDefault
func (r *CardConfig) SetWidthModeDefault() *CardConfig {
	r.WidthMode = WidthDefault
	return r
}

// SetWidthModeFill set CardConfig.WidthMode attribute to WidthFill
func (r *CardConfig) SetWidthModeFill() *CardConfig {
	r.WidthMode = WidthFill
	return r
}

// SetWidthModeAuto set CardConfig.WidthMode attribute to WidthAuto
func (r *CardConfig) SetWidthModeAuto() *CardConfig {
	r.WidthMode = WidthAuto
	return r
}

// SetUseCustomTranslation set CardConfig.UseCustomTranslation attribute
func (r *CardConfig) SetUseCustomTranslation(val bool) *CardConfig {
	r.UseCustomTranslation = val
	return r
}

// SetEnableForwardInteraction set CardConfig.EnableForwardInteraction attribute
func (r *CardConfig) SetEnableForwardInteraction(val bool) *CardConfig {
	r.EnableForwardInteraction = val
	return r
}

// SetStyle set CardConfig.Style attribute
func (r *CardConfig) SetStyle(val *CardConfigStyle) *CardConfig {
	r.Style = val
	return r
}
