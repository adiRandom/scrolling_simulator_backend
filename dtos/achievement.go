package dtos

type Achievement struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	BronzePoints uint   `json:"bronzePoints"`
	SilverPoints uint   `json:"silverPoints"`
	GoldPoints   uint   `json:"goldPoints"`
	IconUrl      string `json:"iconUrl"`
	Unlocked     bool   `json:"unlocked"`
}
