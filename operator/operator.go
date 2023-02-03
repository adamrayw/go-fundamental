package main

import "fmt"

func main() {
	a := 1
	b := 2

	// use of arithmetic operator
	// addiction
	fmt.Printf("a + b = %d", a+b)

	// substraction
	fmt.Printf("a - b = %d", a-b)

	// relational operator
	result := a == b
	fmt.Println(result)

	result1 := a != b
	fmt.Println(result1)

	// Logical Operator
	if a == b {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
