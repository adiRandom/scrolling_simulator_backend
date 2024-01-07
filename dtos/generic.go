package dtos

type Pagination struct {
	Limit int `form:"limit" binding:"required"`
}
