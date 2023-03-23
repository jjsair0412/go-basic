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