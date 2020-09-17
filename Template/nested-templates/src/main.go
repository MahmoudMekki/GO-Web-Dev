package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var temp *template.Template

type person struct {
	Fname string
	Lname string
}

// create a FuncMap to register functions.
// "uc" is what the func will be called in the template
// "uc" is the ToUpper func from package strings
// "ft" is a func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"lc": strings.ToLower,
}

func init() {
	temp = template.Must(template.New("").Funcs(fm).ParseGlob("*gohtml"))
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}
	b := person{
		Fname: "Mahmoud",
		Lname: "Mekki",
	}

	g := person{
		Fname: "ola",
		Lname: "mekki",
	}

	names := []person{b, g}

	err = temp.ExecuteTemplate(nf, "tmplt.gohtml", names)
	if err != nil {
		fmt.Println(err)
	}
}
