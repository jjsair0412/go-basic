package main

import "fmt"

func main() {

	// 익명함수는 아래처럼 사용합니다.
	// js의 익명함수와 동일하게 이름이 없고 , 중괄호가 닫힌 뒤 () 소 괄호를 붙여야합니다.
	func() {
		fmt.Println("ex1 : anonymous func!")
	}()
	// 결과 : ex1 : anonymous func!

	// js처럼 매개변수를 바로 넣기 위해 뒤에 () 소괄호 안에 변수 넣습니다.
	msg := "hello golang"
	func(n string) {
		fmt.Println("ex2 : ", n)
	}(msg)
	// 결과 : ex2 :  hello golang

	// 매개변수를 바로 넣으면서 함수를 사용해도 무관합니다.
	func(x, y int) {
		fmt.Println("ex3 : ", x+y)
	}(10, 20)
	// 결과 : ex3 :  30

	// 실행과 동시에 짧은 선언으로 r 변수 안에 func 리턴값이 들어갑니다.
	r := func(x, y int) int {
		return x * y
	}(10, 100)

	fmt.Println("ex4 : ", r)

	// 결과 : ex4 :  1000
}
