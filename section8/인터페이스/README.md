# GO_Interface
## 인터페이스란 ?
객체의 동작을 표현하는 뼈대라고 보면 됩니다.

메서드(동작) 에 대한 방법만 표시합니다.

추상화를 제공하고 , 인터페이스 메소드를 구현한 타입은 인터페이스로 사용이 가능합니다.

**Go 언어를 유연하게 사용이 가능합니다 - 타언어와 비교할때 독보적이게 좋음**
- ***덕타이핑*** : Go 언어의 독창적인 특성

### 인터페이스를 왜 쓸까 ?
디자인패턴 측면에서 Client 입장은 , 메소드의 구체적인 클래스를 몰라도, 인터페이스에 정의된 메서드를 사용하는 객체가 보장되어있기 때문입니다.
- 클래스간의 결합도를 감소시킬 수 있습니다.
    - 결합도가 감소되면 , 유지보수성이 향상되며 기능 추가가 용이해집니다.
    - 독립적으로 프로그래밍 하는것이 가능해 집니다.

## 1. Go에서 인터페이스를 선언하는 방법
### 1.1 기본 사용 방안
기본적인 사용 방안은 아래와 같습니다.
```golang
//인터페이스를 선언하는 방법
type 인터페이스명 interface{
	메서드1() 반환값(타입 형)
	메서드2() // 반환값 없을 경우
}
```

실 예제
```golang
// interface를 사용할 땐 , 구조체랑 같은 형식으로 만들지만,
// 이름을 interface로 둡니다.
type test interface{} // 빈 인터페이스

//인터페이스를 선언하는 방법
/*
type 인터페이스명 interface{
	메서드1() 반환값(타입 형)
	메서드2() // 반환값 없을 경우
}
*/

func main() {
	var t test
	fmt.Println("ex1 : ", t) // 빈(empty) 인터페이스인 경우 , nil 리턴
}
```

### 1.2 실 사용 예
인터페이스의 실제 사용 예는 다음과 같습니다.
```golang
type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// bite 메서드 구현
func (d Dog) bite() {
	fmt.Println("Dog ", d.name, "bite !!")
}

func (c Cat) bite() {
	fmt.Println("Cat ", c.name, "bite! !")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
}

func main() {
	var inter1 Behavior

	dog1 := Dog{"poll", 10}
	inter1 = dog1
	inter1.bite()

	cat1 := Cat{"jin", 30}
	inter1 = cat1
	inter1.bite()

	fmt.Println()
	// interface를 선언하지 않고 , Behavior 메서드를 구현한 객체가 무엇인지를 넣어주는
	// 아래 방식을 더 많이 사용합니다.
	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)
	inter2.bite()

	fmt.Println()
	// interface with slice
	inters := []Behavior{dog1, dog2}
	for idx, _ := range inters {
		inters[idx].bite()
	}

	fmt.Println()
	// 값 형태로 실행(인터페이스)
	for _, val := range inters {
		// dog1.bite() , dog2.bite() 실행
		val.bite()
	}
	/*
		결과 :
			Dog  poll bite !!
		Cat  jin bite! !

		Dog  marry bite !!

		Dog  poll bite !!
		Dog  marry bite !!

		Dog  poll bite !!
		Dog  marry bite !!
	*/
}
```

### 1.3 덕타이핑
메서드만을 가지고 타입을 판단할 수 있는것을 의미합니다.
- 물고 소리를 내고 뛰고 하면 동물이다.
- 꽥꽥 소리를 내고 날고 수영하면 오리다.
- 오리처럼 걷고 , 소리내고 , 헤엄치고 행동이 같으면 오리라고 볼 수 있다.
	- 덕타이핑은 . 구조체 및 변수의 값이나 타입은 상관하지 않고, 오로지 구현된 메소드로만 판단하는 방식입니다.
	- **Go의 중요한 특징입니다.**

덕타이핑을 사용해봅시다.

Cat과 Dog 구조체의 리시버는,  모두 동일하게 bite(), sounds(), run() 메서드를 구현합니다.

Behavior 인터페이스는 bite(), sounds(), run() 메서드를 선언합니다.

따라서 Cat, Dog는 인터페이스를 구현한 구조체 입니다.

act() 함수는 Behavior 인터페이스를 파라미터로 받기 때문에 , Behavior를 구현한 Cat , Dog 구조체를 넣을 수 있으며 , act()함수에선 들어간 구조체에 맞춰서 메서드를 다르게 실행할 수 있습니다.

```golang
type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 구조체 Dog 메소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, " : Dog bite!")
}

func (d Dog) sounds() {
	fmt.Println(d.name, " : Dog barks!")
}

func (d Dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

// 구조체 Cat 메소드 구현
func (d Cat) bite() {
	fmt.Println(d.name, " : Cat 할퀴다!")
}

func (d Cat) sounds() {
	fmt.Println(d.name, " : Cat cries!")
}

func (d Cat) run() {
	fmt.Println(d.name, " : Cat is running!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()
	sounds()
	run()
}

// 인터페이스를 파라미터로 받는다.
// act 함수를 호출할 때 , interface Behavior를 구현한 메서드를 리시버로 갖고 있는
// 구조체가 들어가면 됩니다.
func act(animal Behavior) {
	animal.bite()
	animal.sounds()
	animal.run()
}

func main() {
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	// 개 행동 실행
	act(dog)

	// 고양이 행동 실행
	act(cat)

	/*
	   결과 :
	   poll  : Dog bite!
	   poll  : Dog barks!
	   poll  : Dog is running!
	   bob  : Cat 할퀴다!
	   bob  : Cat cries!
	   bob  : Cat is running!
	*/
}
```

### 1.4 익명 인터페이스 사용 예제 (즉시 선언 후 사용 가능)
익명 인터페이스를 선언해서 바로 사용합니다.
```golang
type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 구조체 Dog 메소드 구현
func (d Dog) run() {
	fmt.Println(d.name, " : Dog is running!")
}

// 구조체 Cat 메소드 구현
func (d Cat) run() {
	fmt.Println(d.name, " : Cat is running!")
}

// interface를 따로 두지 않고 , 메게변수로 넣으면서 익명 선언을 사용합니다.
func act(animal interface{ run() }) { // 익명 인터페이스. 타입 정의 안함
	animal.run()
}

func main() {
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	// 개 행동 실행
	act(dog)

	// 고양이 행동 실행
	act(cat)

	/*
	   결과  :
	   poll  : Dog is running!
	   bob  : Cat is running!
	*/
}
```

### 1.5 인터페이스 활용 - 빈 인터페이스
빈 인터페이스는 ,, 함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있습니다. -> 만능
- 모든 타입을 지정하는것이 가능합니다.

인터페이스는 모든 자료형을 받을 수 있습니다.

그니까 함수를 정의할때, 그냥 빈 인터페이스를 파라미터로 넣어두면 , 해당 함수에 어떤 타입이 들어오던간에 상관하지 않고 알아서 타입캐스팅을 해줍니다..
```golang
type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex1 : ", s)
}

func main() {
	dog := Dog{"Poll", 10}
	cat := Cat{"bob", 5}

	// 그냥 구조체 타입으로 인터페이스가 변환되어서 구조체 값을 출력합니다.
	printValue(dog)
	printValue(cat)      // struct
	printValue(15)       // int
	printValue("animal") // string
	printValue(25.5)     // float
	printValue([]Dog{})  // 슬라이스
	printValue([5]Dog{}) // 배열
	printValue(true)     // bool
	/*
	   결과 :
	   ex1 :  {Poll 10}
	   ex1 :  {bob 5}
	   ex1 :  15
	   ex1 :  animal
	   ex1 :  25.5
	   ex1 :  []
	   ex1 :  [{ 0} { 0} { 0} { 0} { 0}]
	*/
}
```