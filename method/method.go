package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (p Person) sayHello() {
	fmt.Println("Name :", p.Name)
	fmt.Println("age :", p.Age)
	fmt.Println("gender :", p.Gender)
}

func main() {
	p := Person{
		Name:   "Adam",
		Age:    20,
		Gender: "Male",
	}

	p.sayHello()
}
