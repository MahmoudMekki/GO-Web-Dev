package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/blah/", http.StripPrefix("/blah", http.FileServer(http.Dir("./resources"))))
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--image doesn't serve-->
	<img src="/resources/1.jpg">
	`)
}
