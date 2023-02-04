package main

import "fmt"

func main() {
	cars := [4]string{"BMW", "AUDI", "HONDA"}

	cars[3] = "TOYOTA"

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
}
