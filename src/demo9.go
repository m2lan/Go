package main

<<<<<<< HEAD
import (
	"fmt"
)

func main() {
	fmt.Println("demo9")
=======
import "fmt"

func main() {

	fmt.Println(test(10)(20)(30))

	test1(1, 2, 3, 4, 5, 6)
}

func test(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

func test1(a ...int) {
	fmt.Println(a)
>>>>>>> 18255b2cbf6d68e0fe4d0fe56b2c062bb3bc5b30
}
