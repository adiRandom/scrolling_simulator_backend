package models

import (
	"backend_scrolling_simulator/dtos"
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	name  string `gorm:"not null"`
	emoji string `gorm:"not null"`
}

func (t *Topic) ToDto() dtos.Topic {
	return dtos.Topic{
		Id:    t.ID,
		Name:  t.name,
		Emoji: t.emoji,
	}
}
