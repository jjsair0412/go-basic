package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	// for {
	// 	fmt.Println(" hello golang ")
	// 	fmt.Println(" infinite loop ")
	// }

	// python처럼 range 가능
	loc := []string{"Seoul", "Busan", "Incheon"}
	for _, name := range loc {
		fmt.Println("ex3 : ", name)
	}
}
