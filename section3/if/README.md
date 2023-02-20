# 제어문 ( if )
## if문
if문은 반드시 boolean type으로 검사해야 합니다.
- **go에서는 반드시 if문을 true & false 로 검사해야 합니다.**
- 1, 0으로는 안됨 ( 자동 형변환 불가 )
- 소괄호는 사용하지 않습니다.

아래처럼 사용합니다.

```golang
func main() {
	// 제어문 ( 조건문 )

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15 이상입니다.")
	}

	if b >= 25 {
		fmt.Println("25 이상입니다.")
	}
}
```
### if문 error list
go는 ;를 사용하지 않는 대신에 들여쓰기가 끝나는 시점에 ; 를 자동으로 넣어줍니다.

따라서 아래처럼 { } 를 if문 밖에 사용하면 에러가 발생합니다.
```golang
func main() {
    var a int = 20
	b := 20

	// 에러 발생
	if b >= 25
	{

	}
}
```

go는 괄호를 자동으로 생성해주지 않기 때문에 꼭 { } 를 사용해야 합니다.
```golang
func main() {
    var a int = 20
	b := 20

	// 에러 발생
	if b >= 25
		fmt.Println("25 이상")
}
```

golang은 자동 형변환을 사용할 수 없기 때문에 if문 조건에 1, 0 을 사용할 수 없으며 , 꼭 true false로 비교해야 합니다.
```golang
func main() {
    var a int = 20
	b := 20

	// 에러 발생
	if c := 1; c {
        fmt.println("TRUE")
    }
}
```

## if문과 짧은 선언
if문 내부에 비교하는 값을 짧은 선언으로 변수를 초기화 한 후 , 사용할 수 있습니다.
```golang
func main() {
    	    
    if c := 40; c >= 35 {
		fmt.Println("35 이상")
	}

}
```

## if - else 문
golang에서의 if else 문은 다음과 같이 사용합니다.
```golang
func main() {
	var a int = 50
	b := 70

	// 예제 1
	if a >= 65 {
		fmt.Println("65 이상")
	} else {
		fmt.Println("65 미만")
	}

	// 예제 2
	if b >= 70 {
		fmt.Println("70 이상")
	} else {
		fmt.Println("70 미만")
	}
}
```

## 다중 if문
else if 문은 다음과 같이 작성합니다.
```golang
func main() {
	i := 100

	// if - else if 예제(1)
	if i >= 120 {
		fmt.Println("120 이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100 이상 120 미만")
	} else if i < 100 && i >= 50 {
		fmt.Println("50 이상 100 미만")
	} else {
		fmt.Println("50 미만")
	}

}
```