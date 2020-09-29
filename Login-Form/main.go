package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string // UserName
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbSessionCleaned time.Time

const sessionLength int = 30

func init() {
	tpl = template.Must(template.ParseGlob("templates/*html"))
	dbSessionCleaned = time.Now()
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	// added new route
	http.HandleFunc("/checkUserName", checkUserName)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func index(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	showSessions()
	tpl.ExecuteTemplate(res, "index.html", u)
}
func bar(res http.ResponseWriter, req *http.Request) {
	u := getUser(res, req)
	if !alreadLogedin(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(res, "You must be 007 agent to get to the bar!", http.StatusForbidden)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(res, "bar.html", u)
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadLogedin(res, req) {
		http.Redirect(res, req, "/", http.StatusNotAcceptable)
		return
	}
	var u user

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pass := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		if _, ok := dbUsers[un]; ok {
			http.Error(res, "The username is already taken!", http.StatusForbidden)
			return
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{un, time.Now()}

		bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Server internal Error", http.StatusInternalServerError)
		}
		u = user{un, bs, f, l, r}
		dbUsers[un] = u
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	showSessions()
	tpl.ExecuteTemplate(res, "signup.html", u)
}

func login(res http.ResponseWriter, req *http.Request) {
	if alreadLogedin(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		u, ok := dbUsers[un]
		if !ok {
			http.Error(res, "User name and/or password dont match", http.StatusForbidden)
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(res, "User name and/or password dont match", http.StatusForbidden)
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(res, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "login.html", u)
}

func logout(res http.ResponseWriter, req *http.Request) {
	if !alreadLogedin(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("my-session")
	delete(dbSessions, c.Value)
	c = &http.Cookie{
		Name:   "my-session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, c)
	if time.Now().Sub(dbSessionCleaned) > (time.Second * time.Duration(sessionLength)) {
		go cleanSessions()
	}
	http.Redirect(res, req, "/login", http.StatusSeeOther)
}

func checkUserName(res http.ResponseWriter, req *http.Request) {
	sampleUsers := map[string]bool{
		"test@example.com": true,
		"jame@bond.com":    true,
		"moneyp@uk.gov":    true,
	}
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	fmt.Println("USERNAME: ", sbs)

	fmt.Fprint(res, sampleUsers[sbs])
}
