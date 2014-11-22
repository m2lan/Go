package main

import (
	"fmt"
)

/* 内置类型和自定义类型做为struct匿名字段 */

type Str []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Str
	int
	spec string
}

func main() {
	jane := Student{
		Human: Human{
			"zhangsan",
			20,
			120,
		},
		Str:  []string{"hahaha"},
		spec: "spec",
	}
	fmt.Println(jane.name)
	fmt.Println(jane.age)
	fmt.Println(jane.weight)
	fmt.Println(jane.Str)
	fmt.Println(jane.int)
	fmt.Println(jane.spec)
}
