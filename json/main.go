package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Isadult bool   `json:"isadult"`
}

func main() {
	fmt.Println("Learning JSON in Go")
	p := Person{
		Name:    "John Doe",
		Age:     30,
		Isadult: true,
	}
	fmt.Println(p)
	// convert to json encoding (Marshal)
	jsondata, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsondata))
	// convert to json decoding (Unmarshal)
	var p1 Person
	err = json.Unmarshal(jsondata, &p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)
}
