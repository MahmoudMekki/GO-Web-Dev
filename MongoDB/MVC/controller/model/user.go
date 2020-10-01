package models

// User struct from models package
type User struct {
	Name   string `json : "name"`
	Gender string `json : "gender"`
	Age    int    `json : "age"`
	ID     string `json : "id"`
}
