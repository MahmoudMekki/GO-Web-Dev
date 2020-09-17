package main

import (
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("tpl.gohtml"))
}

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

type state struct {
	Name   string
	Hotels []hotel
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}

	nw := []state{
		state{
			Name: "New York",
			Hotels: []hotel{
				hotel{"Eagle",
					"tayran st.",
					"downtown",
					122111,
					"Central",
				},
				hotel{"Mekki",
					"Khedr st.",
					"downtown",
					122111,
					"blah",
				},
			},
		},
	}
	err = temp.Execute(nf, nw)
	if err != nil {
		os.Exit(1)
	}
}
