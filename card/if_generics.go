// +build go1.18

package card

func If[T any](cond bool, val T, otherwise T) T {
	if cond {
		return val
	}
	return otherwise
}

func IfLay[T any](cond bool, val func() T, otherwise func() T) T {
	if cond {
		return val()
	}
	return otherwise()
}
