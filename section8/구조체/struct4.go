package main

import "fmt"

func main() {

	// 아래처럼 struct를 익명으로 바로 구조체 선언하고 , 사용한다.
	car1 := struct{ name, color string }{"520d", "red"}

	fmt.Println("ex1 : ", car1)
	fmt.Printf("ex1 : %#v", car1)
	/*
	   결과 :
	   ex1 :  {520d red}
	   ex1 : struct { name string; color string }{name:"520d", color:"red"}
	*/

	cars := []struct{ name, color string }{
		{"530i", "red"},
		{"530e", "white"},
		{"530a", "blue"},
	}

	fmt.Println()
	for _, c := range cars {
		fmt.Printf("ex2 : %s, %s --- %#v\n", c.name, c.color, c)
	}

	/*
	   결과 :
	   ex2 : 530i, red --- struct { name string; color string }{name:"530i", color:"red"}
	   ex2 : 530e, white --- struct { name string; color string }{name:"530e", color:"white"}
	   ex2 : 530a, blue --- struct { name string; color string }{name:"530a", color:"blue"}
	*/
}
