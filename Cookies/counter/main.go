package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func foo(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("My-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "My-cookie",
			Value: "0",
		}
	}

	counter, _ := strconv.Atoi(cookie.Value)
	counter++
	cookie.Value = strconv.Itoa(counter)
	http.SetCookie(res, cookie)

	fmt.Fprintf(res, "Thanks for your %s visit to our website! you means alot to us!\n", cookie.Value)
}
