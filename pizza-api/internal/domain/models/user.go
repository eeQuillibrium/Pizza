package models

type User struct {
	UserId   int    `json:"userid"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
