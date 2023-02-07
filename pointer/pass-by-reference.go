package main

import "fmt"

func Add(num *int) {
	result := *num

	fmt.Println(result)
}

func main() {
	num := 8
	Add(&num)
}
