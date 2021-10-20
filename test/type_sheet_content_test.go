package test

import (
	"encoding/json"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func TestSheetValue_MarshalJSON(t *testing.T) {
	as := assert.New(t)

	tests := []struct {
		name       string
		arg        *lark.SheetContent
		want       string
		errContain string
	}{
		{name: "1", arg: &lark.SheetContent{String: ptr.String("string")}, want: `"string"`},
		{name: "2", arg: &lark.SheetContent{Int: ptr.Int64(1)}, want: `1`},
		{name: "3", arg: &lark.SheetContent{Link: &lark.SheetValueLink{Text: "text", Link: "url"}}, want: `{"text":"text","link":"url","type":"url"}`},
		{
			name: "4", arg: &lark.SheetContent{AtUser: &lark.SheetValueAtUser{Text: "aaa@aa.com", TextType: "email", Notify: false, GrantReadPermission: false}},
			want: `{"type":"mention","text":"aaa@aa.com","textType":"email","notify":false,"grantReadPermission":false}`,
		},
		{name: "5", arg: &lark.SheetContent{Formula: &lark.SheetValueFormula{Text: "=A1"}}, want: `{"type":"formula","text":"=A1"}`},
		{name: "6", arg: &lark.SheetContent{AtDoc: &lark.SheetValueAtDoc{Text: "sheet-token", ObjType: "sheet"}}, want: `{"type":"mention","textType":"fileToken","text":"sheet-token","objType":"sheet"}`},
		{name: "7", arg: &lark.SheetContent{MultiValue: &lark.SheetValueMultiValue{Values: []interface{}{1, "2"}}}, want: `{"type":"multipleValue","values":[1,"2"]}`},
	}
	for _, tt := range tests {
		t.Run(tt.name+" marshal", func(t *testing.T) {
			got, err := tt.arg.MarshalJSON()
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
				as.Empty(got, tt.name)
			} else {
				as.Nil(err, tt.name)
				as.NotEmpty(got, tt.name)
				as.Equal(tt.want, string(got), tt.name)
			}
		})
		t.Run(tt.name+" unmarshal", func(t *testing.T) {
			resp := new(lark.SheetContent)
			err := resp.UnmarshalJSON([]byte(tt.want))
			if tt.errContain != "" {
				as.NotNil(err, tt.name)
			} else {
				as.Nil(err, tt.name)
				a, _ := json.Marshal(tt.arg)
				b, _ := json.Marshal(resp)
				as.EqualValues(string(b), string(a), tt.name)
			}
		})
	}
}

func Test_SheetContent_2(t *testing.T) {
	t.Skip()
	as := assert.New(t)

	str := `[
        [
          "字符串",
          [
            {
              "text": "@这是名字",
              "type": "mention"
            },
            {
              "text": " ",
              "type": "text"
            }
          ],
          [
            {
              "link": "https://google.com/",
              "text": "Google",
              "type": "url"
            }
          ],
          "option1"
        ],
        [
          {
            "link": "https%3A%2F%2Frs6qnacjws.feishu.cn%2Fdocs%2FdoccnKzILuBnSjMIY7IEqyvNcPB",
            "text": "Lark SDK Ext 使用指南",
            "type": "mention"
          },
          1,
          2,
          "option2"
        ]
      ]`

	contents := [][]lark.SheetContent{}
	err := json.Unmarshal([]byte(str), &contents)
	as.Nil(err)

	as.Equal("字符串", *contents[0][0].String)
	as.Equal("@这是名字", (*contents[0][1].Children)[0].AtUser.Text)
	as.Equal(" ", *(*contents[0][1].Children)[1].String)
	as.Equal("https://google.com/", (*contents[0][2].Children)[0].Link.Link)
	as.Equal("Lark SDK Ext 使用指南", contents[1][0].AtDoc.Text)
}
