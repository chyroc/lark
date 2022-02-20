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
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_AppLink(t *testing.T) {
	as := assert.New(t)

	t.Run("bot", func(t *testing.T) {
		res := lark.AppLink.OpenBot(&lark.OpenBotReq{AppID: "cli_x"})
		as.Equal("https://applink.feishu.cn/client/bot/open?appId=cli_x", res)
	})

	t.Run("calendar", func(t *testing.T) {
		res := lark.AppLink.OpenCalender(&lark.OpenCalenderReq{})
		as.Equal("https://applink.feishu.cn/client/calendar/open", res)
	})

	t.Run("calender_account", func(t *testing.T) {
		res := lark.AppLink.OpenCalenderAccount(&lark.OpenCalenderAccountReq{})
		as.Equal("https://applink.feishu.cn/client/calendar/account", res)
	})

	t.Run("calender_event_create", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			res := lark.AppLink.OpenCalenderEventCreate(&lark.OpenCalenderEventCreateReq{
				StartTime: ptr.String("1581950880"),
			})
			as.Equal("https://applink.feishu.cn/client/calendar/event/create?startTime=1581950880", res)
		})
		t.Run("", func(t *testing.T) {
			as.Equal("https://applink.feishu.cn/client/calendar/event/create?endTime=1581951000&startTime=1581950880", lark.AppLink.OpenCalenderEventCreate(&lark.OpenCalenderEventCreateReq{
				StartTime: ptr.String("1581950880"),
				EndTime:   ptr.String("1581951000"),
			}))
		})
		t.Run("", func(t *testing.T) {
			res := lark.AppLink.OpenCalenderEventCreate(&lark.OpenCalenderEventCreateReq{
				StartTime: ptr.String("1581950880"),
				Summary:   ptr.String("主题"),
			})
			as.Equal("https://applink.feishu.cn/client/calendar/event/create?startTime=1581950880&summary=%E4%B8%BB%E9%A2%98", res)
		})
	})

	t.Run("calendar_view", func(t *testing.T) {
		as.Equal("https://applink.feishu.cn/client/calendar/view?type=week", lark.AppLink.OpenCalenderView(&lark.OpenCalenderViewReq{
			Type: ptr.String("week"),
		}))
		as.Equal("https://applink.feishu.cn/client/calendar/view?type=meeting", lark.AppLink.OpenCalenderView(&lark.OpenCalenderViewReq{
			Type: ptr.String("meeting"),
		}))
		as.Equal("https://applink.feishu.cn/client/calendar/view?date=1581999948&type=day", lark.AppLink.OpenCalenderView(&lark.OpenCalenderViewReq{
			Type: ptr.String("day"),
			Date: ptr.String("1581999948"),
		}))
	})

	t.Run("chat", func(t *testing.T) {
		as.Equal("https://applink.feishu.cn/client/chat/open?openId=1234567890", lark.AppLink.OpenChat(&lark.OpenChatReq{
			OpenID: ("1234567890"),
		}))
		as.Equal("https://applink.feishu.cn/client/chat/open?openChatId=oc_41e7bdf4877cfc316136f4ccf6c32613", lark.AppLink.OpenChat(&lark.OpenChatReq{
			OpenChatID: "oc_41e7bdf4877cfc316136f4ccf6c32613",
		}))
	})

	t.Run("docs", func(t *testing.T) {
		as.Equal("https://applink.feishu.cn/client/docs/open?URL=https%3A%2F%2Fbytedance.feishu.cn%2Fdocs%2Fdoccn9EOHrnB0r0iEN9HoCPczbf", lark.AppLink.OpenDocs(&lark.OpenDocsReq{
			URL: "https://bytedance.feishu.cn/docs/doccn9EOHrnB0r0iEN9HoCPczbf",
		}))
	})

	t.Run("lark", func(t *testing.T) {
		as.Equal("https://applink.feishu.cn/client/op/open", lark.AppLink.OpenLark(&lark.OpenLarkReq{}))
	})

	t.Run("mini-pro", func(t *testing.T) {
		as.Equal("https://applink.feishu.cn/client/mini_program/open?appId=1234567890&mode=window", lark.AppLink.OpenMiniProgram(&lark.OpenMiniProgramReq{
			AppID: "1234567890",
			Mode:  ptr.String("window"),
		}))
		as.Equal("https://applink.feishu.cn/client/mini_program/open?appId=1234567890&mode=window&path=pages%2Fhome", lark.AppLink.OpenMiniProgram(&lark.OpenMiniProgramReq{
			AppID: "1234567890",
			Mode:  ptr.String("window"),
			Path:  ptr.String("pages/home"),
		}))
		as.Equal("https://applink.feishu.cn/client/mini_program/open?appId=1234567890&mode=window&path=pages%2Fhome%3Fxid%3D123", lark.AppLink.OpenMiniProgram(&lark.OpenMiniProgramReq{
			AppID: "1234567890",
			Mode:  ptr.String("window"),
			Path:  ptr.String("pages/home?xid=123"),
		}))
		as.Equal("https://applink.feishu.cn/client/mini_program/open?appId=1234567890&mode=window&path=pages%2Fhome%3Fxid%3D123&path_pc=pages%2Fpc_home%3Fpid%3D123", lark.AppLink.OpenMiniProgram(&lark.OpenMiniProgramReq{
			AppID:  "1234567890",
			Mode:   ptr.String("window"),
			Path:   ptr.String("pages/home?xid=123"),
			PathPc: ptr.String("pages/pc_home?pid=123"),
		}))
		// todo: min_lk_ver_pc
	})

	t.Run("open-sso", func(t *testing.T) {
		as.Equal("https://applink.feishu.cn/client/passport/sso_login?sso_domain=sso.domain.com&tenant_name=tenant-id",
			lark.AppLink.OpenSSOLogin(&lark.OpenSSOLoginReq{
				SSODomain:  "sso.domain.com",
				TenantName: "tenant-id",
			}))
	})

	t.Run("open-web", func(t *testing.T) {
		as.Equal("https://applink.feishu.cn/client/web_app/open?appId=app-id",
			lark.AppLink.OpenWebApp(&lark.OpenWebAppReq{
				AppID: "app-id",
			}))
	})

	t.Run("open-web-url", func(t *testing.T) {
		as.Equal("https://applink.feishu.cn/client/web_url/open?mode=window&url=google.com",
			lark.AppLink.OpenWebURL(&lark.OpenWebURLReq{
				URL:  "google.com",
				Mode: "window",
			}))
	})
}
