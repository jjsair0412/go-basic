package main

import "fmt"

type Account struct {
	number   string
	balance  float64 //잔액
	interest float64 // 이자
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {

	// 포인터형으로 구조체를 선언할 때 ,
	// new 키워드를 붙여서 아래처럼 선언하려면 * 를 붙여주어야 함
	// 인터페이스에 넘길때는 아래처럼 사용해야함.
	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	// 주소값 넘겨서 선언
	hong := &Account{
		number:   "245-902",
		balance:  1500000,
		interest: 0.04,
	}

	// 짧은 선언으로 new 키워드 사용해서 선언
	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 1300000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", hong)
	fmt.Println("ex1 : ", lee)
	/*
	   결과 :
	   ex1 :  &{245-901 1e+07 0.015}
	   ex1 :  &{245-902 1.5e+06 0.04}
	   ex1 :  &{245-903 1.3e+06 0.025}
	*/

	fmt.Println()

	// %#v 는 구조체 값을 모두 출력
	fmt.Printf("ex2 : %#v\n", kim)
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex2 : %#v\n", lee)
	/*
	   결과 :
	   ex2 : &main.Account{number:"245-901", balance:1e+07, interest:0.015}
	   ex2 : &main.Account{number:"245-902", balance:1.5e+06, interest:0.04}
	   ex2 : &main.Account{number:"245-903", balance:1.3e+06, interest:0.025}
	*/

	fmt.Println()

	fmt.Println("ex3 : ", int(kim.Calculate()))
	fmt.Println("ex3 : ", int(hong.Calculate()))
	fmt.Println("ex3 : ", int(lee.Calculate()))
	/*
	   결과 :
	   ex3 :  10150000
	   ex3 :  1560000
	   ex3 :  1332500
	*/
}
