# Golang 패키지
## Golang 패키지 선언 방법
변수를 선언할때와 동일하게 , 두가지 선언 방법을 사용할 수 있습니다.
```golang

// 선언방법
import "fmt"
import "os"

// 선언방법2 
import (
	"fmt"
	"os"
)

```
## package의 중요성
패키지는 코드 구조화 ( 모듈 ) 및 재 사용율을 높힐 수 있습니다.
결합도를 낮추고 , 응집도를 높힘으로서 가독성 및 유지보수 . 재 사용성이 좋은 클린 코드를 작성할 수 있습니다.

Go에서는 패키지를 단위의 독립이고 작은 단위로 개발할 수 있으며 , 작은 단위로 결합하여 프로그램을 작성할 것을 권고 합니다.

## Golang의 패키지 선언 규칙
1. 패키지 이름은 디렉토리 이름으로 작성합니다.
2. 같은 패키지 내부의 소스 파일들은 , 디렉토리 명을 패키지 명으로 사용합니다.
3. 함수 네이밍 규칙 
- 소문자 : private 
- 대문자 : public
    - ex ) fmt.Println -> 대문자라서 Println은 전역 변수이다.


**그러나 main 함수가 있는 golang 파일은 단 한개여야만 하기 때문에 , main package는 특별하게 인식됩니다. !!**

**이때 컴파일러는 공유 라이브러리가 아닌 프로그램의 시작점으로 인식하기에 main 함수가 있는 golang 파일은 package main으로 해야만 합니다. !!!!**

## Go 표준 패키지
Golang의 기본 패키지 ( ex : fmt 등 ..) 은 , Go를 설치했을 때 같이 설치되는 Go/src 경로에 위치합니다.
- GoRoot 경로 -> src -> 에 위치합니다.

따로 사용자가 생성한 Go package는 GoPath 경로에 위치합니다.

해당 경로들을 확인하기 위해선, go env 명령어를 cmd 및 bash에 입력합니다.

```bash
$ go env
...
set GOPATH=C:\Users\PC\go
set GOROOT=C:\Program Files\Go
...
```

## package 만드는 방법
package는 GOROOT 경로 밑에 /src에 위치해아 합니다.

```bash
$ cd C:\Program Files\Go\src
```

따라서 .go 파일을 생성할 때 import 내부에 내가만들 패키지가 들어갈 경로를
작성해 주어야 합니다.
- gopath는 자동으로 맞춰주기때문에 , gopath 하위 경로만 작성해주면 됩니다.
- 아래 예는 mypkg 밑 lib안에 나의 코드들이 들어간다고 생각하면 됩니다.
```golang
package main

import (
	"fmt"
	"mypkg/lib" // 실제 경로 : C:\Program Files\Go\src\mypkg\lib
)

func main(){

}
```

lib 안에 내가 만들 함수가 들어가는 패키지 golang 코드를 작성합니다.
- 패키지 명은 폴더 이름으로 생성하는것이 key point
- 해당 golang파일이 있는 위치 
    - ```C:\Program Files\Go\src\mypkg\lib```
```golang
// 패키지 명은 폴더 이름으로 생성
package lib

func CheckNum(c int32) bool {
	return c > 10
}
```

main 함수에서 내가만든 CheckNum() 함수를 사용해 봅니다.
```golang
package main

import (
	"fmt"
	"mypkg/lib"
)

func main() {
	fmt.Println("10보다 큰 수? : ", lib.CheckNum(15))
}
```

아래처럼 정상적으로 함수가 call되어 컴파일되는 것을 확인할 수 있습니다.
```bash
$ go run package2.go 
10보다 큰 수? :  true
```
