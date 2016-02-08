package view

import (
	"model"
)
type Home struct {
	Title string
	//Active string
	Threads model.List_ThrList
	CurrentUserInfo
}

func GetHome() Home {
	result := Home{
		Title: "Forum - Homepage",
		//Active: "home",
	}
	return result
}
