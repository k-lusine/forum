package model_test
import (
	"testing"
	"model"
)

//func Test_RequireEmailUnique(t *testing.T){
//	input := model.User{
//		Email: "lusinekarapetian@gmail.com",
//	}
//	input.Re
//	passed, msg := input.RequireUniqueEmail()
//	if passed==true || msg =="" {
//		t.Fail()
//	}
//
//	input.Email="fasfasjshaj"
//	passed, msg := input.RequireUniqueEmail()
//	if passed==false || msg !="" {
//		t.Fail()
//	}
//}



func Test_ConfirmLogin(t *testing.T){
	input := model.User{
		Username: "lusineKarapetyan",
		Password: "88ls77LS99!a",
	}
	passed, msg, id := model.ConfirmLogin(input.Username, input.Password)
		if passed==false || msg != "" || id == ""{
			t.Fail()
		}
	input = model.User{
		Username: "lusineeKarapetyan",
		Password: "88ls77LS99!a",
	}
	passed, msg, id = model.ConfirmLogin(input.Username, input.Password)
	if passed==true || msg == "" || id != ""{
		t.Fail()
	}
	input = model.User{
		Username: "lusineKarapetyan",
		Password: "88ls7777LS99!a",
	}
	passed, msg, id = model.ConfirmLogin(input.Username, input.Password)
	if passed==true || msg == "" || id != ""{
		t.Fail()
	}

}
