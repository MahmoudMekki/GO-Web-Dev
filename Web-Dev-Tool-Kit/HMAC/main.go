package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/auth", auth)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-serssion")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-session",
			Value: "",
		}
		http.SetCookie(res, c)
	}
	if req.Method == http.MethodPost {
		email := req.FormValue("email")
		h := getcode(email)
		value := email + "|" + h
		c.Value = value
	}
	http.SetCookie(res, c)
	tpl.ExecuteTemplate(res, "index.gohtml", c.Value)
}

func auth(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-session")
	if err == http.ErrNoCookie {
		http.Redirect(res, req, "/index", http.StatusSeeOther)
	}
	if c.Value == "" {
		http.Redirect(res, req, "/index", http.StatusSeeOther)
	}
	value := strings.Split(c.Value, "|")
	email := value[0]
	hash := value[1]
	check := getcode(email)
	if hash == check {
		tpl.ExecuteTemplate(res, "auth.gohtml", value)
		return
	}
	io.WriteString(res, "Falied not authorized!!")
}

func getcode(w string) string {
	h := hmac.New(sha256.New, []byte("my-key"))
	io.WriteString(h, w)
	return fmt.Sprintf("%x", h.Sum(nil))
}
