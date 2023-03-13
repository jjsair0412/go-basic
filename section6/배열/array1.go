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
