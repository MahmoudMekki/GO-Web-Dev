package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type book struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgres://mahmoud:12345@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rw, err := db.Query("select * from books;")
	if err != nil {
		panic(err)
	}
	defer rw.Close()
	bks := make([]book, 0)
	for rw.Next() {
		bk := book{}
		_ = rw.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		bks = append(bks, bk)
	}
	fmt.Printf("%-20s %-20s %-20s %s\n", "ISBN", "Title", "Author", "Price")
	for _, v := range bks {
		fmt.Printf("%-20s %-20s %-20s %.2f\n", v.isbn, v.title, v.author, v.price)
	}

}
