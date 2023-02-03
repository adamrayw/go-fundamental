package main

import "fmt"

/*
	Disini kita akan belajar operator logic
	Ada tiga jenis logical operator dalam Go

		&& (and)
		|| (or)
		! (not)

	Studi kasus:
		Buatlah program sederhana
		Jika jenis kelamin P maka cetak perempuan, jika jenis kelamin L maka cetak laki-laki, jika tidak keduanya maka cetak tidak valid
		Jika nilai diantara 85-100 maka cetak = A
		Jika nilai diantara 70-84 maka cetak = B
		Jika nilai diantara 0 - 70 maka cetak = C

	Contoh Input :
		- Case #1
			jenis_kelamin = L
			nilai = 85
		- Case #2
			jenis_kelamin = X
			nilai = 90

	Contoh Output :
		 - Case #1
			laki-laki dan nilai A
		 - Case #2
			tidak valid
*/

func main() {
	jenis_kelamin := "P"
	nilai := 33

	//TODO:
	//start_answer
	if jenis_kelamin != "L" && jenis_kelamin != "P" {
		fmt.Printf("Tidak Valid")
	} else if jenis_kelamin == "L" {
		if nilai >= 85 && nilai < 100 {
			fmt.Printf("Laki - laki dan Nilai A")
		} else if nilai >= 70 && nilai <= 84 {
			fmt.Printf("Laki - laki dan Nilai B")
		} else {
			fmt.Printf("Laki - laki dan Nilai C")
		}
	} else if jenis_kelamin == "P" {
		if nilai >= 85 && nilai < 100 {
			fmt.Printf("Perempuan dan Nilai A")
		} else if nilai >= 70 && nilai <= 84 {
			fmt.Printf("Perempuan dan Nilai B")
		} else {
			fmt.Printf("Perempuan dan Nilai C")
		}
	}
	//end_answer
}
