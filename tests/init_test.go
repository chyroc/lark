package tests

import (
	"os"
)

type App struct {
	AppID     string
	AppSecret string
}

var AppNoPermission = App{
	AppID:     os.Getenv("LARK_APP_NO_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_NO_PERMISSION_APP_SECRET"),
}
