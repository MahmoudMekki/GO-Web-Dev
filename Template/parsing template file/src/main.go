package main

import (
	"os"
	"text/template"
)

func main() {
	temp, err := template.ParseFiles("tmplt.gohtml")
	if err != nil {
		os.Exit(1)
	}
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}
	defer nf.Close()
	err = temp.Execute(nf, temp)
	if err != nil {
		os.Exit(1)
	}
}
