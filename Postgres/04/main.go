package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type book struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://mahmoud:12345@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/books", index)
	http.HandleFunc("/books/show", showBook)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Now Alowed", http.StatusNotAcceptable)
		return
	}

	rw, err := db.Query("select * from books;")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rw.Close()
	bks := make([]book, 0)
	for rw.Next() {
		bk := book{}
		_ = rw.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		bks = append(bks, bk)
	}
	fmt.Fprintf(res, "%-20s %-20s %-20s %s\n", "ISBN", "Title", "Author", "Price")
	for _, v := range bks {
		fmt.Fprintf(res, "%-20s %-20s %-20s %.2f\n", v.isbn, v.title, v.author, v.price)
	}

}

func showBook(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Now Alowed", http.StatusNotAcceptable)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(res, "No isbn provided", http.StatusNoContent)
		return
	}
	rw := db.QueryRow("select * from books where isbn = $1;", isbn)
	bk := book{}

	err := rw.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)

	switch {
	case err == sql.ErrNoRows:
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(res, "%-20s %-20s %-20s %s\n", "ISBN", "Title", "Author", "Price")
	fmt.Fprintf(res, "%-20s %-20s %-20s %.2f\n", bk.isbn, bk.title, bk.author, bk.price)
}
