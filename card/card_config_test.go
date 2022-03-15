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
package card_test

import (
	"testing"

	"github.com/chyroc/lark"
	"github.com/chyroc/lark/card"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name string
		args *lark.MessageContentCardConfig
		want string
	}{
		{"1", card.Config(), `{}`},
		{"2", card.Config().SetEnableForward(true), `{"enable_forward":true}`},
		{"3", card.Config().SetUpdateMulti(true), `{"update_multi":true}`},
		{"4", card.Config().SetEnableForward(true).SetUpdateMulti(true), `{"enable_forward":true,"update_multi":true}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, jsonString(tt.args), tt.name)
		})
	}
}
