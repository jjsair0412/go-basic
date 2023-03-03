package main

// 선언방법
// import "fmt"
// import "os"

// 선언방법2
import (
	"fmt"
	"os"
)

func main() {
	var name string

	fmt.Print("이름은 ? :")

	fmt.Scanf("%s", &name)

	fmt.Fprintf(os.Stdout, "Hi %s\n", name)

}
