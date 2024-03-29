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

// OpenNativeApp 打开原生集成应用
//
// 打开一个原生集成应用。如需接入原生集成应用，请参考：
// - [Android](https://open.feishu.cn/document/uAjLw4CM/ukzMukzMukzM/native-integration/open-scene-introduction/protocol-components/native-integrated-application/andr)
// - [iOS](https://open.feishu.cn/document/uAjLw4CM/ukzMukzMukzM/native-integration/open-scene-introduction/protocol-components/native-integrated-application/ios-)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/uYjL24iN/applink-protocol/supported-protocol/open-a-native-app
func (r *AppLinkService) OpenNativeApp(req *OpenNativeAppReq) string {
	return internal.JoinAppLinkURL("https://applink.feishu.cn/client/native_app/open", req)
}

// OpenNativeAppReq ...
type OpenNativeAppReq struct {
	AppID string `json:"appId,omitempty"` // 原生应用 appId(可从「开发者后台-凭证与基础信息」获取)
}
