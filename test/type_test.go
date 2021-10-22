package test

import (
	"encoding/json"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_Type(t *testing.T) {
	as := assert.New(t)

	t.Run("SheetContent", func(t *testing.T) {
		t.Run("children", func(t *testing.T) {
			children := []*lark.SheetContent{{String: ptr.String("str")}}
			res := lark.SheetContent{
				Children: &children,
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("str", func(t *testing.T) {
			res := lark.SheetContent{
				String: ptr.String("str"),
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("int", func(t *testing.T) {
			res := lark.SheetContent{
				Int: ptr.Int64(1),
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("link", func(t *testing.T) {
			res := lark.SheetContent{
				Link: &lark.SheetValueLink{
					Text: "text",
					Link: "link",
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("user", func(t *testing.T) {
			res := lark.SheetContent{
				AtUser: &lark.SheetValueAtUser{
					Text:                "email",
					TextType:            "email",
					Notify:              true,
					GrantReadPermission: true,
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("formula", func(t *testing.T) {
			res := lark.SheetContent{
				Formula: &lark.SheetValueFormula{
					Text: "formula",
				},
				// AtDoc: &lark.SheetValueAtDoc{
				// 	TextType: "sheet",
				// 	Text:     "sheet",
				// },
				// MultiValue: &lark.SheetValueMultiValue{
				// 	Values: []interface{}{"1", "2"},
				// },
				// EmbedImage: &lark.SheetValueEmbedImage{
				// 	FileToken: "token",
				// 	Height:    100,
				// 	Link:      "link",
				// 	Text:      "text",
				// 	Type:      "image",
				// 	Width:     100,
				// },
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("doc", func(t *testing.T) {
			res := lark.SheetContent{
				AtDoc: &lark.SheetValueAtDoc{
					TextType: "sheet",
					Text:     "sheet",
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("multi", func(t *testing.T) {
			res := lark.SheetContent{
				MultiValue: &lark.SheetValueMultiValue{
					Values: []interface{}{"1", "2"},
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})

		t.Run("image", func(t *testing.T) {
			res := lark.SheetContent{
				EmbedImage: &lark.SheetValueEmbedImage{
					FileToken: "token",
					Height:    100,
					Link:      "link",
					Text:      "text",
					Type:      "image",
					Width:     100,
				},
			}
			bs, err := json.Marshal(res)
			as.Nil(err)
			obj := &lark.SheetContent{}
			as.Nil(json.Unmarshal(bs, obj))
			as.EqualValues(res, *obj)
		})
	})

	t.Run("DocContent", func(t *testing.T) {
		res := lark.DocContent{
			Title: &lark.DocTitle{
				Elements: []*lark.DocParagraphElement{
					{
						Type:             "",
						TextRun:          nil,
						DocsLink:         nil,
						Person:           nil,
						Equation:         nil,
						Reminder:         nil,
						File:             nil,
						Jira:             nil,
						UndefinedElement: nil,
					},
				},
				Location: &lark.DocLocation{
					ZoneID:     "id",
					StartIndex: 1,
					EndIndex:   2,
				},
				LineID: "3",
			},
			Body: &lark.DocBody{
				Blocks: []*lark.DocBlock{
					{
						Type:           "",
						Paragraph:      nil,
						Gallery:        nil,
						File:           nil,
						ChatGroup:      nil,
						Table:          nil,
						HorizontalLine: nil,
						EmbeddedPage:   nil,
						Sheet:          nil,
						Bitable:        nil,
						Diagram:        nil,
						Jira:           nil,
						Poll:           nil,
						Code:           nil,
						DocsApp:        nil,
						Callout:        nil,
						UndefinedBlock: nil,
					},
				},
			},
		}
		bs, err := json.Marshal(res)
		as.Nil(err)

		obj := new(lark.DocContent)
		as.Nil(json.Unmarshal(bs, obj))

		as.Equal(res, *obj)
	})

	t.Run("md-builder", func(t *testing.T) {
		as.Equal(`<at id=all></at>`, lark.MdBuilder.AtAll())
		as.Equal(`[title](url)`, lark.MdBuilder.Link("url", "title"))
		as.Equal(`![hover](key)`, lark.MdBuilder.Image("key", "hover"))
		as.Equal(`<at email=email></at>`, lark.MdBuilder.AtUserEmail("email"))
		as.Equal(`**bold**`, lark.MdBuilder.Bold("bold"))
		as.Equal(`<at id=user></at>`, lark.MdBuilder.AtUserID("user"))
		as.Equal(`*italic*`, lark.MdBuilder.Italic("italic"))
		as.Equal(`~~Strikethrough~~`, lark.MdBuilder.Strikethrough("Strikethrough"))
	})

	t.Run("at-builder", func(t *testing.T) {
		as.Equal(`<at user_id="1"></at>`, lark.AtBuilder.AtOpenID("1"))
		as.Equal(`<at user_id="all"></at>`, lark.AtBuilder.AtAll())
	})
}
