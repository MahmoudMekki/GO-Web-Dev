package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form) // req.Form and req.POSTForm are available after calling ParseForm
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog

	http.ListenAndServe(":8080", d)
}
