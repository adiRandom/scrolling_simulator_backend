package dtos

import "backend_scrolling_simulator/models"

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
