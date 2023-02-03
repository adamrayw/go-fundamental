package main

import "fmt"

/*
	Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.

	Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct
	Dan juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.

	Pada latihan ini kita akan berkenalan dengan Embedded struct
	yaitu mekanisme untuk menempelkan sebuah struct sebagai properti struct lain.
	Agar lebih mudah dipahami, mari kita bahas pada latihan ini.
*/

type Subject struct {
	subjectName  string
	subjectScore float64
}

type Address struct {
	Street string
	City   string
}

type Student struct {
	name     string
	age      int
	address  Address   //baris ini merupakan embedded struct
	subjects []Subject //baris ini merupakan embeded struct, dan kita bisa meng-embed slice of strcut
}

// Untuk membuat method kita bisa menuliskan struct apa yang menempel pada fungsinya, lihat contoh dibawah
// Fungsi dibawah kita akan menghitung nilai rata-rata dari struct student.subjects
func (s *Student) averageScore() float64 {
	var total float64
	var countSubject int

	// disini kita mengakses nilai subjectsScore dari struct student.Subject
	for _, subject := range s.subjects {
		total = total + subject.subjectScore
		countSubject++
	}

	return total / float64(countSubject)
}

// TODO: buatlah method printSubjects yang akan mencetak seluruh subject yang diambil oleh object student
// Output: Kalkulus, Alpro, ASD, Kewarganegaraan
//start_answer


//end_answer

func main() {
	// bagian ini berfungsi untuk membuat object reginaSubject yang merupakan slice of struct
	reginaSubject := []Subject{
		{
			subjectName:  "Kalkulus",
			subjectScore: 81.50,
		},
		{
			subjectName:  "Alpro",
			subjectScore: 71.50,
		},
		{
			subjectName:  "ASD",
			subjectScore: 61.50,
		},
		{
			subjectName:  "Kewarganegaraan",
			subjectScore: 91.50,
		},
	}

	regina := Student{
		name: "Regina Putri",
		age:  20,
		address: Address{
			Street: "Jl kenangan banget deh ah",
			City:   "Jakarta Timur",
		},
		subjects: reginaSubject, // kita meng-embed reginaSubject yang sudah kita buat ke object regina
	}

	fmt.Printf("Nilai rata-rata :%f", regina.averageScore()) // cara pemanggilan method yang menempel dengan cara regina.averageScore()
	fmt.Println()
	regina.printSubjects() // Kalkulus,Alpro,ASD,Kewarganegaraan,
}
