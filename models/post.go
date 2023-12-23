package models

import (
	"backend_scrolling_simulator/dtos"
	"backend_scrolling_simulator/lib/functional"
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

func (p *Post) ToDto() dtos.Post {
	return dtos.Post{
		Id:       p.ID,
		Body:     p.Body,
		Title:    p.Title,
		Subtitle: p.Subtitle,
		Topics:   functional.Map(p.Topics, func(topic Topic) dtos.Topic { return topic.ToDto() }),
	}
}
