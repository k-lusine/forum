package controler

import (
	"html/template"
	"view"
	"net/http"
	"strings"
	"model"
)

type editReplyController struct {
	template *template.Template
	double bool
}
func (this *editReplyController) handle(w http.ResponseWriter, r *http.Request) {
	vm := view.GetEditReply()
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
	if this.double == false  {
		if len(res)>=5{
			if res[4]=="new" {
			vm.Reply.Thread_id = res[3]
		}
		} else{
			vm.Reply.Id=res[3]
			vm.Reply.GetSingleReply()
		}
	}
	this.double=true
	if r.Method == "POST" {
		vm.Reply.UnmarshalHTTP(r)
		isValid, msg := vm.Reply.Validate()
		if isValid==true {
			vm.Reply.UserID = getcurrentUserID(r)
			vm.Reply.Save()
			vm.Reply = model.Reply{}
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


