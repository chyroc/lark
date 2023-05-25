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
	"fmt"
	"strconv"
	"strings"
)

type SheetContentType string

const (
	SheetContentTypeString     SheetContentType = "string"
	SheetContentTypeInt        SheetContentType = "int"
	SheetContentTypeFloat      SheetContentType = "float"
	SheetContentTypeLink       SheetContentType = "link"
	SheetContentTypeAtUser     SheetContentType = "at_user"
	SheetContentTypeFormula    SheetContentType = "formula"
	SheetContentTypeAtDoc      SheetContentType = "at_doc"
	SheetContentTypeMultiValue SheetContentType = "multi_value"
	SheetContentTypeEmbedImage SheetContentType = "embed_image"
	SheetContentTypeList       SheetContentType = "list"
	SheetContentTypeNull       SheetContentType = "null"
)

// SheetContent ...
type SheetContent struct {
	Children   *[]*SheetContent      `json:"children,omitempty"`
	String     *string               `json:"string,omitempty"`      // 字符串, `"`
	Int        *int64                `json:"int,omitempty"`         // 数字, `0-9`
	Float      *float64              `json:"float,omitempty"`       // 小数, `x.x`
	Link       *SheetValueLink       `json:"link,omitempty"`        // 带文本的链接, `{
	AtUser     *SheetValueAtUser     `json:"at_user,omitempty"`     // @人, `{
	Formula    *SheetValueFormula    `json:"formula,omitempty"`     // 公式, `{
	AtDoc      *SheetValueAtDoc      `json:"at_doc,omitempty"`      // @文档, `{
	MultiValue *SheetValueMultiValue `json:"multi_value,omitempty"` // 下拉列表, `{
	EmbedImage *SheetValueEmbedImage `json:"embed_image,omitempty"` // 内嵌图片, `{
}

// SheetValueEmbedImage ...
type SheetValueEmbedImage struct {
	FileToken string `json:"fileToken"`
	Height    int    `json:"height"`
	Link      string `json:"link"`
	Text      string `json:"text"`
	Type      string `json:"type"` // embed-image
	Width     int    `json:"width"`
}

// SheetValueLink ...
type SheetValueLink struct {
	Text string `json:"text"`
	Link string `json:"link"`
	Type string `json:"type"` // url
}

// 只支持@同租户的用户;最多支持同时@50人；
type SheetValueAtUser struct {
	Type                string `json:"type"`                // mention
	Text                string `json:"text"`                // 需要@的人的信息，由textType指定
	TextType            string `json:"textType"`            // 指定text字段的传入的内容，可选email，openId，unionId；
	Notify              bool   `json:"notify"`              // 是否发送飞书消息，没有阅读权限的用户不会收到飞书消息；
	GrantReadPermission bool   `json:"grantReadPermission"` // 是否赋予该用户阅读权限（仅在独立表格中支持该字段）；
}

// 公式
type SheetValueFormula struct {
	Type string `json:"type"` // formula
	Text string `json:"text"` // text字段为对应的公式
}

// @文档
type SheetValueAtDoc struct {
	Type     string `json:"type"`     // mention
	TextType string `json:"textType"` // 固定为fileToken
	Text     string `json:"text"`     // 文档token
	ObjType  string `json:"objType"`  // 文档类型，可选sheet,doc,slide,bitable,mindnote
}

// 下拉列表
type SheetValueMultiValue struct {
	Type   string        `json:"type"`   // multipleValue
	Values []interface{} `json:"values"` // values为数组，可填bool,string,number类型。string类型数据不能包含","。使用前需要先使用设置下拉列表接口设置下拉列表。
}

func (r *SheetContent) Type() SheetContentType {
	switch {
	case r.String != nil:
		return SheetContentTypeString
	case r.Int != nil:
		return SheetContentTypeInt
	case r.Float != nil:
		return SheetContentTypeFloat
	case r.Link != nil:
		return SheetContentTypeLink
	case r.AtUser != nil:
		return SheetContentTypeAtUser
	case r.Formula != nil:
		return SheetContentTypeFormula
	case r.AtDoc != nil:
		return SheetContentTypeAtDoc
	case r.MultiValue != nil:
		return SheetContentTypeMultiValue
	case r.EmbedImage != nil:
		return SheetContentTypeEmbedImage
	case r.Children != nil:
		return SheetContentTypeList
	default:
		return SheetContentTypeNull
	}
}

// UnmarshalJSON ...
func (r *SheetContent) UnmarshalJSON(bytes []byte) error {
	if len(bytes) == 0 {
		return nil
	}
	if bytes[0] == '"' {
		var dest string
		if err := json.Unmarshal(bytes, &dest); err != nil {
			return err
		}
		r.String = &dest
		return nil
	} else if bytes[0] >= '0' && bytes[0] <= '9' {
		str := string(bytes)
		if strings.Contains(str, ".") {
			float, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return err
			}
			r.Float = &float
		} else {
			integer, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				return err
			}
			r.Int = &integer
		}
		return nil
	} else if bytes[0] == 'n' {
		if len(bytes) == 4 && string(bytes) == "null" {
			return nil
		}
		return fmt.Errorf("unsupport sheet value")
	} else if bytes[0] == '{' {
		var obj struct {
			Type     string `json:"type"`
			TextType string `json:"textType"`
		}
		if err := json.Unmarshal(bytes, &obj); err != nil {
			return err
		}
		switch obj.Type {
		case "embed-image":
			resp := new(SheetValueEmbedImage)
			if err := json.Unmarshal(bytes, resp); err != nil {
				return err
			}
			r.EmbedImage = resp
			return nil
		case "url":
			resp := new(SheetValueLink)
			if err := json.Unmarshal(bytes, resp); err != nil {
				return err
			}
			r.Link = resp
			return nil
		case "mention":
			switch obj.TextType {
			case "fileToken":
				resp := new(SheetValueAtDoc)
				if err := json.Unmarshal(bytes, resp); err != nil {
					return err
				}
				r.AtDoc = resp
			default:
				resp := new(SheetValueAtUser)
				if err := json.Unmarshal(bytes, resp); err != nil {
					return err
				}
				r.AtUser = resp
			}
			return nil
		case "formula":
			resp := new(SheetValueFormula)
			if err := json.Unmarshal(bytes, resp); err != nil {
				return err
			}
			r.Formula = resp
			return nil
		case "multipleValue":
			resp := new(SheetValueMultiValue)
			if err := json.Unmarshal(bytes, resp); err != nil {
				return err
			}
			r.MultiValue = resp
			return nil
		case "text":
			resp := new(SheetValueFormula)
			if err := json.Unmarshal(bytes, resp); err != nil {
				return err
			}
			r.String = &resp.Text
			return nil
		default:
			return fmt.Errorf("unsupport sheet value")
		}
	} else if bytes[0] == '[' {
		var resp []*SheetContent
		if err := json.Unmarshal(bytes, &resp); err != nil {
			return err
		}
		r.Children = &resp
		return nil
	} else {
		return fmt.Errorf("unsupport sheet value")
	}
}

// MarshalJSON ...
func (r SheetContent) MarshalJSON() ([]byte, error) {
	if r.String != nil {
		return []byte(fmt.Sprintf("%q", *r.String)), nil
	} else if r.Int != nil {
		return []byte(strconv.FormatInt(*r.Int, 10)), nil
	} else if r.Float != nil {
		return []byte(strconv.FormatFloat(*r.Float, byte('f'), 10, 64)), nil
	} else if r.Link != nil {
		r.Link.Type = "url"
		return json.Marshal(r.Link)
	} else if r.AtUser != nil {
		r.AtUser.Type = "mention"
		return json.Marshal(r.AtUser)
	} else if r.Formula != nil {
		r.Formula.Type = "formula"
		return json.Marshal(r.Formula)
	} else if r.AtDoc != nil {
		r.AtDoc.Type = "mention"
		r.AtDoc.TextType = "fileToken"
		return json.Marshal(r.AtDoc)
	} else if r.MultiValue != nil {
		r.MultiValue.Type = "multipleValue"
		return json.Marshal(r.MultiValue)
	} else if r.EmbedImage != nil {
		r.EmbedImage.Type = "embed-image"
		return json.Marshal(r.EmbedImage)
	} else if r.Children != nil {
		return json.Marshal(r.Children)
	}
	return []byte("null"), nil
}
