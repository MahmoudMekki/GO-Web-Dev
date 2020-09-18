package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

type data struct {
	File string
}

func init() {
	tpl = template.Must(template.ParseGlob("*gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {

		f, h, err := req.FormFile("file")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		s = string(bs)

		des, err := os.Create(filepath.Join("./users/", h.Filename))

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		defer des.Close()
		_, err = des.Write(bs)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
	err := tpl.ExecuteTemplate(res, "index.gohtml", data{s})
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
