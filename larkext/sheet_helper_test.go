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
package larkext

import (
	"testing"
)

func TestItoCol(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		{"1", 1, "A"},
		{"2", 2, "B"},
		{"3", 26, "Z"},
		{"4", 27, "AA"},
		{"5", 52, "AZ"},
		{"6", 53, "BA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItoCol(tt.args); got != tt.want {
				t.Errorf("ItoCol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCellRange(t *testing.T) {
	type args struct {
		sheetID string
		x1      int
		y1      int
		x2      int
		y2      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"st1", 1, 1, 1, 1}, "st1!A1:A1"},
		{"2", args{"st1", 1, 1, 26, 1}, "st1!A1:Z1"},
		{"3", args{"st1", 1, 1, 27, 1}, "st1!A1:AA1"},
		{"4", args{"st1", 1, 1, 27, 200}, "st1!A1:AA200"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CellRange(tt.args.sheetID, tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("CellRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
