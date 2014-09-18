package main

import (
	"fmt"
	"math"
)

/* 使用method特性实现计算面积 */
type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r := Rectangle{10, 20}
	c := Circle{10}

	fmt.Println("Rectangle: ", r.area())
	fmt.Println("circle: ", c.area())
}
