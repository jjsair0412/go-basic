# 자료형 
## 1. 배열
배열은 용량과 길이가 항상 같다.

**배열과 슬라이스의 차이점을 아는것이 중요합니다.**

### 1.1 배열 vs 슬라이스
배열은 길이가 고정되어있는 반면에 , 슬라이스는 길이가 가변적 입니다.

배열은 값을 저장하고 슬라이스는 저장되어잇는 주소값을 참조합니다.

배열은 값을 복사해서 전달하고 , 슬라이스는 값이 저장되어있는 메모리 주소를 전달합니다.
- 따라서 슬라이스는 값을 수정하면 주소값이 변경됩니다.

배열은 비교연산자를 사용 가능 하지만 , 슬라이스는 비교연산자를 사용할 수 없습니다.

**따라서 대부분 슬라이스를 사용합니다.**

### 1.2 함수 종류
1. cap()
배열 , 슬라이스의 용량 반환
2. len() 
배열 , 슬라이스의 값 개수를 반환

## 2. 배열의 사용법
### 2.1 기초 사용법
아래 에처럼 다양하게 사용할 수 있습니다.
```golang
package main

import "fmt"

func main() {
	// 배열
	var arr1 [5]int                         // 선언만
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5} //선언 하면서 초기화
	var arr3 = [5]int{1, 2, 3, 4, 5}

	// 짧은선언 가능
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}         // 선언안한 빈곳은 0 으로 초기화 ( int기 때문에 )
	arr6 := [...]int{1, 2, 3, 4, 5} // ... 으로 선언하면 배열 크기를 자동으로 맞춰준다.
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	} // 2차원배열 . 지정안된건 0 출력

	arr1[2] = 5 // 2번 index를 제외한 나머지는 다 0 출력

	fmt.Printf("%-5T %d ,%v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d ,%v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d ,%v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d ,%v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d ,%v\n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d ,%v\n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T %d ,%v\n", arr7, len(arr7), arr7)

	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee", "park"}
	fmt.Printf("%-5T %d ,%v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d ,%v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d ,%v\n", arr10, len(arr10), arr10)
	/*
	   결과 :
	   [5]int 5 ,[0 0 5 0 0]
	   [5]int 5 ,[1 2 3 4 5]
	   [5]int 5 ,[1 2 3 4 5]
	   [5]int 5 ,[1 2 3 4 5]
	   [5]int 5 ,[1 2 3 0 0]
	   [5]int 5 ,[1 2 3 4 5]
	   [5][5]int 5 ,[[1 2 3 4 5] [6 7 8 9 10] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
	   [5]int 5 ,[1 2 3 4 5]
	   [5]int 5 ,[1 2 3 4 5]
	   [3]string 3 ,[kim lee park]
	*/
}

```

### 2.1 for문과 함께하는 배열
배열 자체를 range 함수 범위로 정해서 사용할 수 있습니다.

아래처럼 사용합니다.
```golang
package main

import "fmt"

func main() {
	arr1 := [5]int{1, 10, 100, 1000, 10000}

	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex1: ", arr1[i])
		/*
			결과 :
			ex1:  1
			ex1:  10
			ex1:  100
			ex1:  1000
			ex1:  10000
		*/
	}

	fmt.Println()

	arr2 := [5]int{7, 77, 777, 7777, 77777}

	for i, v := range arr2 {
		fmt.Println("ex2: ", i, v)
		/*
			결과 :
			ex2:  0 7
			ex2:  1 77
			ex2:  2 777
			ex2:  3 7777
			ex2:  4 77777
		*/

	}

	fmt.Println()

	// 인덱스 길이 생략
	// 앞에는 인덱스 번호 , 뒤에는 값 ( range라서 )
	for _, v := range arr2 {
		fmt.Println("ex3: ", v)
		/*
			결과 :
			ex3:  7
			ex3:  77
			ex3:  777
			ex3:  7777
			ex3:  77777
		*/
	}

	fmt.Println()

	// 인덱스 길이 생략 2
	for v := range arr2 {
		fmt.Println("ex4: ", v)
		/*
			결과 :
			ex4:  0
			ex4:  1
			ex4:  2
			ex4:  3
			ex4:  4
		*/
	}
}

```
## 3. 배열 검증
배열은 값을 참조하는것을 검증해봅시다.

배열안에 배열을 복사해서 , 값만 복사되고 주소값은 다른것을 확인해 봅니다.
- %p 옵션은 , 포인터 주소값을 출력합니다.
- %v 옵션은 , 주소안에 실제 값을 출력합니다.

```golang

package main

import "fmt"

func main() {
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1 // arr2 배열에 arr1 복사

	fmt.Println("ex1 : ", arr1, arr1, &arr1)
	fmt.Println("ex1 : ", arr2, arr2, &arr2)

	fmt.Printf("ex1: %p , %v\n", &arr1, arr1)
	fmt.Printf("ex1: %p , %v\n", &arr2, arr2)
	/*
		결과 :
		ex1 :  [1 10 100 1000 10000] [1 10 100 1000 10000] &[1 10 100 1000 10000]
		ex1 :  [1 10 100 1000 10000] [1 10 100 1000 10000] &[1 10 100 1000 10000]
		ex1: 0xc000012630 , [1 10 100 1000 10000]
		ex1: 0xc000012660 , [1 10 100 1000 10000]

		주소값은 다르지만 , 값은 동일한것을 볼 수 있습니다.

		따라서 값을 복사했습니다.

		배열은 값을 참조하기에 값만 복사합니다.
	*/
}

```