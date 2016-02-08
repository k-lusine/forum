package model

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"

)

func connectDB() (db *sql.DB,err error) {
	db, err = sql.Open("mysql", "klusine:123456@/forum")
	return
}