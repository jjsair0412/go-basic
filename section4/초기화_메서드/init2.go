package main

import (
	"fmt"
)

// init은 main보다 먼저딱 한번 실행됨.
func init() {
	fmt.Println("init method start")
}

func init() {
	fmt.Println("init2 method start")
}

func init() {
	fmt.Println("init3 method start")
}

func init() {
	fmt.Println("init4 method start")
}

func main() {
	fmt.Println("hello i'm main")
}
