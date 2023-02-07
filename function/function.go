package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func showName(name string) {
	fmt.Println("My name is", name)
}

func main() {
	add(2, 3)
	showName("Adam")
}
