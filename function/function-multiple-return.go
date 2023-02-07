package main

import "fmt"

func multiReturn(a int, b int) (int, int) {
	result1 := a + b
	result2 := a*b + a

	return result1, result2
}

func main() {
	showResult1, showResult2 := multiReturn(10, 20)

	fmt.Println(showResult1)
	fmt.Println(showResult2)
}
