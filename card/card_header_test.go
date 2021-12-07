package card_test

import (
	"testing"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/card"
	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	tests := []struct {
		name string
		args *lark.MessageContentCardHeader
		want string
	}{
		{"1", card.Header("# header"), `{"title":{"tag":"plain_text","content":"# header"}}`},
		{"2", card.Header("# header").Red(), `{"template":"red","title":{"tag":"plain_text","content":"# header"}}`},
		{"2", card.Header("# header").Yellow(), `{"template":"yellow","title":{"tag":"plain_text","content":"# header"}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, jsonString(tt.args), tt.name)
		})
	}
}
