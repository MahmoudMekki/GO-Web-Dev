package main

import (
	"fmt"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("tmplt.gohtml"))
}

type user struct {
	Name   string
	Status bool
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}
	u1 := user{
		Name:   "Admin",
		Status: true,
	}
	u2 := user{
		Name:   "Ola",
		Status: false,
	}
	u3 := user{
		Name:   "Farah",
		Status: false,
	}

	users := []user{u1, u2, u3}
	err = temp.Execute(nf, users)
	if err != nil {
		fmt.Println(err)
	}
}
