package main

import "fmt"

// Sekarang kita akan coba untuk membaca huruf pertama dan terakhir dari nama kalian
// Buatlah nama kalian menjadi arrat, contoh: [R, O, B, E, R, T]
// Dengan mengakses indeks 0 untuk huruf pertama, dan indeks 5 untuk huruf terakhir
// Didaptkan hasil R dan T

func main() {
	//TODO: Buatlah nama kalian menjadi array
	//start_answer
	nama := [4]string{"A", "D", "A", "M"}
	//end_answer

	//TODO: Cetak huruf pertama dari nama kalian dengan mengakses indeks 0
	//start_answer
	fmt.Println(nama[0])
	//end_answer

	//TODO: Cetak huruf terakhir dari nama kalian dengan mengakses indeks terakhir
	//start_answer
	fmt.Println(nama[3])
	//end_answer

	//TODO: dengan menggunakan perintah loop cetak nama kalian secara berurutan
	//Hint: gunakan for range
	//Hint: gunakan fmt.printf() untuk mencetak tambah menambahkan baris baru
	//start_answer
	for _, name := range nama {
		fmt.Println("value: ", name)
	}
	//end_answer
}
