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

func Test_Config(t *testing.T) {
	as := assert.New(t)

	as.NotEmpty(AppNoPermission.AppID)
	as.NotEmpty(AppNoPermission.AppSecret)
	as.NotEmpty(AppALLPermission.AppID)
	as.NotEmpty(AppALLPermission.AppSecret)
	as.NotEmpty(UserAdmin.UserID)
}
