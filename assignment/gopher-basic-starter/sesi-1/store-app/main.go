package main

import "fmt"

// Kita akan membuat aplikasi toko sederhana, kita akan belajar menggunakan map, slice, for, if-else dalam project ini
// Kerjakan hanya bagian kode yang ada "TODO:" nya yaa
// Kerjakan step-by-step

func main() {
	//TODO:Buatlah map product:harga, nama variable map adalah "mapProduct" (1)
	/*
		pisang:10000
		jeruk:5000
		semangka:20000
		durian:50000
	*/
	//start_answer
	mapProduct := make(map[string]int)

	mapProduct["pisang"] = 10000
	mapProduct["jeruk"] = 5000
	mapProduct["semangka"] = 20000
	mapProduct["durian"] = 50000
	//end_answer
	fmt.Println(mapProduct)

	//TODO: Buatlah slice of int, yang akan digunakan untuk menampung harga product yang dipilih, nama slice "slicePrices" (2)
	//start_answer
	slicePrices := []int{}
	//end_answer

	// Sekarang kita akan membuat seolah-oleh pembeli membeli pisang, jeruk, dan semangka
	// ambil harga dari map lalu isi harganya ke slice
	//HINT: gunakan append pada slice dan nilai nya diambil dari map
	//TODO: tambahkan harga pisang, jeruk, dan semangka ke slice
	//start_answer
	slicePrices = append(slicePrices, mapProduct["pisang"])
	slicePrices = append(slicePrices, mapProduct["jeruk"])
	slicePrices = append(slicePrices, mapProduct["semangka"])
	//end_answer
	fmt.Println(slicePrices)

	var totalPrice int
	//TODO: hitung total harga di slicePrices, dan taruh total nya ke variable totalPrice (3)
	//HINT: gunakan for range pada slice, dan jumlahkan nilainya
	//start_answer

	for _, value := range slicePrices {
		totalPrice += value
	}
	//end_answer

	//TODO: jika harga lebih besar dari 10000 maka pelanggan akan mendapatkan diskon 1000, lalu cetak total harga nya,
	// jika tidak maka cetak total harga saja
	//start_answer
	if totalPrice > 10000 {
		finalPrice := totalPrice - 1000
		fmt.Println("Selamat kamu mendapatkan diskon 1000!")
		fmt.Println("Total :", finalPrice)
	}
	//end_answer
}
