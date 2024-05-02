package entity

type Command struct {
	Type    string `json:"type" form:"type" binding:"required"`
	Content string `json:"content" from:"content" binding:"required"`
}
