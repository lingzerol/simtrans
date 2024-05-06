package errno

import "fmt"

func NewCodeError(code int, msg string, err error) error {
	return &CodeErr{
		code:   code,
		errmsg: msg,
		err:    err,
	}
}

func WrapCodeError(code int, err error) error {
	return NewCodeError(code, "", err)
}

func WrapCodeErrorf(code int, format string, args ...any) error {
	return NewCodeError(code, fmt.Sprintf(format, args...), nil)
}

func NewDefaultCodeError(code int) error {
	return &CodeErr{
		code:   code,
		errmsg: ErrMsg(code),
	}
}
