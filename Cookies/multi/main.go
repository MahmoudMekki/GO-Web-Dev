package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/group", gorup)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)

}

func set(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "First-cookie",
		Value: "001",
		Path:  "/",
	})
	fmt.Fprintln(res, "Cookie is written check your browser!")
}

func read(res http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("First-cookie")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(res, "Your cookies is %s\n", c1)

	c2, err := req.Cookie("Second-cookie")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(res, "Your cookies is %s\n", c2)

	c3, err := req.Cookie("Third-cookie")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(res, "Your cookies is %s\n", c3)

}

func gorup(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name:  "Second-cookie",
		Value: "002",
	})

	http.SetCookie(res, &http.Cookie{
		Name:  "Third-cookie",
		Value: "003",
	})

	fmt.Fprintln(res, "Your cookies have been written check your browser!!!")
}
