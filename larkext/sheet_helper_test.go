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
