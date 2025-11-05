package main

import (
	"fmt"
)

func Greetings() {
	greetings := []string{"Hello", "Hola", "नमार", "こんにちは", "Привіт"}

	for i, greet := range greetings {
		fmt.Println(i, greet)
	}

	firstSlice := greetings[:2]
	secondSlice := greetings[1:4]
	thirdSlice := greetings[3:]

	fmt.Println(firstSlice, secondSlice, thirdSlice)

}
