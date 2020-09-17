package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	fmt.Fprintln(conn, "Hello there")
	fmt.Fprintln(conn, "Welcome to my server")
}
