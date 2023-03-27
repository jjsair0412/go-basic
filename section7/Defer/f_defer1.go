package main

import "fmt"

func ex_f1() {
	fmt.Println("f1 : start !")
	defer ex_f2() // 함수가 종료되기 직전 마지막에 호출된다.
	fmt.Println("f1 : end !")
}

func ex_f2() {
	fmt.Println("f2 : called !")
}

func main() {
	ex_f1()
	/*
		결과 :
		f1 : start !
		f1 : end !
		f2 : called !

		ex_f2() 함수를 defer로 예약해 두었기 때문에 , ex_f1() 함수가
		종료된 이후( end가 찍히고 ) 나서 f2가 호출됩니다.
	*/
}
