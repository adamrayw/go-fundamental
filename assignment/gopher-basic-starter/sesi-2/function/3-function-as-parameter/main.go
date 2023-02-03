package main

import "fmt"

/*
	Setelah pada latihan sebelumnya kita belajar mengenai fungsi yang mengembalikan nilai balik berupa fungsi,
	kali ini topiknya tidak kalah unik, yaitu fungsi yang digunakan sebagai parameter.

	Di Go, fungsi bisa dijadikan sebagai tipe data variabel.
	Dari situ sangat memungkinkan untuk menjadikannya sebagai parameter juga.
*/

// Terdapat fungsi filterGenap yang mengecek apakah nilai yang diberikan merupakan bilangan genap atau bukan
func filterGenap(number int) bool {
	return number%2 == 0
}

// TODO: buatlah fungsi cetakGenap, yang menerima slice of int, dan mengembalikan hanya bilangan genap
// Gunakan fungsi filterGenap sebagai parameter fungsi cetakGenap
//start_answer

//end_answer

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(cetakGenap(numbers, filterGenap))
}
