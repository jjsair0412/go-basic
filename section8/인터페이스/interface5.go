package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex1 : ", s)
}

func main() {
	dog := Dog{"Poll", 10}
	cat := Cat{"bob", 5}

	// 그냥 구조체 타입으로 인터페이스가 변환되어서 구조체 값을 출력합니다.
	printValue(dog)
	printValue(cat)      // struct
	printValue(15)       // int
	printValue("animal") // string
	printValue(25.5)     // float
	printValue([]Dog{})  // 슬라이스
	printValue([5]Dog{}) // 배열
	printValue(true)     // bool
	/*
	   결과 :
	   ex1 :  {Poll 10}
	   ex1 :  {bob 5}
	   ex1 :  15
	   ex1 :  animal
	   ex1 :  25.5
	   ex1 :  []
	   ex1 :  [{ 0} { 0} { 0} { 0} { 0}]
	*/
}
