package main
import (
	"net/http"
	"regexp"
	"io/ioutil"
	"html/template"
	"os"
	"controler"
)

func main() {
	templates := getTemplates();
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	controler.Register(templates)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
//var store = sessions.NewCookieStore([]byte("something-very-secret"))
//
//func MyHandler(w http.ResponseWriter, r *http.Request) {
//	// Get a session. We're ignoring the error resulted from decoding an
//	// existing session: Get() always returns a session, even if empty.
//	session, err := store.Get(r, "session-name")
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//
//	// Set some session values.
//	session.Values["foo"] = "bar"
//	session.Values[42] = 43
//	// Save it before we write to the response/return from the handler.
//	session.Save(r, w)
//}

func getTemplates() *template.Template{
	result := template.New("templates")
	var validTemplateName=regexp.MustCompile("^(_?[a-zA-Z0-9]+).html$") //regex for template name
	basePath := "templates" //where to look up for templates
	filesRaw,_:=ioutil.ReadDir(basePath) //get all files in the directory
	var filesValid []os.FileInfo
	for _, f := range filesRaw {
		fileNameRaw:=f.Name()
		if m := validTemplateName.FindStringSubmatch(fileNameRaw); m!=nil  {
			filesValid = append(filesValid,f)
		}
	}
	templatePaths := new([]string)
	for _, f := range filesValid {
		fileName :=f.Name()
		*templatePaths = append(*templatePaths,basePath + "/" + fileName )
	}
	result.ParseFiles(*templatePaths...)

	return result
}




