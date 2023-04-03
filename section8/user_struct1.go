package main

import "fmt"

// 구조체
type Car struct {
	// 멤버 변수
	name  string
	color string
	price int64
	tax   int64
}

// 일반 메서드 방법. 객체 지향이라고 하기 힘듬
// 구조체는 메서드 바인딩을 통해 기능 (메서드) 를 정의한다.
// Go는 메서드를 함수로 정의한다.
// 메게변수가 구조체로 들어오게 된다. ( 여기선 Car 구조체가 메게변수로 들어옴 )
func Price(c Car) int64 {
	return c.price + c.tax
}

// 객체 지향의 방법
// 구조체와 메서드를 바인딩할때 쓰는 방식
// 메게변수가 먼저 오고 , 함수 이름이 그다음 온다. 그리고 리턴타입
// 구조체가 해당 메서드에 자동으로 주입됨.
func (c Car) Price() int64 {
	return c.price + c.tax
}

func main() {

	// Car 구조체의 멤버 변수에 접근하여 bmw와 benz 변수를 선언한다.
	bmw := Car{
		name:  "520d",
		price: 5000,
		color: "white",
		tax:   500,
	}

	benz := Car{
		name:  "220d",
		price: 6000,
		color: "white",
		tax:   600,
	}

	hyundai := Car{
		name:  "zenesis",
		price: 1500,
	}

	fmt.Println("ex1 : ", bmw, &bmw)
	fmt.Println("ex1 : ", benz, &benz)

	fmt.Println()
	fmt.Println("ex2 : ", Price(bmw))
	fmt.Println("ex2 : ", Price(benz))

	fmt.Println()
	// 객체지향은 다음과 같이 사용합니다.
	fmt.Println("ex3 : ", bmw.Price())
	fmt.Println("ex3 : ", benz.Price())

	fmt.Println()
	// 정의안된 필드는 기본값 리턴
	fmt.Println("ex4: ", hyundai.Price())
	/*
		결과 :
		ex1 :  {520d white 5000 500} &{520d white 5000 500}
		ex1 :  {220d white 6000 600} &{220d white 6000 600}

		ex2 :  5500
		ex2 :  6600

		ex3 :  5500
		ex3 :  6600
	*/

	// false. 하나의 구조체로 여러 객체를 생성하였기 때문에 메모리 공간이 다름
	fmt.Println("ex4 : ", &bmw == &benz)

}
