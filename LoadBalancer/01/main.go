package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)

}
func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello from AWS!")
}

func ping(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "OK")
}
