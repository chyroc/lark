package helper

import "strings"

func ClearString(s, target, to string) string {
	for {
		tmp := strings.ReplaceAll(s, target, to)
		if tmp == s {
			return tmp
		}
		s = tmp
	}
}
