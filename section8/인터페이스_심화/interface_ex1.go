package main

import "fmt"

func printContents(v interface{}) {
	fmt.Printf("Type : (%T)\n", v) // v의 원본 타입출력 (%T)
	fmt.Println("ex1 : ", v)
	fmt.Println()
}

func main() {
	var a interface{}
	printContents(a)

	a = 7.5
	printContents(a)

	a = "golang"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)
	/*
	   결과 :
	   Type : (<nil>)
	   ex1 :  <nil>

	   Type : (float64)
	   ex1 :  7.5

	   Type : (string)
	   ex1 :  golang

	   Type : (bool)
	   ex1 :  true

	   Type : (<nil>)
	   ex1 :  <nil>
	*/
}
