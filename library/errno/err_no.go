package errno

const (
	StatusOK int = iota
	SystemError
	EmptyCheckFailed
	AuthFailed
	ParamsError
	ConnSendFailed
	ConnReceiveFailed
	DecryptError
	EncryptError
	NilError
	DatabaseError
	RequrestTimeOut
	RepeatRequest
)

var ErrCodeMsg = map[int]string{
	SystemError:   "system error",
	RepeatRequest: "repeat request",
}

func ErrMsg(code int) string {
	if msg, ok := ErrCodeMsg[code]; ok {
		return msg
	}
	return ErrCodeMsg[SystemError]
}
