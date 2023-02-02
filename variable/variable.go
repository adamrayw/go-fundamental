package main

import "fmt"

func main() {
	// Multiple variables in the same type
	var num1, num2, num3 int = 1, 2, 3

	// // Multiple variables of different types
	var num, name, isTrue = 1, "Adam", true

	// In Go language variables are created in two different ways:
	// declare variable with data type
	var firstName string = "Adam"

	// short variable declaration
	firstName2 := "Ray"

	// without explicit type
	var firstName3 = "John"

	fmt.Println(firstName, firstName2, firstName3)
}
