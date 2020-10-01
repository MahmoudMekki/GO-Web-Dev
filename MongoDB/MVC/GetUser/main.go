package main

import (
	"encoding/json"
	"io"
	"net/http"

	models "github.com/Models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getuser)
	http.ListenAndServe("localhost:8080", r)
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `
		<!DOCTYPE html>
		<html>
		<head>
		<meta charset = "UTF-8">
		<title>MEkki</title>
		</head>
		<body>
		<h1><a href = "/user/123456">GOTO locahost:8080/user/123456</a></h1>
		</body>
		</html>
	`
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, s)
}

func getuser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Mahmoud Mekki",
		Gender: "Male",
		Age:    26,
		ID:     p.ByName("id"),
	}
	res.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(res).Encode(u)
}
