package main

import (
	"fmt"
	"strconv"
)

// 함수는 어디에 선언하든 상관없다.
func helloGolang() {
	fmt.Println("ex1 : Hello Golang")
}

func main() {
	helloGolang()

	helloGolang2()

	say_one("hello world")

	result := sum(5, 5)

	fmt.Println(result)

	fmt.Println(sum(5, 5))

	fmt.Println("int -> string ex3 : ", strconv.Itoa(sum(3, 3)))
}

// 매개변수 , return 둘다 없을경우
func helloGolang2() {
	fmt.Println("ex1 : Hello Golang2")
}

// 매개변수만 있을경우
func say_one(m string) {
	fmt.Println("ex2 : ", m)
}

// 매개변수 , return 둘다 있을경우
// return type의 type을 매개변수 괄호 닫은 뒤 작성해줍니다.
func sum(x int, y int) int {
	return x + y
}
