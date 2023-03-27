package main

import "fmt"

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func a() {
	defer end(start("b"))
	fmt.Println("in a")
}

func main() {
	a()
	/*
	   결과 :
	   start :  b
	   in a
	   end :  b

	   Defer 함수는 , defer가 붙은 함수 한개만 걸린다.
	   따라서 해당 예제처럼 defer end(start("b"))
	   이렇게 해주면 , defer end() 만 defer가 걸리고 start()는
	   defer로 걸리지 않는다.
	*/
}
