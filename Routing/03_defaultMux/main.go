package main

import (
	"io"
	"net/http"
)

type hCat int
type hDog int

func (handler hCat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Cat CAT cAt")
}

func (handler hDog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog DOG dOg")
}

func main() {
	var cat hCat
	var dog hDog

	http.Handle("/cat/", cat)
	http.Handle("/dog/", dog)

	http.ListenAndServe(":8080", nil) // uses the default mux
}
