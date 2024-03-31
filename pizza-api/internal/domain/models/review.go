package models

type Review struct {
	UserId int    `json:"userid"`
	Text   string `json:"text"`
}
