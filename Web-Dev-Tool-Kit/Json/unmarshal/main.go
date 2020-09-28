package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type data []struct {
	Precision string
	Latitude  float64
	Longitude float64
	Address   string
	City      string
	State     string
	Zip       string
	Country   string
}

func main() {
	var d data
	rcvd := `[
        {
            "precision": "zip",
            "Latitude":  37.7668,
            "Longitude": -122.3959,
            "Address":   "",
            "City":      "SAN FRANCISCO",
            "State":     "CA",
            "Zip":       "94107",
            "Country":   "US"
        },
        {
            "precision": "zip",
            "Latitude":  37.371991,
            "Longitude": -122.026020,
            "Address":   "",
            "City":      "SUNNYVALE",
            "State":     "CA",
            "Zip":       "94085",
            "Country":   "US"
        }
	]`

	err := json.Unmarshal([]byte(rcvd), &d)
	if err != nil {
		log.Fatal(err)
	}
	for i := range d {
		fmt.Println(d[i])
	}
	fmt.Println(d[1].Country)
	fmt.Println(d[0].Country)

}
