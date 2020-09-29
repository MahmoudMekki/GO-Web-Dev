package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(res http.ResponseWriter, req *http.Request) user {

	//get cookie
	c, err := req.Cookie("my-session")
	if err == http.ErrNoCookie {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "my-session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(res, c)

	// if the user is already exists
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func alreadLogedin(res http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("my-session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]

	//refresh the session
	c.MaxAge = sessionLength
	http.SetCookie(res, c)
	return ok
}

func showSessions() {
	fmt.Println("***************")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}

func cleanSessions() {
	fmt.Println("Before deleting sessions")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * time.Duration(sessionLength)) {
			delete(dbSessions, k)
		}
	}
	dbSessionCleaned = time.Now()
	fmt.Println("After deleting sessions")
	showSessions()
}
