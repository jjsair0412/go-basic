# 변수 및 상수
## 변수 및 상수 사용방법
golang은 var ,const를 사용합니다.

변수는 선언과 초기화가 필요합니다.
- 선언시 var , const를 사용하여 선언 합니다.

## 변수명 조건
- 첫글자로 숫자는 올 수 없습니다.
- 대소문자를 구분합니다.
- 문자 , 숫자 , 밑줄 , 특수기호를 사용할 수 있습니다.

## 변수와 상수의 차이
변수와 상수의 차이점은 , 함수 내부 또는 함수 외부에서 사용하기 위해 사용합니다.
- 상수 : 불변 

## golang에서 값을 지정하지 않고 변수를 할당하면 
golang에서 값을 지정하지 않고 변수를 할당하면 , 

```golang
int= 0 
string= null 
```

이 들어갑니다 

## go언어 변수 여러개 선언방법
golang은 변수를 선언할 때 , 아래 예제 코드처럼 묶어두고 선언할 수 있습니다.

어떤 특징적인 상태를 표현하는 변수들을 한번에 묶어둠으로써 , 코드를 이해할 때 더 편하게 이해할 수 있다는 장점이 있습니다.

```go
func main() {
	// 변수 여러개로 선언
	var (
		name string ="machine"
		height int32
		weight float32
		isRunning bool
	)

    var (
        food_name string="떡볶이"
        price int32=5000
    )

}
```

## go언어 변수 여러개 선언방법 2
golang은 여러개 연속으로 선언할 때 , 각 다른 타입이 와도 에러가 발생하지 않습니다.
```golang
func main() {
    
    const c, d, e = true, 0.84, "TEST"

}
```

## **golang의 짧은 선언 ! - 중요**
golang은 짧은 선언이 있습니다.

주로 제한된 범위의 함수 내에서 사용할 경우 코드 가독성을 높일 수 있는 golang의 굉장한 특징입니다.
- 예를들어서 , main 함수에서만 사용하는 a 변수를 할당할 수 있습니다.

짧은 선언은 함수 내부에서만 사용이 가능합니다.
- 전역으로 초기화나 선언은 불가능합니다.

또한 선언한 뒤 값을 재 할당하면 에러가 발생합니다.

아래처럼 사용합니다.

```golang
func main() {
	// 짧은 선언

	// := 으로 사용하고 , 타입을 지정하지 않습니다.
	shortVar1 := 3
	shortVar2 := "TEST"
	shortVar3 := false
}
```

응용하면 , 아래처럼 golang은 if문 안에서만 조건이 되는 변수를 선언할 수 있습니다.
```golang
func main() {
	// 짧은 선언

	// EXAMPLE
	// 짧은 선언을 이용해서 , if문안에서만 사용할 수 있는 변수를 선언하고 , if가 끝나면 사라지는 변수를 만들 수 있습니다.
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success")
	}

	// undefined: i 에러 발생. 짧은 선언으로 if안에서만 사용
	// fmt.Println("i가 존재하나요 ? : ", i)
}
```

## 상수
불변값을 선언할 때 사용합니다.
```golang
func main(){

	const a string = "test1"		

}
```

상수는 한번 선언되면 , 값을 변경하는것이 금지됩니다. 고정값을 통해서 관리하는 용도로 사용합니다.
```golang
func main(){
    // 선언하고나서 다음에 또 선언하면 에러 발생
	const a string = "test1"		
    a = "test2"
}
```


함수를 사용해서 , 리턴값을 받을 때 상수를 사용하면 , 해당 함수가 항상 같은 값을 반환한다는 보장이 없기 때문에 에러가 발생합니다.
```golang
func main(){

	const d = getHeight()	

}
```

