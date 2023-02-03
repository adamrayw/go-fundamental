package main

import "fmt"

/*
	Sekarang kalian akan mencoba membuat program sederhana konversi nilai
	Jika nilai 100 = Lulus Sempurna
	Jika nilai 80 - 99 = Lulus
	Jika nilai < 80 = Tidak Lulus

	Buatlah menggunakan if-else
*/

func main() {
	nilai := 100

	//TODO:
	//start_answer
	if nilai == 100 {
		fmt.Printf("Sempurna")
	} else if nilai >= 80 || nilai == 99 {
		fmt.Printf("Lulus")
	} else {
		fmt.Printf("Tidak Lulus")
	}
	//end_answer
}
