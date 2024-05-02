package entity

const (
	MaxBufferSize             = 65536
	AuthTimeOut        int    = 300
	DefaultAuthMessage string = "simtrans auth success"
)

type AuthParams struct {
	CommonField
}

type AuthResonse struct {
	CommonField
	DeviceName string `json:"device_name" form:"device_name"`
	Message    string `json:"message" form:"message"`
}
