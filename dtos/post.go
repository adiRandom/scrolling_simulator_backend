package dtos

type Post struct {
	Id       uint    `json:"id"`
	Author   *User   `json:"author"`
	Body     string  `json:"body"`
	Title    string  `json:"title"`
	Subtitle string  `json:"subtitle"`
	Topics   []Topic `json:"topics"`
}

type Topic struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Emoji string `json:"emoji"`
}

type PostIdPathParam struct {
	PostId uint `uri:"postId" binding:"required" json:"post_id"`
}

type ReactResponse struct {
	Ratio float64 `json:"ratio"`
	Text  string  `json:"text"`
}
