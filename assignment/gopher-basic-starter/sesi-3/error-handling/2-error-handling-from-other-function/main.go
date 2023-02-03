package main

import (
	"fmt"
)

/*
	Error handling selain bisa dibuat sendiri, juga bisa dibuat dari fungsi lain.
	Teman-teman akan mencoba untuk membuat error handling dari fungsi lain.

	Terdapat function ConvertToInt(number string) (int, error) yang menerima satu parameter bertipe string dan mengembalikan hasil konversi string ke int dan error.
	Pada fungsi ConvertToInt terdapat error handling yang mengembalikan error jika string yang dikonversi tidak bisa dikonversi ke int.

	Error berasal dari fungsi lain dapat ditangkap dengan cara menambahkan parameter error pada fungsi yang akan menangkap error tersebut.
*/

func ConvertToInt(number string) (int, error) {
	var returnInt int
	var err error

	//TODO: isi variabel returnInt dan err dengan hasil konversi number ke int menggunakan strconv.Atoi
	//start_answer

	//end_answer
}

func main() {
	fmt.Println(ConvertToInt("10"))  // 10 <nil>
	fmt.Println(ConvertToInt("10a")) // 0 strconv.Atoi: parsing "10a": invalid syntax
}
