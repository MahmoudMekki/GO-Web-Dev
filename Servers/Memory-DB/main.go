package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8081")
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
	defer conn.Close()

	io.WriteString(conn, "Use:\n"+
		"GET [Key]\n"+
		"SET [key][value]\n"+
		"DEL [key]\n\n"+
		"EXAMPLE:\n"+
		"SET fav chocolate\n"+
		"GET fav\n"+
		"********************\n")
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		switch words[0] {
		case "GET":
			k := words[1]
			v := data[k]
			fmt.Fprintf(conn, "Your %s is %s\n", k, v)
		case "SET":
			k := words[1]
			v := words[2]
			data[k] = v
		case "DEL":
			k := words[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "Invalid entry")
		}
	}

}
