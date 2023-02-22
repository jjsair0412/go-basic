package main

import "fmt"

func main() {
	sum, i := 0, 0

	for i <= 100 {
		// sum += i++ // 반환값이 없어서 에러 발생

		sum += i
		i++
	}

	fmt.Println("sum : ", sum)

}
