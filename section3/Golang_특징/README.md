# GO 언어의 특징
## 1. Go의 특징 
go는 모호하거나 , 애매한 문법 자체를 금지했습니다.

### 1.1 연산자 관련
후치 연산자는 허용했지만 , 전치 연산자는 비 허용 합니다.

또한 증감 연산자는 반환값이 없습니다.
```golang
func main(){
	sum, i := 0 , 0

	for i <= 100 {
		sum += i++ // 반환값이 없어서 에러 발생 

    }
}
```

따라서 위 코드를 정상화 시키기 위해선 아래와 같이 코딩해야 합니다.

```golang
func main() {
	sum, i := 0, 0

	for i <= 100 {
		// sum += i++ // 반환값이 없어서 에러 발생

		sum += i
		i++
	}

}
```

### 1.2 Pointer 
Golang은 Pointer는 허용하지만 , 포인터를 사용한 연산은 비 허용 합니다.

### 1.3 주석
주석 사용법은 타 언어와 동일합니다.

### 1.4 세미콜론
golang은 컴파일러가 자동으로 끝에 세미콜론을 삽입합니다.

그래서 두 문장을 한 문장에 표현할 경우 , 명시적으로 세미콜론을 사용해야 합니다.
- 그러나 해당 방법은 권장 x

반복문 및 제어문 을 사용할 때 , 세미콜론을 주의 깊게 사용해야 합니다.
```golang
func main() {
    // 에러 발생
	for j := 10; j >= 0; j-- 
    {
		fmt.Println("ex2 : ", j)
	}
}
```

위 예제 코드가 컴파일되면 , 아래처럼 변하기 때문에 에러가 발생합니다.
```golang
func main() {
    // 에러 발생
	for j := 10; j >= 0; j--; // 문장 끝에 세미콜론 생성 . 에러 발생
    {
		fmt.Println("ex2 : ", j)
	}
}
```

### 1.5 go 코드 서식 지정
golang은 한 사람이 코딩한것 같은 일관성을 유지할 수 있게끔 도와줍니다.

golang은 IDE 툴을 사용한다면 , 코드 스타일을 유지시켜줍니다.

1. gofmt
gofmt 명령어를 통해서 코드 스타일을 유지시킵니다.
```bash
$ gofmt -h # 사용법

$ gofmt -w # 원본 파일에 반영
```

vscode같은 곳에서 golang 코딩을 하고 . 저장하면 사용안하는 외부 패키지나 변수가 지워지는 이유가 , 얘가 작동해서 코드 스타일을 통일시키기 때문입니다.

띄어쓰기나 들여쓰기도 통일시킵니다.

아래와 같은 golang code를 gofmt로 변환하게 되면 ,
```golang
package main

import "fmt"

func main(){
	for i:=0; i<=100; i++{
		fmt.Println("ex 1 : ", i)
	}
}
```

gofmt 명령어 수행
```bash
$ gofmt -w test.go
```

아래와 같이 일관성있게 코드 서식이 바뀝니다.
- 띄어쓰기 등 변경됨
```golang
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println("ex 1 : ", i)
	}
}
```