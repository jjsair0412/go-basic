package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// golang에서 다음과 같이 경로를 작성하면 escape 문이 겹쳐서 , 에러난다.
	// var str1 string = "c:\go_study\src\"

	// 아래 경로처럼 작성해야 아래 경로가 됩니다.
	// c:\go_study\src\
	var str1 string = "c:\\go_study\\src\\"

	str2 := `c:\go_study\src\`

	fmt.Println("str1 : ", str1)
	fmt.Println("str2 : ", str2)
	/*
	   결과 :
	   str1 :  c:\go_study\src\
	   str2 :  c:\go_study\src\
	*/

	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요."

	fmt.Println("str3 : ", str3)
	fmt.Println("str4 : ", str4)

	fmt.Println("str3 크기 : ", len(str3))
	fmt.Println("str4 크기 : ", len(str4))

	fmt.Println("str3 길이 : ", utf8.RuneCountInString(str3))
	fmt.Println("str4 길이 : ", utf8.RuneCountInString(str4))
	fmt.Println("str4 길이 : ", len([]rune(str4)))
}
