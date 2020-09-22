package main

import (
	"net/http"
	"strings"
	"text/template"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
}

type userin struct {
	UserName string
	Password []byte
}

var dbSession = map[string]string{}
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
	p1, _ := bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.MinCost)
	dbUsers["mahmoudmekki3@gmail.com"] = user{"mahmoudmekki3@gmail.com", p1, "mahmoud", "mekki"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":80", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.gohtml", u)
}

func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "bar.gohtml", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("last")

		if _, ok := dbUsers[un]; ok {
			http.Error(res, "Username is already taken!", http.StatusForbidden)
			return
		}
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "Session",
			Value: sID.String(),
		}

		http.SetCookie(res, c)
		dbSession[c.Value] = un
		bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)

		u := user{un, bs, f, l}
		dbUsers[un] = u

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var u userin
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		username = strings.ToLower(username)
		password := req.FormValue("password")
		bs := []byte(password)
		u = userin{username, bs}

		if _, ok := dbUsers[u.UserName]; ok {
			if passCheck(dbUsers[username].Password, bs) {
				sID, _ := uuid.NewV4()
				c := &http.Cookie{
					Name:  "Session",
					Value: sID.String(),
				}
				http.SetCookie(res, c)
				dbSession[c.Value] = username
				http.Redirect(res, req, "/bar", http.StatusSeeOther)
				return
			}
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}

		http.Redirect(res, req, "/signup", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(res, "login.gohtml", nil)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("Session")
	delete(dbSession, c.Value)
	c = &http.Cookie{
		Name:   "Session",
		MaxAge: -1,
	}
	http.SetCookie(res, c)
	http.Redirect(res, req, "/login", http.StatusSeeOther)
	return
}
