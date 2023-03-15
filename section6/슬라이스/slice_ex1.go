package main

import "fmt"

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

	s4 := make([]int, 0, 5)
	for i := 0; i <= 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex4 -> len : %d, cap : %d, value : %v\n", len(s4), cap(s4), s4)
	}
}
