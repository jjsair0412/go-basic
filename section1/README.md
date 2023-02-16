# section 1
## go를 실행하는 방법들
1. go run
cmd에서 go 파일 을 go run 명령어로 실행하는 방법
- 메모리에서 바로 실행
- 단위테스트
```
$ go run helloworld.go
```

2. go build
go build 명령어로 .exe파일처럼 실행 가능한 파일로 바꿔줍니다.
- 단위테스트용
```bash
$ go build helloworld.go
```

3. go install
바이너리 파일을 생성하는 방법
- .exe파일이 생성되어 bin 폴더로 이동
- main 함수 명으로 생성되어 bin 폴더로 이동
- 외부 pkg를 가져다 gofile에 썻다면 , install하면 의존관계가 있는 ( 모든 패키지를 ) 파일들을 묶어서 .exe 파일 ( 실행 가능한 )을 생성합니다.
```
$ go install hellowrold.go

$ cd bin

$ ls | grep hello
helloworld.exe*
```