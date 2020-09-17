package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/dogs/", dogs)
	http.ListenAndServe(":8080", nil)

}

func dogs(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "File not found", 404)
	}
}
