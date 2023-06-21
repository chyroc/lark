// Code generated by lark_sdk_gen. DO NOT EDIT.
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

package lark

import (
	"github.com/chyroc/lark/internal"
)

// OpenWebApp 打开网页应用
//
// 打开一个已安装的H5应用
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-an-h5-app
// new doc: https://open.feishu.cn/document/common-capabilities/applink-protocol/supported-protocol/open-an-h5-app
func (r *AppLinkService) OpenWebApp(req *OpenWebAppReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/web_app/open", req)
}

// OpenWebAppReq ...
type OpenWebAppReq struct {
	AppID       string  `json:"appId,omitempty"`         // H5应用的 appId(可从「开发者后台-凭证与基础信息」获取)
	Mode        *string `json:"mode,omitempty"`          // 打开H5应用的容器模式, 枚举值包括  `appCenter`: 在工作台打开, 3.20版本开始支持（缺省值）   `window`: 在独立窗口打开, 3.20版本开始支持   `sidebar`: 在侧边栏打开, 3.40版本开始支持   `window-semi`: 在独立窗口以小屏形式打开, 5.10版本开始支持
	Height      *string `json:"height,omitempty"`        // 自定义独立窗口高度（仅当`mode`为`window`时生效）, 飞书5.12版本开始支持 最小值: 480 最大值: 屏幕的高度 默认值: 飞书窗口的高度
	Width       *string `json:"width,omitempty"`         // 自定义独立窗口宽度（仅当`mode`为`window`时生效）, 飞书5.12版本开始支持 最小值: 640 最大值: 屏幕的宽度 默认值: 飞书窗口的宽度
	Path        *string `json:"path,omitempty"`          // 指定要打开网页应用的某个页面路径。配置path参数后, 此参数将会被添加到或者替换网页应用原始url中的path部分（原始url是在开发者后台的网页应用配置页进行配置的）, 生成要打开页面的最终url。    注意: 1.path 的含义详见[Applink 的结构](https://open.feishu.cn/document/uYjL24iN/ucjN1UjL3YTN14yN2UTN), 不能出现'#'和'?'这样的字符（即不能携带query参数和fragment）, 否则会造成要打开的页面url结构异常, 导致页面打开的表现不符合预期  2.如果需要携带query参数或fragment（query和fragment的含义详见 Applink 的结构）, 推荐使用下面的lk_target_url参数   3.可以使用 path_android、path_ios、path_pc 参数对不同的客户端指定不同的path   4.该参数从 3.20版本 开始支持
	PathAndroid *string `json:"path_android,omitempty"`  // 同 path 参数, Android 端会优先使用该参数, 如果该参数不存在, 则会使用 path 参数。  3.20版本开始支持
	PathIos     *string `json:"path_ios,omitempty"`      // 同 path 参数, iOS 端会优先使用该参数, 如果该参数不存在, 则会使用 path 参数  3.20版本开始支持
	PathPc      *string `json:"path_pc,omitempty"`       // 同 path 参数, PC 端会优先使用该参数, 如果该参数不存在, 则会使用 path 参数  3.20版本开始支持
	LkTargetURL *string `json:"lk_target_url,omitempty"` // 指定要打开网页应用的某个页面完整url   注意: 1.此url参数中的scheme和domain必须和网页应用在开发者后台配置的url的scheme和domain一致   2.此url参数必须进行encode处理, 如果url携带query参数, 并且query参数中也包含 url 或 中文 等非ASCII码字符, 必须先对query参数中的非ASCII码字符进行encode后, 作为参数的值, 然后对url整体进行encode。具体用法详见实例3  3. 该参数从 5.12版本 开始支持, 低版本无法解析此参数, 如果在低版本使用, 将打开开发者后台配置的网页应用首页  4. 该参数和path 系列参数互斥, 同时传递时lk_target_url 优先级更高
	Reload      *string `json:"reload,omitempty"`        // 如果网页应用当前所处的页面路径和 applink 要打开的页面路径相同时: true: 重新加载页面   false: 保留原页面状态   默认值: false   5.20版本开始支持
}
