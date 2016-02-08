package view

import(
	"model"
)

type Thread struct {
	Title string
	//Active string
	Thread model.Thread
	CurrentUserInfo
}

func GetThread() Thread {
	result := Thread{
		Title: "View",
		//Active: "home",
	}
	return result
}
