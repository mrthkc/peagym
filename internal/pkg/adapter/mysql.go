package adapter

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewClient(host string, port string, dbName string, uName string, pass string) (db *sql.DB) {
	db, err := sql.Open("mysql", uName+":"+pass+"@tcp("+host+":"+port+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
