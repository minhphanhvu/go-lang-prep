package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"sort"
)

func main() {
	s := make([]int, 0, 3)
	var stop bool = false
	scanner := bufio.NewScanner(os.Stdin)

	for stop != true {
		fmt.Printf("Please enter a new number:\n")
		scanner.Scan()
		if scanner.Text() == "X" {
			stop = true
		} else {
			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			num := int(input)
			s = append(s, num)
			sort.Ints(s)
			fmt.Println(s)
		}
	}
}