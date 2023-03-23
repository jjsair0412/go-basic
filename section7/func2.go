package main

import "fmt"

// 매개변수로 int형 i와 , 매개변수로 int , int 타입 두개를 가지는 함수를 가진다.
func sum(i int, f func(int, int)) {
	f(i, 10)
}

// 매개변수 타입이 동일할땐 , 하나를 지우고 뒤에만써도 된다.
func add(x, y int) {
	fmt.Println("ex1 : ", x+y)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i *= 10 // *i = *i * 10 과 동일
}

func main() {
	sum(100, add)

	a := 100
	multi_value(a)
	fmt.Println(a) // 실제 변수에 영향 x (영향주려면 포인터로 주소값 참조해야됨)

	multi_reference(&a)
	fmt.Println(a) // 실제 변수에 영향 o (실제주소값 참조) . 결과 : 1000
}
