package main

import (
	"fmt"
)

func main() {
	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex1 : ", str1 == str2) // false
	fmt.Println("ex2 : ", str1 != str2) // true

	fmt.Println("ex3 : ", str1 > str2) // false !
	fmt.Println("ex3 : ", str1 < str2) // true

}
