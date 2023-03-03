package main

import (
	"fmt"
)

// init은 main보다 먼저딱 한번 실행됨.
func init() {
	a := "hi"

	fmt.Println("init method start")
	fmt.Println(a)
}

func main() {
	fmt.Println("main method start")
}
