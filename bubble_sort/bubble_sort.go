package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func BubbleSort(numbers []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(numbers) - 1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				swapped = true
			}
		}
	}
}

func Swap(numbers []int, idx int) {
	temp := numbers[idx]
	numbers[idx] = numbers[idx+1]
	numbers[idx+1] = temp
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := make([]int, 0, 0)

	fmt.Println("Please type your numbers, hit enter after each number to enter another.")
	fmt.Println("Press enter up to 10 numbers.")
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		numbers = append(numbers, num)
		if len(numbers) == 10 {
			break
		}
	}

	BubbleSort(numbers)
	fmt.Println(numbers)
}