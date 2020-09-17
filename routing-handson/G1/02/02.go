package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func me(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("temp.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(res, "temp.gohtml", "Mahmoud Hamdi Mekki")
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Doggy Doggy Doggy")
}

func cat(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Kitty Kitty Kitty")
}

func main() {
	http.HandleFunc("/", cat)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/dog/", dog)

	http.ListenAndServe(":8080", nil)
}
