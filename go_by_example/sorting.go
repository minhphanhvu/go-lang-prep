package main

import (
		"fmt"
		"sort"
)

func main() {
		strs := []string{"z", "b", "c"}
		sort.Strings(strs)
		fmt.Println("strs:", strs)

		ints := []int{6, -5, 3}
		sort.Ints(ints)
		fmt.Println("ints:", ints)

		s := sort.IntsAreSorted(ints)
		fmt.Println("Sorted?:", s)
}
