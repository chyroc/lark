package lark

type MethodOption struct {
	userAccessToken string
}

type MethodOptionFunc func(*MethodOption)

func WithUserAccessToken(token string) MethodOptionFunc {
	return func(option *MethodOption) {
		option.userAccessToken = token
	}
}

func newMethodOption(options []MethodOptionFunc) *MethodOption {
	opt := new(MethodOption)
	for _, v := range options {
		v(opt)
	}
	return opt
}
