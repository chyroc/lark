/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package test

import (
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

type App struct {
	AppID        string
	AppSecret    string
	CustomURL    string
	CustomSecret string
}

func (r *App) Ins() *lark.Lark {
	if IsInCI() {
		return lark.New(
			lark.WithAppCredential(r.AppID, r.AppSecret),
			lark.WithTimeout(time.Second*20),
		)
	}
	return lark.New(
		lark.WithAppCredential(r.AppID, r.AppSecret),
		lark.WithTimeout(time.Second*20),
		lark.WithLogger(lark.NewLoggerStdout(), lark.LogLevelTrace),
	)
}

func (r *App) CustomBot() *lark.Lark {
	return lark.New(
		lark.WithCustomBot(r.CustomURL, r.CustomSecret),
		lark.WithTimeout(time.Second*20),
		// lark.WithLogger(lark.NewLoggerStdout(), lark.LogLevelDebug),
	)
}

type Helpdesk struct {
	AppID     string
	AppSecret string
	ID        string
	Token     string
}

func (r *Helpdesk) Ins() *lark.Lark {
	return lark.New(
		lark.WithAppCredential(r.AppID, r.AppSecret),
		lark.WithHelpdeskCredential(r.ID, r.Token),
		lark.WithTimeout(time.Second*20),
	)
}

var HelpdeskAllPermission = Helpdesk{
	AppID:     os.Getenv("LARK_APP_ALL_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_ALL_PERMISSION_APP_SECRET"),
	ID:        os.Getenv("LARK_HELPDESK_ALL_PERMISSION_ID"),
	Token:     os.Getenv("LARK_HELPDESK_ALL_PERMISSION_TOKEN"),
}

type User struct {
	UserID       string
	OpenID       string
	Name         string
	AccessToken  map[string]string
	RefreshToken map[string]string
}

func (r User) OneAccessToken() string {
	for _, v := range r.AccessToken {
		return v
	}
	return ""
}

type Chat struct {
	ChatID string
	Name   string
}

var AppNoPermission = App{
	AppID:     os.Getenv("LARK_APP_NO_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_NO_PERMISSION_APP_SECRET"),
}

var AppAllPermission = App{
	AppID:     os.Getenv("LARK_APP_ALL_PERMISSION_APP_ID"),
	AppSecret: os.Getenv("LARK_APP_ALL_PERMISSION_APP_SECRET"),
}

var AppCustomBotNoValid = App{
	CustomURL: os.Getenv("LARK_APP_CUSTOM_BOT_NO_VALID_WEBHOOK_URL"),
}

var AppCustomBotCheckCanSendWord = App{
	CustomURL: os.Getenv("LARK_APP_CUSTOM_BOT_CHECK_CAN_SEND_WEBHOOK_URL"),
}

var AppCustomBotCheckSign = App{
	CustomURL:    os.Getenv("LARK_APP_CUSTOM_BOT_CHECK_SIGN_WEBHOOK_URL"),
	CustomSecret: os.Getenv("LARK_APP_CUSTOM_BOT_CHECK_SIGN_SIGN"),
}

var UserAdmin = User{
	UserID: os.Getenv("LARK_USER_ADMIN_USER_ID"),
	OpenID: os.Getenv("LARK_USER_ADMIN_OPEN_ID"),
	Name:   os.Getenv("LARK_USER_ADMIN_NAME"),
	AccessToken: map[string]string{
		os.Getenv("LARK_APP_ALL_PERMISSION_APP_ID"): os.Getenv("LARK_ACCESS_TOKEN_ALL_PERMISSION_APP"),
	},
	RefreshToken: map[string]string{
		os.Getenv("LARK_APP_ALL_PERMISSION_APP_ID"): os.Getenv("LARK_REFRESH_TOKEN_ALL_PERMISSION_APP"),
	},
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

var ChatForSendMessage = Chat{
	ChatID: os.Getenv("LARK_CHAT_FOR_SEND_MSG_CHAT_ID"),
	Name:   "for-send-message",
}

type File struct {
	Key string
}

var File1 = File{
	Key: os.Getenv("LARK_FILE_KEY_TEST_FILE_1_PNG"), // this is file of ./test/file_1.png
}

var File2 = File{
	Key: os.Getenv("LARK_FILE_KEY_TEST_FILE_2_DOC"), // ./test/file_2.docx
}

type Message struct {
	MessageID string
	ChatID    string
}

var MessageAdminSendTextInChatContainAllPermissionApp = Message{
	MessageID: os.Getenv("LARK_MESSAGE_ADMIN_SEND_TEST_IN_CHAT_CONTAINS_ALL_PERMISSION_APP"),
	ChatID:    os.Getenv("LARK_CHAT_FOR_SEND_MSG_CHAT_ID"),
}

var MessageAdminSendImageInChatContainAllPermissionApp = Message{
	MessageID: os.Getenv("LARK_MESSAGE_ADMIN_SEND_IMAGE_IN_CHAT_CONTAINS_ALL_PERMISSION_APP"),
	ChatID:    os.Getenv("LARK_CHAT_FOR_SEND_MSG_CHAT_ID"),
}

var MessageAllPermissionAppSendTextInChatContainAllPermissionApp = Message{
	MessageID: os.Getenv("LARK_MESSAGE_ALL_PERMISSION_APP_SEND_TEXT_IN_CHAT_CONTAINS_ALL_PERMISSION_APP"),
	ChatID:    os.Getenv("LARK_CHAT_CONTAINS_APP_PERMISSION_APP_CHAT_ID"),
}

type Approval struct {
	Code string `json:"code"`
}

var ApprovalALLField = Approval{
	Code: os.Getenv("LARK_APPROVAL_ALL_FIELD"),
}

func Test_Config(t *testing.T) {
	as := assert.New(t)

	as.NotEmpty(AppNoPermission.AppID)
	as.NotEmpty(AppNoPermission.AppSecret)
	as.NotEmpty(AppAllPermission.AppID)
	as.NotEmpty(AppAllPermission.AppSecret)
	as.NotEmpty(UserAdmin.UserID)
	as.NotEmpty(UserAdmin.OpenID)
	as.NotEmpty(ChatContainALLPermissionApp.ChatID)
	as.NotEmpty(ChatNotContainALLPermissionApp.ChatID)
	as.NotEmpty(ChatForSendMessage.ChatID)
	as.NotEmpty(File1.Key)
	as.NotEmpty(File2.Key)
	as.NotEmpty(MessageAdminSendTextInChatContainAllPermissionApp.ChatID)
	as.NotEmpty(MessageAdminSendTextInChatContainAllPermissionApp.MessageID)
	as.NotEmpty(MessageAdminSendImageInChatContainAllPermissionApp.ChatID)
	as.NotEmpty(MessageAdminSendImageInChatContainAllPermissionApp.MessageID)
	as.NotEmpty(MessageAllPermissionAppSendTextInChatContainAllPermissionApp.ChatID)
	as.NotEmpty(MessageAllPermissionAppSendTextInChatContainAllPermissionApp.MessageID)
}

type fakeHTTPWriter struct {
	header http.Header
	code   int
	lock   sync.Mutex
	data   []byte
}

func newFakeHTTPWriter() *fakeHTTPWriter {
	return &fakeHTTPWriter{
		header: map[string][]string{},
	}
}

func (r *fakeHTTPWriter) Header() http.Header {
	return r.header
}

func (r *fakeHTTPWriter) Write(bytes []byte) (int, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.data = append(r.data, bytes...)
	return len(bytes), nil
}

func (r *fakeHTTPWriter) WriteHeader(statusCode int) {
	r.code = statusCode
}

func (r *fakeHTTPWriter) str() string {
	return string(r.data)
}
