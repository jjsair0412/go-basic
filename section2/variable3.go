package main

import "fmt"

func main() {
	// 짧은 선언

	// := 으로 사용하고 , 타입을 지정하지 않습니다.
	shortVar1 := 3
	shortVar2 := "TEST"
	shortVar3 := false

	// 초기화를 한번 더하면 에러가 발생합니다.
	// no new variables on left side of :=
	// shortVar1 := 10

	fmt.Println("shortVar1 = ", shortVar1, " shortVar2 = ", shortVar2, " shortVar3 = ", shortVar3)

	// EXAMPLE
	// 짧은 선언을 이용해서 , if문안에서만 사용할 수 있는 변수를 선언하고 , if가 끝나면 사라지는 변수를 만들 수 있습니다.
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success")
	}

	// undefined: i 에러 발생. 짧은 선언으로 if안에서만 사용
	// fmt.Println("i가 존재하나요 ? : ", i)
}
