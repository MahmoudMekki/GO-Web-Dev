package main

import (
	"io"
	"net/http"
)

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Cat CAT cAt")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog DOG dOg")
}

func main() {
	http.HandleFunc("/cat/", c)
	http.HandleFunc("/dog/", d)

	http.ListenAndServe(":8080", nil)
}
