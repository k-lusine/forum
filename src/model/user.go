package model
import (
	"net/http"
	"general"
	"crypto/sha256"
	"encoding/hex"
	"model/validation"
	//"fmt"
	"time"
)

type User struct{
	Id string
	Email string
	Username string
	Status string
	Status_id string
	passwordhashed [32]byte
	Password string
	Name string
	Surname string
}
type Session struct {
	Id string
	UserId string
	SessionId string
}
func (U *User) GetCurrentUserData() {
	U.Id = getCurrentUserId()
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	rows, err := db.Query(`SELECT user.username, user.email, user.name, user.surname, userstatus.status
	FROM user
	INNER JOIN userstatus
	ON user.userstatus_id=userstatus.Id
	WHERE user.Id=?`, U.Id)
	for rows.Next() {
		err = rows.Scan(&U.Username, &U.Email, &U.Name, &U.Surname, &U.Status)
		general.CheckErr(err)
	}
}
//controller
func (U *User) UnmarshalHTTP(r *http.Request) (err error) {
	U.Username= r.FormValue("username")
	U.Email = r.FormValue("email")
	U.Password = r.FormValue("password")
	U.Status = r.FormValue("status")
	U.Name = r.FormValue("name")
	U.Surname = r.FormValue("surname")
	//fmt.Println(R.thread_id)
	return
}

func (U *User) Validate() (isValid bool, msg string){
	isValid, msg = validation.RequireAlphaNumeric(U.Username)
	if isValid==false{
		msg = "username should have minumun 8 letters, should be alphanumeric and start with an english letter"
		goto ReturnPart
	}
	isValid, msg = validation.RequireFirstLetter(U.Username)
	if isValid==false{
		msg = "username should have minumun 8 letters, should be alphanumeric and start with an english letter"
		goto ReturnPart
	}
	isValid, msg = validation.RequireMinCount(U.Username,6)
	if isValid==false {
		msg = "username should have minumun 8 letters, should be alphanumeric and start with an english letter"
		goto ReturnPart
	}
	isValid, msg = validation.RequireEmail(U.Email)
	if isValid==false{
		goto ReturnPart
	}
	isValid, msg = validation.RequireMinCount(U.Password,11)
	if isValid==false{
		msg = "your password is not strong enough. Enter at least 11 characters, including 1 number, 1 symbol and 1 uppercase"
		goto ReturnPart
	}
	isValid, msg = validation.RequireStrongPassword(U.Password)
	if isValid==false{
		msg = "your password is not strong enough. Enter at least 11 characters, including 1 number, 1 symbol and 1 uppercase"
		goto ReturnPart
	}
	isValid, msg = validation.RequireNotEmpty(U.Name)
	if isValid==false{
		goto ReturnPart
	}
	isValid, msg = validation.RequireNotEmpty(U.Surname)
	if isValid==false{
		goto ReturnPart
	}
	ReturnPart: return

}

func (U *User) requireUniqueEmail() (passed bool, errorMessage string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	rows, err := db.Query(`SELECT Id FROM user WHERE email="?"`, U.Email)
	general.CheckErr(err)
	var id string
	var count int
	for rows.Next() {
		err = rows.Scan(&id)
		general.CheckErr(err)
		count++
	}
	if count ==0{
		passed=true
	} else {
		errorMessage = "This is duplicate entry. Email is already registered"
	}
	return
}
func (U *User) requireUniqueUsername() (passed bool, errorMessage string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	rows, err := db.Query(`SELECT Id FROM user WHERE username="?"`, U.Username )
	general.CheckErr(err)
	var id string
	var count int
	for rows.Next() {
		err = rows.Scan(&id)
		general.CheckErr(err)
		count++
	}
	if count ==0{
		passed=true
	} else {
		errorMessage = "This is duplicate entry. Username is already registered"
	}
	return
}
func (U *User) Update(){
	U.Id = getCurrentUserId()
	//U.status_id = getUserStatusId(U.Status)
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("UPDATE user SET username=?,email=?, password=?, name=?, surname=? WHERE id=?")
	general.CheckErr(err)
	_, err = stmt.Exec(U.Username, U.Email, hex.EncodeToString(U.passwordhashed[:]),U.Name, U.Surname, U.Id)
	general.CheckErr(err)
}
func (U *User) Create(){
	//Saves new reply to database
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("INSERT user SET username=?,email=?, password=?, name=?, surname=?, userstatus_id=?")
	general.CheckErr(err)
	_, err = stmt.Exec(U.Username, U.Email, hex.EncodeToString(U.passwordhashed[:]), U.Name, U.Surname, "2")
	general.CheckErr(err)
}
func (U *User) Save(){
	U.Id = getCurrentUserId()
	U.passwordhashed = sha256.Sum256([]byte(U.Password))
	//checks to see if create or update request is needed
	if U.Id=="" {
		U.Create()
	} else {
		U.Update()
	}
}

func getUserStatusId(status string) (id string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT Id FROM userstatus WHERE status=?",status)
	for rows.Next() {
		err = rows.Scan(&id)
		general.CheckErr(err)
	}
	return
}

func getCurrentUserId() string{
	id := ""
	//WRITE THIS FUNCTION
	return id
}

func ConfirmLogin(username string, password string)(passed bool, errorMessage string, id string){
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	var passwordhashed [32]byte
	passwordhashed	= sha256.Sum256([]byte(password))
	rows, err := db.Query(`SELECT Id FROM user WHERE username="` + username + `" AND password="` + hex.EncodeToString(passwordhashed[:]) + `"`)
	general.CheckErr(err)
	var count int
	for rows.Next() {
		err = rows.Scan(&id)
		general.CheckErr(err)
		count++
	}
	if count > 0{
		passed=true
	} else {
		errorMessage = "Sorry, you can't login, smth. is wrong with your username or password"
	}
	return
}


func (U *User) CreateSession() Session {
	result := Session{}
	result.UserId = U.Id
	sessionId := sha256.Sum256([]byte(U.Username + time.Now().Format("12:00:00")))
	result.SessionId = hex.EncodeToString(sessionId[:])

	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("INSERT session SET user_id=?,sessionID=?")
	general.CheckErr(err)
	_, err = stmt.Exec(result.UserId, result.SessionId)
	general.CheckErr(err)
	return result
}

func GetMemberBySessionId(sessionId string) (result User, err error) {
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	rows, err := db.Query(`SELECT user.Id, user.username, user.name, user.surname, user.userstatus_id
	FROM session
	INNER JOIN user
	ON session.user_id=user.Id
	WHERE session.sessionID=?`, sessionId)
	for rows.Next() {
		err = rows.Scan(&result.Id, &result.Username, &result.Name, &result.Surname, &result.Status_id)
		general.CheckErr(err)
	}
	return
}
func DeleteSession(sessionId string) {
	db, err :=connectDB()
	general.CheckErr(err)
	defer db.Close()
	stmt, err := db.Prepare("delete from session where sessionID=?")
	general.CheckErr(err)
	_, err = stmt.Exec(sessionId)
	general.CheckErr(err)
}