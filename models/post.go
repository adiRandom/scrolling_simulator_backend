package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string  `gorm:"not null"`
	Body     string  `gorm:"not null"`
	Subtitle string  `gorm:"default:''"`
	ImageUrl string  `gorm:"not null"`
	Topics   []Topic `gorm:"many2many:post_topics;"`
	Tags     []Tag   `gorm:"many2many:post_tags;"`
}
