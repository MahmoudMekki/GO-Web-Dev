package main

import (
	"io"
	"net/http"
)

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "CAt cAt Cat cat")
}

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "DOG dOg Dog dog")
}

func main() {
	http.Handle("/cat/", http.HandlerFunc(c))
	http.Handle("/dog/", http.HandlerFunc(d))

	http.ListenAndServe(":8080", nil)
}
