package main

import (
	"fmt"
	"sort"
)

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
