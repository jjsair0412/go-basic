// 변수 1
package main

// fmt 패키지를 import시켜놓고 사용하지 않으면 , 저장시 지워집니다.
import "fmt"

func main() {
	//변수 선언은 초기화가 필요함
	//정수 타입 : 0. 실수 (소수점) :0.0 , 문자열 : "", Boolean : true , False 등 타입이 많음

	var a int // 선언만
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3 // 선언과 동시에 초기화 가능
	var i float32 = 11.4      // 실수형
	var j string = "Hi Golang !"
	var k = 4.75       // 자료형 타입을 지정하지 않아도 , 크기에 따라 자료형을 선언시켜준다.
	var l = "hi seoul" // string으로 지정
	var m = true       // boolean

	a = 30 // 30으로 a 변수 초기화

	// golang은 변수 선언을 해놓고 사용하지 않으면 에러가 발생합니다.
	// ~ declared and not usedcompiler

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)

}
