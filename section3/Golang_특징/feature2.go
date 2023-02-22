package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print("ex1 : ")
		fmt.Println(i)
	}

	fmt.Println()

	for j := 10; j >= 0; j-- {
		fmt.Println("ex2 : ", j)
	}
}
