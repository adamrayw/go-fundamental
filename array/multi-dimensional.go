package main

import "fmt"

func main() {
	multiDimen := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < 3; i++ {
		for x := 0; x < 3; x++ {
			fmt.Println(multiDimen[i][x])
		}
	}
}
