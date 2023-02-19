package main

import "fmt"

func main() {
	// 변수 여러개로 선언
	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name : ", name)
	fmt.Println("height : ", height)
	fmt.Println("weight : ", weight)
	fmt.Println("isRunning ? ", isRunning)

	// 결과
	/**
	name :  machine
	height :  250
	weight :  350.56
	isRunning ?  true
	**/
}
