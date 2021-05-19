package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var name string
	var address string
	person := make(map[string]string)
	fmt.Println("Please enter your name:")
	fmt.Scan(&name)
	fmt.Println("Please enter your address:")
	fmt.Scan(&address)
	
	person["name"] = name
	person["address"] = address
	j, _ := json.Marshal(person)
	fmt.Println(string(j))
}