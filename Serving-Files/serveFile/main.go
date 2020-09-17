package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog/", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--image doesn't serve-->
	<img src="/toby.jpg">
	`)
}

func dogpic(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "toby.jpg")

}
