package main

import "fmt"

func main() {
	// Sama hal nya array, kita bisa merubah nilai dari indeks tertentu
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(sliceInt)

	//TODO: ubah nilai indeks 3 dan 4, menjadi 0
	//start_answer
	sliceInt[3] = 0
	sliceInt[4] = 0
	//end_answer
	fmt.Println(sliceInt)

	// Kita juga dapat menggabungkan slice ke slice lain
	sliceIntNew := []int{11, 12, 13}

	//TODO: gabungkan sliceIntNew ke sliceInt
	//start_answer
	sliceInt = append(sliceInt, sliceIntNew...)
	//end_answer
	fmt.Println(sliceInt)
}
