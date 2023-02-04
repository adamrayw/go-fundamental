package main

import "fmt"

/*
	Kalian sudah melihat penggunaan loop pada materi sebelumnya
	Tetapi disini kita akan coba loop lebih dalam lagi
*/

func main() {
	//TODO: Buatlah kode looping sederhana untuk mencetak nilai 0-10
	//start_answer
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	//end_answer

	//TODO: Buatlah kode looping dengan kondisi, jika nilai genap cekak nilai tersebut, nilai 0-10
	//Hint: Nilai genap adalah nilai yang habis dibagi 2, gunakan moduluse (%), 4%2=0
	//start_answer
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	//end_answer

	//TODO: Buatlah kode looping yang hanya mencetak nilai ganjil, nilai 0-20
	//Hint: gunakan 'continue' untuk langsung melanjutkan proses loop tanpa menjalankan eksekusi dibawahnya
	//start_answer
	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
	//end_answer
}
