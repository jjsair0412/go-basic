package main

import (
	"fmt"
)

func main() {

	// 정수형 문자

	// 예제 1번
	// 아스키코드 ( 영문 )
	// uint8 = byte는 같다.
	var char1 uint8 = 72  // 10진수
	var char2 byte = 0110 // 8진수
	var char3 byte = 0x48 // 16진수

	// 예제 2번
	// 유니코드 ( 한글 )
	var char4 rune = 50556   // 10진수
	var char5 rune = 0142574 // 44032(8진수)
	var char6 rune = 0xc57c  // 44032(16진수)

	// %c : 문자형으로 출력
	fmt.Printf("%c %c %c\n", char1, char2, char3)
	// %d : 숫자형으로 출력
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	// %o : 8진수 , %x : 16진수
	fmt.Printf("%d %o %x\n", char1, char2, char3)

	fmt.Printf("%c %c %c\n", char4, char5, char6)
	fmt.Printf("%d %d %d\n", char4, char5, char6)
	fmt.Printf("%d %o %x\n", char4, char5, char6)

	/*
		결과 :
		H H H
		72 72 72
		72 110 48
		야 야 야
		50556 50556 50556
		50556 142574 c57c
	*/
}
