package main

import "fmt"

/*
	Go tidak memiliki class yang ada di bahasa-bahasa strict OOP lain. Tapi Go memiliki tipe data struktur yang disebut dengan Struct.

	Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method),
	yang dibungkus sebagai tipe data baru dengan nama tertentu.

	Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal,
	dan tipe data tiap itemnya bisa berbeda.
*/

//TODO: buatlah struct Student dengan spesifikasi berikut
/*
	name string
	age int
	score float64
*/
//start_answer


//end_answer

func main() {
	var asti Student // merupakan cara inisiasi varibale dari struct
	asti.name = "Asti"
	asti.age = 20
	asti.score = 85.75

	fmt.Println(asti)

	// cara lain untuk inisasi varibale dari struct
	vania := Student{
		name:  "Vania",
		age:   19,
		score: 70.25,
	}

	// cara dibawah merupakan cara untuk mengakses property dari struct yang sudah dibuat
	fmt.Println("Nama: ", vania.name)
	fmt.Println("Umur: ", vania.age)
	fmt.Println("Score: ", vania.score)

}
