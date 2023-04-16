package main

import "fmt"

// 이런식으로 포인터형 인지 아닌지를 체크해서 로직을 따로 짤 수 도 있습니다.
// func checkType(arg interface{}) {
// 	switch arg.(type) {
// 	case Account:
// 	case *Acccount:
// 	}
// }

func checkType(arg interface{}) {
	// 형변환 방법 : arg.(type) 으로 현재 데이터타입 변환
	switch arg.(type) {
	case bool:
		fmt.Println("This is a bool : ", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int : ", arg)
	case float64:
		fmt.Println("This is a float : ", arg)
	case nil:
		fmt.Println("This is a nil : ", arg)
	case string:
		fmt.Println("This is a string : ", arg)
	default:
		fmt.Println("What is this Type ?", arg)
	}

}

func main() {
	checkType(true)
	checkType(1)
	checkType(22.542)
	checkType(nil)
	checkType("hello golang")
}
