package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

// (임원도 직원이다.)
type Executives struct {
	Employee     // 구조체를 멤버변수로 받음 -> is a 관계라 한다.
	specialBonus float64
}

func (e Employee) Caculate() float64 {
	return e.salary + e.bonus
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
	/*
	   결과 :
	   ex1 :  2150000
	   ex1 :  350000
	*/

	//Employee 구조체 통해서 메소드 호출
	// 상위 구조체 리시버를 그냥 접근할 수 있습니다.
	fmt.Println("ex1 : ", int(ex.Caculate())+int(ex.specialBonus))
	/*
		결과 :
		ex1 :  7000000
	*/
}
