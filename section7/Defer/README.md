# Defer
## Defer 함수란 ?
지연 실행 함수라고 합니다.

Defer 함수는 , Defer 함수를 호출한 함수가 종료되기 직전에 호출됩니다.

타 언어의 Finally문과 비슷합니다.

주로 리소스 반환 (JDBC Connection close) , 열린 파일 닫기, Mutex 잠금 해제 등 마지막에 해야하는 작업을 해야 할 때 , Defer 함수를 사용합니다.

## Defer 함수 사용법
### 1. Defer 기본 사용법
Defer 함수의 기본적인 사용법은 간단합니다.
```golang
import "fmt"

func ex_f1() {
	fmt.Println("f1 : start !")
	defer ex_f2() // 함수가 종료되기 직전 마지막에 호출된다.
	fmt.Println("f1 : end !")
}

func ex_f2() {
	fmt.Println("f2 : called !")
}

func main() {
	ex_f1()
	/*
		결과 :
		f1 : start !
		f1 : end !
		f2 : called !

		ex_f2() 함수를 defer로 예약해 두었기 때문에 , ex_f1() 함수가
		종료된 이후( end가 찍히고 ) 나서 f2가 호출됩니다.
	*/
}

```
### 2. 익명함수에서의 Defer
```golang
import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi !")
	}()
}

func main() {
	sayHello("Golang")
	/*
	   결과 :
	   Hi !
	   Golan

	   함수가 끝나기 직전에 미뤄놨던 defer func() 익명함수가 실행됩니다.
	*/
}

```

### 3. Defer로 명령어 수행하기
Defer로 간단하게 stack 구조를 생성할 수 있습니다.
```golang
func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 : ", i)
	}
}

func main() {
	stack()
	/*
	   결과 :
	   ex1 :  10
	   ex1 :  9
	   ex1 :  8
	   ex1 :  7
	   ex1 :  6
	   ex1 :  5
	   ex1 :  4
	   ex1 :  3
	   ex1 :  2
	   ex1 :  1

	   defer 함수로 i를 찍기 때문에 , 선입선출이 된다.
	   1부터 10까지 누적되며 예약한 뒤 , 함수가 끝나기 직전 (여기선 명령어)에
	   마지막에 예약됐던 10번부터 출력된다.
	*/
}
```

### 4. Defer 사용 시 주의사항
defer함수를 사용할 때 중첩 함수를 Defer로 등록할 때 , 주의할 사항이 있습니다.

```golang
func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func a() {
	defer end(start("b"))
	fmt.Println("in a")
}

func main() {
	a()
	/*
	   결과 :
	   start :  b
	   in a
	   end :  b

	   Defer 함수는 , defer가 붙은 함수 한개만 걸린다.
	   따라서 해당 예제처럼 defer end(start("b"))
	   이렇게 해주면 , defer end() 만 defer가 걸리고 start()는
	   defer로 걸리지 않는다.
	*/
}

```