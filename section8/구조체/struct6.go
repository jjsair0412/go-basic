package main

import "fmt"

// 패키지 외부에서 접근가능. 첫글자 대문자라서
type Car struct {
	name    string "차량명" // 필드 태그 생성
	color   string "차량 색상"
	company string "재조사"
	detail  spec   "차량 상세 제원" // 중첩 구조체 생성
}

// 패키지 외부에서 접근불가. 첫글자 소문자라서
type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	car1 := Car{
		"520d",
		"silver",
		"bmw",
		spec{
			4000,
			1000,
			2000,
		},
	}

	fmt.Println("ex1 : ", car1.name)
	fmt.Println("ex1 : ", car1.color)
	fmt.Println("ex1 : ", car1.company)
	fmt.Printf("ex1 : %#v\n", car1.detail)
	/*
	   결과 :
	   ex1 :  520d
	   ex1 :  silver
	   ex1 :  bmw
	   ex1 : main.spec{length:4000, height:1000, width:2000}
	*/

	fmt.Println()

	fmt.Println("ex2 : ", car1.detail.height)
	fmt.Println("ex2 : ", car1.detail.length)
	fmt.Println("ex2 : ", car1.detail.width)
	/*
		결과 :
		   ex2 :  1000
		   ex2 :  4000
		   ex2 :  2000
	*/
}
