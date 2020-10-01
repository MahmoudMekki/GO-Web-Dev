package main

import (
	controller "controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controller.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.Create)
	r.DELETE("/user/:id", uc.Delete)
	http.ListenAndServe("localhost:8080", r)
}
