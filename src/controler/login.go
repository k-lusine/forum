package controler
import (
	"html/template"
	"net/http"
	"view"
	"model"
	"time"
)

type loginController struct {
	template *template.Template
}

func (this *loginController) handle (w http.ResponseWriter, r *http.Request){
	vm := view.GetLogin()
	IsLoggedIn,IsAdmin, user,_ := checkLoginStatus(r)

	vm.IsLoggedIn=IsLoggedIn
	vm.IsAdmin=IsAdmin
	vm.Username=user.Username
	if stat := user.Status_id; stat=="2"{

		vm.Status="user"
		vm.WhatCanDo="Hello, you are ordinary and you have restricted preveliges in this forum"
	} else if stat=="1" {

		vm.Status="admin"
		vm.WhatCanDo="Hello, you are admin and you have full preveliges in this forum"
	}

	if vm.IsLoggedIn==true{

		http.Redirect(w, r, "/home", http.StatusFound)
	} else{
		w.Header().Add("Content Type", "text/html")
		if r.Method=="POST"{
			vm.User.UnmarshalHTTP(r)
			passed, msg, userid := model.ConfirmLogin(vm.User.Username, vm.User.Password)
			if passed == false {
				vm.Error = msg
			} else{
				vm.User.Id = userid
				cookieData :=vm.User.CreateSession()
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = cookieData.SessionId
				cookie.Expires = time.Now().Add(300 * time.Second)
				cookie.Path="/"
				w.Header().Add("Set-Cookie", cookie.String())
				http.Redirect(w, r, "/home", http.StatusFound)
			}
		}
		this.template.Execute(w,vm)
	}


}

func checkLoginStatus(r *http.Request)(isLoggedIn bool, isAdmin bool, user model.User, err error ){
	sessId, err :=r.Cookie("sessionId")
	if err == nil{
		sess := sessId.Value
		user, err = model.GetMemberBySessionId(sess)

		if err != nil{
			user = model.User{}

		} else {
			isLoggedIn = true
			if user.Status_id == "1"{
				isAdmin = true
			}
		}
	}
	return
}

func getcurrentUserID (r *http.Request) (id string){
	_, _, user, err := checkLoginStatus(r)
	if err==nil{
		id = user.Id
	}
		return
}