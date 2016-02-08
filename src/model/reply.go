package model
import (
	"general"

	"net/http"
	"model/validation"

	"time"
	"strconv"
	"fmt"
)

type Reply struct{
	Id string
	Post
	Thread_id string
	when
	userData
}

//controller
func (R *Reply) UnmarshalHTTP(r *http.Request) (err error) {
	R.Content= r.FormValue("body")
	R.Name= r.FormValue("title")
	R.Thread_id=r.FormValue("thr_id")
	R.Id=r.FormValue("reply_id")
	//fmt.Println(R.thread_id)
	return
}
func (R *Reply) Validate() (isValid bool, msg string){
	isValid, msg = validation.RequireNotEmpty(R.Content)
	if isValid==false{
		goto ReturnPart
	}
	isValid, msg = validation.RequireNotEmpty(R.Name)
	if isValid==false{
		goto ReturnPart
	}
	ReturnPart: return
}

func (R *Reply) Update(){
	//Updates edited reply to database
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("UPDATE reply SET title=?,content=? WHERE id=?")
	general.CheckErr(err)
	_, err = stmt.Exec(R.Name, R.Content, R.Id)
	general.CheckErr(err)
}

func (R *Reply) Delete() (err error){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("delete from reply where id=?")
	general.CheckErr(err)
	_, err = stmt.Exec(R.Id)
	general.CheckErr(err)
	return
}
func deleteAllRepliesByThread(thr_id string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("delete from reply where thread_id=?")
	general.CheckErr(err)
	_, err = stmt.Exec(thr_id)
	general.CheckErr(err)
}
func (R *Reply) Create() (id int64){
	//Saves new reply to database
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("INSERT reply SET title=?,content=?, thread_id=?, user_id=?")
	general.CheckErr(err)
	res, err := stmt.Exec(R.Name, R.Content, R.Thread_id, R.UserID)
	id, err = res.LastInsertId()
	//fmt.Println(Id, "sgsdg")
	general.CheckErr(err)
	return id
}
func (R *Reply) Save(){
	//checks to see if create or update request is needed
	var id string
	if R.Id=="" {
		ID := R.Create()
		id = strconv.Itoa(int(ID))
	} else {
		R.Update()
		id=R.Id
	}
	time.Sleep(time.Second * 3)
	fmt.Println(id, "hey id")
	updateThreadTimestamp(id)
}

func (R *Reply) GetSingleReply(){
	//opens and closes connection
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	rows, err := db.Query(`SELECT reply.title, reply.content, reply.user_id, user.username
	FROM reply
	INNER JOIN user
	ON reply.user_id=user.Id
	WHERE reply.id=?`,R.Id)
	general.CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&R.Name, &R.Content, &R.UserID, &R.Username)
		general.CheckErr(err)
	}
}

func updateThreadTimestamp(id string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	rows, err := db.Query(`SELECT reply.updated, reply.thread_id, reply.user_id
	FROM reply
	WHERE id=?`,id)
	general.CheckErr(err)
	var upd, thrID, userID string
	fmt.Println(upd, thrID, "sgsdg")
	for rows.Next() {
		err = rows.Scan(&upd, &thrID, &userID)
		general.CheckErr(err)
	}
	fmt.Println(upd, thrID, userID, "fgfgfgf")
	stmt, err := db.Prepare("UPDATE thread SET updated=?, user_id=? WHERE id=?")
	general.CheckErr(err)
	_, err = stmt.Exec(upd, userID, thrID)
	general.CheckErr(err)
}
func (tl *Thread) GetReplies(){
	//opens and closes connection
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	rows, err := db.Query(`SELECT reply.id, reply.title, reply.content,reply.thread_id, user.username, reply.user_id, reply.updated
	FROM reply
	INNER JOIN user
	ON reply.user_id=user.Id
	WHERE reply.thread_id=?`,tl.Id)
	general.CheckErr(err)
	for rows.Next() {
		var r Reply
		err = rows.Scan(&r.Id, &r.Name, &r.Content,&r.Thread_id, &r.Username, &r.UserID, &r.Updated)
		general.CheckErr(err)
		tl.Reply=append(tl.Reply, r)
	}
}
