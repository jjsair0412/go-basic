# How_To_Using_Module ?
해당 문서는 , $GOPATH 외부에서 Go module을 사용하는 방법에 대해 기술하며 , go init과 tidy에 대해 설명합니다.

## Go mod .. Go에서 패키지 관리
기존에 Go는 패키지를 관리할 때 , go env 명령어 출력 결과중 **$GOPATH 내부에서만 모듈을 관리**할 수 있었습니다.

그러나 go 1.11 버전 이상부터는 , $GOPATH 외부에서 패키지를 go.mod 파일로 관리할 수 있으며 (js의 package.json과 비슷) , 1.17부터는 $GOPATH가 사라질 예정입니다.

기존 $GOPATH에서만 패키지를 관리하게 된다면 , **외부에서 import시킨 패키지가 많아지면 많아질수록 복잡해지고 , 유지보수가 힘들 것** 입니다.

또한 코딩하는 프로그래머 입장에서도 , go env에서 $GOPATH를 계속 수정해줘야 하기에 귀찮습니다.

따라서 해당 문서의 예시대로 , $GOPATH외부에서 module을 사용하여 package를 관리한다면 , 더 좋은 클린 코드를 작성할 수 있을것 입니다.

## Go module ?
- Go module은 ```go.mod``` 파일과 ```go.sum``` 파일로 관리되며 , 설명은 다음과 같습니다.
### 1. How to work ?
Go module은 , go env의 ```GO111MODULE``` 로 관리됩니다.

아래 옵션값을 가질 수 있으며 , 옵션별 동작은 다음과 같습니다.
1. on
- $GOPATH에 관계 없이 , Go module에 있는 패키지를 사용
2. off
- Go module에 관계 없이 , $GOPATH에 있는 패키지 사용
3. auto
- 현재 디렉터리가 $GOPATH 외부라면 , 모듈 사용하고 , $GOPATH 내부라면 , $GOPATH 패키지 사용

### 2. introduce go.mod file
Go에서는 패키지를 트리 형식으로 관리하며 , 성격이 비슷한 패키지들을 ```go.mod``` 파일로 관리하게 됩니다.
- js의 package.json과 비슷
- UTF-8로 인코딩된 텍스트 파일

***한개의 모듈은 다수의 패키지를 가지고있을 수 있으며 , 패키지 종속성을 효율적으로 관리할 수 있게끔 도와줍니다.***

```go.mod``` 파일은 다음과 같이 생겨먹었습니다.

```golang
module example.com/my/thing

go 1.12

require example.com/other/thing v1.0.2
require example.com/new/thing/v2 v2.3.4

exclude example.com/old/thing v1.2.3

replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5

retract [v1.9.0, v1.9.5]
```

```go.mod``` 파일은 다음 다섯가지 키워드를 사용합니다.
**1. module**
- 모듈의 경로를 지정합니다.
- 해당 모듈이 필요하다면 , go.mod 파일의 module 경로를 import 시키면 됩니다.

**2. require**
- 해당 모듈에 반드시 필요한 종속성 패키지 정보를 가지고있습니다.
- 모듈을 사용하여 빌드 시 , 필요한 종속성 패키지를 다운로드 하고 , require에 패키지 경로 및 버전정보가 추가됩니다.

**3. replace**
- 모듈 특정 버전 교체를 지시할 때 , 사용합니다. 
- require에 정의된 모듈들을 로컬에서 빌드시킬 때 , => 키워드 우측에 위치한 패키지 버전 및 경로로 대체시킵니다.

**4. exclude**
- 패키지 특정 버전을 사용하지 않도록 할 때 사용합니다.

**5. retract**
- go get으로 가져온 패키지를 로컬 모듈 캐시에서 제거할 때 사용합니다.
    - 예를들어 , 아래와 같이 사용할 수 있습니다.
```bash
$ go get example.com/myproject@v1.0.0
$ go get retract example.com/myproject@v1.0.0
```
    - 1. example.com/myproject@v1.0.0 를 로컬 모듈 캐시에 추가
    - 2. retract로 로컬 모듈 캐시에서 제거
- 해당 명령어는 로컬 캐시에서만 제거된다는점을 주의해야 합니다.
- 원격 레포지토리에 버전을 제거하거나 , 다른 로컬에 영향을주진 않습니다.

### 3. introduce go.sum file
```go.sum``` 파일은 , go.mod에 종속성 패키지가 추가될 때 생성됩니다.

이 파일은 Go 프로젝트에서 go.mod 파일에 나열된 모듈에 대해 예상되는 암호화 체크섬을 기록하는 데 사용됩니다.
- 모듈의 올바른 버전이 프로젝트에서 다운로드되고 사용되도록 보장하는 보안 조치 역할을 합니다.
- SHA-256 알고리즘으로 암호화된 해시값이 들어가게 됩니다.

```bash
cat go.sum
github.com/pkg/errors v0.9.1 h1:42aQCYUrMQpGxrT1kxE0mTbZHvTJGifehRpc3zZ3TNg=
github.com/pkg/errors v0.9.1/go.mod h1:WXkYYxI+T1R/TfC5Y5P5bM9KjAHnYYb3Pq3NeGzCUxk=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 h1:8xUWnvTj+9KGqcBaExzlHRUsYYB+tJg51lAGLlNfNcQ=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:vhb+rWXDqo3trpXKKtg0nZX1A7dYJ5+i2x8az1LlJ5I=
...
```

## Go module을 사용해 보자 !!
golang에서는 다음 명령어로 모듈을 생성할 수 있습니다.
```bash
$ go mod init module_name
```

이때 모듈 이름은 github.com 주소를 쓰는것이 바람직합니다.
- 유일해야하며 깃에있는 모듈을 받아오는것이 일반적이므로
```bash
$ go mod init github.com/jjsair0412    
```
### 1. 폴더 및 파일 생성
먼저 **$GOPATH 외부에** 아래 구조로 폴더를 생성합니다.
- 예제는 해당 경로에 위치함.
```bash
$ ls -al
Test\test.go
main.go
```

각 .go 파일은 아래와 같이 생겨먹었습니다.
- test.go를 main.go에서 import하여 , test.go의 func를 사용할 예정입니다.
***test.go***
```golang
package test // package 이름은 자신이 위치한 디렉토리 이름이여야만 한다. (맨앞 소문자)

import "fmt"

func Test(s ...string) { // 외부에서 해당 함수를 접근하기 위해서 , public 접근 제어자인 맨앞 대문자를 사용합니다.
	for i, v := range s {
		fmt.Printf("index : %d , value : %s", i, v)
        fmt.Println()
	}
}
```

***main.go***
- 가져올 하위경로를 작성합니다.
- go init시 github.com/jjsair0412 로 모듈을 생성하기에, root 경로 뒤 /test에 Test func를 가진 .go 파일이 있다는 의미
```golang
package main

import (
	"fmt"

	test "github.com/jjsair0412/test" // test로 AS 줌. 
)

func main() {
	test.Test("hi", "im", "jinseong")
	fmt.Println("success~~")
}
```

### 2. go.mod 파일 생성
golang 프로젝트 root 경로. 즉 **main.go가 위치한 최 하단 경로** 에서 go init을 수행합니다.

```bash
$ ls -al
total 13
-rw-r--r-- 1 PC 197121   54 Apr 17 20:21 main.go
-rw-r--r-- 1 PC 197121 6163 Apr 17 21:10 README.md
drwxr-xr-x 1 PC 197121    0 Apr 17 20:19 Test/
```

go init 수행
```bash
$ go mod init github.com/jjsair0412
```

```go.mod``` 파일이 생성된것을 확인합니다.
```bash
$ ls -al
total 14
-rw-r--r-- 1 PC 197121   43 Apr 17 21:11 go.mod
-rw-r--r-- 1 PC 197121   54 Apr 17 20:21 main.go
-rw-r--r-- 1 PC 197121 6561 Apr 17 21:11 README.md
drwxr-xr-x 1 PC 197121    0 Apr 17 20:19 Test/

$ cat go.mod 
module github.com/jjsair0412

go 1.20
```

지금은 패키지가 한개라서 깔끔하지만, 많아지면 꼭 아래 명령어를 수행합니다.
```bash
$ go mod tidy
```

### 3. go run !
이제 테스트해봅니다.
```bash
$ go run main.go
index : 0 , value : hi
index : 1 , value : im
index : 2 , value : jinseong
success~~
```

성공적으로 $GOPATH 외부에서 모듈을 꺼내온것을 확인할 수 있습니다.

