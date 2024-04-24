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
//
// docs: https://open.larkoffice.com/document/server-docs/im-v1/message-content-description/message_content
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
			case messageContentPostItemTypeAt:
				r.Content[vvIdx][vIdx] = MessageContentPostAt{UserID: item.UserID, UserName: item.UserName}
			case messageContentPostItemTypeImage:
				r.Content[vvIdx][vIdx] = MessageContentPostImage{ImageKey: item.ImageKey, Height: item.Height, Width: item.Width}
			case messageContentPostItemTypeMedia:
				r.Content[vvIdx][vIdx] = MessageContentPostMedia{FileKey: item.FileKey, ImageKey: item.ImageKey}
			case messageContentPostItemTypeEmotion:
				r.Content[vvIdx][vIdx] = MessageContentPostEmotion{EmojiType: item.EmojiType}
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
	messageContentPostItemTypeText    = "text"
	messageContentPostItemTypeLink    = "a"
	messageContentPostItemTypeAt      = "at"
	messageContentPostItemTypeImage   = "img"
	messageContentPostItemTypeMedia   = "media"
	messageContentPostItemTypeEmotion = "emotion"
)

type messageContentPostAll struct {
	Tag string `json:"tag"`

	Text     string `json:"text"`      // 文本内容
	UnEscape bool   `json:"un_escape"` //
	Href     string `json:"href"`      // 默认的链接地址

	UserID   string `json:"user_id"`   // open_id
	UserName string `json:"user_name"` // 用户姓名

	ImageKey string `json:"image_key"` // 图片的唯一标识, 视频封面图片的唯一标识
	Height   int    `json:"height"`    // 图片的高
	Width    int    `json:"width"`     // 图片的宽

	FileKey string `json:"file_key"` // 适配文件的唯一标识

	EmojiType string `json:"emoji_type"` // 表情类型， 部分可选值请参见表情文案
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

// MessageContentPostMedia 视频
type MessageContentPostMedia struct {
	messageContentPostInterfaceDefaultImpl
	FileKey  string `json:"file_key"`  // 视频文件的唯一标识
	ImageKey string `json:"image_key"` // 视频封面图片的唯一标识
}

// MarshalJSON ...
func (r MessageContentPostMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"file_key":  r.FileKey,
		"image_key": r.ImageKey,
		"tag":       messageContentPostItemTypeMedia,
	})
}

// MessageContentPostEmotion 标签
type MessageContentPostEmotion struct {
	messageContentPostInterfaceDefaultImpl
	EmojiType string `json:"emoji_type"` // 标签类型
}

// MarshalJSON ...
func (r MessageContentPostEmotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"emoji_type": r.EmojiType,
		"tag":        messageContentPostItemTypeEmotion,
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
		if jsonTag == "-" {
			continue
		}
		vvf := vv.Field(i)
		if vvf.IsZero() {
			continue
		}
		m[jsonTag] = vv.Field(i).Interface()
	}
	return json.Marshal(m)
}
