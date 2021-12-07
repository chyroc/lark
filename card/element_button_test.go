package card_test

import (
	"testing"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/card"
	"github.com/stretchr/testify/assert"
)

func TestButton(t *testing.T) {
	tests := []struct {
		name string
		args *lark.MessageContentCardElementButton
		want string
	}{
		{"1", card.Button("文字", map[string]interface{}{"key": "val"}), `{"tag":"button","text":{"tag":"plain_text","content":"文字"},"type":"default","value":{"key":"val"}}`},
		{"2", card.Button("文字", map[string]interface{}{"key": "val"}).Primary(), `{"tag":"button","text":{"tag":"plain_text","content":"文字"},"type":"primary","value":{"key":"val"}}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, jsonString(tt.args), tt.name)
		})
	}
}
