package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

type Variables struct {
	acc float64
	initVel float64
	initDis float64
}

func PromptInput(variable string) float64 {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter your %s\n", variable)
	var result float64
	for scanner.Scan() {
		input, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Please enter the %s again\n", variable)
			continue
		}
		result = input
		break
	}
	return result
}

func GenDisplaceFn(acc float64, initVel float64, initDis float64) func (float64) float64 {
	return func (time float64) float64 {
		return 0.5 * acc * math.Pow(time, 2) + initVel * time + initDis
	}
}

func main() {
	var variables Variables
	var time float64
	variables.acc = PromptInput("acceleration")
	variables.initVel = PromptInput("velocity")
	variables.initDis = PromptInput("displacement")
	time = PromptInput("time")
	fn := GenDisplaceFn(variables.acc, variables.initVel, variables.initDis)
	fmt.Println(fn(time))
}