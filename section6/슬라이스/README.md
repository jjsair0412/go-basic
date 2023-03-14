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