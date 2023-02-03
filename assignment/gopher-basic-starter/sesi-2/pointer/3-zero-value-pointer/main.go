package main

import "fmt"

// Pointer memiliki zero value berupa nil. Ketika kita mencoba mengakses atribut dari nil, maka program akan terjadi panic.

func squarePanic(number *int) {
	*number *= *number
}

//TODO: buatlah fungsi square yang meneima pointer int, dan ubah nilainya menjadi hasil pangkat, cek apakah valunya nil atau tidak
// jika nil, cetak "nilai nil"
//start_answer


//end_answer

func main() {
	// Baris kode dibawah akan menghasilkan panic
	var numberEmpty *int

	// Silahkan comment/uncomment baris dibawah untuk mencoba solusi yang dibuat
	//squarePanic(numberEmpty)

	// ------- //
	square(numberEmpty) // nilai nil
	number := 10
	numberEmpty = &number
	square(numberEmpty)
	fmt.Println(*numberEmpty) // 100

}
