package main

import (
	"fmt"
)

func main() {
	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("ex1 : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println("ex2 : ", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println("ex3 : ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5], str3[6])

	fmt.Printf("ex1 : %c %c %c %c %c %c \n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("ex2 : %c %c %c %c %c \n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("ex3 : %c %c %c %c %c %c %c \n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5], str3[6])
	/*
	   결과
	   ex1 :  71 111 108 97 110 103
	   ex2 :  87 111 114 108 100
	   ex3 :  234 179 160 237 148 132 235
	   ex1 : G o l a n g
	   ex2 : W o r l d
	   ex3 : ê ³   í   ë
	*/
	conStr := []rune(str3)
	fmt.Printf("ex3 : %c %c %c %c %c %c \n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])
	/*
	   ex1 :  71 111 108 97 110 103
	   ex2 :  87 111 114 108 100
	   ex3 :  234 179 160 237 148 132 235
	   ex1 : G o l a n g
	   ex2 : W o r l d
	   ex3 : ê ³   í   ë
	   ex3 : 고 프 로 그 래 밍
	*/

	for i, char := range str1 {
		fmt.Printf("ex3 : %c(%d)\t", char, i)
	}

	fmt.Println()

	for i, char := range str2 {
		fmt.Printf("ex4 : %c(%d)\t", char, i)
	}
}
