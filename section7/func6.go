package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {

	// 0번 인덱스에 multiply , 1번 인덱스에 sum을 넣습니다.
	// func 타입의 값을 가진 슬라이스를 선언하고 , int형 두개를 매개변수로
	// 가지며 return type이 int형 한개를 가지는 애들을 묶을 수 있습니다.
	f := []func(int, int) int{multiply, sum}

	// 0번째 , 1번째 함수에 10. 10을 매개변수로 넣음
	a := f[0](10, 10) // a := multiply(10,10) 과 동일
	b := f[1](10, 10) // b := sum(10,10) 과 동일

	fmt.Println("ex1 : ", a, f[0](10, 10))
	fmt.Println("ex1 : ", b, f[1](10, 10))
	/*
	   결과 :
	   ex1 :  100 100
	   ex1 :  20 20
	*/

	fmt.Println()

	// 변수에 할당하는 방법입니다.
	// int형 매개변수 두개와 , return type을 int로 갖는 함수인 multiply를
	// f1에 넣습니다.
	var f1 func(int, int) int = multiply
	f2 := sum // 이렇게 선언해도 무관합니다.

	fmt.Println("ex2 : ", f1(10, 10))
	fmt.Println("ex2 : ", f2(10, 10))

	fmt.Println()

	// map에 할당하는 방법입니다.
	// key로 string을 갖고 , value로 매개변수 int 두개를 갖고 return값이 int 하나인
	// 함수를 가지는 맵을 선언합니다.
	m := map[string]func(int, int) int{
		"mul_func": multiply,
		"sum_func": sum,
	}

	// map의 key를 통해서 함수를 꺼내고 매개변수를 넣어서 사용합니다.
	fmt.Println("ex3 : ", m["mul_func"](10, 10))
	fmt.Println("ex3 : ", m["sum_func"](10, 10))
	/*
	   결과 :
	   ex3 :  100
	   ex3 :  20
	*/

}
