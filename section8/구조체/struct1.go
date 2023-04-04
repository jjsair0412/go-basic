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
	kim := Account{
		number:   "245-901",
		balance:  1000000,
		interest: 0.015,
	}

	lee := Account{
		number:  "245-902",
		balance: 1200000,
	}

	park := Account{
		number:   "245-903",
		interest: 0.025,
	}

	// 필드 순서대로 그냥 입력해줘도 괜찮음
	cho := Account{"245-904", 1500000, 0.03}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println("ex1 : ", park)
	fmt.Println("ex1 : ", cho)
	/*
		결과 :
		ex1 :  {245-901 1e+06 0.015}
		ex1 :  {245-902 1.2e+06 0}
		ex1 :  {245-903 0 0.025}
		ex1 :  {245-904 1.5e+06 0.03}
	*/

	fmt.Println()

	fmt.Println("ex2 : ", int(kim.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))
	fmt.Println("ex2 : ", int(park.Calculate()))
	fmt.Println("ex2 : ", int(cho.Calculate()))
	/*
	   결과 :
	   ex2 :  1015000
	   ex2 :  1200000
	   ex2 :  0
	   ex2 :  1545000
	*/

	fmt.Println()
}
