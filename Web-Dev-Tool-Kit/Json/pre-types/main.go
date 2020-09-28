package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var data []byte

	rcvd := "null"
	json.Unmarshal([]byte(rcvd), &data)

	fmt.Println(data)
	fmt.Println(len(data))
	fmt.Println(cap(data))
}
