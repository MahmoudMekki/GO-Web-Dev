package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

var db *sql.DB
var tpl *template.Template

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

	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/books", showBooks)
	http.HandleFunc("/books/show", showBook)
	http.HandleFunc("/books/create", createForm)
	http.HandleFunc("/books/create/process", create)
	http.HandleFunc("/books/update", updateForm)
	http.HandleFunc("/books/update/process", update)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/books", http.StatusSeeOther)
	return
}

func showBooks(res http.ResponseWriter, req *http.Request) {
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
		_ = rw.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		bks = append(bks, bk)
	}
	tpl.ExecuteTemplate(res, "books.gohtml", bks)

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

	err := rw.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)

	switch {
	case err == sql.ErrNoRows:
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	case err != nil:
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(res, "show.gohtml", bk)
}

func createForm(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "create.gohtml", nil)
}

func create(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		isbn := req.FormValue("isbn")
		title := req.FormValue("title")
		author := req.FormValue("author")
		p := req.FormValue("price")
		bk := book{}
		if isbn == "" || title == "" || author == "" || p == "" {
			http.Error(res, "Bad request", http.StatusBadRequest)
			return
		}
		pricef64, err := strconv.ParseFloat(p, 32)
		if err != nil {
			http.Error(res, "Bad price!", http.StatusBadRequest)
			return
		}
		price := float32(pricef64)
		bk = book{isbn, title, author, price}
		_, err = db.Exec("insert into books (isbn,title,author,price) values ($1,$2,$3,$4);", bk.Isbn, bk.Title, bk.Author, bk.Price)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.ExecuteTemplate(res, "created.gohtml", bk)
	} else {
		http.Redirect(res, req, "/create", http.StatusSeeOther)
		return
	}
}

func updateForm(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "update.gohtml", nil)
}

func update(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		isbn := req.FormValue("isbn")
		title := req.FormValue("title")
		author := req.FormValue("author")
		p := req.FormValue("price")
		bk := book{}
		if isbn == "" || title == "" || author == "" || p == "" {
			http.Error(res, "Bad request", http.StatusBadRequest)
			return
		}
		pricef64, err := strconv.ParseFloat(p, 32)
		if err != nil {
			http.Error(res, "Bad price!", http.StatusBadRequest)
			return
		}
		price := float32(pricef64)
		bk = book{isbn, title, author, price}
		_, err = db.Exec("update books set isbn=$1,title=$2,author=$3,price=$4 where isbn=$1;", bk.Isbn, bk.Title, bk.Author, bk.Price)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.ExecuteTemplate(res, "updated.gohtml", bk)

	} else {
		http.Redirect(res, req, "/update", http.StatusSeeOther)
		return
	}
}
