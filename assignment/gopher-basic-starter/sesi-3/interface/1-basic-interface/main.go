package main

import "fmt"

/*
	Teman-teman sudah belajar interface di materi sebelumnya, dan sekarang saatnya untuk mengimplementasikan interface tersebut.
	Terdapat sebuah struct Employee yang memiliki method CalculateTax yang mengembalikan nilai float64 dan error.

	Buatlah sebuah Interface EmployeeInterface yang memiliki method CalculateTax yang mengembalikan nilai float64 dan error.

	Implementasikan interface TaxCalculator pada struct Employee.
*/

//TODO : Buatlah interface EmployeeInterface yang memiliki method CalculateTax yang mengembalikan nilai float64 dan error
//start_answer
type EmployeeInterface interface {
	CalculateTax() (float64, error)
}

//end_answer

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e *Employee) CalculateTax() (float64, error) {
	//TODO : implementasikan calculate tax
	/*
		Jika salary < 1000 maka tax = 0.1 * salary
		Jika salary >= 1000 maka tax = 0.2 * salary
		Jika salary < 0 maka return error "salary cannot be negative"
	*/
	//start_answer

	//end_answer
}

func main() {
	var budi, joko EmployeeInterface

	budi = &Employee{
		ID:     1,
		Name:   "Budi",
		Salary: 1000,
	}

	joko = &Employee{
		ID:     2,
		Name:   "Joko",
		Salary: 2000,
	}

	budiTax, err := budi.CalculateTax()
	if err != nil {
		fmt.Println(err)
	}

	jokoTax, err := joko.CalculateTax()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(budiTax) // 100
	fmt.Println(jokoTax) // 400
}
