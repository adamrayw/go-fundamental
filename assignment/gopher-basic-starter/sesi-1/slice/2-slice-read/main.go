package main

import "fmt"

// Untuk membaca slice itu sebenarnya sama seperti pada array.
// Kita cukup mengambil index nya
func main() {
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)
	fmt.Println(slice[0])
	fmt.Println(slice[9])

	//TODO: cetak nilai slice menggunakan for-range
	//Output:
	/*
		indek:0 nilai:1
		indek:1 nilai:2
		...
	*/
	//start_answer

	//end_answer
}
