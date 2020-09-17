package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

type h int

var tpl *template.Template

func (handler h) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		URL         *url.URL
		Submissions url.Values
		Header      http.Header
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var handler h
	http.ListenAndServe(":8080", handler)
}
