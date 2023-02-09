package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type StudentInterface interface {
	getStudentName() string
}

func (s *Student) getStudentName() string {
	return s.Name
}

func (s *Student) getStudentAge() int {
	return s.Age
}

func main() {
	s := Student{Name: "Adam", Age: 20}

	fmt.Println(s.getStudentName())
	fmt.Println(s.getStudentAge())
}
