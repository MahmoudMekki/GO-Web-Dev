package main

import (
	"html/template"
	"os"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("tpl.gohtml"))
}

type items struct {
	Main string
	Side string
}

type meals struct {
	Breakfast []items
	Lunch     []items
	Dinner    []items
}
type restaurant struct {
	Name string
	Food meals
}

func main() {
	nf, err := os.Create("index.html")
	if err != nil {
		os.Exit(1)
	}
	r1 := []restaurant{
		restaurant{

			Name: "MM",
			Food: meals{
				Breakfast: []items{
					items{
						Main: "fool, tamya",
						Side: "Salad",
					},
				},
				Lunch: []items{
					items{
						Main: "Pane, rice",
						Side: "Cola, fries",
					},
				},
				Dinner: []items{
					items{
						Main: "Ommlette ,cheese",
						Side: "Juice",
					},
				},
			},
		},
	}

	err = temp.Execute(nf, r1)
	if err != nil {
		os.Exit(1)
	}
}
