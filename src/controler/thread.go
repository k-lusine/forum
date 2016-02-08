package controler

import(

	"html/template"
	"net/http"
	"view"
)
type threadController struct {
	template *template.Template
	}

func (this *threadController) get(w http.ResponseWriter, r *http.Request) {
	vm := view.GetThread()
	IsLoggedIn,IsAdmin, user,_ := checkLoginStatus(r)
	vm.IsLoggedIn=IsLoggedIn
	vm.IsAdmin=IsAdmin
	vm.Username=user.Username
	vm.UserId = user.Id
	if stat := user.Status_id; stat=="2"{
		vm.Status="user"
		vm.WhatCanDo="Hello, you are ordinary and you have restricted preveliges in this forum"
	} else if stat=="1" {

		vm.Status="admin"
		vm.WhatCanDo="Hello, you are admin and you have full preveliges in this forum"
	}
	vm.Thread.Id = r.URL.Path[len("/view/"):]
	vm.Thread.Main.Thr_id = r.URL.Path[len("/view/"):]
	vm.Thread.Main.GetMainData()
	vm.Thread.GetReplies()
	vm.Title = vm.Title + " | " + vm.Thread.Main.Name
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w,vm)

}

//88Ando123!@#