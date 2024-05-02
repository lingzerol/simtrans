package entity

const (
	HearBeatTimeSpan int64 = 3600
)

type CommonField struct {
	MessageType string `json:"message_type" form:"message_type" default:"required"`
	Timestamp   int64  `json:"timestamp" form:"timestamp" default:"required"`
}
