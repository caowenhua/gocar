package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func SetupDb() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@/car?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

func CloseDb() {
	Db.Close()
}

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}
