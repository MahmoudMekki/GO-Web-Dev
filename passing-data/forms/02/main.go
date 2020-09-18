package main

import (
	"net/http"
	"text/template"
)

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*gohtml"))
}
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon/", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("sub") == "on"

	err := tpl.ExecuteTemplate(res, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(res, "Error", 505)
	}
}
