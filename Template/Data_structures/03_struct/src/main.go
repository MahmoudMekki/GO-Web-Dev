package main

import (
	"os"
	"text/template"
)

var temp *template.Template

type person struct {
	Name string
	ID   int
}

func init() {
	temp = template.Must(template.ParseGlob("*gohtml"))
}

func main() {
	mekki := person{
		Name: "Mahmoud Hamdi Mekki",
		ID:   25002,
	}

	nf, err := os.Create("index.html")
	defer nf.Close()
	err = temp.Execute(nf, mekki)
	if err != nil {
		os.Exit(1)
	}
}
