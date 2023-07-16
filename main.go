package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := []byte(`{"name":"Alice", "age":30}`)

	// var person Person
	var person map[string]any
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// fmt.Println(person.Name)
	fmt.Println(person["name"])
}
