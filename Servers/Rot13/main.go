package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn, "%s - %s\n", ln, r)
	}
	conn.Close()
}

func rot13(bs []byte) []byte {
	re := make([]byte, len(bs))
	for _, v := range bs {
		if v <= 109 {
			re = append(re, v+13)
		} else {
			re = append(re, v-13)
		}
	}
	return re
}
