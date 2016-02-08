package controler

import(
	"view"
	"net/http"
	"html/template"
	//"model"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, r *http.Request) {
	vm := view.GetHome()
	IsLoggedIn,IsAdmin, user,_ := checkLoginStatus(r)
	vm.IsLoggedIn=IsLoggedIn
	vm.IsAdmin=IsAdmin
	vm.Username=user.Username
	if stat := user.Status_id; stat=="2"{
		vm.Status="user"
		vm.WhatCanDo="Hello, you are an ordinary user and you have restricted preveliges in this forum"
	} else if stat=="1" {
		vm.Status="admin"
		vm.WhatCanDo="Hello, you are admin and you have full preveliges in this forum"
	}

	vm.Threads.GetThreadsByCategory()
	w.Header().Add("Content Type", "text/html")

	this.template.Execute(w,vm)
	//	responseWriter := util.GetResponseWriter(w, r)
	//	defer responseWriter.Close()
//	this.template.Execute(responseWriter,vm)
}

