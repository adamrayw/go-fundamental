package main

import "fmt"

type person struct {
	name string
	age  int
}

type school struct {
	schoolName string
	person
}

func main() {
	john := school{}

	john.name = "John Adam"
	john.age = 20
	john.schoolName = "Harvard School"

	fmt.Println(john.name)
	fmt.Println(john.age)
	fmt.Println(john.person.age)
}
