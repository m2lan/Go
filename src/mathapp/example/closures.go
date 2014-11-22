package main

import (
	"fmt"
)

/**
 * 闭包
 */
func c() func() int {
	i := 1
	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	newInt := c()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	newInt2 := c()
	fmt.Println(newInt2())
}
