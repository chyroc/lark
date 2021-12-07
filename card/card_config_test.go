package card_test

import (
	"testing"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/card"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name string
		args *lark.MessageContentCardConfig
		want string
	}{
		{"1", card.Config(), `{}`},
		{"2", card.Config().WithEnableForward(true), `{"enable_forward":true}`},
		{"3", card.Config().WithUpdateMulti(true), `{"update_multi":true}`},
		{"4", card.Config().WithEnableForward(true).WithUpdateMulti(true), `{"enable_forward":true,"update_multi":true}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, jsonString(tt.args), tt.name)
		})
	}
}
