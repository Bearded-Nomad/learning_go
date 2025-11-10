package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intSlice := make([]int, 100)

	for i := 0; i < intSlice[len(intSlice)-1]; i++ {
		intSlice = append(intSlice, rand.Intn(100))
	}

	for i := range intSlice {
		if (i%2 == 0) && (i%3 == 0) && (i != 0) {
			fmt.Println(i, "Six!")
		} else if i%2 == 0 {
			fmt.Println(i, "Two!")
		} else if i%3 == 0 {
			fmt.Println(i, "Three!")
		} else {
			fmt.Println(i, "Never mind!")
		}
	}

	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	// Prints 0 cause out of scope
	fmt.Println(total)
}
