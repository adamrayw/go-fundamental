package main

import "fmt"

func main() {
	num := 10

	pointerNumber := &num

	fmt.Println(pointerNumber)
	fmt.Println(*pointerNumber)
}
