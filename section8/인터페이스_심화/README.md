# Go 인터페이스 심화
## 1. 빈 인터페이스 확장
빈 인터페이스는 , **함수의 매게변수도 되고 , 리턴값도 되고, 구조체 필드로도 사용이 가능합니다.**
- 어떤 타입으로든 변환하는것이 가능

모든 타입을 나타내기 위해, 빈 인터페이스를 사용합니다.
- 동적 타입으로 생각하면 됨. 
- Object와 비슷
```golang
import "fmt"

func printContents(v interface{}) {
	fmt.Printf("Type : (%T)\n", v) // v의 원본 타입출력 (%T)
	fmt.Println("ex1 : ", v)
	fmt.Println()
}

func main() {
	var a interface{}
	printContents(a)

	a = 7.5
	printContents(a)

	a = "golang"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)
	/*
	   결과 :
	   Type : (<nil>)
	   ex1 :  <nil>

	   Type : (float64)
	   ex1 :  7.5

	   Type : (string)
	   ex1 :  golang

	   Type : (bool)
	   ex1 :  true

	   Type : (<nil>)
	   ex1 :  <nil>
	*/
}
```

## 2. Golang Type Assertion - 타입 형변환
실행 ( 런타임 ) 시에는 인터페이스에 할당한 변수는 , 실제 타입으로 변환 후 사용해야하는 경우가 필요합니다.

이때 인터페이스.(타입) 이렇게 씁니다. -> 형 변환
```golang
func main(){
    interfaceVal.(type)
}
```
### 2.1 데이터타입 변환시 에러 발생 조건
빈 인터페이스에 최초 등록된 타입이 있다면 , 그 최초등록 타입으로만 변환이 가능합니다.
```golang

import "fmt"

func main() {
	// 최초 데이터타입은 int. 다른 데이터타입으로 변환 못함
	var a interface{} = 15

	b := a
	c := a.(int)
	d := a.(float64)

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", d)

	fmt.Println()
	/*
	   에러발생 : interface conversion: interface {} is int, not float64
	   빈 인터페이스일때 int값이 들어갔기 때문에 , 다른 형으로는 변환할 수 없습니다.
	*/
}
```

reflect.TypeOf로 해당 변수의 타입을 찍어보면 , 다 int인것을 알 수 있습니다.
- **그러나 실제 연산시에는, interface.(타입) 으로 타입 변환을 실행해야 합니다.!!**
```golang
import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = 15

	b := a
	c := a.(int) // 다 int지만 이렇게 타입을 변환해서 사용해야 합니다.

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", reflect.TypeOf(a))
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", reflect.TypeOf(b))
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", reflect.TypeOf(c))

	fmt.Println()
	/*
	   결과 :
	   ex1 :  15
	   ex1 :  int
	   ex1 :  15
	   ex1 :  int
	   ex1 :  15
	   ex1 :  int
	*/
    
	// 데이터타입이 맞는지 true, false로 검사
	if v, ok := a.(int); ok {
		fmt.Println("ex2 : ", v, ok)
		// 결과 : ex2 :  15 true
	}
}
```

## 3. 실제 타입 검사 - switch문 사용
빈 인터페이스는 어떠한 자료형도 전달받을 수 있으므로 , 타입 체크를 통해 형 변환 후 사용 가능합니다.

```golang
import "fmt"

// 이런식으로 포인터형 인지 아닌지를 체크해서 로직을 따로 짤 수 도 있습니다.
func checkType(arg interface{}) {
	switch arg.(type) {
	case Account:
	case *Acccount:
	}
}

func checkType(arg interface{}) {
	// 형변환 방법 : arg.(type) 으로 현재 데이터타입 변환
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool : ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int : ", arg)
	case float64:
		fmt.Println("This is a float : ", arg)
	case nil:
		fmt.Println("This is a nil : ", arg)
	case string:
		fmt.Println("This is a string : ", arg)
	default:
		fmt.Println("What is this Type ?", arg)
	}

}

func main() {
	checkType(true)
	checkType(1)
	checkType(22.542)
	checkType(nil)
	checkType("hello golang")
}

```