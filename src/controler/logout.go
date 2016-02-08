package controler
import (
	"net/http"
	"model"
	"time"
)

type logoutController struct {}
func (this *logoutController) get(w http.ResponseWriter, r *http.Request) {
	cookie, err :=r.Cookie("sessionId")
	if err == nil{
		sessId :=cookie.Value
		model.DeleteSession(sessId)
		cookie.HttpOnly = true
		cookie.Expires = time.Now()
		cookie.Path= "/"
		cookie.MaxAge= -1
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}