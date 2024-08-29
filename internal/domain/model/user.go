package model

type User struct {
	Id       string
	UserName string
	Email    string
	Passhash []byte
}
