package main

import (
	"fmt"
)

func main() {
	var str1 string = "Golang"
	var str2 string = "world"

	fmt.Println("ex1 : ", str1[0:3], str1[0]) // str1의 0~2번째 까지 추출 , 0번째 추출
	fmt.Println("ex1 : ", str2[3:], str2[0])  // 빈칸은 끝까지라는 의미. str2의0부터 끝까지 , 0번째 추출
	fmt.Println("ex1 : ", str2[:4])           // 처음부터 4번째까지 출력
	fmt.Println("ex1 : ", str1[1:3])          // 1번 인덱스부터 3번인덱스까지 출력
}
