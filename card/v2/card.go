package card

import "encoding/json"

func Card(elements ...Component) *ComponentCard {
	return &ComponentCard{
		Elements: elements,
	}
}

// /go:generate generate_set_attrs -type=ComponentCard
//
//go:generate generate_iface_unmarshal -type=ComponentCard
type ComponentCard struct {
	// 标题组件 JSON 代码。详细字段说明参考标题组件。
	Header *ComponentHeader `json:"header,omitempty"`

	// i18n_header
	I18nHeader map[Language]*ComponentHeader `json:"i18n_header,omitempty"`

	// config 用于配置卡片的全局行为, 包括是否允许被转发、是否为共享卡片等。
	Config *CardConfig `json:"config,omitempty"`

	// card_link 字段用于指定卡片整体的点击跳转链接。
	// 你可以配置一个默认链接, 也可以分别为 PC 端、Android 端、iOS 端配置不同的跳转链接。
	CardLink *ObjectMultiURL `json:"card_link,omitempty"`

	// 在 elements 字段中添加各个组件的 JSON 数据, 组件将按数组顺序纵向流式排列。了解各个组件, 参考组件概述。
	Elements []Component `json:"elements,omitempty" unmarshal:"unmarshalComponent"`

	// 飞书卡片支持多语言设置。设置多语言后, 卡片将根据用户的飞书客户端语言, 自动展示对应语言的卡片内容, 满足国际化业务需求。详情参考配置卡片多语言。
	I18nElements map[Language][]Component `json:"i18n_elements,omitempty" unmarshal:"unmarshalComponent"`

	// fallback 用于为卡片添加全局降级规则。触发降级时, 卡片将全局展示“请升级客户端至最新版本后查看”占位图。
	//
	// 注意：该字段要求飞书客户端的版本为 V7.7 及以上。
	Fallback *CardFallback `json:"fallback,omitempty"`
}

type CardFallback struct {
	// 触发降级的条件数组。满足其中的任一条件就触发降级。
	TriggerConditions []*FallbackCondition `json:"trigger_conditions,omitempty"`
}

type FallbackCondition struct {
	//
	// 条件类型。可选值：
	//
	// min_client_version：设置最低客户端版本, 当用户的客户端版本低于该设置时, 触发降级;
	// element_tags：指定组件。当用户的飞书客户端版本低于这些组件支持的最低客户端版本时, 触发降级。
	Type string `json:"type,omitempty"`

	// 最低飞书客户端版本的值, 写法需符合以下格式之一：
	//
	// 7.4、7.4.1、7.4.1-xxx
	// v7.4、v7.4.1、v7.4.1-xxx
	// v7.3.7（0.103）
	// v7.4.0-dev.d2666af5（0.14）
	Value string `json:"value,omitempty"`

	// 指定组件。当用户的飞书客户端版本低于这些组件支持的最低客户端版本时, 触发降级。
	ElementTags []string `json:"element_tags,omitempty"`
}

type CardConfigStyle struct {
	// 分别为移动端和桌面端添加自定义字号。用于在普通文本组件和富文本组件 JSON 中设置字号属性。支持添加多个自定义字号对象。详情参考普通文本组件和富文本组件。
	TextSize map[string]*CustomTextSize `json:"text_size,omitempty"`

	// 分别为飞书客户端浅色主题和深色主题添加 RGBA 语法。用于在组件 JSON 中设置颜色属性。支持添加多个自定义颜色对象。详情参考颜色枚举值。
	Color map[string]*CustomColor `json:"color,omitempty"`
}

type CustomTextSize struct {
	// 在无法差异化配置字号的旧版飞书客户端上, 生效的字号属性。选填。
	Default string `json:"default,omitempty"`

	// 移动端的字号。
	Mobile string `json:"mobile,omitempty"`

	// 桌面端的字号。
	PC string `json:"pc,omitempty"`
}

type CustomColor struct {
	// 浅色主题下的自定义颜色语法
	LightMode string `json:"light_mode,omitempty"`

	// 深色主题下的自定义颜色语法
	DarkMode string `json:"dark_mode,omitempty"`
}

func (r ComponentCard) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

type unmarshalComponentTag struct {
	Tag string `json:"tag"`
}

// SetHeader set ComponentCard.Header attribute
func (r *ComponentCard) SetHeader(val *ComponentHeader) *ComponentCard {
	r.Header = val
	return r
}

// SetI18nHeader set ComponentCard.I18nHeader attribute
func (r *ComponentCard) SetI18nHeader(val map[Language]*ComponentHeader) *ComponentCard {
	r.I18nHeader = val
	return r
}

// SetConfig set ComponentCard.Config attribute
func (r *ComponentCard) SetConfig(val *CardConfig) *ComponentCard {
	r.Config = val
	return r
}

// SetCardLink set ComponentCard.CardLink attribute
func (r *ComponentCard) SetCardLink(val *ObjectMultiURL) *ComponentCard {
	r.CardLink = val
	return r
}

// SetElements set ComponentCard.Elements attribute
func (r *ComponentCard) SetElements(val []Component) *ComponentCard {
	r.Elements = val
	return r
}

// SetI18nElements set ComponentCard.I18nElements attribute
func (r *ComponentCard) SetI18nElements(val map[Language][]Component) *ComponentCard {
	r.I18nElements = val
	return r
}

// SetFallback set ComponentCard.Fallback attribute
func (r *ComponentCard) SetFallback(val *CardFallback) *ComponentCard {
	r.Fallback = val
	return r
}

// unmarshalComponentCard generated to unmarshal ComponentCard iface
type unmarshalComponentCard struct {
	Header       *ComponentHeader              `json:"header,omitempty"`
	I18nHeader   map[Language]*ComponentHeader `json:"i18n_header,omitempty"`
	Config       *CardConfig                   `json:"config,omitempty"`
	CardLink     *ObjectMultiURL               `json:"card_link,omitempty"`
	Elements     []json.RawMessage             `json:"elements,omitempty"`
	I18nElements map[string][]json.RawMessage  `json:"i18n_elements,omitempty"`
	Fallback     *CardFallback                 `json:"fallback,omitempty"`
}

// UnmarshalJSON generated to unmarshal ComponentCard iface
func (r *ComponentCard) UnmarshalJSON(bs []byte) error {
	obj := new(unmarshalComponentCard)
	err := json.Unmarshal(bs, obj)
	if err != nil {
		return err
	}
	r.Header = obj.Header
	r.I18nHeader = obj.I18nHeader
	r.Config = obj.Config
	r.CardLink = obj.CardLink
	r.Elements = make([]Component, len(obj.Elements))
	for i, v := range obj.Elements {
		r.Elements[i], err = unmarshalComponent(v)
		if err != nil {
			return err
		}
	}
	r.I18nElements = make(map[Language][]Component, len(obj.I18nElements))
	for k, v := range obj.I18nElements {
		for _, vv := range v {
			tmp, err := unmarshalComponent(vv)
			if err != nil {
				return err
			}
			r.I18nElements[k] = append(r.I18nElements[k], tmp)
		}
	}
	r.Fallback = obj.Fallback
	return nil
}
