package main

import (
	"fmt"
	"project-golang/student"
)

func main() {
	// inisialisasi struct student
	ton := student.Student{
		Name:  "Adam",
		Age:   20,
		Score: 80,
	}

	fmt.Println(ton.GetName())
	fmt.Println(ton.GetAge())
	fmt.Println(ton.GetScore())
}
