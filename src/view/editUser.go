package view

import(
	"model"
)

type EditUser struct {
	Title string
	//Active string
	User model.User
	Error string
	CurrentUserInfo
}

func GetEditUser() EditUser {
	result := EditUser {
	Title: "Edit",
	//Active: "home",
}
	return result
}