package models

import (
	"backend_scrolling_simulator/dtos"
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Emoji string `gorm:"not null"`
}

func (t *Topic) ToDto() dtos.Topic {
	return dtos.Topic{
		Id:    t.ID,
		Name:  t.Name,
		Emoji: t.Emoji,
	}
}
