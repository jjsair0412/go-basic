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

// bite 메서드 구현
func (d Dog) bite() {
	fmt.Println("Dog ", d.name, "bite !!")
}

func (c Cat) bite() {
	fmt.Println("Cat ", c.name, "bite! !")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {
	var inter1 Behavior

	dog1 := Dog{"poll", 10}
	inter1 = dog1
	inter1.bite()

	cat1 := Cat{"jin", 30}
	inter1 = cat1
	inter1.bite()

	fmt.Println()
	// interface를 선언하지 않고 , Behavior 메서드를 구현한 객체가 무엇인지를 넣어주는
	// 아래 방식을 더 많이 사용합니다.
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)
	inter2.bite()

	fmt.Println()
	// interface with slice
	inters := []Behavior{dog1, dog2}
	for idx, _ := range inters {
		inters[idx].bite()
	}

	fmt.Println()
	// 값 형태로 실행(인터페이스)
	for _, val := range inters {
		// dog1.bite() , dog2.bite() 실행
		val.bite()
	}
}
