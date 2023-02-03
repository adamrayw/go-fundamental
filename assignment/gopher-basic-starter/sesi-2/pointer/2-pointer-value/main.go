package main

import "fmt"

/*
	Pointer bisa di-passing sebagai parameter pada function
	Apabila parameter merupakan sebuah pointer, fungsi tersbut dapat langsung memodifikasi nilai pointer tanpa harus mengembalikan nilai
*/

// pada fungsi nonPointer perubahan nilai a tidak akan berefek pada kode yang memanggil fungsi nonPointer
func nonPointer(a int) {
	a = 10
}

//TODO: Buatlah fungsi pointer yang menerima pointer int dan ubah nilainya
//start_answer

//end_answer

func main() {
	number := 20
	//Nilai tidak berubah jika memanggil fungsi nonPointer
	nonPointer(number) // tetap 20
	fmt.Println(number)

	//Nilai berubah jika memanggil fungsi pointer
	pointer(&number)
	fmt.Println(number)
}
