package main

import (
	"fmt"
	"mypkg/lib"
	"mypkg/lib2"
)

func main() {
	fmt.Println("10보다 큰 수 ? : , ", lib.CheckNum(11))

	fmt.Println("100보다 큰 수? : ", lib2.CheckNum1(101))
}
