package main

import (
	"fmt"
)

/**
 * 带可变参数函数
 */
func sum(num ...int) {
	fmt.Print(num)
	total := 0
	for _, v := range num {
		total += v
	}
	fmt.Println(total)
}

func main() {
	sum(10, 20, 30)
	sum(5, 10, 15)
}
