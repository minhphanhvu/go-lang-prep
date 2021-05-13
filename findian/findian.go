package main

import "fmt"
import "strings"

func main() {
	var strInput string

	fmt.Printf("Enter a string:\n")
	fmt.Scan(&strInput)
	strInput = strings.ToLower(strInput)

	if strings.Index(strInput, "i") == 0 &&
		strings.Contains(strInput, "a") &&
		strInput[len(strInput)-1:] == "n" {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}