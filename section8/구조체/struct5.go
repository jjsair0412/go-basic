package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	name    string "차량명" // 필드 태그 생성
	color   string "차량 색상"
	company string "재조사"
}

func main() {

	tag := reflect.TypeOf(Car{})

	// tag 필드 갯수만큼 for문 돌면서 태그 뽑아옴
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("ex1 : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
		/*
			결과 :
			태그이름 , 필드명 , 필드 타입 꺼내옴
			ex1 :  차량명 name string
			ex1 :  차량 색상 color string
			ex1 :  재조사 company string
		*/
	}
}
