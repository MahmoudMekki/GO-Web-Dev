package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "My name is mahmoud hamdi mekki,\"Egyptian\""
	//encodingstd := "ABCSFDASD0123456789%$^#"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s)
	fmt.Println(s64)

	a, _ := base64.StdEncoding.DecodeString(s64)
	fmt.Println(string(a))

}
