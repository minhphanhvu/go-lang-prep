package main

import (
	"os"
	"log"
	"fmt"
	"bufio"
	"strings"
)

const MaxLength = 20

type Name struct {
	fname string
	lname string
}

func (n *Name) Set(firstName string, lastName string) {
	n.fname = firstName
	n.lname = lastName

	if len(firstName) > MaxLength {
		n.fname = firstName[:MaxLength]
	}

	if len(lastName) > MaxLength {
		n.lname = lastName[:MaxLength]
	}
}

func (n *Name) FullName() string {
	return n.fname + " " + n.lname
}


func main() {
	s := make([]Name, 0, 0)

	f, err := os.Open("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		firstName := names[0]
		lastName := names[1]
		newName := Name{}
		newName.Set(firstName, lastName)
		s = append(s, newName)
	}

	for _, name := range s {
		fmt.Println(name.FullName())
	}
}