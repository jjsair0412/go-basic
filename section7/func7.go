package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2 : (", n, ")", "hi!")
	prtHello(n - 1)
}

func main() {
	x := fact(5)

	fmt.Println("ex1 : ", x)
	// 결과 : ex1 :  120

	prtHello(10)
	/*
	   결과 :
	   ex2 : ( 10 ) hi!
	   ex2 : ( 9 ) hi!
	   ex2 : ( 8 ) hi!
	   ex2 : ( 7 ) hi!
	   ex2 : ( 6 ) hi!
	   ex2 : ( 5 ) hi!
	   ex2 : ( 4 ) hi!
	   ex2 : ( 3 ) hi!
	   ex2 : ( 2 ) hi!
	   ex2 : ( 1 ) hi!
	*/
}
