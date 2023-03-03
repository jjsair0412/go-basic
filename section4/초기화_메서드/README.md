# 초기화 메서드
init 메서드 , 즉 초기화 메서드는 프로그램 시작 시 main method보다 먼저 실행되는 함수를 의미합니다.

## 초기화메서드가 무엇인가 ?
init함수는 패키지 로드 시에 가장 먼저 호출되며 , 가장 먼저 초기화시켜야되는 작업을 작성 시에 유용합니다.

## 사용법
init() 함수를 지정합니다.

init() 함수는 main보다 먼저 딱 한번 수행됩니다.
```golang
package main

import (
	"fmt"
)

// init은 main보다 먼저딱 한번 실행됨.
func init() {
	a := "hi"

	fmt.Println("init method start")
	fmt.Println(a)
}

func main() {
	fmt.Println("main method start")
}
```

위의 코드의 실행 모습은 다음과 같습니다.
```bash
$ go run init1.go 
init method start
hi
main method start
```

## main이 아닌 package에서 init
init 메서드를 가지고있는 함수를 불러오는 ( import 하는 ) main 함수가 있는 **go file의 main() 메서드가 실행되기 전에 , import안에 있던 init이 먼저 실행됩니다.**

## init메서드의 실행 순서
init 메서드의 개수는 상관 없고 위에서부터 아래 순서대로 init 메서드가 실행됩니다.

따라서 만약 , import한 패키지 내부에 init() 함수가 있다면 , 그 init부터 실행한 뒤 차례대로 실행됩니다.

![initialization][initialization]
  
[initialization]:./images/initialization.PNG

```golang
package main

import (
	"fmt"
)

// init은 main보다 먼저딱 한번 실행됨.
func init() {
	fmt.Println("init method start")
}

func init() {
	fmt.Println("init2 method start")
}

func init() {
	fmt.Println("init3 method start")
}

func init() {
	fmt.Println("init4 method start")
}

func main() {
	fmt.Println("hello i'm main")
}
```

위 예제를 수행한 결과는 다음과 같습니다.
```bash
$ go run init2.go 
init method start
init2 method start
init3 method start
init4 method start
hello i'm main
```
