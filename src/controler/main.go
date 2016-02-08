package controler

import (
	"html/template"
	"net/http"
)
	func Helloworld(w http.ResponseWriter, r *http.Request){
		t,_  := template.ParseFiles("templates/test.html")
		t.Execute(w, nil)
	}

func Register(templates *template.Template) {
//	http.HandleFunc("/", Helloworld)
	//HOME
	hc := new(homeController)
	hc.template = templates.Lookup("home.html")
	http.HandleFunc("/", hc.get)

	//THREAD VIEW
	tc := new(threadController)
	tc.template = templates.Lookup("thread.html")
	http.HandleFunc("/view/", tc.get)

	//EDIT AND SAVE MAIN
	emc := new(editMainController)
	emc.template = templates.Lookup("editMain.html")
	http.HandleFunc("/edit/topic/", emc.handle)

	//EDIT AND SAVE REPLY
	erc := new(editReplyController)
	erc.template = templates.Lookup("editReply.html")
	http.HandleFunc("/edit/reply/", erc.handle)

	//Delete main, it's the same as to delete thread
	dt := new(deleteController)
	http.HandleFunc("/delete/", dt.delete)

	//Register user and edit profile info
	euc := new(editUserController)
	euc.template = templates.Lookup("register.html")
	http.HandleFunc("/edit/user/", euc.handle)

	//Login Controller
	lc := new(loginController)
	lc.template = templates.Lookup("login.html")
	http.HandleFunc("/login/", lc.handle)
	//Logout Controller
	loc := new(logoutController)
	http.HandleFunc("/logout/", loc.get)
}