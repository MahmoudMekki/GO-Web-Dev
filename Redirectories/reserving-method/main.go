package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at foo: " + req.Method)
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at bar: " + req.Method)
	http.Redirect(res, req, "/", http.StatusTemporaryRedirect)

}
