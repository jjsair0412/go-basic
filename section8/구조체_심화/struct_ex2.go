package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	kim := Account{
		"245-901",
		100000,
		0.015,
	}

	lee := Account{
		"245-901",
		100000,
		0.015,
	}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	/*
		결과 :
		ex1 :  {245-901 100000 0.015}
		ex1 :  {245-901 100000 0.015}
	*/

	fmt.Println()

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println("ex2 : ", int(kim.balance))
	fmt.Println("ex2 : ", int(lee.balance))
	/*
		결과 :
		ex2 :  100000
		ex2 :  101500
	*/
}
