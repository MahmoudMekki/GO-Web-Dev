package main

import (
	"html/template"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/setmulti", setMulti)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}

func set(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-session")
	if err == http.ErrNoCookie {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		http.SetCookie(res, c)
	}
	http.Redirect(res, req, "/index", http.StatusSeeOther)
}

func setMulti(res http.ResponseWriter, req *http.Request) {
	s := appending(res, req)

	c, err := req.Cookie("my-session")
	if err == http.ErrNoCookie {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
	} else {
		c = &http.Cookie{
			Name:  "my-session",
			Value: s,
		}
		http.SetCookie(res, c)
		http.Redirect(res, req, "/read", http.StatusSeeOther)
	}

}

func read(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-session")
	if err == http.ErrNoCookie {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
	} else {
		sb := strings.Split(c.Value, "|")
		tpl.ExecuteTemplate(res, "index.gohtml", sb)
	}

}

func appending(res http.ResponseWriter, req *http.Request) string {
	c, _ := req.Cookie("my-session")
	s := c.Value

	p1 := "Mahmoud"
	p2 := "Hamdi"
	p3 := "Mekki"

	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}

	return s
}
