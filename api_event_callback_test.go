package lark

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventCallbackService_parserAllCallbackRequest(t *testing.T) {
	as := assert.New(t)
	r := &EventCallbackService{
		cli: New(),
	}
	got, err := r.parserAllCallbackRequest(context.Background(), http.Header{}, []byte(`{
"open_id": "ou_sdfimx9948345","user_id": "eu_sd923r0sdf5","open_message_id": "om_abcdefg1234567890",
"tenant_key": "d32004232","token": "c-xxxxx"}`))
	as.Nil(err)
	as.Equal("c-xxxxx", got.EventCardCallback.Token)
}
