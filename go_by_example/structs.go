package main

import "fmt"

type person struct {
		name string
		age int
}

func newPerson(name string) *person {
		p := person{name: name}
		p.age = 42 // default age
		return &p
}

func main() {
		fmt.Println(newPerson("Jon"))
		new_person := newPerson("Doe")
		fmt.Println(new_person) 
}
