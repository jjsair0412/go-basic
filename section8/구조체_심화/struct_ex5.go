package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

type Executives struct {
	Employee
	specialBonus float64
}

func (e Employee) Caculate() float64 {
	return e.salary + e.bonus
}

// salary와 bonus는 최상위 구조체인 Employee의 멤버변수 인데,
// Employee 객체에 오버라이딩되어 접근합니다. -> 재정의
// 리시버의 이름은 동일해야 합니다.
func (e Executives) Caculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

func main() {
	ep1 := Employee{
		"kim",
		2000000,
		150000,
	}

	ep2 := Employee{
		"park",
		150000,
		200000,
	}

	ex := Executives{
		Employee{
			"lee",
			5000000,
			1000000,
		},
		1000000,
	}
	fmt.Println("ex1 : ", int(ep1.Caculate()))
	fmt.Println("ex1 : ", int(ep2.Caculate()))

	fmt.Println("오버라이딩 시작 ! ")

	fmt.Println("ex1 : ", int(ex.Caculate()))

	/*
		결과 :
		ex1 :  2150000
		ex1 :  350000

		오버라이딩 시작 !
		메서드가 오버라이딩 되어 , 계산됩니다.
		ex1 :  7000000
	*/
	fmt.Println()

}
