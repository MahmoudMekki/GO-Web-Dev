package main

import (
	"html/template"
	"io"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		http.Error(res, "File not found", 404)
	}
	tpl.ExecuteTemplate(res, "dog.gohtml", nil)
}
func dogpic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "toby.jpg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dogpic/", dogpic)
	http.ListenAndServe(":8080", nil)
}
