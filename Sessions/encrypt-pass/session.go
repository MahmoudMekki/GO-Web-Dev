package main

import "net/http"

func getUser(res http.ResponseWriter, req *http.Request) user {
	var u user
	c, err := req.Cookie("Session")
	if err != nil {
		return u
	}
	if un, ok := dbSession[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("Session")
	if err != nil {
		return false
	}
	un := dbSession[c.Value]
	_, ok := dbUsers[un]
	return ok
}
