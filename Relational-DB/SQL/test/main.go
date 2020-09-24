package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:maxes2312733@tcp(database-1.chqdh8gfrj4n.us-east-2.rds.amazonaws.com:3306)/Mekki?charset=utf8")
	check(err)

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(res, "Successed")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
