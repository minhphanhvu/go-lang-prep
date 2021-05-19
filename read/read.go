package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, _ := ioutil.ReadAll(f)
	fmt.Println(string(b))
}