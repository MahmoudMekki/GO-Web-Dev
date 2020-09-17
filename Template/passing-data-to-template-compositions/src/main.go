package main

import (
	"html/template"
	"os"
)

var temp *template.Template

type person struct {
	Name string
	Age  int
}

type zero struct {
	person
	Permission bool
}

func init() {
	temp = template.Must(temp.ParseGlob("*gohtml"))
}
func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}
	p := zero{
		person{
			Name: "Mahmoud",
			Age:  25,
		},
		false,
	}

	err = temp.Execute(nf, p)
}
