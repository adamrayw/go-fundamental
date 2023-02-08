package main

import "fmt"

/*
	Pointer di golang memiliki dua operator yaitu address operator dan indirect operator.
	Address operator (&) digunakan untuk mendapatkan alamat memori dari suatu variabel.
	Indirect operator (*) digunakan untuk mendapatkan nilai dari alamat memori yang ditunjuk.
*/

func main() {
	name := "Vania"
	age := 20
	isMarried := true

	addressName := &name
	addressAge := &age
	addressIsMarried := &isMarried

	//TODO: Dari variable yang telah diberikan, kamu dapat mencoba untuk menggunakan address operator dan indirect operator.
	//start_answer
	fmt.Println("")
	fmt.Println("Mendapatkan Memory Address dari ketiga variable")
	fmt.Println("Name :", addressName)
	fmt.Println("Age :", addressAge)
	fmt.Println("isMarried :", addressIsMarried)
	fmt.Println("")
	fmt.Println("Mendapatkan value dari Memory Address")
	fmt.Println("Name :", *addressName)
	fmt.Println("Age :", *addressAge)
	fmt.Println("isMarried :", *addressIsMarried)
	//end_answer
}
