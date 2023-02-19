package main

import "fmt"

func main() {
	const a string = "test1"
	const b = "test2"
	const c int32 = 10 * 10

	// 에러 발생
	// const d = getHeight()

	const e = 35.6
	const f = false

	fmt.Println(
		"a : ", a,
		"b : ", b,
		"c : ", c,
		"e : ", e,
		"f : ", f)
}
