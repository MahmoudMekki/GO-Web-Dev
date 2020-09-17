package main

import (
	"io"
	"net/http"
)

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Mahmoud Hamdi Mekki")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Doggy Doggy Doggy")
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Kitty Kitty Kitty")
}

func main() {
	http.HandleFunc("/me/", me)
	http.HandleFunc("/cat/", cat)
	http.HandleFunc("/dog/", dog)

	http.ListenAndServe(":8080", nil)
}
