package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset =utf-8")
	fmt.Fprintln(res, `<h> <a href= "/set">Set cookie</a></h><br>`)
	fmt.Fprintln(res, `<h> <a href= "/read">Read cookie</a></h><br>`)
	fmt.Fprintln(res, `<h> <a href= "/expire">Delete cookie</a></h><br>`)

}

func set(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset =utf-8")
	http.SetCookie(res, &http.Cookie{
		Name:  "Session",
		Value: "some value",
	})
	fmt.Fprintln(res, `<h> <a href= "/read">Read cookie</a></h>`)
}

func read(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset =utf-8")
	c, err := req.Cookie("Session")
	if err == http.ErrNoCookie {
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintln(res, "The saved cookies are! :")
	fmt.Fprintf(res, "<h> cookie name is %s </h><br><h>cookie value is %s</h><br>", c.Name, c.Value)
	fmt.Fprintln(res, `<h> <a href= "/set">Set cookie</a></h><br>`)
	fmt.Fprintln(res, `<h> <a href= "/expire">Delete cookie</a></h><br>`)

}

func expire(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset =utf-8")

	c, err := req.Cookie("Session")
	if err == http.ErrNoCookie {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(res, c)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
