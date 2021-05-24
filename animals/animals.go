package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func makeRequest(animals *map[string]Animal) {
	request := bufio.NewScanner(os.Stdin)
	var animal, information string
	fmt.Println("Please enter animal name and the information you need.")
	fmt.Println("Two inputs need to be separated by a space.")
	fmt.Printf("> ")
	for request.Scan() {
		inputs := strings.Split(request.Text(), " ")
		if len(inputs) != 2 {
			fmt.Println("Please enter two inputs separated by a space")
		} else {
			animal = inputs[0]
			information = inputs[1]
			processInputs(&*animals, animal, information)
			fmt.Printf("> ")
		}
	}
}

func processInputs(animals *map[string]Animal, name string, information string) {
	if animal, exist := (*animals)[name]; exist {
		if information == "eat" {
			animal.Eat()
		} else if information == "move" {
			animal.Move()
		} else if information == "speak" {
			animal.Speak()
		} else {
			fmt.Println("There is no such information")
		}
	} else {
		fmt.Println("There is no such animal")
	}
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	animals := make(map[string]Animal)
	animals["cow"] = cow
	animals["bird"] = bird
	animals["snake"] = snake
	makeRequest(&animals)
}