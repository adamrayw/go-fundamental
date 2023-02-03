package main

import "fmt"

// Sekarang kita bisa menghapus data pada map menggunakan key nya.

func main() {
	var kamus = map[string]string{}
	kamus["sekolah"] = "school"
	kamus["senang"] = "happy"
	kamus["sedih"] = "sad"

	// Kita akan menghapuskan data pada map dengan key "sekolah"
	for k, v := range kamus {
		fmt.Println(k, ":", v)
	}
	//TODO: hapus key "sekolah"
	//start_answer
	
	//end_answer
	fmt.Println()
	for k, v := range kamus {
		fmt.Println(k, ":", v)
	}
}
