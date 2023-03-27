package main

import "fmt"

// 매개변수를 선언할 때 , 아래처럼 ...을 넣어준다면 ,
// int형 몇개가 들어와도 상관이 없다는 의미이다.
// 가변형으로 매개변수를 받을 수 있다.
func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}

	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}

	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2 : ", value)
	}
}

func main() {
	// 1x2x3 의 결과 6 출력
	x := multiply(1, 2, 3)
	fmt.Println(x)

	fmt.Println()

	// 1~10까지 더한 결과 55 출력
	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(y)

	fmt.Println()

	/*
		결과 :
		ex2 :  a
		ex2 :  apple
		ex2 :  test
		ex2 :  golang
		ex2 :  seoul
	*/
	prtWord("a", "apple", "test", "golang", "seoul")

	fmt.Println()
	a := []int{1, 2, 3, 4, 5}

	// 슬라이스 a의 인덱스 0번부터 차례대로 들어간다.
	m := multiply(a...)
	n := sum(a...)

	/*
		결과 :
		ex3 :  120
		ex3 :  15
	*/
	fmt.Println("ex3 : ", m)
	fmt.Println("ex3 : ", n)
}
