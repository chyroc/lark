package lark

import (
	"encoding/json"
)

type SheetContent interface {
	IsSheetContent()
}

type SheetContentString string

func (r SheetContentString) IsSheetContent() {}

type SheetContentInt64 string

func (r SheetContentInt64) IsSheetContent() {}

type SheetContentFloat64 string

func (r SheetContentFloat64) IsSheetContent() {}

type SheetContentLink struct {
	Text string `json:"text,omitempty"`
	Link string `json:"link,omitempty"`
}

func (r SheetContentLink) IsSheetContent() {}

func (r SheetContentLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(sheetContentLink{SheetContentLink: r, Type: "url"})
}

type SheetContentEmail string

func (r SheetContentEmail) IsSheetContent() {}

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
	return json.Marshal(sheetContentUser{SheetContentUser: r, Type: "mention"})
}

type SheetContentFormula struct {
	Text string `json:"text,omitempty"`
}

func (r SheetContentFormula) IsSheetContent() {}

func (r SheetContentFormula) MarshalJSON() ([]byte, error) {
	return json.Marshal(sheetContentFormula{SheetContentFormula: r, Type: "formula"})
}

type SheetContentDoc struct {
	Text    string `json:"text,omitempty"`
	ObjType string `json:"objType,omitempty"`
}

// text为对应文档的token，
// objType为"sheet","doc","slide","bitable","mindnote"之一
func (r SheetContentDoc) IsSheetContent() {}

func (r SheetContentDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(sheetContentDoc{
		SheetContentDoc: r,
		Type:            "mention",
		TextType:        "fileToken",
	})
}

type SheetContentMultipleValue struct {
	Values []interface{} `json:"values"`
}

func (r SheetContentMultipleValue) IsSheetContent() {}

func (r SheetContentMultipleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(sheetContentMultipleValue{SheetContentMultipleValue: r, Type: "multipleValue"})
}

type sheetContentLink struct {
	SheetContentLink
	Type string `json:"type,omitempty"`
}

type sheetContentUser struct {
	SheetContentUser
	Type string `json:"type,omitempty"`
}

type sheetContentFormula struct {
	SheetContentFormula
	Type string `json:"type"`
}

type sheetContentMultipleValue struct {
	SheetContentMultipleValue
	Type string `json:"type"`
}

type sheetContentDoc struct {
	SheetContentDoc
	Type     string `json:"type,omitempty"`
	TextType string `json:"textType,omitempty"`
}
