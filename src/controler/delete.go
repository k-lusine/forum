package controler
import (
	"net/http"
	"strings"
	"model"
)

type deleteController struct {}
type replyDeleteController struct {}

func (this *deleteController) delete(w http.ResponseWriter, r *http.Request){
	res:=strings.Split(r.URL.Path,"/")
	var err error
	if res[2]=="thread"{
		data := model.Thread{}
		data.Id = res[3]
		err = data.Delete()
	} else if res[2]=="reply"{
		data := model.Reply{}
		data.Id = res[3]
		err = data.Delete()
	}
	if err == nil{
		http.Redirect(w, r, "/", http.StatusFound)
	}
}