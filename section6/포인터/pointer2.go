package main

import "fmt"

func main() {

	i := 7
	p := &i // i의 주소값이 p에 들어감

	fmt.Println("ex 1 : ", i, *p, &i, p)
	// 결과 : ex 1 :  7 7 0xc0000160a8 0xc0000160a8

	*p++ // 1 증가
	fmt.Println("ex 1 : ", i, *p, &i, p)
	// 결과 : ex 1 :  8 8 0xc00009e058 0xc00009e058

	*p = 77777 // 포인터 결과 역 참조값 변경
	fmt.Println("ex 1 : ", i, *p, &i, p)
	// 결과 : ex 1 :  77777 77777 0xc0000160a8 0xc0000160a8

	i = 77 // 원본 변경
	fmt.Println("ex 1 : ", i, *p, &i, p)
	// 결과 : ex 1 :  77 77 0xc00009e058 0xc00009e058
}
