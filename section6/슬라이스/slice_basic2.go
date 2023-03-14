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

	for i, v := range slice2 {
		fmt.Printf("%d, %d\n", i, v)
	}
}
