package main

import "fmt"

/*
Kalian sudah tau tentang array kan, sekarang kita akan mencoba slice
Slice bisa dibilang mirip dengan array, perbedaannya adalah ukuran slice fleksibel, dan tidak butuh ditetapkan ukurannya
mudah nya kita pikirkan kalau slice ini berupa array yang dinamis
*/

func main() {
	/*
	   	Untuk inisialisasi slice kita bisa menggunakan dengan cara berikut
	   	var slice1 []int
	  	var slice2 = []int{}
	*/
	//TODO: Buatlah sebuah slice of int dan isi d
	//start_answer
	sliceInt := []int{}
	//end_answer

	//TODO: isikan int sebanyak 10 buah pada slice yang dibuat
	//HINT: gunakan append
	//start_answer
	sliceInt = append(sliceInt, 1)
	sliceInt = append(sliceInt, 2)
	sliceInt = append(sliceInt, 3)
	sliceInt = append(sliceInt, 4)
	sliceInt = append(sliceInt, 5)
	sliceInt = append(sliceInt, 6)
	sliceInt = append(sliceInt, 7)
	sliceInt = append(sliceInt, 8)
	sliceInt = append(sliceInt, 9)
	sliceInt = append(sliceInt, 10)
	//end_answer

	//TODO: cetak seluruh nilai slice menggunakan fmt.println
	//start_answer
	fmt.Println(sliceInt)
	//end_answer
}
