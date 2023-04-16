package main

import "fmt"

// interface를 사용할 땐 , 구조체랑 같은 형식으로 만들지만,
// 이름을 interface로 둡니다.
type test interface{} // 빈 인터페이스

//인터페이스를 선언하는 방법
/*
type 인터페이스명 interface{
	메서드1() 반환값(타입 형)
	메서드2() // 반환값 없을 경우
}
*/

func main() {
	var t test
	fmt.Println("ex1 : ", t) // 빈(empty) 인터페이스인 경우 , nil 리턴
}
