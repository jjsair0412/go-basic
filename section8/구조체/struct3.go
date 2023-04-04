package main

import "fmt"

type car struct {
	color string
	name  string
}

func main() {
	c1 := car{"red", "220d"}

	c2 := new(car)
	c2.color, c2.name = "white", "sonata"

	c3 := &car{}

	c4 := &car{"black", "520d"}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	/*
	   결과 :
	   {red 220d}
	   &{white sonata}
	   &{ }
	   &{black 520d}
	*/
}
