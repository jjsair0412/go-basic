package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Instructions for downloading and installing Go." +
		"A brief Hello, World tutorial to get started. Learn a bit about Go code, tools, packages, and modules." +
		"Tutorial: Create a module"
	str2 := "A tutorial of short topics introducing functions, error handling, arrays, maps, unit testing, and compiling."

	fmt.Println("ex1 : ", str1+str2)

	strSet := []string{} // 슬라이스형 선언

	// strSet에 append로 추가할 문자열 하나씩 넣어줌,
	// 파라미터는 슬라이스형의 string 변수명과 합칠문자 순서대로 들어감
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	// 사용할 땐 , strings.Join() 메서드 사용
	// 파라미터로 슬라이스형 변수와 , 각 문자열 중간중간에 넣을 문자를 넣어줌.
	// 넣을 문자는 빈칸으로 둬도 상관없음.
	fmt.Println("ex2 : ", strings.Join(strSet, "---"))
}
