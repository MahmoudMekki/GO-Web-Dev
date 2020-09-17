package main

import (
	"fmt"
	"net/http"
)

type h int

func (handler h) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Mekki-Key", "This is from Mahmoud Mekki 123456789")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h>Welcome to wonderland!</h>")
}

func main() {
	var handler h
	http.ListenAndServe(":8080", handler)
}
