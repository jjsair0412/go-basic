# Go 데이터 타입
## 1. Boolean ( 참 / 거짓 )
### 1.1 Bool type이란 ?
Boolean은 참과 거짓을 판단할 때 사용
- 조건부 논리 연산자와 주로 사용됨.
    - ! , || (or) , && (and)

**암묵적 형 변환이 안된다.**
- 0 , Nil (null과 동일) 은 false로 취급되지 않는다 !!!

```golang
func main() {
	// 암묵적 형 변환은 인정되지 않는다.
	b4 := 1

	if b4 {
		fmt.Println("ex 8 : TRUE ") 
        // 에러발생 . 1은 true가 아니다.
	}
}

```

## 1.2 boolean 예제
true와 false를 반환하는 예

```golang
package main

import (
	"fmt"
)

func main() {
	// 예제 1번
	fmt.Println("ex1 : ", true && true)
	fmt.Println("ex1 : ", true && false)
	fmt.Println("ex1 : ", false && false)
	fmt.Println("ex1 : ", true || true)
	fmt.Println("ex1 : ", true || false)
	fmt.Println("ex1 : ", false || false)
	fmt.Println("ex1 : ", !true)
	fmt.Println("ex1 : ", !false)

	// 예제 2번
	num1 := 15
	num2 := 37

	fmt.Println("ex2 : ", num1 < num2)
	fmt.Println("ex2 : ", num1 > num2)
	fmt.Println("ex2 : ", num1 <= num2)
	fmt.Println("ex2 : ", num1 <= num2)
	fmt.Println("ex2 : ", num1 != num2)
	fmt.Println("ex2 : ", num1 == num2)
	fmt.Println("ex2 : ", num1 != num2)
}
```

## 2. 숫자형 ( numeric )
### 2.1 golang에서 지원하는 숫자들
1. 정수 
    - 8진수(o)
    - 16진수(ox)
    - 10진수
2. 실수
3. 복소수
4. 32bit
5. 64bit
6. unsigned (양수)

를 지원합니다.

**tip**
golang은 데이터타입에 대해 엄격합니다.

만약 int형을 지정할 때 , 메모리소모를 그렇게까지 신경쓰지 않아도 되는 상황이라면 ,,, 아래처럼 그냥 int를 두고 코딩합니다..

엄격하다면 , int32 , int64를 사용하여 메모리 누수를 최소화 합니다.
```golang
package main

import (
	"fmt"
)

func main(){
	var num1 int = 17
	var num2 int = -68
}
```

### 2.2 예제
int형 선언 예제입니다.

```golang
package main

import (
	"fmt"
)

func main() {
	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println("ex1 : ", num1)
	fmt.Println("ex2 : ", num2)
	fmt.Println("ex3 : ", num3)
	fmt.Println("ex4 : ", num4)

	/*
		결과 :
		ex1 :  17
		ex2 :  -68
		ex3 :  409
		ex4 :  855256181
	*/

}
```

### 2.2 golang에서 출력 포멧 방식 - %d , %c 등
아래 예제처럼 출력시키면 됩니다.

```golang
package main

import (
	"fmt"
)

func main() {

	// 정수형 문자

	// 예제 1번
	// 아스키코드 ( 영문 )
	// uint8 = byte는 같다.
	var char1 uint8 = 72  // 10진수
	var char2 byte = 0110 // 8진수
	var char3 byte = 0x48 // 16진수

	// 예제 2번
	// 유니코드 ( 한글 )
	var char4 rune = 50556   // 10진수
	var char5 rune = 0142574 // 44032(8진수)
	var char6 rune = 0xc57c  // 44032(16진수)

	// %c : 문자형으로 출력
	fmt.Printf("%c %c %c\n", char1, char2, char3)
	// %d : 숫자형으로 출력
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	// %o : 8진수 , %x : 16진수
	fmt.Printf("%d %o %x\n", char1, char2, char3)

	fmt.Printf("%c %c %c\n", char4, char5, char6)
	fmt.Printf("%d %d %d\n", char4, char5, char6)
	fmt.Printf("%d %o %x\n", char4, char5, char6)

	/*
		결과 :
		H H H
		72 72 72
		72 110 48
		야 야 야
		50556 50556 50556
		50556 142574 c57c
	*/
}
```

## 3. 부동소수점 ( 실수형 )
### 3.1 golang의 부동소수점 종류
1. float32 ( 7자리 )
2. float64 ( 15자리 )

### 3.2 부동소수점 사용 예
아래 예에서 , 가장 주의할 부분은 형변환에대한 부분입니다.

ex7번의 float64로 형 변환 할 경우 , 부동소수점 오차가 발생해서 5 , 6 번 ex와 다른 값이 출력됩니다.
```golang
package main

import (
	"fmt"
)

func main() {
	var num1 float32 = 0.14
	var num2 float32 = .7 // .7로 하면 , 0.7로 인식된다.
	var num3 float32 = 442.0378373
	var num4 float32 = 10.0

	// 지수표기법
	var num5 float32 = 14e6
	var num6 float32 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println("ex1 : ", num1)
	fmt.Println("ex2 : ", num2)
	fmt.Println("ex3 : ", num3)
	fmt.Println("ex4 : ", num4)
	fmt.Println("ex5 : ", num4-0.1)
	fmt.Println("ex6 : ", float32(num4-0.1)) // 형변환 .
	fmt.Println("ex7 : ", float64(num4-0.1)) // float64로 형변환 . 값이 이상해짐
	fmt.Println("ex8 : ", num5)
	fmt.Println("ex9 : ", num6)
	fmt.Println("ex10 : ", num7)

	/*
		결과 :
		ex1 :  0.14
		ex2 :  0.7
		ex3 :  442.03784
		ex4 :  10
		ex5 :  9.9
		ex6 :  9.9
		ex6 :  9.9
		ex7 :  9.899999618530273
		ex8 :  1.4e+07
		ex9 :  156.875
		ex10 :  5.32521e-10
	*/
}


```
## 4. golang의 숫자형 연산자
### 4.1 연산자의 종류
golang에서의 연산자 종류는 아래와같습니다.
1. 산술
2. 비교

주의할 점은 , **data type이 같이야 연산자를 사용할 수 있습니다.**
다른 타입끼리는 반드시 형 변환 후 연산해야 합니다.
형 변환이 없을경우 예외가 발생합니다.
- 타 언어와 동일

**형 변환하여 계산할 시 , 데이터 손실을 주의해야하기 때문에 작은 크기의 데이터 타입을 큰 크기의 데이터 타입으로 변환하여 계산하는것을 추천합니다.**

아래 예제처럼 데이터 타입을 지정하지 않고 그냥 정수를 넣은 int type 변수와 , int16으로 형변환이되어 data type이 int16인 n6 변수를 더하면 에러가 납니다.
```golang

package main

import (
	"fmt"
)

func main() {

	n5 := 100000 // data type : int
	n6 := int16(10000) // 형변환됨 . 데이터 타입은 int16
	n7 := int8(100)

	fmt.Println("ex2 : ", n5 + n6)  //에러 발생 . 다른타입이라.
} 
```
### 4.1.1 golang 형변환 방법
golang에서 형변환은 다음과 같이 진행합니다.

```golang
package main

import (
	"fmt"
)

func main() {

	// 변수 초기화시 형변환
	n6 := int16(10000) // 형변환됨 . 데이터 타입은 int16
	n7 := int8(100)

	// 변수 사용시 형변환
	fmt.Println("ex3 : ", float32(n3)+n4)

} 

```

