# Golang zero value
Golang에서 0 또는 빈값 , 그리고 nil을 출력하거나 사용하고 싶을 땐, 일반적으로 다음과같은 사용법을 사용합니다.

아래 코드와같이 작성하면 , 선언된 a 와 b 는 var 선언이후 정의된 변수타입을 통한 0 또는 "공백" 이 저장되며, 출력됩니다.
```golang
package main

import "fmt"

func main() {

	var a int
	fmt.Println(g)
    // 0 출력

	var b string
	fmt.Println(b)
    // "" 공백 출력

	var c float
	fmt.Println(c)
    // 0.0 출력 

}
```

Zero Value 출력될 수 있는 목록은 다음과 같습니다.

|`변수 타입`|`저장되는 Zero Value 값`|`비고`|
|--|--|--|
|boolean|false||
|int|0||
|float|0.0||
|string|""|공백|
|pointers|nil||
|functions|nil||
|interfaces|nil||
|slices|nil||
|channels|nil||
|maps|nil||

