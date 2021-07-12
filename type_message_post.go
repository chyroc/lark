package lark

import (
	"encoding/json"
)

type MessageContentPostAll struct {
	ZhCn *MessageContentPost `json:"zh_cn,omitempty"`
	JaJp *MessageContentPost `json:"ja_jp,omitempty"`
	EnUs *MessageContentPost `json:"en_us,omitempty"`
}

func (r *MessageContentPostAll) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

type MessageContentPost struct {
	Title   string                     `json:"title,omitempty"`
	Content [][]MessageContentPostItem `json:"content,omitempty"`
}

type MessageContentPostItem interface {
	// json.Marshaler
	IsMessageContentPostItem() bool
}

type MessageContentPostText struct {
	messageContentPostInterfaceDefaultImpl
	Text     string `json:"text"`      // 文本内容
	UnEscape bool   `json:"un_escape"` // 表示是不是 unescape 解码，默认为 false ，不用可以不填
}

func (r MessageContentPostText) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"text":      r.Text,
		"un_escape": r.UnEscape,
		"tag":       "text",
	})
}

type MessageContentPostLink struct {
	messageContentPostInterfaceDefaultImpl
	Text     string `json:"text"`      // 文本内容
	UnEscape bool   `json:"un_escape"` //
	Href     string `json:"href"`      // 默认的链接地址
}

func (r MessageContentPostLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"text":      r.Text,
		"un_escape": r.UnEscape,
		"href":      r.Href,
		"tag":       "a",
	})
}

type MessageContentPostAt struct {
	messageContentPostInterfaceDefaultImpl
	UserID   string `json:"user_id"`   // open_id
	UserName string `json:"user_name"` // 用户姓名
}

func (r MessageContentPostAt) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"user_id": r.UserID,
		"tag":     "at",
	}
	if r.UserName != "" {
		v["user_name"] = r.UserName
	}
	return json.Marshal(v)
}

type MessageContentPostImage struct {
	messageContentPostInterfaceDefaultImpl
	ImageKey string `json:"image_key"` // 图片的唯一标识
	Height   int    `json:"height"`    // 图片的高
	Width    int    `json:"width"`     // 图片的宽
}

func (r MessageContentPostImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"image_key": r.ImageKey,
		"height":    r.Height,
		"width":     r.Width,
		"tag":       "img",
	})
}

type messageContentPostInterfaceDefaultImpl struct{}

func (r messageContentPostInterfaceDefaultImpl) IsMessageContentPostItem() bool {
	return true
}
