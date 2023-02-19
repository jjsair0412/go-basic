package main

import "fmt"

func main() {

	// 열거형
	// 상수를 사용하며 일정 규칙을 가지고 , 숫자가 증가된다.
	const (
		Jan = 1
		Feb = 2
		Mar = 3
		Apr = 4
		May = 5
		Jun = 6
	)

	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)

}
