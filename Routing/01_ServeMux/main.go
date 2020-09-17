package main

import (
	"io"
	"net/http"
)

type h int

func (handler h) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var data string
	switch req.URL.Path {
	case "/dog":
		data = "Doggy Doggy Doggy Doggy"
		io.WriteString(w, data)

	case "/cat":
		data = "kitty kitty kitty kitty"
		io.WriteString(w, data)

	}

}

func main() {
	var handler h
	http.ListenAndServe(":8080", handler)
}
