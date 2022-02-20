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

	"github.com/chyroc/lark"
	"github.com/stretchr/testify/assert"
)

func Test_Error(t *testing.T) {
	as := assert.New(t)

	t.Run("", func(t *testing.T) {
		as.Equal("TRACE", lark.LogLevelTrace.String())
		as.Equal("DEBUG", lark.LogLevelDebug.String())
		as.Equal("INFO", lark.LogLevelInfo.String())
		as.Equal("WARN", lark.LogLevelWarn.String())
		as.Equal("ERROR", lark.LogLevelError.String())
		as.Equal("", (lark.LogLevel(0)).String())
	})

	t.Run("", func(t *testing.T) {
		var err error = &lark.Error{
			Scope:    "",
			FuncName: "",
			Code:     0,
			Msg:      "",
		}
		as.Equal("", err.Error())
	})
}
