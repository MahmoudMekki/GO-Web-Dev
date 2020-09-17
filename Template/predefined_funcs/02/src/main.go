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

	scores := struct {
		Score1 int
		Score2 int
	}{
		2,
		2,
	}

	err = temp.Execute(nf, scores)
	if err != nil {
		fmt.Println(err)
	}
}
