package dtos

type Pagination struct {
	Limit  int  `form:"limit" binding:"required"`
	Offset *int `form:"offset"`
}

type PaginatedResponse[T any] struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
	Data       []T `json:"data"`
}

func NewInfinitePaginatedResponse[T any](data []T) PaginatedResponse[T] {
	return PaginatedResponse[T]{Page: 0, PerPage: len(data), Total: len(data) * 2, TotalPages: 1, Data: data}
}
