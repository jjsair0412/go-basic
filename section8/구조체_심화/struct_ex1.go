package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account { // 포인터 반환 아닌경우 , 값 복사
	return &Account{ // 구조체 인스턴스를 생성한 후 , 주소값을 리턴
		number,
		balance,
		interest,
	}
}

func main() {
	// 구조체 선언과 동시에 입력
	// 생성자 활용 ( 맴버변수에 접근 )
	kim := Account{
		number:   "245-901",
		balance:  100000,
		interest: 0.015,
	}

	// 포인터로 객체 생성
	// 생성자가 없는 객체를 인스턴스화
	// getter , setter와 비슷
	var lee *Account = new(Account)
	lee.number = "245-900"
	lee.balance = 100000
	lee.interest = 0.015

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	/*
		결과 :
		ex1 :  {245-901 100000 0.015}
		ex1 :  &{245-900 100000 0.015}
	*/

	// 마치 자바의 class를 인스턴스화 하는것과 비슷
	park := NewAccount("245-903", 1700000, 0.04)
	fmt.Println("ex2 : ", park)

	/*
		결과 :
		ex2 :  &{245-903 1.7e+06 0.04}
	*/

}
