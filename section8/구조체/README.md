# Go에서의 구조체
## 1. 구조체 기본
Go 필드들의 집합체 또는 컨테이너

class의 형태를 갖습니다.

필드는 갖지만 , 메서드는 갖고있지 않습니다.
- 그 대신 메서드는 리시버 형식으로 구조체와 메서드를 연결해서 사용합니다.

객체지향 방식을 지원하며 , 리시버를 통해 메서드와 연결합니다.

상속 , 객체 , 클래스 개념이 없지만 구조체를 통해서 객체지향 특성을 사용합니다.

구조체를 선언하고 구조체의 인스턴스를 생성하여 리시버와 연결하여 사용합니다.

## 2. 구조체 사용법
기본적인 구조체 사용법은 다음과 같습니다.

구조체 필드에 값을 초기화시켜주지 않으면 , 기본자료형의 default값이 입력됩니다.
```golang
type Account struct {
	number   string
	balance  float64 //잔액
	interest float64 // 이자
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	kim := Account{
		number:   "245-901",
		balance:  1000000,
		interest: 0.015,
	}

	lee := Account{
		number:  "245-902",
		balance: 1200000,
	}

	park := Account{
		number:   "245-903",
		interest: 0.025,
	}

	// 필드 순서대로 그냥 입력해줘도 괜찮음
	cho := Account{"245-904", 1500000, 0.03}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)
	fmt.Println("ex1 : ", park)
	fmt.Println("ex1 : ", cho)
	/*
		결과 :
		ex1 :  {245-901 1e+06 0.015}
		ex1 :  {245-902 1.2e+06 0}
		ex1 :  {245-903 0 0.025}
		ex1 :  {245-904 1.5e+06 0.03}
	*/

	fmt.Println()

	fmt.Println("ex2 : ", int(kim.Calculate()))
	fmt.Println("ex2 : ", int(lee.Calculate()))
	fmt.Println("ex2 : ", int(park.Calculate()))
	fmt.Println("ex2 : ", int(cho.Calculate()))
	/*
	   결과 :
	   ex2 :  1015000
	   ex2 :  1200000
	   ex2 :  0
	   ex2 :  1545000
	*/

	fmt.Println()
}
```

## 3. 구조체의 다양한 선언 방법
Golang의 구조체는 선언만 하면 자동으로 참조형을 넘깁니다.
- 메서드가 리시버일때만 자동으로 참조타입이 넘어감
- 만약 그냥 함수에 넣어주면 에러발생
```golang
import "fmt"

type Account struct {
	number   string
	balance  float64 //잔액
	interest float64 // 이자
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {

	// 포인터형으로 구조체를 선언할 때 ,
	// new 키워드를 붙여서 아래처럼 선언하려면 * 를 붙여주어야 함
	// 인터페이스에 넘길때는 아래처럼 사용해야함.
	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	// 주소값 넘겨서 선언
	hong := &Account{
		number:   "245-902",
		balance:  1500000,
		interest: 0.04,
	}

	// 짧은 선언으로 new 키워드 사용해서 선언
	lee := new(Account)
	lee.number = "245-903"
	lee.balance = 1300000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", hong)
	fmt.Println("ex1 : ", lee)
	/*
	   결과 :
	   ex1 :  &{245-901 1e+07 0.015}
	   ex1 :  &{245-902 1.5e+06 0.04}
	   ex1 :  &{245-903 1.3e+06 0.025}
	*/

	fmt.Println()

	// %#v 는 구조체 값을 모두 출력
	fmt.Printf("ex2 : %#v\n", kim)
	fmt.Printf("ex2 : %#v\n", hong)
	fmt.Printf("ex2 : %#v\n", lee)
	/*
	   결과 :
	   ex2 : &main.Account{number:"245-901", balance:1e+07, interest:0.015}
	   ex2 : &main.Account{number:"245-902", balance:1.5e+06, interest:0.04}
	   ex2 : &main.Account{number:"245-903", balance:1.3e+06, interest:0.025}
	*/

	fmt.Println()

	fmt.Println("ex3 : ", int(kim.Calculate()))
	fmt.Println("ex3 : ", int(hong.Calculate()))
	fmt.Println("ex3 : ", int(lee.Calculate()))
	/*
	   결과 :
	   ex3 :  10150000
	   ex3 :  1560000
	   ex3 :  1332500
	*/
}

```
단순하게 작성하면 , 4가지 구조체 선언 방식이 있습니다.
```golang
type car struct {
	color string
	name  string
}

func main() {
	c1 := car{"red", "220d"}

	c2 := new(car)
	c2.color, c2.name = "white", "sonata"

	c3 := &car{}

	c4 := &car{"black", "520d"}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	/*
	   결과 :
	   {red 220d}
	   &{white sonata}
	   &{ }
	   &{black 520d}
	*/
}

```
## 4. 익명함수와 구조체
구조체 익명 선언 및 활용 방법은 다음과 같습니다.

익명함수와 같이 구조체 이름을 빼버리고 선언하면 됩니다.
```golang
import "fmt"

func main() {

	// 아래처럼 struct를 익명으로 바로 구조체 선언하고 , 사용한다.
	car1 := struct{ name, color string }{"520d", "red"}

	fmt.Println("ex1 : ", car1)
	fmt.Printf("ex1 : %#v", car1)
	/*
	   결과 :
	   ex1 :  {520d red}
	   ex1 : struct { name string; color string }{name:"520d", color:"red"}
	*/

	cars := []struct{ name, color string }{
		{"530i", "red"},
		{"530e", "white"},
		{"530a", "blue"},
	}

	fmt.Println()
	for _, c := range cars {
		fmt.Printf("ex2 : %s, %s --- %#v\n", c.name, c.color, c)
	}

	/*
	   결과 :
	   ex2 : 530i, red --- struct { name string; color string }{name:"530i", color:"red"}
	   ex2 : 530e, white --- struct { name string; color string }{name:"530e", color:"white"}
	   ex2 : 530a, blue --- struct { name string; color string }{name:"530a", color:"blue"}
	*/
}
```

## 5. 구조체 필드 태그
구조체를 선언할 때 태그를 넣어줘서 코드 가시성을 높힐 수 있습니다.

reflect의 typeof 필드를 통해서 가져올 수 있습니다.
```golang
import (
	"fmt"
	"reflect"
)

type Car struct {
	name    string "차량명" // 필드 태그 생성
	color   string "차량 색상"
	company string "재조사"
}

func main() {

	tag := reflect.TypeOf(Car{})

	// tag 필드 갯수만큼 for문 돌면서 태그 뽑아옴
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("ex1 : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
		/*
			결과 :
			태그이름 , 필드명 , 필드 타입 꺼내옴
			ex1 :  차량명 name string
			ex1 :  차량 색상 color string
			ex1 :  재조사 company string
		*/
	}
}

```
## 6. 중첩 구조체 - 구조체 속의 구조체
구조체 속의 구조체. 중첩 구조체를 생성할 수 있습니다.

구조체를 선언할 때 , 데이터 타입을 중첩 구조체를 넣어줄 구조체로 주면 됩니다.

```golang
import "fmt"

// 패키지 외부에서 접근가능. 첫글자 대문자라서
type Car struct {
	name    string "차량명" // 필드 태그 생성
	color   string "차량 색상"
	company string "재조사"
	detail  spec   "차량 상세 제원" // 중첩 구조체 생성
}

// 패키지 외부에서 접근불가. 첫글자 소문자라서
type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {
	car1 := Car{
		"520d",
		"silver",
		"bmw",
		spec{
			4000,
			1000,
			2000,
		},
	}

	fmt.Println("ex1 : ", car1.name)
	fmt.Println("ex1 : ", car1.color)
	fmt.Println("ex1 : ", car1.company)
	fmt.Printf("ex1 : %#v\n", car1.detail)
	/*
	   결과 :
	   ex1 :  520d
	   ex1 :  silver
	   ex1 :  bmw
	   ex1 : main.spec{length:4000, height:1000, width:2000}
	*/

	fmt.Println()

	fmt.Println("ex2 : ", car1.detail.height)
	fmt.Println("ex2 : ", car1.detail.length)
	fmt.Println("ex2 : ", car1.detail.width)
	/*
		결과 :
		   ex2 :  1000
		   ex2 :  4000
		   ex2 :  2000
	*/
}

```
