package view

import(
	"model"
)

type EditMain struct {
	Title string
	//Active string
	Main model.Main
	Error string
	CurrentUserInfo
}

func GetEditMain() EditMain {
	result := EditMain {
	Title: "Edit",
	//Active: "home",
}
	return result
}