# 구조체 심화 - 각종 패턴
## 구조체 생성자 패턴 예
마치 자바의 class로 인스턴스화 시키는 것 처럼 , golang에서도 함수를 만들어서 특정 구조체 인스턴스의 주소값을 리턴 시키면서 기존 객체지향 언어처럼 사용할 수 있습니다.
- 아래 예시처럼 구조체를 사용하는 방법엔 다양한 방법이 있지만 , 함수를 통해 기존 객체지향처럼 사용하는 패턴이 가장 많이 사용됩니다.
```golang
type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account { // 포인터 반환 아닌경우 , 값 복사
	return &Account{ // 구조체 인스턴스를 생성한 후 , 주소값을 리턴
		number,
		balance,
		interest,
	}
}

func main() {
	// 구조체 선언과 동시에 입력
	// 생성자 활용 ( 맴버변수에 접근 )
	kim := Account{
		number:   "245-901",
		balance:  100000,
		interest: 0.015,
	}

	// 포인터로 객체 생성
	// 생성자가 없는 객체를 인스턴스화
	// getter , setter와 비슷
	var lee *Account = new(Account)
	lee.number = "245-900"
	lee.balance = 100000
	lee.interest = 0.015

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	/*
		결과 :
		ex1 :  {245-901 100000 0.015}
		ex1 :  &{245-900 100000 0.015}
	*/

	// 마치 자바의 class를 인스턴스화 하는것과 비슷
	park := NewAccount("245-903", 1700000, 0.04)
	fmt.Println("ex2 : ", park)

	/*
		결과 :
		ex2 :  &{245-903 1.7e+06 0.04}
	*/

}
```
## 구조체 리시버 구현
포인터형 리시버는 실제 값을 참조할 수 있기 때문에 , 값이 변화합니다.
```golang
type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	// 정리 : 구조체 인스턴스 값 변경시 -> 포인터 전달, 보통의 경우 -> 값 전달
	kim := Account{
		"245-901",
		100000,
		0.015,
	}

	lee := Account{
		"245-901",
		100000,
		0.015,
	}

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	/*
		결과 :
		ex1 :  {245-901 100000 0.015}
		ex1 :  {245-901 100000 0.015
	*/
	fmt.Println()

	lee.CalculateD(100000)
	kim.CalculateP(100000)

	fmt.Println("ex2 : ", int(lee.balance))
	fmt.Println("ex2 : ", int(kim.balance))

	/*
		결과 :
		ex2 :  100000
		ex2 :  201500
		값의 의한 전달은 , 실제 객체(kim 인스턴스) 의 주소값을 참조할 수 없기 때문에
		값이 변하지 않았지만 ,

		포인터 리시버는(참조에 의한 전달) 실제 객체의 주소값을  참조할 수 있기 때문에
		값이 변합니다.
	*/
}
```
## 구조체 상속 -> 임베디드 패턴
- 구조체 임베디드 패턴
상속을 허용하지 않는 Go 언어에서 메서드 상속 활용을 위한 패턴입니다.

구조체를 멤버변수로 받아서 상속처럼 사용합니다.
```golang
type Employee struct {
	name   string
	salary float64
	bonus  float64
}

// (임원도 직원이다.)
type Executives struct {
	Employee     // 구조체를 멤버변수로 받음 -> is a 관계라 한다.
	specialBonus float64
}

func (e Employee) Caculate() float64 {
	return e.salary + e.bonus
}

func main() {
	ep1 := Employee{
		"kim",
		2000000,
		150000,
	}

	ep2 := Employee{
		"park",
		150000,
		200000,
	}

	ex := Executives{
		Employee{
			"lee",
			5000000,
			1000000,
		},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Caculate()))
	fmt.Println("ex1 : ", int(ep2.Caculate()))
	/*
	   결과 :
	   ex1 :  2150000
	   ex1 :  350000
	*/

	//Employee 구조체 통해서 메소드 호출
	// 상위 구조체 리시버를 그냥 접근할 수 있습니다.
	fmt.Println("ex1 : ", int(ex.Caculate())+int(ex.specialBonus))
	/*
		결과 :
		ex1 :  7000000
	*/
}
```

## 구조체 임베디드 메소드 오버라이딩 패턴
최상위 객체의 멤버변수에 접근해서 , 자신만의 리시버 객체를 만들 수 있습니다.

```golang

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

type Executives struct {
	Employee
	specialBonus float64
}

func (e Employee) Caculate() float64 {
	return e.salary + e.bonus
}

// salary와 bonus는 최상위 구조체인 Employee의 멤버변수 인데,
// Employee 객체에 오버라이딩되어 접근합니다. -> 재정의
// 리시버의 이름은 동일해야 합니다.
func (e Executives) Caculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

func main() {
	ep1 := Employee{
		"kim",
		2000000,
		150000,
	}

	ep2 := Employee{
		"park",
		150000,
		200000,
	}

	ex := Executives{
		Employee{
			"lee",
			5000000,
			1000000,
		},
		1000000,
	}
	fmt.Println("ex1 : ", int(ep1.Caculate()))
	fmt.Println("ex1 : ", int(ep2.Caculate()))

	fmt.Println("오버라이딩 시작 ! ")

	fmt.Println("ex1 : ", int(ex.Caculate()))

	/*
		결과 :
		ex1 :  2150000
		ex1 :  350000

		오버라이딩 시작 !
		메서드가 오버라이딩 되어 , 계산됩니다.
		ex1 :  7000000
	*/
	fmt.Println()

}

```