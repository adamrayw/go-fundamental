package main

import "fmt"

//Fungsi HelloWorld tanpa parameter dan tanpa return value
//TODO: Buatlah fungsi untuk mencetak "Hello World"
//start_answer
func HelloWorld() {
	fmt.Println("Hello World")
}

//end_answer

//Fungsi HelloName dengan single parameter tanpa return value
//TODO: Buatlah fungsi yang menerima nama(string) sebagai input dan cetak "Hello Nama", eg input:rinal output:Hello rinal
//start_answer
func HelloName(nama string) {
	fmt.Println(nama)
}

//end_answer

//Fungsi Add dengan multiple parameter dengan return value
//TODO: Buatlah fungsi yang menerima dua buah bilangan, dan balikan nilai perkaliannya
//start_answer
func Add(a int, b int) int {
	return a * b
}

//end_answer

func main() {
	HelloWorld()
	HelloName("Rinal")
	fmt.Println("Hasil Perkalian adalah ", Add(5, 4))
}
