package view
import "model"
//import "controler"

type CurrentUserInfo struct {
	IsLoggedIn bool
	IsAdmin bool
	Username string
	Status string
	WhatCanDo string
	UserId string
}
type Login struct {
	Title string
	Error string
	User model.User
	CurrentUserInfo
}

//type loginInfo struct {
//	Username string
//	Password string
//}

func GetLogin() Login {
	result := Login{
		Title: "Login",
		//Active: "home",
	}
	return result
}