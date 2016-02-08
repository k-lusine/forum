package controler

import (
	"html/template"
	"view"
	"net/http"
	"strings"
	"model"
)

type editUserController struct {
	template *template.Template
	double bool
}
func (this *editUserController) handle(w http.ResponseWriter, r *http.Request) {
	vm := view.GetEditUser()
	res:=strings.Split(r.URL.Path,"/")
	vm.User.Id= res[3]
	if res[3]!="new"&& this.double==false{
		vm.User.GetCurrentUserData()
	}
	this.double=true
	if r.Method == "POST" {
		vm.User.UnmarshalHTTP(r)
		isValid, msg := vm.User.Validate()
		if isValid==true {

			vm.User.Save()
			vm.User = model.User{}
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