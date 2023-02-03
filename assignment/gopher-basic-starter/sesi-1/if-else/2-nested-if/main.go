package main

import "fmt"

/*
	Disini kita akan belajar nested-if, yaitu if didalam if

	Studi kasus:
		Buatlah program sederhana
		Jika jenis kelamin P maka cetak perempuan, jika jenis kelamin L maka cetak laki-laki
		Jika nilai >= 80 cetak lulus, jika nilai < 80 cetak tidak lulus

	Contoh Input :
		jenis_kelamin = L
		nilai = 85

	Contoh Output :
		laki-laki dan lulus
*/

func main() {
	jenis_kelamin := "L"
	nilai := 77

	//TODO:
	//start_answer
	if jenis_kelamin == "L" {
		if nilai >= 80 {
			fmt.Printf("Laki - laki dan Lulus")
		} else {
			fmt.Printf("Laki - laki dan Tidak Lulus")
		}
	} else {
		if nilai >= 80 {
			fmt.Printf("Perempuan dan Lulus")
		} else {
			fmt.Printf("Perempuan dan Tidak Lulus")
		}
	}
	//end_answer
}
