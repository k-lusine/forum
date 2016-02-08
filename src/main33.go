package main
import (

	"os/exec"
	"net/http"
	//"controler"
"html/template"
//	"regexp"
//	"io/ioutil"
//	"os"
)



func main() {
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	//Register(templates)
	http.HandleFunc("/", Helloworld)
	http.HandleFunc("/action", Action)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
func Helloworld(w http.ResponseWriter, r *http.Request){
	t,_  := template.ParseFiles("src/test.html")
	t.Execute(w, nil)
}
func Action(w http.ResponseWriter, r *http.Request){
	t,_  := template.ParseFiles("src/action.html")

	path, err := exec.LookPath(`C:\Users\user\Desktop\ScreenToGif 1.4.2.exe`)
	var result string
	if err != nil {
		result = "installing fortune is in your future"
	} else{
		result = "fortune is available at %s\n" +  path
	}
	var a exec.Cmd
	a.Path = path
	a.Run()
	t.Execute(w, result)
}
