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
	"fmt"
)

// AtBuilder ...
var AtBuilder atBuilder

type atBuilder string

// AtAll ...
func (r atBuilder) AtAll() string {
	return `<at user_id="all"></at>`
}

// AtOpenID ...
func (r atBuilder) AtOpenID(openID string) string {
	return fmt.Sprintf(`<at user_id=%q></at>`, openID)
}
