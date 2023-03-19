package main

import "fmt"

func main() {
	map1 := map[string]string{
		"daum":   "https:daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
	}

	fmt.Println("ex1 : ", map1["google"])
	fmt.Println("ex1 : ", map1["daum"])
	fmt.Println("ex1 : ", map1["naver"])

	fmt.Println()

	// 순회할 때 map은 순서가 없으므로 랜덤으로 찍힌다.
	for k, v := range map1 {
		// k 에 key , v 에 value
		fmt.Printf("ex2 : %s, %s\n", k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		// k 에 key , v 에 value
		fmt.Printf("ex2 : %s\n", v)
	}

	fmt.Println()

}
