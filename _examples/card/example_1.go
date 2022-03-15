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
package card

import (
	"fmt"

	"github.com/chyroc/lark/card"
)

func Example1() {
	content := card.Card(
		card.ModuleImage("img_v2_fddd29cd-2846-4a03-aaed-d22878e503fg"),
		card.Div().SetText(card.MarkdownText("圣诞老人赶着麋鹿在平安夜悄悄光临办公楼，为大家带来了丰盛的下午茶～🎅\nBUT...圣诞老人走得太急，忘记带礼物了！！😢\n\n为活跃办公室气氛，增加同事间情谊，我们将举办圣诞礼物交换派对～！🥂")),
		card.Div().SetText(card.MarkdownText("**🎄 圣诞派对时间：**12月24日 14:00-17:00\n\n**🎁 礼物交换方式：**请各位小伙伴们在包装好你准备的礼物之后，附上卡片祝福语，在 12 月 23 日下午 14:00 前交给前台，我们会统一交到圣诞老人手里～")),
	).
		SetHeader(card.Header("🔔 叮～你有一封圣诞邀请函待查收 🎁"))
	fmt.Println(content.String())
}
