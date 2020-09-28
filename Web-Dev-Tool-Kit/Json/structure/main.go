package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name     string
	Lname    string
	Nickname []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mash", msh)
	http.HandleFunc("/encode", enc)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
func foo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html")

	s := `
	<!DOCTYPE html>
	<html>
	<header>
	<meta charset = "en">
	<title>json</title>
	</header>
	<body>
	<p> hello to json </p>
	</p>
	</body>
	</html>
	`
	res.Write([]byte(s))
}

func msh(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	p1 := person{
		"Mahmoud",
		"Mekki",
		[]string{
			"Makook",
			"Maxes",
		},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	res.Write(json)

}

func enc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	p1 := person{
		"Mahmoud",
		"Mekki",
		[]string{
			"Makook",
			"Maxes",
		},
	}
	json.NewEncoder(res).Encode(p1)
}
