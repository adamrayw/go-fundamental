package main

import "fmt"

type student struct {
	name    string
	school  string
	age     int
	isLulus bool
}

func main() {
	jonathan := student{name: "Jonathan", school: "Harvard School", age: 20, isLulus: false}
	fmt.Println(jonathan.name)
}
