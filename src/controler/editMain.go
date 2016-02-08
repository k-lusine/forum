package controler

import (
	"html/template"
	"view"
	"net/http"
	"strings"
	"model"
)

type editMainController struct {
	template *template.Template
	double bool
}
func (this *editMainController) handle(w http.ResponseWriter, r *http.Request) {
	vm := view.GetEditMain()

	IsLoggedIn,IsAdmin, user,_ := checkLoginStatus(r)
	vm.IsLoggedIn=IsLoggedIn
	vm.IsAdmin=IsAdmin
	vm.Username=user.Username
	if stat := user.Status_id; stat=="2"{
		vm.Status="user"
		//vm.WhatCanDo="Hello, you are ordinary user and you have restricted preveliges in this forum"
	} else if stat=="1" {
		vm.Status="admin"
		//vm.WhatCanDo="Hello, you are admin and you have full preveliges in this forum"
	}
	res:=strings.Split(r.URL.Path,"/")
	vm.Main.Thr_id = res[3]
	if res[3]!="new"&& this.double==false{
		vm.Main.GetMainData()
	}
	if r.Method == "POST" {
		vm.Main.UnmarshalHTTP(r)
		//fmt.Println(vm.Main)
		isValid, msg := vm.Main.Validate()
		if isValid==true {
			vm.Main.UserID = getcurrentUserID(r)
			vm.Main.Save()
			vm.Main = model.Main{}
			vm.Error = ""
			http.Redirect(w, r, "/", http.StatusFound)
			this.double=false
		} else {
			vm.Error=msg
			this.double=true
		}
	}
	w.Header().Add("Content Type", "text/html")
	this.template.Execute(w,vm)
	//	responseWriter := util.GetResponseWriter(w, r)
	//	defer responseWriter.Close()
	//	this.template.Execute(responseWriter,vm)
}
