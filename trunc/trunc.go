/*Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.*/
package main

import "fmt"

func main() {
	var floatNum float32
	var intNum int32

	fmt.Printf("Please enter a floating number: ")
	fmt.Scan(&floatNum)
	intNum = int32 (floatNum)
	fmt.Printf("Your integer number is always rounded down: %d\n", intNum)
	fmt.Printf("Float number does not change: %f\n", floatNum)
}