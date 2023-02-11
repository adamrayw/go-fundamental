package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello World", "World")) // true

	fmt.Println(strings.ToLower("HELLO WORLD")) // hello world

	replaceWord := strings.Replace("Hello World", "World", "Golang", -1)
	fmt.Println(replaceWord)

	splitString := strings.Split("Hello World", " ")
	fmt.Println(splitString)

	jointString := strings.Join([]string{"Hello", "World"}, " ")
	fmt.Println(jointString)
}
