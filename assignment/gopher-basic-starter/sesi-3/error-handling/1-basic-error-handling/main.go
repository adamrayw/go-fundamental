package main

import "fmt"

/*
	Teman-teman sudah belajar error handling pada sesi sebelumnya.
	Sekarang, teman-teman akan mencoba untuk membuat error handling sendiri.

	Terdapat function Div(a int, b int) int yang menerima dua parameter bertipe int dan mengembalikan hasil pembagian a dengan b.
	Terdapat juga function DivWithErr(a int, b int) (int, error) yang menerima dua parameter bertipe int dan mengembalikan hasil pembagian a dengan b dan error.

	Tugas teman-teman adalah membuat fungsi DivWithErrSendiri yang menerima dua parameter bertipe int dan mengembalikan hasil pembagian a dengan b dan error.
	Fungsi DivWithErrSendiri akan mengembalikan error jika pembagi b bernilai 0.
*/

func Div(a int, b int) int {
	return a / b
}

func DivWithErr(a int, b int) (int, error) {
	//TODO: buatlah fungsi DivWithErrSendiri yang menerima dua parameter bertipe int dan mengembalikan hasil pembagian a dengan b dan error.

	//start_answer

	//end_answer
}

func main() {
	fmt.Println(DivWithErr(10, 0)) // 0 pembagi tidak boleh 0
	fmt.Println(DivWithErr(10, 2)) // 5 <nil>
}
