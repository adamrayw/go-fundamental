package main

import "fmt"

func main() {
	score := 10

	// in the go language, the difference is only in brackets, in go you don't need brackets
	if score == 10 {
		fmt.Printf("Sempurna")
	} else if score >= 5 {
		fmt.Printf("Nilai Cukup")
	} else {
		fmt.Printf("Nilai Kurang")
	}

}
