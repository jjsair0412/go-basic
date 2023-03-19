package main

import (
	"fmt"
)

func main() {

	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1, ok1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]

	fmt.Println("ex1 : ", value1, ok1)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, ok)
	/*
		결과 :

		없는 key를 리턴시켰지만 에러가 발생하지 않고 , 기본 자료형의 default값을 넣어줍니다.
		int : 0 , string : "", float : 0.0

		그런데 만약 위 예처럼 map의 value가 기본 자료형의 default 값이 들어가 있다면 , 이게 진짜
		값이 들어간건지 , 기본자료형이 들어간건지 구분하기 힘듭니다.

		따라서 , 선언할 때 두번째 리턴 값으로 키 존재의 유무를 확인하게 됩니다.

		ex1 :  0 true
		ex1 :  0
		ex1 :  0 false

		lemon은 있으니까 true를 반환하고 , kiwi는 map1에 없으니까 false를 반환합니다.
	*/

	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 :", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist! ")
	}
	/*
	   결과 :
	   ex2 : kiwi is not exist!

	*/

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 :", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist! ")
	}
	/*
	   결과 :
	   ex2 : ex2 : 115

	*/

	// key가 있는지 없는지만 확인
	if _, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : kiwi is not exist!")
	} else {
		fmt.Println("ex2 : kiwi is not exist! ")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exist! ")
	}

	/*
		결과 :
		ex2 : kiwi is not exist!
		ex2 : kiwi is not exist!
	*/
}
