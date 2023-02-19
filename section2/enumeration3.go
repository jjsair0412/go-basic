package main

import "fmt"

func main() {
	const (
		_ = iota // 0
		A        // 1
		B        // 2
		C        // 3
		_        // 스킵
		D        // 5
	)

	const (
		_        = iota + 0.75*2 // 0 + 0.75 * 2
		DEFAULT                  // 1 + 0.75 * 2
		SIVER                    // 2 + 0.75 * 2
		_                        // 생략
		PLATINUM                 // 4 + 0.75 * 2

	)

	fmt.Println("Default : ", DEFAULT)
	fmt.Println("Slver : ", SIVER)
	fmt.Println("platinum : ", PLATINUM)
}
