package model
import(

	"general"
)
type Thread struct{
	Id string
	Title string
	Main
	Reply []Reply
	when
	userData
}
type ThrList struct {
	Tl []Thread
	CatName string
}

type List_ThrList struct{
	List []ThrList
}


func getCategory() (catList map[string]string) {
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	//query to get list of categories` max 4
	rows, err := db.Query(`SELECT id, name
	FROM category
	ORDER BY RAND()
	LIMIT 3`)
	general.CheckErr(err)
	//var catList [string]string
	catList=make(map[string]string, 3)
	for rows.Next() {
		var cat_id, cat_name string
		err = rows.Scan(&cat_id, &cat_name)
		general.CheckErr(err)
		//catList[cat_id] =
		catList[cat_id]=cat_name
	}
	return
}
func (tl *List_ThrList) GetThreadsByCategory(){
	db, err :=connectDB()
	general.CheckErr(err)
	catList:=getCategory()
	//fmt.Println(catList)
	for cat_id, cat_name:= range catList{
		rows, err := db.Query(`SELECT thread.title, thread.id, thread.updated, thread.user_id, user.username
		FROM thread
		INNER JOIN user
		ON thread.user_id=user.Id
		WHERE category_id=?
		ORDER BY updated
		LIMIT 8`, cat_id)
		general.CheckErr(err)
		var tlSingle ThrList
		tlSingle.CatName=cat_name
		for rows.Next() {
			var newThread Thread
			err = rows.Scan(&newThread.Title, &newThread.Id, &newThread.Updated, &newThread.UserID, &newThread.Username)
			general.CheckErr(err)
			tlSingle.Tl=append(tlSingle.Tl, newThread)
		}
		tl.List=append(tl.List,tlSingle)
	}
	db.Close()
}
//this is model
//func (tl *Thread) GetThreadData() {
//	//opens and closes connection
//	db, err :=connectDB()
//	general.CheckErr(err)
//	defer db.Close()
//	//query to get data from db
//	rows, err := db.Query(`SELECT  category.name
//	FROM thread
//	INNER JOIN category
//	ON thread.category_id=category.id
//	WHERE thread.id=?`,tl.Id)
//	general.CheckErr(err)
//	//saves values to thread
//	//tl.Id=id
//	for rows.Next() {
//		err = rows.Scan(&tl.Main.Category, &tl.Main.ID)
//		general.CheckErr(err)
//	}
//	tl.Main.GetMainData()
//}

func (T *Thread) Delete() (err error){
		db, err :=connectDB()
		general.CheckErr(err)
		defer db.Close()
		deleteAllRepliesByThread(T.Id)
		var mainID string
		rows, err := db.Query(`SELECT main_id
		FROM thread
		WHERE id=?`,T.Id)
		general.CheckErr(err)
		for rows.Next() {
			err = rows.Scan(&mainID)
			general.CheckErr(err)
		}
		stmt, err := db.Prepare("delete from thread where id=?")
		general.CheckErr(err)
		_, err = stmt.Exec(T.Id)
		general.CheckErr(err)
		deleteMain(mainID)
	return
}
