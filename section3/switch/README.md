# Golang switch문
## switch 문
조건문의 일종입니다.
- go에선 if문처럼 switch case를 굉장히 유연하게 사용 가능 합니다.

Switch 뒤 표현식(Expression) 생략 가능

case 뒤 표현식(Expression) 사용 가능

자동 break 때문에 fallthrouth가 존재합니다.
- go의 switch문은 break를 작성하지 않아도 , 조건에 맞는다면 break가 걸리고 아래 case는 실행되지 않습니다.

Type으로 분기가 가능합니다.
- 값이 아닌 , 변수 타입으로 분기가 가능합니다. ( int , string .. )

## switch의 다양한사용법
switch case문은 아래와 같이 사용합니다.


```golang
func main() {
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}
}
```

switch 안에서 짧은 선언을 사용할 수 도 있습니다.
```golang
func main() {
	switch b := 27; {
	case b < 0:
		fmt.Println(a, "는 음수")
	case b == 0:
		fmt.Println(a, "는 0")
	case b > 0:
		fmt.Println(a, "는 양수")
	}
}
```

switch 안에서 표현식을 자동 선언해 줄 수 있습니다.
- 아래 예제에서 , 짧은 선언으로 c를 "go"로 초기화 한 후 , 뒤에 한번 더 작성함으로써 , case문의 조건을 **같다** 로 통일시킬 수 있습니다.
- 또한 switch 내부에서 문자비교도 가능합니다.
```golang
func main() {
	switch c := "go"; c {
	case "go":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치하는 값 없음")
	}
}
```

switch 내부에서 , 연산자 ( + - * / ) 를 사용할 수 있습니다.
- 아래 예시처럼 **go 문자열을 짧은 선언으로 초기화** 한 후 , **lang을 더해서 같다면 이라는 조건으로 통일**시킬 수 있습니다.
```golang
func main(){
    switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("일치하는 값 없음")
	}
}
```

switch 내부에서 짧은 선언 변수를 여러개 초기화할 수 있습니다.
```golang
func main(){
    switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i와 j는 같다")
	case i > j:
		fmt.Println("i가 j보다 크다")
	}
}
```

아래처럼 switch 내부에 함수를 사용하고 , case문 내부에서 if 조건문처럼 조건문을 사용할 수 있습니다.
```golang
func main() {
	//예제 1
	rand.Seed(time.Now().UnixNano()) // 현재 시간에 맞게 랜덤값 생성
	switch i := rand.Intn(100); {    // 0 ~ 100 까지 랜덤값 생성
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " 50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " 25 이상 50 미만")
	default:
		fmt.Println("i -> ", i, " 기본 값")
	}
}
```

switch case문에 여러가지 값을 나열해서 비교할 수 도 있습니다.
```golang
func main() {
	a := 30 / 15

	switch a {
	case 2, 4, 6: // a가 2 , 4 , 6일 경우 true
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: // a가 1 , 3 , 5일 경우 true
		fmt.Println("a ->", a, "는 홀수")
	}

}
```

if문과 조합도 가능합니다.
```golang
func main() {
    a := 300 / 2

    switch a {
	case 2, 4, 6: // a가 2 , 4 , 6일 경우 true
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: // a가 1 , 3 , 5일 경우 true
		if a >= 20 {
			fmt.Println("a 가 20 이상입니다. ")
		}
		fmt.Println("a ->", a, "는 홀수")
	}
}
```

## fallthrough
golang은 자동 break 때문에 , 조건에 맞으면 case 내부로 들어가고 , 아래 case들에는 들어가지 않습니다.

그러나 fallthrough 예약어를 사용하면 , 아래 case도 실행되게 됩니다.
- java에서 break 안쓴것처럼 작동
- 마지막 case에 fallthrough를 쓰는것은 의미가 없기 때문에 에러가 발생합니다.
```golang
func main() {
	//예제 2
	switch e := "go"; e {
	case "java":
		fmt.Println("Java")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
		fallthrough
	case "c":
		fmt.Println("c")
		// fallthrough
	}

    /*
    결과 : 
    go
    python
    ruby
    c
    */
}

```