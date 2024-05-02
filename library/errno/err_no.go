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
)
