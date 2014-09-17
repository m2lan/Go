package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	str string
}

func main() {
	mark := Student{Human{"zhangsan", 20, 120}, "xxxx"}
	fmt.Println(mark.name)
	fmt.Println(mark.age)
	fmt.Println(mark.weight)
	fmt.Println(mark.str)

	mark.str = "yyyy"
	fmt.Println(mark.str)
}
