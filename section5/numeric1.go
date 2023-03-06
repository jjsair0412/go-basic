package main

import (
	"fmt"
)

func main() {
	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println("ex1 : ", num1)
	fmt.Println("ex2 : ", num2)
	fmt.Println("ex3 : ", num3)
	fmt.Println("ex4 : ", num4)

	/*
		결과 :
		ex1 :  17
		ex2 :  -68
		ex3 :  409
		ex4 :  855256181
	*/

}
