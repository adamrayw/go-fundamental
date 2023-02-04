package main

import "fmt"

// Array adalah suatu tipe data yang dapat menampung suatu variable dengan tipe yang sama, dan ukuran array sudah ditentukan diawal

func main() {
	/*
		ada beberapa cara melakukan inisiallisasi array dalam Go
		var namaArray [size]type
		- var array1 [10]int	<- Data kosong
		- array1 := [10]int{1,2,3,4,5,6,7,8,9,10} <- Data ada
	*/

	//TODO: Buatlah array of int dengan ukuran 10 lalu cetak isi nya dengan perintah fmt.println()
	//start_answer
	var array1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	//begin_answer
}
