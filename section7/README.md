# Golang에서의 함수
## 1. 함수 선언 방법
func 키워드로 함수를 선언합니다.

함수 선언 패턴은 아래와 같은 종류가 있습니다.
```golang
1.
func 함수명(매개변수) (반환타임 & 반환 값 변수 명) : 번환값이 존재할 때

2.
func 함수명() (반환타입 & 반환값 변수명) : 매개변수 없음, 반환값 존재

3.
func 함수명(매개변수) : 매개변수 존재 , 반환 값 없음

4.
func 함수명() : 매개변수 없음 , 반환값 없음
```

**타 언어와 달리 반환 값 (return value)이 여러개 반환하는것이 가능합니다.**

또한 함수 선언의 위치는 어디든 가능합니다.

아래 예제처럼 함수를 선언합니다.

```golang
import (
	"fmt"
	"strconv"
)

// 함수는 어디에 선언하든 상관없다.
func helloGolang() {
	fmt.Println("ex1 : Hello Golang")
}

func main() {
	helloGolang()

	helloGolang2()

	say_one("hello world")

	result := sum(5, 5)

	fmt.Println(result)

	fmt.Println(sum(5, 5))

	fmt.Println("int -> string ex3 : ", strconv.Itoa(sum(3, 3)))
}

// 매개변수 , return 둘다 없을경우
func helloGolang2() {
	fmt.Println("ex1 : Hello Golang2")
}

// 매개변수만 있을경우
func say_one(m string) {
	fmt.Println("ex2 : ", m)
}

// 매개변수 , return 둘다 있을경우
// return type의 type을 매개변수 괄호 닫은 뒤 작성해줍니다.
func sum(x int, y int) int {
	return x + y
}
```

## 2. 함수의 전달 방법
golang에서는 함수 전달 방법은 세가지로 나뉩니다..
1. 함수(콜백) 전달
2. 값 전달(call by value)
3. 참조 전달(call by reference)


### 2.1 함수(콜백) 전달 방법
매개변수로 함수를 전달할 수 있습니다..

아래 예시는 , return type이 없고 매개변수로 int 두개를 가지는 함수를 매개변수로 가진 sum 함수를 정의해서 사용하는 예제입니다..
```golang

// 매개변수로 int형 i와 , 매개변수로 int , int 타입 두개를 가지는 함수를 가진다.
func sum(i int, f func(int, int)) {
	f(i, 10)
}

// 매개변수 타입이 동일할땐 , 하나를 지우고 뒤에만써도 된다.
func add(x, y int) {
	fmt.Println("ex1 : ", x+y)
}

func main() {
	sum(100, add)
}
```

### 2.2 값 전달
변수 실제 값만 전달하는 예시입니다.

변수가 할당되어잇는 메모리 주소값을 가르키지 않고 , 값만 전달했기에 실제 값에 영향을 주지 못합니다.
```golang

func multi_value(i int) {
	i = i * 10
}

func main() {

	a := 100
	multi_value(a)
	fmt.Println(a) // 실제 변수에 영향 x (영향주려면 포인터로 주소값 참조해야됨)
}
```

### 2.3 참조 전달(call by reference)
변수가 할당된 메모리 주소값을 전달했기 때문에 , 실제 변수에 영향을 줄 수 있습니다.
```golang
func multi_reference(i *int) {
	*i *= 10 // *i = *i * 10 과 동일
}

func main() {
	a := 100

	multi_reference(&a)
	fmt.Println(a) // 실제 변수에 영향 o (실제주소값 참조) . 결과 : 1000
}

```

## 3. 함수 다중 값 반환
함수 return할때 여러개의 값을 반환할 수 있습니다.

```golang
// 매개변수 타입이 int로 동일하기에 int 하나 생략
// return할 변수가 두개라서 int, int 두개 작성
func multiply(x, y int) (int, int) {
	// x * 10은 첫번쨰 int , y * 10은 두번째 int
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func main() {
	// a에 100 담기고 , b에 50 담김.
	a, b := multiply(10, 5)

	// c := multiply(10, 5) 함수 리턴타입 개수 안맞아서 에러발생
	c, _ := multiply(10, 5) // _로 뒤에 리턴되는 y * 10 생략
	_, d := multiply(10, 5)

	fmt.Println(a, b)
	fmt.Println(c)
	fmt.Println(d)

	/*
		결과 :
		100 50
		100
		50
	*/

	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, _, y3, _, y5 := arrayMultiply(1, 2, 3, 4, 5)

	fmt.Println(x1, x2, x3, x4, x5)
	fmt.Println(y1, y3, y5)
	/*
		결과 :
		1 4 9 16 25
		1 9 25
	*/
}
```

또한 return type을 지정할 때 , return 변수를 선언하여 사용할 수 있습니다.

```golang
// return 매개변수에서 변수 선언 가능
func multiply(x, y int) (r1 int, r2 int) {
	// r1, r2를 명시적 선언하여 가독성 증대
	r1 = x * 10
	r2 = y * 10

	return r1, r2
}

func main() {
	a, b := multiply(10, 5)

	fmt.Println(a, b)
}
```

## 4. 가변인자
함수의 매개변수 개수가 동적으로 변할 때 가변인자 함수를 사용합니다.

### 4.1 기본 사용방법
기본적으로 가변인자 함수를 선언하고 사용하는 방법은 , ...을 사용합니다.
```golang
import "fmt"

// 매개변수를 선언할 때 , 아래처럼 ...을 넣어준다면 ,
// int형 몇개가 들어와도 상관이 없다는 의미이다.
// 가변형으로 매개변수를 받을 수 있다.
func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}

	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2 : ", value)
	}
}

func main() {
	// 1x2x3 의 결과 6 출력
	x := multiply(1, 2, 3)
	fmt.Println(x)

	fmt.Println()

	// 1~10까지 더한 결과 55 출력
	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(y)

	fmt.Println()

	/*
		결과 :
		ex2 :  a
		ex2 :  apple
		ex2 :  test
		ex2 :  golang
		ex2 :  seoul
	*/
	prtWord("a", "apple", "test", "golang", "seoul")

	fmt.Println()
}
```

### 4.2 가변인자 함수 매개변수에 슬라이스 및 배열 넣기
슬라이스나 배열을 가변인자 함수에 집어넣을 수도 있습니다.

아래 예시처럼 ... 을 사용합니다.
```golang
func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func main() {
	a := []int{1, 2, 3, 4, 5}

	// 슬라이스 a의 인덱스 0번부터 차례대로 들어간다.
	m := multiply(a...)
	n := sum(a...)

	/*
		결과 :
		ex3 :  120
		ex3 :  15
	*/
	fmt.Println("ex3 : ", m)
	fmt.Println("ex3 : ", n)
}
```

## 5. 함수를 변수에 할당해서 사용하기
### 5.1 슬라이스에 할당
슬라이스 인덱스안에 함수를 할당하는 방법입니다.

```golang
import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {

	// 0번 인덱스에 multiply , 1번 인덱스에 sum을 넣습니다.
	// func 타입의 값을 가진 슬라이스를 선언하고 , int형 두개를 매개변수로
	// 가지며 return type이 int형 한개를 가지는 애들을 묶을 수 있습니다.
	f := []func(int, int) int{multiply, sum}

	// 0번째 , 1번째 함수에 10. 10을 매개변수로 넣음
	a := f[0](10, 10) // a := multiply(10,10) 과 동일
	b := f[1](10, 10) // b := sum(10,10) 과 동일

	fmt.Println("ex1 : ", a, f[0](10, 10))
	fmt.Println("ex1 : ", b, f[1](10, 10))
	/*
	   결과 :
	   ex1 :  100 100
	   ex1 :  20 20
	*/

}
```

### 5.2 변수에 할당
변수에 함수를 할당하는 방법입니다.
```golang

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	// 변수에 할당하는 방법입니다.
	// int형 매개변수 두개와 , return type을 int로 갖는 함수인 multiply를
	// f1에 넣습니다.
	var f1 func(int, int) int = multiply
	f2 := sum // 이렇게 선언해도 무관합니다.

	fmt.Println("ex2 : ", f1(10, 10))
	fmt.Println("ex2 : ", f2(10, 10))
	
	/*
	   결과 :
	   ex3 :  100
	   ex3 :  20
	*/

}
```


### 5.3 맵에 할당
map의 각 key에 value로 함수를 할당하는 방법입니다.

```golang

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	// map에 할당하는 방법입니다.
	// key로 string을 갖고 , value로 매개변수 int 두개를 갖고 return값이 int 하나인
	// 함수를 가지는 맵을 선언합니다.
	m := map[string]func(int, int) int{
		"mul_func": multiply,
		"sum_func": sum,
	}

	// map의 key를 통해서 함수를 꺼내고 매개변수를 넣어서 사용합니다.
	fmt.Println("ex3 : ", m["mul_func"](10, 10))
	fmt.Println("ex3 : ", m["sum_func"](10, 10))
	/*
	   결과 :
	   ex3 :  100
	   ex3 :  20
	*/
}
```


## 6. 재귀함수 (Recursion)
재귀함수를 사용하면 , 프로그램이 보기쉽고 코드가 간결해지며 오류 수정이 용이해집니다.

그러나 코드를 이해하기 어렵고 , 메모리를 많이 사용합니다.
- 무한루프에 빠질 가능성이 있음

```golang
import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2 : (", n, ")", "hi!")
	prtHello(n - 1)
}

func main() {
	x := fact(5)

	fmt.Println("ex1 : ", x)
	// 결과 : ex1 :  120

	prtHello(10)
	/*
	   결과 :
	   ex2 : ( 10 ) hi!
	   ex2 : ( 9 ) hi!
	   ex2 : ( 8 ) hi!
	   ex2 : ( 7 ) hi!
	   ex2 : ( 6 ) hi!
	   ex2 : ( 5 ) hi!
	   ex2 : ( 4 ) hi!
	   ex2 : ( 3 ) hi!
	   ex2 : ( 2 ) hi!
	   ex2 : ( 1 ) hi!
	*/
}
```

## 7. 익명함수
주로 선언과 동시에 실행해야 할 때 사용합니다.

```golang
import "fmt"

func main() {

	// 익명함수는 아래처럼 사용합니다.
	// js의 익명함수와 동일하게 이름이 없고 , 중괄호가 닫힌 뒤 () 소 괄호를 붙여야합니다.
	func() {
		fmt.Println("ex1 : anonymous func!")
	}()
	// 결과 : ex1 : anonymous func!

	// js처럼 매개변수를 바로 넣기 위해 뒤에 () 소괄호 안에 변수 넣습니다.
	msg := "hello golang"
	func(n string) {
		fmt.Println("ex2 : ", n)
	}(msg)
	// 결과 : ex2 :  hello golang

	// 매개변수를 바로 넣으면서 함수를 사용해도 무관합니다.
	func(x, y int) {
		fmt.Println("ex3 : ", x+y)
	}(10, 20)
	// 결과 : ex3 :  30

	// 실행과 동시에 짧은 선언으로 r 변수 안에 func 리턴값이 들어갑니다.
	r := func(x, y int) int {
		return x * y
	}(10, 100)

	fmt.Println("ex4 : ", r)

	// 결과 : ex4 :  1000
}
```