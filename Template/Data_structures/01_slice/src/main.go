package main

import (
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("*gohtml"))
}

func main() {
	sname := []string{
		"Mahmoud",
		"Hamdi",
		"Mekki",
	}
	nf, err := os.Create("index.html")
	defer nf.Close()
	err = temp.ExecuteTemplate(nf, "tmplt.gohtml", sname)
	if err != nil {
		os.Exit(1)
	}
}
