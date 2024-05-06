package entity

const (
	CommandExpireTime = 600
)

type Command struct {
	ID        uint64 `json:"id" form:"id" binding:"required"`
	Type      string `json:"type" form:"type" binding:"required"`
	Content   string `json:"content" from:"content" binding:"omitempty"`
	TimeStamp int64  `json:"timestamp" form:"timestamp" binding:"required"`
}

type CommandParams struct {
	Command
}
