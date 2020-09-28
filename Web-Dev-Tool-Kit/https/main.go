package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello tls!")
}
