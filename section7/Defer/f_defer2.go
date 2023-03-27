package main

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi !")
	}()
}

func main() {
	sayHello("Golang")
	/*
	   결과 :
	   Hi !
	   Golan

	   함수가 끝나기 직전에 미뤄놨던 defer func() 익명함수가 실행됩니다.
	*/
}
