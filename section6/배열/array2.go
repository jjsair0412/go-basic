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
