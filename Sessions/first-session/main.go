package main

import (
	"net/http"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName  string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Session")
	if err == http.ErrNoCookie {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "Session",
			Value: sID.String(),
			//Secure: true,
		}
		http.SetCookie(res, c)
	}

	u := user{}
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{
			un,
			f,
			l,
		}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Session")
	if err == http.ErrNoCookie {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
	un := dbSessions[c.Value]
	u := dbUsers[un]

	tpl.ExecuteTemplate(res, "bar.gohtml", u)

}
