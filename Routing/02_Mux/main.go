package main

import (
	"io"
	"net/http"
)

type hDog int
type hCat int

func (handler hDog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Doggy Doggy Doggy Doggy")
}

func (handler hCat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Kitty Kitty Kitty Kitty")
}

func main() {
	var cat hCat
	var dog hDog

	mux := http.NewServeMux()
	mux.Handle("/cat/", cat)
	mux.Handle("/dog/", dog)

	http.ListenAndServe(":8080", mux)
}
