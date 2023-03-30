# Golang에서의 Closure
## Closure란 ?
자바나 js에서 지원하는 클로저와 같은것을 Golang에서도 지원합니다.

**익명함수를 사용할 경우** , 함수를 변수에 할당해서 사용이 가능한것을 의미합니다.

함수 안에서 함수를 선언및 정의를 하는것이 가능합니다.
이때 외부 함수에 선언된 변수에 접근 가능한 함수를 뜻합니다.

함수가 선언되는 순간에 , 함수가 실행될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체) 입니다.

***왜 사용하냐 ?***
- 함수를 호출할 때 , 이전에 존재했던 값을 유지하기 위해서 사용합니다.
- 비동기 , 누적 count 등 무분별하게 전역변수를 남발하는 것을 감소시킬 수 있습니다.

***단점***
- 객체들이 메모리에 계속 존재하기 때문에 , 심하면 오버플로우 현상이 나올 수 있습니다.
- 따라서 클로저를 사용할 땐 , 정확히 이해하고 사용하는것이 중요합니다.

## Go에서의 Closure 예제
### 아래 예는 **클로저를 사용하지 않을 때**의 예제입니다.

익명함수를 multiply 함수에 선언해 놓고 , r1 변수에 해당 함수를 이용해서 나온 결과값을 초기화하는 예제 입니다.
```golang
import "fmt"

func main() {
	multiply := func(x, y int) int {
		return x * y
	}

	r1 := multiply(5, 10)

	fmt.Println(r1)
    // 결과 : 50
}
```

### 아래 예는 **클로저를 사용한 예시**입니다.

**익명함수를 사용할 때, 익명함수 외부에 선언된 변수를 , 익명함수 내부에서 접근하여 연산합니다.**

m , n은 sum 변수에 초기화된 익명함수안에 정의되어있지 않지만 , m , n 변수를 캡쳐해 두기 때문에 , 익명함수 내부에서 외부에 있는 변수에 접근할 수 있는것 입니다.

**또한 지역변수 (m,n) 가 소멸되지 않기 때문에 함수호출마다 사용이 가능합니다.**
```golang

import "fmt"

func main() {
	m, n := 5, 10 // 변수가 캡쳐되는 순간
	sum := func(c int) int {
		return m + n + c // 지역변수 가 소멸되지 않는다 .
	}

	r2 := sum(10)

	fmt.Println("ex2 : ", r2)
	// 결과 : ex2 :  25
}
```

## closure를 이용한 증감연산 예제
1부터 5까지 클로저로 누적되어 더해지는 예제 입니다.
```golang
import "fmt"

func main() {

	cnt := increaseCnt()

	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())
	fmt.Println("ex 1 : ", cnt())

	/*
	   결과 :
	   ex 1 :  1
	   ex 1 :  2
	   ex 1 :  3
	   ex 1 :  4
	   ex 1 :  5

	   함수 increaseCnt()에서 외부 n을 캡쳐해서 저장하고 있기 때문에 ,
	   누적되어 저장되는것이다.
	*/
}

func increaseCnt() func() int {
	n := 0 // 지역변수 0 . 캡쳐 대상
	return func() int {
		n += 1
		return n
	}
}

```