package main

import (
	"html/template"
	"os"
)

type person struct {
	Name string
	ID   int
}

type car struct {
	brand string
	model int
}

type ident struct {
	data      []person
	transport []car
}

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("*gohtml"))
}
func main() {
	a := person{
		Name: "Mahmoud",
		ID:   1234,
	}
	b := person{
		Name: "Ola",
		ID:   4567,
	}

	d := car
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}
	defer nf.Close()

}
