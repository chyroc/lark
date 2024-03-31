package card

func OpenURLBehavior(url, androidURL, iosURL, pcURL string) *ObjectBehavior {
	return &ObjectBehavior{
		Type:       "open_url",
		DefaultURL: url,
		AndroidURL: androidURL,
		IOSURL:     iosURL,
		PCURL:      pcURL,
	}
}

func CallbackBehavior(value any) *ObjectBehavior {
	return &ObjectBehavior{
		Type:  "callback",
		Value: value,
	}
}

func SubmitFormBehavior() *ObjectBehavior {
	return &ObjectBehavior{
		Type:     "form_action",
		Behavior: "submit",
	}
}

func ResetFormBehavior() *ObjectBehavior {
	return &ObjectBehavior{
		Type:     "form_action",
		Behavior: "reset",
	}
}

// ObjectBehavior behaviors 字段用于配置交互组件的交互类型和交互行为。behaviors 字段支持跳转链接交互、服务端回传交互、和表单事件交互属性。本小节介绍各类交互对应的字段说明。
type ObjectBehavior struct {
	// open_url: 配置跳转链接交互
	// callback: 配置服务端回传交互
	// form_action: 要配置表单事件交互
	Type string `json:"type,omitempty"`

	// 声明交互类型。要配置跳转链接交互，取固定值 open_url
	//
	// 兜底跳转地址。
	DefaultURL string `json:"default_url,omitempty"`
	// Android 端的链接地址。可配置为 lark://msgcard/unsupported_action 声明当前端不允许跳转。
	AndroidURL string `json:"android_url,omitempty"`
	// iOS 端的链接地址。可配置为 lark://msgcard/unsupported_action 声明当前端不允许跳转。
	IOSURL string `json:"ios_url,omitempty"`
	// PC 端的链接地址。可配置为 lark://msgcard/unsupported_action 声明当前端不允许跳转。
	PCURL string `json:"pc_url,omitempty"`

	// 声明交互类型。要配置服务端回传交互，取固定值 callback。
	//
	// 你可在交互事件中自定义回传数据，支持 string 或 object 数据类型。
	Value any `json:"value,omitempty"`

	// 声明交互类型。要配置表单事件交互，取固定值 form_action。
	//
	// 表单事件类型。可取值：
	//
	// submit：提交整个表单
	// reset：重置整个表单
	Behavior string `json:"behavior,omitempty"`
}
