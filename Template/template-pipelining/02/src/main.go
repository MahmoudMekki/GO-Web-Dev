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

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}
	xs := []string{
		"zero", " one",
		"two", "three", "four",
	}
	err = temp.Execute(nf, xs)
	if err != nil {
		fmt.Println(err)
	}
}
