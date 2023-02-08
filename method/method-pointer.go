package main

import "fmt"

type Student struct {
	Name  string
	Score int
	Index string
}

func (s Student) modifyIndex1(score int) {
	if score > 80 {
		s.Index = "A"
	} else {
		s.Index = "B"
	}
}

func (s *Student) modifyIndex2(score int) {
	if score > 80 {
		s.Index = "A"
	} else {
		s.Index = "B"
	}
}

func main() {
	s := Student{"Adam", 100, ""}

	s.modifyIndex1(s.Score)
	fmt.Println(s.Index)

	s.modifyIndex2(s.Score)
	fmt.Println(s.Index)

}
