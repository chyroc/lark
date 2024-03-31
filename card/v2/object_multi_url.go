package card

const URLDisabled = "lark://msgcard/unsupported_action"

func MultiURL(url string) *ObjectMultiURL {
	return &ObjectMultiURL{
		URL: url,
	}
}

//go:generate generate_set_attrs -type=ObjectMultiURL
type ObjectMultiURL struct {
	// 兜底的跳转链接。
	URL string `json:"url,omitempty"`

	// Android 端的跳转链接。可配置为 lark://msgcard/unsupported_action 声明当前端不允许跳转。
	AndroidURL string `json:"android_url,omitempty"`

	// iOS 端的跳转链接。可配置为 lark://msgcard/unsupported_action 声明当前端不允许跳转。
	IOSURL string `json:"ios_url,omitempty"`

	// PC 端的跳转链接。可配置为 lark://msgcard/unsupported_action 声明当前端不允许跳转。
	PCURL string `json:"pc_url,omitempty"`
}

// SetURL set ObjectMultiURL.URL attribute
func (r *ObjectMultiURL) SetURL(val string) *ObjectMultiURL {
	r.URL = val
	return r
}

// SetAndroidURL set ObjectMultiURL.AndroidURL attribute
func (r *ObjectMultiURL) SetAndroidURL(val string) *ObjectMultiURL {
	r.AndroidURL = val
	return r
}

// SetIOSURL set ObjectMultiURL.IOSURL attribute
func (r *ObjectMultiURL) SetIOSURL(val string) *ObjectMultiURL {
	r.IOSURL = val
	return r
}

// SetPCURL set ObjectMultiURL.PCURL attribute
func (r *ObjectMultiURL) SetPCURL(val string) *ObjectMultiURL {
	r.PCURL = val
	return r
}
