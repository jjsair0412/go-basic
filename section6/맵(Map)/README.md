# golang의 Map
## golang에서의 Map 
맵은 해시테이블 , 딕셔너리(python과 동일) 라고 부르기도 합니다.

key-value 로 자료를 저장합니다.

**Map은 레퍼런스 타입(참조값을 전달)** 입니다.

레퍼런스 타입이기에
- 비교연산자 사용 불가능
- 참조타입(key)로 사용이 불가능하며 , 값(value)로는 모든 타입을 사용 가능합니다.
    - key로 value를 조회하는데 , key에는 참조타입을 넣을 수 없지만, value에서는 참조타입을 넣을 수 있습니다..
- make함수 및 축약(리터럴) 로 초기화가 가능합니다.
- 삽입된 순서가 없습니다.
    - 슬라이스는 append등 되면서 순서가 있지만 , 얘는 없음

## map 사용법
### 1. 다양한 선언 방법
아래처럼 map을 선언하고 사용할 수 있습니다.

```golang
func main() {
	// key는 문자열 , 값은 int 사용하겠다 라고 map을 정의
	var map1 map[string]int = make(map[string]int) // 정석
	var map2 = make(map[string]int)                // 자료형 생략
	map3 := make(map[string]int)                   // 축약형

	fmt.Println("ex1 : ", map1)
	fmt.Println("ex1 : ", map2)
	fmt.Println("ex1 : ", map3)

	fmt.Println()

	map4 := map[string]int{} // {}를 붙이면 형태만 Json 형태 로 정의할 수 있다.
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	// 아래처럼 json 형태로 key-value를 넣어서 map을 선언할 수 있다.
	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 23,
	}

	// 메모리용량 10으로 선언
	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	// map type에서는 순서가 없기 때문에 값 순서가 바뀔 수 있음
	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["apple"])
	fmt.Println("ex2 : ", map6["banana"])
	fmt.Println("ex2 : ", map6["orange"])

}
```

### 2. 맵 조회 및 순회 ( for문 )
map type을 조회 및 순회하는 방법은 다음과 같습니다.

```golang
func main() {
	map1 := map[string]string{
		"daum":   "https:daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
	}

	fmt.Println("ex1 : ", map1["google"])
	fmt.Println("ex1 : ", map1["daum"])
	fmt.Println("ex1 : ", map1["naver"])

	fmt.Println()

	// 순회할 때 map은 순서가 없으므로 랜덤으로 찍힌다.
	for k, v := range map1 {
		// k 에 key , v 에 value
		fmt.Printf("ex2 : %s, %s\n", k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		// k 에 key , v 에 value
		fmt.Printf("ex2 : %s\n", v)
	}

	fmt.Println()
}

```

### 3. 맵 값 변경 및 삭제
map의 값을 수정 & 삭제하는 방법은 간단합니다.
- 아래와 같이 사용합니다.
```golang
func main() {

	// 맵 값 변경 및 삭제

	map1 := map[string]string{
		"daum":   "https:daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
		"home1":  "https://test1.com",
	}

	fmt.Println("ex1 : ", map1)

	// 추가할땐 그냥 아래처럼 key와 value에 맞게 값을 대입해서 초기화시키면된다.
	// append이런거 할필요 없음
	map1["home2"] = "https://test2.com"

	fmt.Println("ex1 : ", map1)

	// 수정할땐 map에 key가 있다면 그 값을 넣어준 value로 수정합니다.
	map1["home2"] = "https://test2-2.com"

	fmt.Println("ex1 : ", map1)

	fmt.Println()

	// 삭제메소드를 제공한다
	// delete(대상 map이름 , key이름) 넣어주면 삭제됩니다.
	delete(map1, "home2")

	fmt.Println("ex1 : ", map1)

	delete(map1, "google")

	fmt.Println("ex1 : ", map1)

}
```

### **4. 맵을 조회할 경우 주의할 점 - 중요 !**
map을 조회할 경우에 주의할 점이 있습니다.

```golang
func main() {

	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1, ok1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]

	fmt.Println("ex1 : ", value1, ok1)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, ok)
	/*
		결과 :

		없는 key를 리턴시켰지만 에러가 발생하지 않고 , 기본 자료형의 default값을 넣어줍니다.
		int : 0 , string : "", float : 0.0

		그런데 만약 위 예처럼 map의 value가 기본 자료형의 default 값이 들어가 있다면 , 이게 진짜
		값이 들어간건지 , 기본자료형이 들어간건지 구분하기 힘듭니다.

		따라서 , 선언할 때 두번째 리턴 값으로 키 존재의 유무를 확인하게 됩니다.

		ex1 :  0 true
		ex1 :  0
		ex1 :  0 false

		lemon은 있으니까 true를 반환하고 , kiwi는 map1에 없으니까 false를 반환합니다.
	*/

	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 :", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist! ")
	}
	/*
	   결과 :
	   ex2 : kiwi is not exist!

	*/

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 :", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist! ")
	}
	/*
	   결과 :
	   ex2 : ex2 : 115

	*/

	// key가 있는지 없는지만 확인
	if _, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : kiwi is not exist!")
	} else {
		fmt.Println("ex2 : kiwi is not exist! ")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exist! ")
	}

	/*
		결과 :
		ex2 : kiwi is not exist!
		ex2 : kiwi is not exist!
	*/
}
```
