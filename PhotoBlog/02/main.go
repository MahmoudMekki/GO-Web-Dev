package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}
func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	id, _ := uuid.NewV4()
	http.SetCookie(res, &http.Cookie{
		Name:  "my-session",
		Value: id.String(),
	})

	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-session")
	if err == http.ErrNoCookie {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(res, "index.gohtml", c.Value)
}
