package larkext

func itoCol(index int) string {
	var (
		columnBase = 26
		digitMax   = 7 // ceil(log26(Int32.Max))
		digits     = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	)

	if index <= 0 {
		panic("index <= 0")
	}

	if index <= columnBase {
		return string([]byte{digits[index-1]})
	}

	current := index
	offset := digitMax
	res := make([]byte, digitMax)
	for current > 0 {
		offset -= 1
		current -= 1
		res[offset] = digits[current%columnBase]
		current /= columnBase
	}

	return string(res[offset:])
}
