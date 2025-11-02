package main

import "fmt"

const value = 100

func main() {
	var i int = value
	var f float32 = float32(i)
	fmt.Println(i, f)

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println(b, smallI, bigI)
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)

}
