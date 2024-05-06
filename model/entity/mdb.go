package entity

const (
	DefaultPageSize int64 = 10
)

type PageOrder struct {
	PageNo   int64  `json:"page_no" form:"page_no"`
	PageSize int64  `json:"page_size" form:"page_size"`
	OrderBy  string `json:"order_by" form:"order_by"`
	OrderSeq string `json:"order_seq" form:"order_seq" binding:"oneof=asc desc"`
}
