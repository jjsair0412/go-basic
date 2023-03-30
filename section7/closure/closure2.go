package main

import "fmt"

func main() {

	cnt := increaseCnt()

	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())

	/*
	   결과 :
	   ex 1 :  1
	   ex 1 :  2
	   ex 1 :  3
	   ex 1 :  4
	   ex 1 :  5

	   함수 increaseCnt()에서 외부 n을 캡쳐해서 저장하고 있기 때문에 ,
	   누적되어 저장되는것이다.
	*/
}

func increaseCnt() func() int {
	n := 0 // 지역변수 0 . 캡쳐 대상
	return func() int {
		n += 1
		return n
	}
}
