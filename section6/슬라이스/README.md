# 슬라이스
## 슬라이스 ?
배열과 비슷하지만 차이가 있다.
1. 가변형
동적으로 크기가 늘어납니다.
2. 참조 타입
값 자체를 참조하지 않고 , 값이 저장된 주소값을 참조합니다.

따라서 슬라이스를 배열보다 많이 사용합니다.

슬라이스는 **길이**와 **용량**의 개념이 있습니다.

슬라이스는 참조 값 자체를 전달합니다.

참조 타입이기에 비교연산자를 사용하지 못합니다.

## 슬라이스 사용방안
슬라이스의 사용 방안은 배열처럼 선언하는 방법과 , make 함수를 사용하는 방법 두 가지가 있습니다.

make 함수를 사용하면 , 아래처럼 파라미터를 넣어줍니다.
```golang
make(자료형 , 길이 , 용량)
```

이때 용량을 생략하면 길이가 가변적으로 늘어나거나 줄어듭니다.
### 1. 배열처럼 사용하는 방법
아래처럼 사용하면 됩니다.
```golang
package main

import (
	"fmt"
)

func main() {
	// 가변길이기 때문에 길이를 할당하지 않습니다.
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}

	// 2차원 슬라이스
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	// slice2[0] = 1 에러발생 . 길이가 가변이기 때문에 배열에 원소를 넣지 않는다면 index out of range 에러가 발생합니다.
	slice3[4] = 10 // 값 수정 가능 . 짧은선언 사용 불가

	fmt.Printf("%-5T %d, %d ,%v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d, %d ,%v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d, %d ,%v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d, %d ,%v\n", slice4, len(slice4), cap(slice4), slice4)
	/*
	   결과 :
	   []int 0, 0 ,[]
	   []int 0, 0 ,[]
	   []int 5, 5 ,[1 2 3 4 10]
	   [][]int 2, 2 ,[[1 2 3 4 5] [6 7 8 9 10]]

	*/
}
```

### 2. make 사용 방안
make 함수를 사용하여 slice를 만드는 예 입니다.
- 주석으로 설명을 대신합니다.
```golang
package main

import (
	"fmt"
)

func main() {
	// 가변길이기 때문에 길이를 할당하지 않습니다.
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}

	// 2차원 슬라이스
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	// slice2[0] = 1 에러발생 . 길이가 가변이기 때문에 배열에 원소를 넣지 않는다면 index out of range 에러가 발생합니다.
	slice3[4] = 10 // 값 수정 가능 . 짧은선언 사용 불가

	fmt.Printf("%-5T %d, %d ,%v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d, %d ,%v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d, %d ,%v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d, %d ,%v\n", slice4, len(slice4), cap(slice4), slice4)
	/*
	   결과 :
	   []int 0, 0 ,[]
	   []int 0, 0 ,[]
	   []int 5, 5 ,[1 2 3 4 10]
	   [][]int 2, 2 ,[[1 2 3 4 5] [6 7 8 9 10]]

	*/

	// 자료형 int , 길이 5 , 총 용량 10 확보.
	// 슬라이스에 원소를 11개 넣으면 , 11개로 늘어난다.
	// 그러나 메모리공간을 재 할당 해 주어야하기 때문에 , 속도저하 발생
	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	// make 함수로 미리 용량을 정해준 구간에 원소를 넣어줬기때문에 , 2번 원소를 제외한
	// 나머지 공간은 0으로 초기화 됩니다.
	slice6[2] = 7

	fmt.Printf("%-5T %d, %d ,%v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d, %d ,%v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d, %d ,%v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d, %d ,%v\n", slice8, len(slice8), cap(slice8), slice8)
	/*
	   결과 :
	   []int 5, 10 ,[0 0 0 0 0]
	   []int 5, 5 ,[0 0 7 0 0]
	   []int 5, 100 ,[0 0 0 0 0]
	   []int 5, 5 ,[0 0 0 0 0]

	   배열과 다르게 미리 용량을 정해주었기 때문에 , 초기화되지 않은 메모리 공간은 0으로 초기화 됩니다.
	*/

```

## 슬라이스가 진짜 참조 타입인지 증명
슬라이스와 배열을 비교하며 메모리 참조 타입인지 증명합니다.

```golang
package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Println("ex1 : ", arr1)
	fmt.Println("ex1 : ", arr2)

	/*
	   배열이라 값 자체를 복사하기에 아래와 같은 결과 출력
	   결과 :
	   ex1 :  [1 2 3]
	   ex1 :  [7 2 3]
	*/

	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7

	fmt.Println("ex1 : ", slice1)
	fmt.Println("ex1 : ", slice2)
	/*
	   주소값 자체를 복사하기에 , 같은 메모리를 가르킨다.
	   따라서 아래와 같은 결과가 출력
	   - 메모리 번지수 자체를 slice2에 넘김
	   ex1 :  [7 2 3]
	   ex1 :  [7 2 3]
	*/

	slice3 := make([]int, 50, 100)
	fmt.Println("ex3 : ", slice3[4])
	// fmt.Println("ex3 : ", slice3[50]) index out of range 에러발생. 50개만 사용할 수 있기 때문에. 0~49
}
```

## 슬라이스 추가 및 병합
슬라이스를 copy하면 메모리값이 복사되는 단점이 있습니다.

이것을 실제로 값이 복사되는것 처럼 하는 방안에 대해 기술합니다.

### 1. 추가
1. append 함수
첫째 파라미터로 오는 슬라이스에 , 두번째부터 순차적으로 값을 추가합니다.

추가된 값 만큼 메모리공간을 가변적으로 늘립니다.

```golang
func main() {
	s1 := []int{1, 2, 3, 4, 5}

	// append : 첫째 파라미터로 오는 슬라이스에 , 두번째부터 순차적으로 값을 추가한다.
	s1 = append(s1, 6, 76)

	// 결과 : ex1 :  [1 2 3 4 5 6 76]
	fmt.Println("ex1 : ", s1)

}
```

그런데 만약 , 5개 용량과 10개 최대 용량을 가진 slice를 make 함수를 사용해서 만들었다고 생각해 봅시다.

그리고 10개 원소가 들어간 slice에 append로 새로운 원소를 집어넣으면 , **메모리공간이 11이 되는것이 아니라 , 2배인 20이 할당됩니다.**

따라서 make()함수를 사용하여 slice를 선언할 때 , 주의해야합니다.

증명 예제 :
```golang
func main() {
	s4 := make([]int, 0, 5)
	for i := 0; i <= 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex4 -> len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4)
	}
}
```

결과 : 
```golang
ex4 -> len : 1, cap : 5, value : [0]
ex4 -> len : 2, cap : 5, value : [0 1]
ex4 -> len : 3, cap : 5, value : [0 1 2]
ex4 -> len : 4, cap : 5, value : [0 1 2 3]
ex4 -> len : 5, cap : 5, value : [0 1 2 3 4]
ex4 -> len : 6, cap : 10, value : [0 1 2 3 4 5]
ex4 -> len : 7, cap : 10, value : [0 1 2 3 4 5 6]
ex4 -> len : 8, cap : 10, value : [0 1 2 3 4 5 6 7]
ex4 -> len : 9, cap : 10, value : [0 1 2 3 4 5 6 7 8]
ex4 -> len : 10, cap : 10, value : [0 1 2 3 4 5 6 7 8 9]
ex4 -> len : 11, cap : 20, value : [0 1 2 3 4 5 6 7 8 9 10]
ex4 -> len : 12, cap : 20, value : [0 1 2 3 4 5 6 7 8 9 10 11]
ex4 -> len : 13, cap : 20, value : [0 1 2 3 4 5 6 7 8 9 10 11 12]
ex4 -> len : 14, cap : 20, value : [0 1 2 3 4 5 6 7 8 9 10 11 12 13]
ex4 -> len : 15, cap : 20, value : [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14]
ex4 -> len : 16, cap : 20, value : [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
```
**최대 용량이 10인데 , 11이 넘어가면 최대 용량이 10x2 인 20이 되는것을 확인할 수 있습니다.**

### 1.1 slice에 slice를 append로 삽입할 경우
아래처럼 사용합니다.

병합할 슬라이스에 ... 구문을 추가해야 합니다.
```golang
func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	// append : 첫째 파라미터로 오는 슬라이스에 , 두번째부터 순차적으로 값을 추가한다.
	s1 = append(s1, 6, 7)
	//  append에서 slice안에 slice를 append 시킬땐 ... 구문을 추가해야 합니다.
	// s1에다가 s2를 추가해서 s2 슬라이스에 대입
	s2 = append(s1, s2...)
	s3 = append(s2, s3[0:3]...)

	fmt.Println("ex1 : ", s1)
	fmt.Println("ex1 : ", s2)
	fmt.Println("ex1 : ", s3)

	/*
	   결과 :
	   ex1 :  [1 2 3 4 5 6 7]
	   ex1 :  [1 2 3 4 5 6 7 8 9 10 11 12]
	   ex1 :  [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
	*/
}
```

## 슬라이스 추출 및 정렬
### 1. 추출
- slice[i:j] 라는 의미는 , i부터 j-1 까지 추출하라는 의미입니다.

- slice[i:] 라는 의미는 , i부터 끝까지 추출하라는 의미입니다.

- slice[:i] 라는 의미는 , 처음부터 i-1번째까지 추출하라는 의미입니다.

- slice[:] 라는 의미는 , 처음부터 끝까지 추출하라는 의미입니다.

```golang
func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("ex1 : ", slice1[:])
	fmt.Println("ex1 : ", slice1[0:])
	fmt.Println("ex1 : ", slice1[:5])
	fmt.Println("ex1 : ", slice1[:len(slice1)]) // slice1의 길이까지
	fmt.Println("ex1 : ", slice1[3:])
	fmt.Println("ex1 : ", slice1[:3])
	fmt.Println("ex1 : ", slice1[1:3])
	/*
	   결과 :
	   ex1 :  [1 2 3 4 5 6 7 8 9 10]
	   ex1 :  [1 2 3 4 5 6 7 8 9 10]
	   ex1 :  [1 2 3 4 5]
	   ex1 :  [1 2 3 4 5 6 7 8 9 10]
	   ex1 :  [4 5 6 7 8 9 10]
	   ex1 :  [1 2 3]
	   ex1 :  [2 3]
	*/
}

```
### 2. 정렬
golang에서 정렬하기위해선 **sort package**를 사용해야 합니다.

sort 정렬 메서드를 이용해서 , int 숫자와 string 문자를 자동으로 정렬시켜줍니다.

예제 :

```golang
package main

import (
	"fmt"
	"sort"
)

func main() {
	slice2 := []int{3, 6, 10, 9, 1, 4, 5, 8, 2, 7}
	slice3 := []string{"b", "f", "a", "c", "e"}

	fmt.Println("ex2 : ", sort.IntsAreSorted(slice2)) // 정렬되있냐 확인 . false 추출
	sort.Ints(slice2)                                 // 바로 int 정렬
	fmt.Println(slice2)
	// 결과 : [1 2 3 4 5 6 7 8 9 10]

	fmt.Println()

	fmt.Println("ex2 : ", sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println("ex2 : ", slice3)
}
```

결과확인 :
```golang
ex2 :  false
[1 2 3 4 5 6 7 8 9 10]

ex2 :  false
ex2 :  [a b c e f]
```

## 슬라이스 복사
슬라이스를 복사해서 , 실제 값을 가져오는 방법에 대한 예제입니다.

copy(복사대상 , 원본) 메서드를 사용합니다.

**반드시 make로 공간을 할당한 이후 복사해야 복사가 먹힙니다.**

**복사된 slice 값을 변경해도 , 원본은 영향이 없습니다.**
- 메모리공간이 다르기 때문에 !!
- 핵심

```golang
func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1) //slice2번에 slice1번 복사
	copy(slice3, slice1) //slice3번에 slice1번 복사

	fmt.Println("ex1 : ", slice2)
	fmt.Println("ex1 : ", slice3)

	/*
	   결과 :
	   ex1 :  [1 2 3 4 5]
	   ex1 :  []
	   make로 길이와 공간을 모두 5로만 잡았기에 , 값이 다섯개만 들어감
	   make로 공간을 할당하지 않았기 때문에 복사 안먹힘
	*/
	fmt.Println()

}
```

또한 중요한 특징인 , 메모리공간이 다르기에 값이 다르다는 특징에 대한 예 :
```golang
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)
	b[0] = 7
	b[4] = 10

	fmt.Println("ex2 : ", a)
	fmt.Println("ex2 : ", b)

	/*
	   결과 :
	   a와 b가 복사됐더라도 메모리공간이 다르기 때문에 , 값이서로 다름
	   ex2 :  [1 2 3 4 5]
	   ex2 :  [7 2 3 4 10]
	*/

}

```
**copy 주의**
슬라이스한 값 ( 부분추출 ) 을 다른 슬라이스에 복사하면 , 메모리 참조기 때문에 같은 메모리공간을 바라봅니다 !!!
```golang
func main() {
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3]

	d[1] = 7

	fmt.Println("ex2 : ", c)
	fmt.Println("ex2 : ", d)
	/*
	   슬라이스 부분추출해서 가지고오면 , 참조타입으로 갖고온다 !!
	   따라서 메모리 주소가 동일해서 , 같은 곳을 바라보게 된다.

	   ex2 :  [1 7 3 4 5]
	   ex2 :  [1 7 3]
	*/

}

```

**slice 특징**
슬라이싱 할 때 , 용량을 지정해서 복사할 수 있습니다.
- 물론 용량을 지정해줫다해서 , 슬라이싱해서 넣은거기 때문에 메모리 참조 타입으로 복사됩니다.

```golang
func main() {
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] //0부터 4까지 슬라이싱하고 용량은 7로 변경
	f[2] = 8

	fmt.Println("ex4 : ", len(f), cap(f))
	fmt.Println("ex4 : ", f)
}
```

결과 :
```golang
ex4 :  5 7
ex4 :  [1 2 8 4 5 6 7 8 9 10]
ex4 :  [1 2 8 4 5]
```
