# 포인터
## go에서의 포인터
Golang에서는 포인터를 지원합니다.

포인터를 왜 사용할까 ?
- 변수의 지역성
- 함수 외부에서 사용한 변수를 핸들링하기 위해서
- 연속된 메모리를 참조하기 위해서 등 ..

Golang은 포인터를 완전히 지원하지는 않습니다.
- 리스크가 있는 연산은 에러발생
- **주소의 값은 직접 변경이 불가능합니다.**
    - 잘못된 코딩으로 인한 버그를 방지
    - *(에스터리스크) 로 포인터를 사용합니다.
    - nill 로 초기화 됩니다. -> **nill은 0이 아닙니다.**

## 1. 포인터 사용법
golang의 포인터는 아래와 같이 사용할 수 있습니다.
```golang
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
```

## 2. golang의 포인터 기초
Golang에서 포인터는 아래 예제로 정리할 수 있습니다.

```golang
func main() {

	i := 7
	p := &i // i의 주소값이 p에 들어감

	fmt.Println("ex 1 : ", i, *p, &i, p)
	// 결과 : ex 1 :  7 7 0xc0000160a8 0xc0000160a8

	*p++ // 1 증가
	fmt.Println("ex 1 : ", i, *p, &i, p)
	// 결과 : ex 1 :  8 8 0xc00009e058 0xc00009e058

	*p = 77777 // 포인터 결과 역 참조값 변경
	fmt.Println("ex 1 : ", i, *p, &i, p)
	// 결과 : ex 1 :  77777 77777 0xc0000160a8 0xc0000160a8

	i = 77 // 원본 변경
	fmt.Println("ex 1 : ", i, *p, &i, p)
	// 결과 : ex 1 :  77 77 0xc00009e058 0xc00009e058
}
```

## 3. 함수에서의 포인터
golang에서 왜 포인터를 사용할까 ?

아래 예제에서 답을 찾을 수 있다.
- 함수에서 역 참조로 실제 변수를 참조하여 값을 변경하기 위해
```golang
func main() {
	// 포인터 값 전달
	// 함수나 메서드 호출 시에 매게변수 값을 복사 전달
	// 함수나 메서드 내에서는 원본 값을 변경할 수 없습니다.
	// 주로 원본 값을 변경하기 위해 포인터로 전달합니다.
	// 특히 크기가 큰 배열일 경우 값 복사시에 시스템에 부담을 주기 때문에 ,
	// 포인터로 전달합니다. -> 슬라이스나 맵 참조 전달

	var a int = 10
	var b int = 10

	fmt.Println("ex1 : ", a, b)

	fmt.Println()

	rptc(&a) // 원본 값 변경 ( a 주소값 전달하였기 때문에 역참조로 실제 a 주소 참조하여 값 변경됨 )
	vptc(b)  // 원본 값 유지 ( b 안바뀜 . )

	fmt.Println("ex1 : ", a, b)

	/*
	   결과 :
	   ex1 :  10 10

	   ex1 :  77 10

	   rptc() 함수에서 실제 a 변수의 주소값을 받아와서 , 파라미터로 받은 변수의 주소값을
	   역 참조하여 실제 값을 변경시켯기 때문에 rptc() 함수는 실제 변수인 a를 변경시켰으며,
	   vptc()는 함수가 실제 변수를 참조할 수 없기 때문에 함수가 종료되면 그냥 소멸됩니다.
	*/

}
```