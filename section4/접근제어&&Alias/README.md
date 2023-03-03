# 접근제어 및 Alias
패키지에서 접근 제어자와 Alias를 공부합니다.

## golang package 접근제어
변수 , 상수 , 함수 , 메서드 , 구조체 등 식별자들은 아래 규칙을 따릅니다.
- 첫 글자가 대문자일 경우
    - 패키지 외부에서 접근이 가능
- 첫 글자가 소문자인 경우
    - 패키지 외부 접근 불가. 해당 패키지 내에서만 접근가능 


## bean 식별자 및 package 접근제어 사용법
package를 가져온다음 , 해당 패키지가 어떤 의미를 가지는지 별칭을 지정해줄 수 있습니다.

다음처럼 앞에 그냥 작성해주면 됩니다.

사용할 때도 패키지명이 아닌 Alias로 불러올 수 있습니다.
```golang
package main

import (
	"fmt"
	more10check "mypkg/lib"
	more100check "mypkg/lib2"
)

func main(){
	fmt.Println("10보다 큰 수 ? : ",more10check.CheckNum(11))

	fmt.Println("100보다 큰 수 ? : ",more100check.CheckNum1(101))
}
```

추가로 , 밑줄 ( _)을 사용하면 golang에서 사용하지 않은 패키지는 지워버리는것을 생략할 수 있는데 , 이것을 **빈 식별자** 라 합니다.
- 아래 예에서 원래는 mypkg/lib2가 golang이 컴파일되는순간 지워져야 하지만 , 빈 식별자 ( _ ) 를 사용해서 생략할 수 있습니다. 
```golang
package main

import (
	"fmt"
	more10check "mypkg/lib"
	_ "mypkg/lib2"
)

func main(){
	fmt.Println("10보다 큰 수 ? : ",more10check.CheckNum(11))

}
```