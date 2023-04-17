package test

import "fmt"

func Test(s ...string) {
	for i, v := range s {
		fmt.Printf("index : %d , value : %s", i, v)
		fmt.Println()
	}
}
