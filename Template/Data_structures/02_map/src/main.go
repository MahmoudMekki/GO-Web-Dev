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

	mnames := map[string]string{
		"fname": "Mahmoud",
		"mname": "Hamdi",
		"lname": "Mekki",
	}
	nf, err := os.Create("index.html")
	defer nf.Close()
	err = temp.ExecuteTemplate(nf, "tmplt.gohtml", mnames)

	if err != nil {
		os.Exit(1)
	}
}
