package main

import (
	"fmt"
)

func Runes() {
	message := "Hi 👩 and 👨"
	fmt.Println(message)
	runes := []rune(message)
	fmt.Println(string(runes[3]))
}
