package main

import "fmt"

func main() {

	// 맵 값 변경 및 삭제

	map1 := map[string]string{
		"daum":   "https:daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
		"home1":  "https://test1.com",
	}

	fmt.Println("ex1 : ", map1)

	// 추가할땐 그냥 아래처럼 key와 value에 맞게 값을 대입해서 초기화시키면된다.
	// append이런거 할필요 없음
	map1["home2"] = "https://test2.com"

	fmt.Println("ex1 : ", map1)

	// 수정할땐 map에 key가 있다면 그 값을 넣어준 value로 수정합니다.
	map1["home2"] = "https://test2-2.com"

	fmt.Println("ex1 : ", map1)

	fmt.Println()

	// 삭제메소드를 제공한다
	// delete(대상 map이름 , key이름) 넣어주면 삭제됩니다.
	delete(map1, "home2")

	fmt.Println("ex1 : ", map1)

	delete(map1, "google")

	fmt.Println("ex1 : ", map1)

}
