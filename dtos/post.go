package dtos

import (
	"backend_scrolling_simulator/lib/functional"
	"backend_scrolling_simulator/models"
)

type Post struct {
	Id       uint         `json:"id"`
	Author   *models.User `json:"author"`
	Body     string       `json:"body"`
	Title    string       `json:"title"`
	Subtitle string       `json:"subtitle"`
	Topics   []Topic      `json:"topics"`
}

type Topic struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

func NewPostFromModel(post models.Post) Post {
	return Post{
		Id:       post.ID,
		Author:   nil,
		Body:     post.Body,
		Title:    post.Title,
		Subtitle: post.Subtitle,
		Topics: functional.Map(post.Topics,
			func(topic models.Topic) Topic { return Topic{Id: topic.ID, Name: topic.Name, Emoji: topic.Emoji} }),
	}
}
