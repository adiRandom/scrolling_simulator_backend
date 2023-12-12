package models

import "backend_scrolling_simulator/lib"

type ApiResponse[T any] struct {
	Success bool
	Error   lib.Error
	Data    *T
}

func NewErrorApiResponse(err lib.Error) ApiResponse[interface{}] {
	return ApiResponse[interface{}]{false, err, nil}
}
