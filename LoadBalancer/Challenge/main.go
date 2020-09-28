package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "admin:123456789@tcp(database-1.chqdh8gfrj4n.us-east-2.rds.amazonaws.com:3306)/RDS1?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/amigos", amigos)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":80", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello from AWS!!")
}

func instance(res http.ResponseWriter, req *http.Request) {
	s := getinstance()
	io.WriteString(res, s)
}

func amigos(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM Amigos;`)
	check(err)
	defer rows.Close()

	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(res, s)
}

func ping(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "OK")
}

func getinstance() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		log.Fatalln(err)
	}

	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	return string(bs)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
