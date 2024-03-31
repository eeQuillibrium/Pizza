package models

type User struct {
	UserId   int    `json:"userid"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
