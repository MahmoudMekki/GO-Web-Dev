package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.gohtml"))
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at foo: " + req.Method)
}

func bar(res http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at bar: " + req.Method)
	http.Redirect(res, req, "/barred", http.StatusSeeOther)
}

func barred(res http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at barred: " + req.Method)
	http.Redirect(res, req, "/favicon.ico", http.StatusSeeOther)
	tpl.Execute(res, nil)
}
