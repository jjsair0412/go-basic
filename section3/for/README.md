# for문
go는 반복문이 for밖에 없습니다.
- do while , while 등이 없다.

go에서 유일하게 제공되는 반복문 이기에 다양한 사용법을 숙지하는것이 중요합니다.

golang에서 후치 연산자는 반환값이 없기 때문에 , 변수에 담지 못합니다.
- [golang 후치 연산자 특이점](#golang-후치-연산자-특이점)

## golang for 사용법
기본적인 사용법은 타 언어와 같습니다.
```golang
func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}
}
```

## golang for문을 사용할 때 에러 발생 경우들
if문과 같은 이유로 , 아래와 같은 코드는 에러가 발생합니다.
```golang
func main() {

    // ; 를 끝에 자동으로 찍히기에
	for i := 0; i<5; i++
	{ // 얘가 뭔지 모름
		fmt.Println("ex1 : ", i)
	}

    // 자동 중괄호 대입 불가 { }
	for i := 0; i<5; i++
		fmt.Println("ex1 : ", i)
}
```

## golang 무한루프
golang에서 무한루프는 for문 조건 자체를 빼버리면 됩니다.
```golang
func main() {

	for {
		fmt.Println(" hello golang ")
		fmt.Println(" infinite loop ")
	}
}
```

## golang range 예약어
golang에서 python처럼 배열의 range 함수를 이용해서 , 배열을 순회할 수 있습니다.
```golang
func main(){
    // python처럼 range 가능
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}

    /*
    결과 
    ex3 :  0 Seoul
    ex3 :  1 Busan
    ex3 :  2 Incheon
    */
}
```

range 함수의 순서는 , 
1. 배열 index 값 
2. 실제 value
입니다.

index값을 생략하면서 변수에 value만 초기화 시키고 싶다면 , **_** 를 사용하면 됩니다.
```golang
func main(){
    // python처럼 range 가능
	loc := []string{"Seoul", "Busan", "Incheon"}
	for _, name := range loc {
		fmt.Println("ex3 : ", name)
	}

    /*
    결과 
    ex3 :  Seoul
    ex3 :  Busan
    ex3 :  Incheon
    */
}
```

## for문의 다양한 방법
1부터 100까지 더하는 for문은 , 아래처럼 두가지로 작성할 수 있습니다.

해외 코드는 예제 2번처럼 작성하는 경우가 더 많습니다.
```golang
func main() {
	//예제1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}

	fmt.Println("ex1 : ", sum1)

	//예제2
	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
	}

	fmt.Println("ex2 : ", sum2)

}
```

## golang 후치 연산자 특이점
이때 예제 2번에서 특이사항이 있는데 , **golang에서 후치 연산자 는 반환 값이 없기 때문에 변수에 담지 못합니다.**
```golang
func main(){
    sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
	    j := i++ // i++ 은 후치연산 . golang에서는 반환값 없어서 변수에 담지 못한다.
	}

	fmt.Println("ex2 : ", sum2)
}
```

## while문과 비슷하게 for문을 쓰는 방법
for 무한루프문 내부에 if 조건 break를 달아서 , while문과 비슷하게 사용할 수 있습니다.
- i가 100일때 break가 걸리면서 무한루프문 (가장 가까운 반복문) 탈출
```golang
func main(){
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3 : ", sum3)
}
```

## 두 개 변수 이상을 통해 for 증감식 활용
for문 내부 변수 여러개를 짧은 선언으로 초기화하여 증감식을 사용할 수 도 있습니다.
```golang
func main(){
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex 4 : ", i, j)
	}
    /*
    결과 
    ex 4 :  0 0
    ex 4 :  1 10
    ex 4 :  2 20
    ex 4 :  3 30
    ex 4 :  4 40
    ex 4 :  5 50
    ex 4 :  6 60
    ex 4 :  7 70
    ex 4 :  8 80
    ex 4 :  9 90
    ex 4 :  10 100
    */
}
```

해당 패턴으로 작성할 때 , 주의할 점은 후치 연산자가 섞이면 안된다는 것 입니다.

후치연산자는 반환 값이 없기 때문에 에러가 발생합니다.
- i++ 에 반환값이 없기에 에러 발생
```golang
func main(){
    for i, j := 0, 0; i <= 10; i++, j +=10 {
		fmt.Println("ex 4 : ", i, j)
	}
}
```

## Loop 함수
다중 for문에서 , 내부적으로 라벨을 사용할 수 있습니다.

다중 for문에서, **라벨(Label)**을 사용할 수 있습니다. 

만약 break만 사용한다면 가장 가까운 for문만 탈출하고, 바깥쪽 for문은 그대로 계속 실행됩니다. 하지만 라벨과 함께 break를 사용하면, 해당 라벨이 붙은 반복문을 완전히 종료할 수 있습니다.
- 해당 예제에서 ```Loop1``` 은 함수명이 아닌 **라벨(Label)** 입니다.

결과값을 확인해 보면 , i가 2 , j가 4일 때 for문 전체에서 탈출하는것을 확인할 수 있습니다.

**Loop label 밑에 관련없는 소스코드 ( println ) 이 있으면 에러 발생합니다.**


```golang
func main(){
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex1 : ", i, j)
		}
	}

/*
결과 
ex1 :  0 0
ex1 :  0 1
ex1 :  0 2
ex1 :  0 3
ex1 :  0 4
ex1 :  1 0
ex1 :  1 1
ex1 :  1 2
ex1 :  1 3
ex1 :  1 4
ex1 :  2 0
ex1 :  2 1
ex1 :  2 2
ex1 :  2 3

2 4 이상 안감
*/
}
```

## continue 함수
for문 내부에서 continue를 통해 조건에 맞으면 스킵할 수 있습니다.
```golang
func main(){
    for i := 0; i < 10; i++ {
		if i%2 == 0 { // 짝수일경우 skip
			continue
		}
		fmt.Println("ex2 : ", i)
	}
/*
결과:
ex2 :  1
ex2 :  3
ex2 :  5
ex2 :  7
ex2 :  9
*/
}

```

## continue와 Loop
Loop와 continue를 사용해서 , break에는 아예 for문을 탈출했다면 , 특정 값을 skip하고 다음 값이 작동되도록 할 수 있습니다.

예제에서는 i가 1, j가 2일 경우를 넘어가고 다음 값 인 i가 2 , j가 0 을 실행합니다.
```golang
func main(){
Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3 : ", i, j)
		}
	}

/*
결과
ex3 :  0 0
ex3 :  0 1
ex3 :  0 2
ex3 :  1 0
ex3 :  1 1
ex3 :  2 0
ex3 :  2 1
ex3 :  2 2
*/
}

```