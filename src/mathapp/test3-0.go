package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var p person
	p.name, p.age = "zhangsan", 20

	p1 := person{"wangwu", 18}
	p2 := person{"zhaoliu", 88}

	fmt.Println(Older(p, p1))
	fmt.Println(Older(p1, p2))
	fmt.Println(Older(p, p2))
}
