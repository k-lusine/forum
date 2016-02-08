package view

import(
	"model"
)

type EditReply struct {
	Title string
	//Active string
	Reply model.Reply
	Error string
	CurrentUserInfo
}

func GetEditReply() EditReply {
	result := EditReply {
	Title: "Edit",
	//Active: "home",
}
	return result
}