package card_test

import (
	"encoding/json"
	"testing"

	"github.com/chyroc/lark/card"
)

func TestName(t *testing.T) {
	// header
	card.Header("# 报警").Red()

	// config
	card.Config()

	// module
	card.Action()
	card.HR()
	card.Note(
		card.ElementImage("image-key"),
		card.Text("名称"),
	)
	card.ModuleImage("img-key")
	card.ModuleImage("img-key").CropCenter().WithPreview(false)

	card.Confirm("title", "text")
	card.Confirm("", "text").WithTitle(card.Text("title"))

	card.LinkOption("百度", "https://baidu.com")
	card.PersonOption("open_id_1")
	card.SelectOption("选项1", "value_1")

	card.StaticSelectMenu(
		card.SelectOption("选项1", "value_1"),
		card.SelectOption("选项2", "value_2"))

	card.PersonSelectMenuForIDs("open_id_1", "open_id_2")

	card.Overflow(
		card.SelectOption("选项1", "value_1"),
		card.SelectOption("选项2", "value_2"))

	card.Button("确认", map[string]interface{}{"action_key": "uuid"})
	card.LinkButton("去官网", "https://baidu.com")
	card.LinkButton("去官网", "https://baidu.com").WithIOS("ios://url")
	card.LinkButton("去官网", "https://baidu.com").WithConfirm(card.Confirm("确认吗", "这会跳转到外部链接"))
	card.LinkButton("去官网", "https://baidu.com").WithConfirm(card.Confirm("确认吗", "这会跳转到外部链接")).Danger()

	card.Text("hi").Line(1)
	card.URL("https://www.baidu.com").WithIOS("ios://url")
}

func jsonString(v interface{}) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
