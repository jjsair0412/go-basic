# Go basic
Golang study repository 입니다.

## 0. godoc 에러
```
'godoc'은 내부 또는 외부 명령, 실행할 수 있는 프로그램, 또는 배치 파일이 아닙니다.
```
위 에러 발생하면 , 아래 명령어 입력

```bash
$ go env -w GO111MODULE=off

godoc 패키지 설치
$ go get golang.org/x/tools/cmd/godoc 
```

## 1. file path
**golang 과거에는 세 폴더가 필수였지만 , 지금은 VSC에서 Go extension을 설치하기만 하면 됨**
1. **src**
    - golang 실제 소스코드 위치
2. **pkg**
    - golang 패키지 모이는 경로 = 외부 패키지
3. **bin**
    - 소스코드 컴파일하여 실행 가능한 파일 ( 바이너리 ) 로 소스코드가 변하면 , bin으로 들어감

### 1.1 godoc
bash에 아래 명령어 작성 한 뒤 . 8080 포트로 들어가면 go document 페이지가 열립니다.

```
$ godoc -http=:8080
```