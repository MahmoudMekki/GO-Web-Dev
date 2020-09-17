package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog/", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--image doesn't serve-->
	<img src="/toby.jpg">
	`)
}

func dogpic(res http.ResponseWriter, req *http.Request) {
	dog, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(res, "File not found", 404)
	}
	io.Copy(res, dog)
	defer dog.Close()

}
