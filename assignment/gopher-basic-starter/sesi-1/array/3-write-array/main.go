package main

import "fmt"

/*
	Teman-teman sudah belajar mengenai array
	Pada sesi ini, teman-teman akan mencoba untuk membuat array sendiri dan menulis data ke dalam array tersebut
*/

func main() {
	var namaKota [5]string

	//TODO: isi array namaKota dengan nama kota
	// urutan nama kota: Jakarta, Bandung, Surabaya, Semarang, dan Makassar
	//start_answer
	namaKota[0] = "Jakarta"
	namaKota[1] = "Bandung"
	namaKota[2] = "Surabaya"
	namaKota[3] = "Semarang"
	namaKota[4] = "Makassar"
	//end_answer

	fmt.Println(namaKota) // [Jakarta Bandung Surabaya Medan Makassar]
}
