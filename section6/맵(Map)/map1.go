package main

import (
	"fmt"
)

func main() {
	// key는 문자열 , 값은 int 사용하겠다 라고 map을 정의
	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)                // 자료형 생략
	map3 := make(map[string]int)                   // 축약형

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)

	fmt.Println()

	map4 := map[string]int{} // {}를 붙이면 형태만 Json 형태 로 정의할 수 있다.
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	// 아래처럼 json 형태로 key-value를 넣어서 map을 선언할 수 있다.
	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 23,
	}

	// 메모리용량 10으로 선언
	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	// map type에서는 순서가 없기 때문에 값 순서가 바뀔 수 있음
	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["apple"])
	fmt.Println("ex2 : ", map6["banana"])
	fmt.Println("ex2 : ", map6["orange"])

}
