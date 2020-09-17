package main

import (
	"fmt"
	"math"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.New("").Funcs(fm).ParseGlob("*gohtml"))
}

var fm = template.FuncMap{
	"fd":   double,
	"fsq":  square,
	"fsqr": squareRoot,
}

func double(x int) int {
	return x * 2
}
func square(x int) int {
	return x * x
}
func squareRoot(x int) float64 {
	return math.Sqrt(float64(x))
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}

	err = temp.ExecuteTemplate(nf, "tmplt.gohtml", 3)
	if err != nil {
		fmt.Println(err)
	}
}
