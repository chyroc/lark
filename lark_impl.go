package lark

import (
	"net/http"
)

type apiImpl struct {
	appID      string
	appSecret  string
	httpClient *http.Client
}
