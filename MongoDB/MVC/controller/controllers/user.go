package controller

import (
	"encoding/json"
	"io"
	models "model"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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

func (uc UserController) GetUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Mahmoud Mekki",
		Gender: "Male",
		Age:    26,
		ID:     p.ByName("id"),
	}
	res.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(res).Encode(u)
}

func (uc UserController) Create(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(req.Body).Decode(&u)

	u.ID = "007"
	json.NewEncoder(res).Encode(u)
}

func (uc UserController) Delete(res http.ResponseWriter, req *http.Request, p httprouter.Params) {

	io.WriteString(res, "Write any code to delete user !")
}
