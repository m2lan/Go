package main

import (
	"fmt"
)

/* 计算长方形面积 */

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	fmt.Println(area(Rectangle{10, 20}))
}
