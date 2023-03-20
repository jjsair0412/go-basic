package main

import (
	"fmt"
)

func main() {
	// 변수 a의 포인터 형으로 할당
	var a *int            // 방법 1
	var b *int = new(int) // 방법 2 -> 정석

	fmt.Println(a)
	fmt.Println(b)
	/*
	   결과 :
	   <nil>
	   0xc0000160a8

	   new로 초기화시킨 변수는 메모리공간을 할당받습니다.
	*/

	i := 7

	a = &i // i변수 메모리 공간(주소값) 을 a에 전달
	b = &i

	fmt.Println("ex1 : ", a, &i)
	fmt.Println("ex1 : ", &a)
	fmt.Println("ex1 : ", *a)

	fmt.Println()

	fmt.Println("ex1 : ", b, &i)
	fmt.Println("ex1 : ", &b)
	fmt.Println("ex1 : ", *b)
	/*
		   결과 :
		   ex1 :  0xc00009e078 0xc00009e078
		   ex1 :  0xc0000c0018
		   ex1 :  7

		   ex1 :  0xc00009e078 0xc00009e078
		   ex1 :  0xc0000c0020
		   ex1 :  7

		   i의 주소값을 a와 b에 저장했기 때문에 출력시키면 주소값이 동일함.
			- 그냥 출력시키면 저장된 주소값이 출력됨

		   그러나 a와 b의 메모리 주소값은 다름. 왜냐면 i의 주소값을 저장할
		   공간이 또 필요하기 때문에 i의 주소값이 저장된 a, b의 주소가 나온거임.

		   *(에스터리스크) 를 사용해서 a와 b를 출력시키면 , 저장한 주소값이 가르키고있는 값을 역으로 참조합니다.
		   - 역참조
	*/

	var c = &i
	d := &i

	*d = 10 // d에 저장된 주소가 가르키는곳을 역참조 하여 10으로 초기화.

	fmt.Println("ex1 : ", c, &i)
	fmt.Println("ex1 : ", &c)
	fmt.Println("ex1 : ", *c) // 역참조

	fmt.Println()

	fmt.Println("ex1 : ", d, &i)
	fmt.Println("ex1 : ", &d)
	fmt.Println("ex1 : ", *d) // 역참조
	/*
	   결과 :
	   ex1 :  0xc0000160c8 0xc0000160c8
	   ex1 :  0xc00000a040
	   ex1 :  10

	   ex1 :  0xc0000160c8 0xc0000160c8
	   ex1 :  0xc00000a048
	   ex1 :  10

	   역참조해서 10으로 변경시켰기 때문에 , 가르키고있는곳이 c d 모두 동일하기에
	   다 변경됩니다.
	*/
}
