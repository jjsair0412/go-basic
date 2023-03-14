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

	var slice9 []int // nil 슬라이스 ( 길이와 용량이 0 )
	if slice9 == nil {
		fmt.Println("this is nil slice")
	}
}
