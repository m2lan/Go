package main

import (
	"fmt"
)

/**
 * 函数的定义和返回值
 */
func sum(a, b int) int {
	return a + b
}

func main() {
	num := sum(10, 20)
	fmt.Println("10+20=", num)
}
