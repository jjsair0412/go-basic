package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 구조체 Dog 메소드 구현
func (d Dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

// 구조체 Cat 메소드 구현
func (d Cat) run() {
	fmt.Println(d.name, " : Cat is running!")
}

// interface를 따로 두지 않고 , 메게변수로 넣으면서 익명 선언을 사용합니다.
func act(animal interface{ run() }) { // 익명 인터페이스. 타입 정의 안함
	animal.run()
}

func main() {
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	// 개 행동 실행
	act(dog)

	// 고양이 행동 실행
	act(cat)

	/*
	   결과  :
	   poll  : Dog is running!
	   bob  : Cat is running!
	*/
}
