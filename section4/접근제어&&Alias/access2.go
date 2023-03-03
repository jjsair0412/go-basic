package main

import (
	"fmt"
	more10check "mypkg/lib"
	more100check "mypkg/lib2"
)

func main() {
	fmt.Println("10보다 큰 수 ? : ", more10check.CheckNum(11))

	fmt.Println("100보다 큰 수 ? : ", more100check.CheckNum1(101))
}
