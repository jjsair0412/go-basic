# 사용자 정의 타입 
## Golang에서의 객체 지향
**Golang에서는 클래스나 상속 개념이 없고 , 구조체로 객체 지향을 정의합니다.**

자바의 객체 지향은 클래스의 상태나 속성(맴버 변수) 이나 기능(메서드) 으로 이루어져 있지만 , Go는 다릅니다.

Go는 전형적인 객체 지향의 특징을 가지고 있지 않지만 , 인터페이스를 통한 다형성 지형이나 구조체를 통한 class 형태의 코딩이 가능하게 합니다.
- **따라서 Golang은 객체 지향의 기본을 가지고 있으며 , 객체 지향 언어입니다.**

### Go 객체지향의 특징
- 상태와 메서드를 분리해서 정의하여 결합성이 없습니다.
- 사용자 정의 타입 : 구조체 , 인터페이스 , 기본 타입 (int, float, string , 함수 등) 의 타입을 막 재정의하여 사용할 수 있습니다.
- 구조체와 메서드 연결을 통해 타 언어의 클래스 형식 처럼 사용하는것이 가능합니다.

## 1. Golang에서 구조체와 메서드 연결
Go는 구조체와 메서드를 한번에 정의하지 않고 , 따로 분리시켜서 정의한 다음 연결합니다.

다음과 같이 메서드와 구조체가 자동으로 연결되는 방법이 객체 지향에 가깝습니다.
- 실행 방법에 있어서 자동 주입이 객체 지향 !

아래 예제로 개념 학습

```golang
import "fmt"

// 구조체
type Car struct {
	// 멤버 변수
	name  string
	color string
	price int64
	tax   int64
}

// 일반 메서드 방법. 객체 지향이라고 하기 힘듬
// 구조체는 메서드 바인딩을 통해 기능 (메서드) 를 정의한다.
// Go는 메서드를 함수로 정의한다.
// 메게변수가 구조체로 들어오게 된다. ( 여기선 Car 구조체가 메게변수로 들어옴 )
func Price(c Car) int64 {
	return c.price + c.tax
}

// 객체 지향의 방법
// 구조체와 메서드를 바인딩할때 쓰는 방식
// 메게변수가 먼저 오고 , 함수 이름이 그다음 온다. 그리고 리턴타입
// 구조체가 해당 메서드에 자동으로 주입됨.
func (c Car) Price() int64 {
	return c.price + c.tax
}

func main() {

	// Car 구조체의 멤버 변수에 접근하여 bmw와 benz 변수를 선언한다.
	bmw := Car{
		name:  "520d",
		price: 5000,
		color: "white",
		tax:   500,
	}

	benz := Car{
		name:  "220d",
		price: 6000,
		color: "white",
		tax:   600,
	}

	hyundai := Car{
		name:  "zenesis",
		price: 1500,
	}

	fmt.Println("ex1 : ", bmw, &bmw)
	fmt.Println("ex1 : ", benz, &benz)

	fmt.Println()
	fmt.Println("ex2 : ", Price(bmw))
	fmt.Println("ex2 : ", Price(benz))

	fmt.Println()
	// 객체지향은 다음과 같이 사용합니다.
	fmt.Println("ex3 : ", bmw.Price())
	fmt.Println("ex3 : ", benz.Price())

	fmt.Println()
	// 정의안된 필드는 기본값 리턴
	fmt.Println("ex4: ", hyundai.Price())
	/*
		결과 :
		ex1 :  {520d white 5000 500} &{520d white 5000 500}
		ex1 :  {220d white 6000 600} &{220d white 6000 600}

		ex2 :  5500
		ex2 :  6600

		ex3 :  5500
		ex3 :  6600
	*/

	// false. 하나의 구조체로 여러 객체를 생성하였기 때문에 메모리 공간이 다름
	fmt.Println("ex4 : ", &bmw == &benz)

}
```

## 2. 기본 자료형 타입을 사용자 정의 타입으로 바꾸는 방법
Golang에선 기본 자료형과 같은 기본 타입들을 사용자 정의 타입으로 재 정의하여 사용할 수 있습니다.

사용자 정의 타입과 기본 타입은 매게변수 전달 시에 변환해야 사용이 가능합니다.
- int를 재정의해서 cnt를 만들었다면 , 둘은 다른 타입으로 구분됩니다.

```golang
import (
	"fmt"
)

// int를 cnt라는 이름으로 재 정의
type cnt int

func testConvertT(i int) {
	fmt.Println("ex3 : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("ex3 : (Custom Type) : ", i)
}

func main() {
	a := cnt(15)
	fmt.Println("ex1 : ", a)

	var b cnt = 15
	fmt.Println("ex1 : ", b)

	// testConvertT(b) 에러 발생. int를 넣어야지 왜 cnt 타입으로 넣냐 !! 라는 에러..
	// 에러가 발생하지 않으려면 아래처럼 형변환 해야 함.
	testConvertT(int(b))
	testConvertD(a)

	/*
	   결과 :
	   ex1 :  15
	   ex1 :  15
	   ex3 : (Default Type) :  15
	   ex3 : (Custom Type) :  15
	*/
}
```

## 3. 기본 함수를 사용자 정의 타입으로 바꾸는 방법
func 함수의 매게변수 및 리턴타입을 사용자 입맛에 따라 재 정의하여 고정시켜두고, 유기적으로 함수를 사용하는 예제 입니다.

totCost라는 함수를 메게변수와 리턴타입을 고정시키고 나서 , 재 정의한 함수를 매게변수로 받는 함수를 만들어갖구 실수를 줄일 수 있는 유기적인 코드 입니다.
```golang
// int두개를 매게변수로 받아서 int 한개를 리턴시키는 함수를
// totCost라는 이름으로 재 정의
type totCost func(int, int) int

// totCost 함수도 메게변수로 받음
// 이렇게하면 totCost를 받지 않으면 해당 함수가 실행되지 않기에 실수를 줄일 수 있음.
func describe(cnt, price int, fn totCost) {
	fmt.Printf("cnt : %d, price : %d, orderPrice : %d", cnt, price, fn(cnt, price))

}

func main() {
	var orderPrice totCost
	// orderPrice라는 변수에 totCost라는 함수 타입으로 정의해서 , 함수를 고정시킴
	orderPrice = func(cnt, price int) int {
		return (cnt * price) + 10000
	}

	describe(3, 5000, orderPrice)
    // 결과 : cnt : 3, price : 5000, orderPrice : 25000
}
```

## 4. Golang 구조체에서 포인터
함수는 기본적으로 값 전달 방식입니다. 따라서 원본 값 수정이 안됩니다.

그러나 맵이나 슬라이스는 참조 전달이기에 원본 값 수정이 됩니다.

리시버(구조체) 또한 포인터를 활용해서 메서드 내에서 원본 수정이 가능합니다.

구조체는 & 연산자같은걸 넣어주지 않아도 자동으로 주소값이 전달되지만 , * 형으로 받아야 합니다.
- 아래 예제 참고
```golang
type shoppingBasket struct {
	cnt,
	price int
}

// 결제 함수
func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본 수정(참조 전달 형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
	/*
		실제 로 발생하는 동작은 아래처럼 동작함
		*b.cnt += cnt
		*b.price += price

	*/

}

// 원본 수정 안됨(값 전달 형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	bs1 := shoppingBasket{
		3,
		5000,
	}

	fmt.Println("ex1 Total price : ", bs1.purchase())
	// 결과 : ex1 Total price :  15000

	// 원본 값 ( 3에 7 더하고 5000에 5000 더함)
	bs1.rePurchaseP(7, 5000)
	fmt.Println("ex2 Total price : ", bs1.purchase())
	// 결과 : ex1 Total price :  100000
	// 원본값이 수정됐음.
	// 주소값을 넣어주는 & 기호같은걸 넣어주지 않아도 , 자동으로 주소값이 넘어감

	bs1.rePurchaseD(10, 0)
	fmt.Println("ex3 Total price : ", bs1.purchase())
	// 결과 : ex1 Total price :  100000
	// 원본값이 수정되지 않음. 값 전달 방식으로 진행.
	// 받는게 포인터가 아니니까 값 전달 방식이 됨.
}

```