package card_test

import (
	"testing"

	"github.com/chyroc/lark/card"
)

func TestName(t *testing.T) {
	card.Confirm("title", "text")
	card.Confirm("", "text").WithTitle(card.Text("title"))
	card.Text("hi").Line(1)
	card.URL("https://www.baidu.com").IOS("ios://url")
}
