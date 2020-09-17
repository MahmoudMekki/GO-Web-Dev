package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	/*name := "Mahmoud Mekki"
	str := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang= "en">
	<head>
	<meta charset = "UTF-8">
	<title>Hello World</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)
	fmt.Println(str)*/
	name := os.Args[1]
	str := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang= "en">
	<head>
	<meta charset = "UTF-8">
	<title>Hello World</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	nf, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Error creating file !")
		os.Exit(1)
	}
	io.Copy(nf, strings.NewReader(str))
	defer nf.Close()

}
