package models

import (
	"backend_scrolling_simulator/dtos"
	"gorm.io/gorm"
)

type ReactText struct {
	gorm.Model
	Ratio float64 `gorm:"not null"`
	Text  string  `gorm:"not null"`
}

func (r *ReactText) ToDto() dtos.ReactResponse {
	return dtos.ReactResponse{
		Ratio: r.Ratio,
		Text:  r.Text,
	}
}
