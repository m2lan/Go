package main

import (
	"fmt"
)

func test() (int, int) {
	return 3, 7
}

func test1() int {
	return 1
}

func test2() (a int, b int) {
	a = 2
	b = 3

	return a, b
}

func main() {
	a, b := test()
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(test1())

	a1, b1 := test2()
	fmt.Println(a1)
	fmt.Println(b1)
}
