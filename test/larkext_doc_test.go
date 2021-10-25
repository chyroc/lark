package test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/chyroc/lark/larkext"
	"github.com/stretchr/testify/assert"
)

func Test_LarkExt_Doc(t *testing.T) {
	as := assert.New(t)

	larkClient := AppAllPermission.Ins()
	folderClient, err := larkext.NewRootFolder(ctx, larkClient)
	as.Nil(err)

	docClient, err := folderClient.NewDoc(ctx, "doc title")
	as.Nil(err)
	defer func() {
		as.Nil(docClient.Delete(ctx))
	}()

	t.Run("meta", func(t *testing.T) {
		meta, err := docClient.Meta(ctx)
		as.Nil(err)
		as.NotNil(meta)
	})

	t.Run("raw-content", func(t *testing.T) {
		text, err := docClient.RawContent(ctx)
		as.Nil(err)
		as.Equal("", strings.TrimSpace(text))
	})

	t.Run("content", func(t *testing.T) {
		content, err := docClient.Content(ctx)
		as.Nil(err)
		bs, err := json.Marshal(content)
		as.Nil(err)
		as.Equal(`{"title":{"elements":[],"location":{"zoneId":"0","startIndex":0,"endIndex":0},"lineId":""},"body":{"blocks":[{"type":"paragraph","paragraph":{"style":null,"elements":[],"location":{"zoneId":"0","startIndex":1,"endIndex":1},"lineId":""},"file":null,"chatGroup":null,"undefinedBlock":null}]}}`, string(bs))
		fmt.Println(string(bs))
	})
}
