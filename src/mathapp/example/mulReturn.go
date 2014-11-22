package main

import (
	"fmt"
)

/**
 * 函数多返回值
 * @return int,int
 */
func cc() (int, int) {
	return 10, 20
}

func main() {
	a, b := cc()
	fmt.Println(a, "----", b)
}
