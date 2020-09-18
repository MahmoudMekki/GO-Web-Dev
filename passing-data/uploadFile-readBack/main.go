package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type data struct {
	File string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favion.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		f, _, err := req.FormFile("q")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

		s = string(bs)
	}

	err := tpl.ExecuteTemplate(res, "index.gohtml", data{s})
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

}
