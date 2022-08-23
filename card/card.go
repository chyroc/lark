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
	"github.com/chyroc/lark"
)

func Card(modules ...lark.MessageContentCardModule) *lark.MessageContentCard {
	return &lark.MessageContentCard{
		Header:      nil,
		Config:      Config(),
		Modules:     modules,
		I18NModules: nil,
	}
}

func I18NCard(module *lark.MessageContentCardI18NModule) *lark.MessageContentCard {
	return &lark.MessageContentCard{
		Header:      nil,
		Config:      Config(),
		Modules:     nil,
		I18NModules: module,
	}
}
