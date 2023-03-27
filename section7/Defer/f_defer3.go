package main

import "fmt"

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 : ", i)
	}
}

func main() {
	stack()
	/*
	   결과 :
	   ex1 :  10
	   ex1 :  9
	   ex1 :  8
	   ex1 :  7
	   ex1 :  6
	   ex1 :  5
	   ex1 :  4
	   ex1 :  3
	   ex1 :  2
	   ex1 :  1

	   defer 함수로 i를 찍기 때문에 , 선입선출이 된다.
	   1부터 10까지 누적되며 예약한 뒤 , 함수가 끝나기 직전 (여기선 명령어)에
	   마지막에 예약됐던 10번부터 출력된다.
	*/
}
