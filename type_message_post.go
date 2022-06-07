/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package lark

import (
	"encoding/json"
	"reflect"
)

// MessageContentPostAll ...
type MessageContentPostAll struct {
	ZhCn *MessageContentPost `json:"zh_cn,omitempty"`
	JaJp *MessageContentPost `json:"ja_jp,omitempty"`
	EnUs *MessageContentPost `json:"en_us,omitempty"`
}

// String ...
func (r *MessageContentPostAll) String() string {
	bs, _ := json.Marshal(r)
	return string(bs)
}

// MessageContentPost ...
type MessageContentPost struct {
	Title   string                     `json:"title,omitempty"`
	Content [][]MessageContentPostItem `json:"content,omitempty"`
}

func (r *MessageContentPost) UnmarshalJSON(bytes []byte) error {
	var post struct {
		Title   string              `json:"title,omitempty"`
		Content [][]json.RawMessage `json:"content,omitempty"`
	}
	if err := json.Unmarshal(bytes, &post); err != nil {
		return err
	}

	r.Title = post.Title

	r.Content = make([][]MessageContentPostItem, len(post.Content))
	for vvIdx, vv := range post.Content {
		r.Content[vvIdx] = make([]MessageContentPostItem, len(vv))
		for vIdx, v := range vv {
			var item messageContentPostAll
			if err := json.Unmarshal(v, &item); err != nil {
				return err
			}
			switch item.Tag {
			case messageContentPostItemTypeText:
				r.Content[vvIdx][vIdx] = MessageContentPostText{Text: item.Text, UnEscape: item.UnEscape}
			case messageContentPostItemTypeLink:
				r.Content[vvIdx][vIdx] = MessageContentPostLink{Text: item.Text, UnEscape: item.UnEscape, Href: item.Href}
			case messageContentPostItemTypeImage:
				r.Content[vvIdx][vIdx] = MessageContentPostImage{ImageKey: item.ImageKey, Height: item.Height, Width: item.Width}
			case messageContentPostItemTypeAt:
				r.Content[vvIdx][vIdx] = MessageContentPostAt{UserID: item.UserID, UserName: item.UserName}
			}
		}
	}
	return nil
}

// MessageContentPostItem ...
type MessageContentPostItem interface {
	IsMessageContentPostItem() bool
}

const (
	messageContentPostItemTypeText  = "text"
	messageContentPostItemTypeLink  = "a"
	messageContentPostItemTypeImage = "img"
	messageContentPostItemTypeAt    = "at"
)

type messageContentPostAll struct {
	Tag string `json:"tag"`

	Text     string `json:"text"`      // 文本内容
	UnEscape bool   `json:"un_escape"` //
	Href     string `json:"href"`      // 默认的链接地址

	UserID   string `json:"user_id"`   // open_id
	UserName string `json:"user_name"` // 用户姓名

	ImageKey string `json:"image_key"` // 图片的唯一标识
	Height   int    `json:"height"`    // 图片的高
	Width    int    `json:"width"`     // 图片的宽
}

// MessageContentPostText 	文本内容
type MessageContentPostText struct {
	messageContentPostInterfaceDefaultImpl
	Text     string `json:"text"`      // 文本内容
	UnEscape bool   `json:"un_escape"` // 表示是不是 unescape 解码，默认为 false ，不用可以不填
}

// MarshalJSON ...
func (r MessageContentPostText) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"text":      r.Text,
		"un_escape": r.UnEscape,
		"tag":       messageContentPostItemTypeText,
	})
}

// MessageContentPostLink 链接
type MessageContentPostLink struct {
	messageContentPostInterfaceDefaultImpl
	Text     string `json:"text"`      // 文本内容
	UnEscape bool   `json:"un_escape"` //
	Href     string `json:"href"`      // 默认的链接地址
}

// MarshalJSON ...
func (r MessageContentPostLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"text":      r.Text,
		"un_escape": r.UnEscape,
		"href":      r.Href,
		"tag":       messageContentPostItemTypeLink,
	})
}

// MessageContentPostAt At用户
type MessageContentPostAt struct {
	messageContentPostInterfaceDefaultImpl
	UserID   string `json:"user_id"`   // open_id
	UserName string `json:"user_name"` // 用户姓名
}

// MarshalJSON ...
func (r MessageContentPostAt) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"user_id": r.UserID,
		"tag":     messageContentPostItemTypeAt,
	}
	if r.UserName != "" {
		v["user_name"] = r.UserName
	}
	return json.Marshal(v)
}

// MessageContentPostImage 图片
type MessageContentPostImage struct {
	messageContentPostInterfaceDefaultImpl
	ImageKey string `json:"image_key"` // 图片的唯一标识
	Height   int    `json:"height"`    // 图片的高
	Width    int    `json:"width"`     // 图片的宽
}

// MarshalJSON ...
func (r MessageContentPostImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"image_key": r.ImageKey,
		"height":    r.Height,
		"width":     r.Width,
		"tag":       messageContentPostItemTypeImage,
	})
}

type messageContentPostInterfaceDefaultImpl struct{}

// IsMessageContentPostItem ...
func (r messageContentPostInterfaceDefaultImpl) IsMessageContentPostItem() bool {
	return true
}

func marshalJSONWithMap(v interface{}, m map[string]interface{}) ([]byte, error) {
	vv := reflect.ValueOf(v)
	vt := reflect.TypeOf(v)
	for i := 0; i < vt.NumField(); i++ {
		jsonTag := vt.Field(i).Tag.Get("json")
		if len(jsonTag) > 10 && jsonTag[len(jsonTag)-10:] == ",omitempty" {
			jsonTag = jsonTag[:len(jsonTag)-10]
		}
		vvf := vv.Field(i)
		if vvf.IsZero() {
			continue
		}
		m[jsonTag] = vv.Field(i).Interface()
	}
	return json.Marshal(m)
}
