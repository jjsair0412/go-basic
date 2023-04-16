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
func (d Dog) bite() {
	fmt.Println(d.name, " : Dog bite!")
}

func (d Dog) sounds() {
	fmt.Println(d.name, " : Dog barks!")
}

func (d Dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

// 구조체 Cat 메소드 구현
func (d Cat) bite() {
	fmt.Println(d.name, " : Cat 할퀴다!")
}

func (d Cat) sounds() {
	fmt.Println(d.name, " : Cat cries!")
}

func (d Cat) run() {
	fmt.Println(d.name, " : Cat is running!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
	sounds()
	run()
}

// 인터페이스를 파라미터로 받는다.
// act 함수를 호출할 때 , interface Behavior를 구현한 메서드를 리시버로 갖고 있는
// 구조체가 들어가면 됩니다.
func act(animal Behavior) {
	animal.bite()
	animal.sounds()
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
	   결과 :
	   poll  : Dog bite!
	   poll  : Dog barks!
	   poll  : Dog is running!
	   bob  : Cat 할퀴다!
	   bob  : Cat cries!
	   bob  : Cat is running!
	*/
}
