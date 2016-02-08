package controler
//
//import (
//	"net/http"
//	"html/template"
//	topic "model"
//	"strings"
//	"fmt"
//)
//
////DONE
////func HomeHandler(w http.ResponseWriter, r *http.Request){
////	//t, _ := template.ParseFiles("templates/_threads.html", "templates/header.gtpl", "templates/footer.gtpl")
////	var thr_list topic.List_ThrList
////	thr_list.GetThreadsByCategory()
//////	t.ExecuteTemplate(w,"header",nil)
//////	t.ExecuteTemplate(w,"footer",nil)
//////	t.ExecuteTemplate(w,"home", thr_list)
////}
//
////this is controller
////func  ViewHandler(w http.ResponseWriter, r *http.Request)  {
////	//gets the current thread id from url and gets the info from db
////	var tl topic.Thread
////	tl.Id=r.URL.Path[len("/view/"):]
////	if tl.Id==""{
////	http.Redirect(w, r, "/", http.StatusFound)
////	} else {
////	tl.GetThreadData()
////	tl.GetReplies()
////	//passes the data to view
////	t, _ := template.ParseFiles("templates/view.gtpl","templates/header.gtpl", "templates/footer.gtpl")
////	t.ExecuteTemplate(w,"thread",tl)
////	//		t.ExecuteTemplate(w,"header",nil)
////	//		t.ExecuteTemplate(w,"footer",nil)
////	}
////}
////var tempVal topic.Main
////
////func ThreadCreateUpdateHandler (w http.ResponseWriter, r *http.Request){
////	t, _ := template.ParseFiles("templates/header.gtpl", "templates/footer.gtpl", "templates/_editMain.html")
////	//var value topic.Thread
////	res:=strings.Split(r.URL.Path,"/")
////	if res[4]!="new"{
////		tempVal.Thr_id=res[4]
////		tempVal.GetMainData()
////	}
////	t.ExecuteTemplate(w,"header",nil)
////	t.ExecuteTemplate(w,"footer",nil)
////	t.ExecuteTemplate(w,"editMain",tempVal)
////	}
////	func TopicSaveHandler(w http.ResponseWriter, r *http.Request){
////	tempVal.UnmarshalHTTP(r)
////	var a bool
////	if a=tempVal.Validate(); a==true {
////	tempVal.Save()
////	tempVal=topic.Main{}
////	http.Redirect(w, r, "/", http.StatusFound)
////	} else if a==false {
////	http.Redirect(w, r, "/edit/topic/thread/", http.StatusFound)
////	}
////}
//
//var tempValReply topic.Reply
//
//func ReplyCreateUpdateHandler(w http.ResponseWriter, r *http.Request){
//t, _ := template.ParseFiles("templates/header.gtpl", "templates/footer.gtpl", "templates/_editReply.html")
//res:=strings.Split(r.URL.Path,"/")
//fmt.Println(len(res))
//if len(res)==6{
//if res[5]=="new"{
//fmt.Println("teste")
//tempValReply.Thread_id=res[4]
//}
//} else{
//fmt.Println("testeererere")
//tempValReply.Id=res[4]
//tempValReply.GetSingleReply()
//}
//t.ExecuteTemplate(w,"header",nil)
//t.ExecuteTemplate(w,"footer",nil)
//t.ExecuteTemplate(w,"reply",tempValReply)
//fmt.Println("rr", tempValReply)
//}
//
//func ReplySaveHandler(w http.ResponseWriter, r *http.Request){
//tempValReply.UnmarshalHTTP(r)
//fmt.Println("nn", tempValReply)
//var a bool
//if a=tempValReply.Validate(); a==true {
//tempValReply.Save()
//tempValReply=topic.Reply{}
////fmt.Println("aaa", tempValReply)
//http.Redirect(w, r, "/", http.StatusFound)
//} else if a==false {
//http.Redirect(w, r, "/edit/topic/reply/", http.StatusFound)
//}}