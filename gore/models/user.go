package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
