package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("public/pics"))))
	http.ListenAndServe(":8080", nil)

}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		http.Error(res, "File not found", 404)
	}
}
