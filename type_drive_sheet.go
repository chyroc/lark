package lark

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type SheetContent interface {
	IsSheetContent()
}

type SheetContentString struct {
	Text string `json:"content"`
}

func (r SheetContentString) IsSheetContent() {}

func (r SheetContentString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", r.Text)), nil
}

type SheetContentNumber struct {
	Number string `json:"content"`
}

func (r SheetContentNumber) IsSheetContent() {}

func (r SheetContentNumber) MarshalJSON() ([]byte, error) {
	return []byte(r.Number), nil
}

type SheetContentLink struct {
	Text string `json:"text,omitempty"`
	Link string `json:"link,omitempty"`
}

func (r SheetContentLink) IsSheetContent() {}

func (r SheetContentLink) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"type": "url"})
}

type SheetContentEmail struct {
	Email string `json:"content"`
}

func (r SheetContentEmail) IsSheetContent() {}

func (r SheetContentEmail) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", r.Email)), nil
}

// 	{"type":"mention","text": "text","textType":"email","notify": true,"grantReadPermission":true }
// 	只支持同租户@，type固定为mention，
// 	notify为是否发送lark消息
// 	grantReadPermission为是否赋予该用户阅读权限，没有阅读权限的用户不会收到lark消息
// 	textType为"email"时，text为对应用户的邮箱
// 	textType为"openId"时，text为对应用户的openId
// 	textType为"unionId"时，text为对应用户的unionId
// 	单次请求最多@50个用户
type SheetContentUser struct {
	TextType            string `json:"text_type,omitempty"`
	Text                string `json:"text,omitempty"`
	Notify              bool   `json:"notify,omitempty"`
	GrantReadPermission bool   `json:"grantReadPermission,omitempty"`
}

func (r SheetContentUser) IsSheetContent() {}

func (r SheetContentUser) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"type": "mention"})
}

type SheetContentFormula struct {
	Text string `json:"text,omitempty"`
}

func (r SheetContentFormula) IsSheetContent() {}

func (r SheetContentFormula) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"type": "formula"})
}

type SheetContentDoc struct {
	Text    string `json:"text,omitempty"`
	ObjType string `json:"objType,omitempty"`
}

// text为对应文档的token，
// objType为"sheet","doc","slide","bitable","mindnote"之一
func (r SheetContentDoc) IsSheetContent() {}

func (r SheetContentDoc) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"type": "mention", "textType": "fileToken"})
}

type SheetContentMultipleValue struct {
	Values []interface{} `json:"values"`
}

func (r SheetContentMultipleValue) IsSheetContent() {}

func (r SheetContentMultipleValue) MarshalJSON() ([]byte, error) {
	return marshalJSONWithMap(r, map[string]interface{}{"type": "multipleValue"})
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
