package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["alice"] = 10
	ages["adam"] = 20

	value, isExist := ages["alice"]

	if isExist {
		fmt.Println("Alice available", value, isExist)
	} else {
		fmt.Println("Item not available")
	}

	// for key, value := range ages {
	// 	fmt.Println("age of", key, "is", value)
	// }
}
