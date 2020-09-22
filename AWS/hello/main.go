package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":80", nil)

}
func foo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from AWS!!!")
}
