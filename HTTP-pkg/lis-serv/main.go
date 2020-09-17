package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) { // now hotdog is an interface og handler
	fmt.Fprintln(w, "Blah blah blah")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
