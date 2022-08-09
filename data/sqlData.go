package data

import (
	"database/sql"
)

var Db *sql.DB
var err error

func OpenDb() {
	Db, err = sql.Open("mysql", "root:test123@tcp(172.19.0.2:3306)/teams_test")
	if err != nil {
		panic(err.Error())
	}
}
