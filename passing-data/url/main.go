package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon/", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	v := req.FormValue("m")
	io.WriteString(res, "Hello "+v)
}
