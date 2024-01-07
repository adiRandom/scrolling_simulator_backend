package dtos

type User struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatarUrl"`
}
