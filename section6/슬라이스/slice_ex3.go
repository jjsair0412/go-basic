package main

import "fmt"

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
	   make로 공간을 할당하지 않았기 때문에 복사 안먹힘.
	*/
	fmt.Println()

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
	fmt.Println()

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
	fmt.Println()
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] //0부터 4까지 슬라이싱하고 용량은 7로 변경
	f[2] = 8

	fmt.Println("ex4 : ", len(f), cap(f))
	fmt.Println("ex4 : ", f)

}
