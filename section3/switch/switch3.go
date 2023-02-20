package main

import "fmt"

func main() {
	a := 30 / 15

	switch a {
	case 2, 4, 6: // a가 2 , 4 , 6일 경우 true
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: // a가 1 , 3 , 5일 경우 true
		fmt.Println("a ->", a, "는 홀수")
	}

	//예제 2
	switch e := "go"; e {
	case "java":
		fmt.Println("Java")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
		fallthrough
	case "ruby":
		fmt.Println("ruby")
		fallthrough
	case "c":
		fmt.Println("c")
		// fallthrough
	}
}
