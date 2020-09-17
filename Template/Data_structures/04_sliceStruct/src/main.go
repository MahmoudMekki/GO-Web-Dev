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

	ola := person{
		Name: "Ola Hamdi Mekki",
		ID:   25004,
	}
	farah := person{
		Name: "Farah hamdi Mekki",
		ID:   25006,
	}
	abdo := person{
		Name: "Abdo Hamdi Mekki",
		ID:   25008,
	}
	names := []person{mekki, ola, abdo, farah}

	nf, err := os.Create("index.html")
	defer nf.Close()
	err = temp.Execute(nf, names)
	if err != nil {
		os.Exit(1)
	}
}
