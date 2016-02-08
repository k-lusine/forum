package model
import (
	"general"
	"net/http"
	"model/validation"
)

//DATA STRUCTURE
type  Post struct{
	Name string
	Content string
}
type when struct {
	Created string
	Updated string
}
type userData struct {
	UserID string
	Username string
}
type Main struct {
	ID string
	Post
	Category string
	Thr_id string
	when
	userData
}


func (m *Main) GetMainData() {
	//opens and closes connection
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	//query to get data from db
	rows, err := db.Query(`SELECT category.name, main.title, main.content, main.id, main.user_id, user.username, main.created, main.updated
	FROM thread
	INNER JOIN category
	ON thread.category_id=category.id
	INNER JOIN main
	ON thread.main_id=main.id
    INNER JOIN user
    ON main.user_id=user.Id
	WHERE thread.id=?`, m.Thr_id)
	general.CheckErr(err)
	//saves values to thread
	//tl.Id=id
	for rows.Next() {
		err = rows.Scan(&m.Category,&m.Name, &m.Content, &m.ID, &m.UserID, &m.Username, &m.Created, &m.Updated)
		general.CheckErr(err)
	}
}

func (M *Main) Create(cat_id string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	//insert new entry into the table main
	stmt, err := db.Prepare("INSERT main SET title=?,content=?, user_id=?")
	general.CheckErr(err)
	res, err := stmt.Exec(M.Name, M.Content, M.UserID)
	general.CheckErr(err)
	main_id, err := res.LastInsertId()
	general.CheckErr(err)
	//insert new entry into the table thread
	stmt, err = db.Prepare("INSERT thread SET title=?, category_id=?, main_id=?, user_id=?")
	general.CheckErr(err)
	res, err = stmt.Exec(M.Name, cat_id, main_id, M.UserID)
	general.CheckErr(err)
}

func (M *Main) Update(cat_id string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("UPDATE main SET title=?,content=? WHERE id=?")
	general.CheckErr(err)
	res, err := stmt.Exec(M.Name, M.Content, M.ID)
	general.CheckErr(err)
	_, err = res.RowsAffected()
	general.CheckErr(err)
	stmt, err = db.Prepare("UPDATE thread SET title=?, category_id=? WHERE main_id=?")
	general.CheckErr(err)
	res, err = stmt.Exec(M.Name, cat_id, M.ID)
	general.CheckErr(err)
}
func (M *Main) Save(){
	//Saves new main post and creates new corresponding thread in the database
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	//gets the category Id based on the category name
	cat_result, err := db.Query("SELECT id FROM category WHERE `name`='"+M.Category+"'")
	var cat_id string
	for cat_result.Next(){
		err = cat_result.Scan(&cat_id)
		general.CheckErr(err)
	}
	general.CheckErr(err)
	//checks to see if create or update request is needed
	if M.ID=="" {
		M.Create(cat_id)
	} else {
		M.Update(cat_id)
	}
}

func  deleteMain(mainID string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("delete from main where id=?")
	general.CheckErr(err)
	_, err = stmt.Exec(mainID)
	general.CheckErr(err)
}
//controller
func (m *Main) UnmarshalHTTP(r *http.Request) (err error) {
	m.Content= r.FormValue("body")
	m.Name=r.FormValue("title")
	m.Category=r.FormValue("category")
	return
}

func (M *Main) Validate() (isValid bool, msg string){
	isValid, msg = validation.RequireNotEmpty(M.Content)
	if isValid==false{
		goto ReturnPart
	}
	isValid, msg = validation.RequireNotEmpty(M.Name)
	if isValid==false{
		goto ReturnPart
	}
	isValid, msg = validation.RequireNotEmpty(M.Category)
	if isValid==false{
		goto ReturnPart
	}
	ReturnPart: return
}
