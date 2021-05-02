package lark

func GetErrorCode(err error) int {
	if err != nil {
		if e, ok := err.(*Error); ok {
			return e.Code
		}
	}
	return -1
}
