package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

type App struct {
	AppID     string
	AppSecret string
}

func (r *App) Ins() lark.API {
	return lark.New(r.AppID, r.AppSecret)
}

type User struct {
	UserID string
	Name   string
}

type Chat struct {
	ChatID string
	Name   string
}

var AppNoPermission = App{
	AppID:     os.Getenv("LARK_APP_NO_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_NO_PERMISSION_APP_SECRET"),
}

var AppALLPermission = App{
	AppID:     os.Getenv("LARK_APP_ALL_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_ALL_PERMISSION_APP_SECRET"),
}

var UserAdmin = User{
	UserID: os.Getenv("LARK_USER_ADMIN_USER_ID"),
	Name:   os.Getenv("LARK_USER_ADMIN_NAME"),
}

// 这个群公共，必须设置「群公共」三个字
var ChatContainALLPermissionApp = Chat{
	ChatID: os.Getenv("LARK_CHAT_CONTAINS_APP_PERMISSION_APP_CHAT_ID"),
	Name:   "包含「lark-sdk」的群",
}

var ChatNotContainALLPermissionApp = Chat{
	ChatID: os.Getenv("LARK_CHAT_NOT_CONTAINS_APP_PERMISSION_APP_CHAT_ID"),
	Name:   "不包含「lark-sdk」的群",
}

func Test_Config(t *testing.T) {
	as := assert.New(t)

	as.NotEmpty(AppNoPermission.AppID)
	as.NotEmpty(AppNoPermission.AppSecret)
	as.NotEmpty(AppALLPermission.AppID)
	as.NotEmpty(AppALLPermission.AppSecret)
	as.NotEmpty(UserAdmin.UserID)
	as.NotEmpty(ChatContainALLPermissionApp.ChatID)
	as.NotEmpty(ChatNotContainALLPermissionApp.ChatID)
}
