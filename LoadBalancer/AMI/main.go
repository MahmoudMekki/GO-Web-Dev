package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instances)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":80", nil)

}
func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello from AWS!!")
}
func ping(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "OK")
}
func instances(res http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		http.Error(res, err.Error(), 404)
	}
	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	resp.Body.Close()
	io.WriteString(res, string(bs))
}
