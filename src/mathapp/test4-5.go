package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
}

type Employee struct {
	Human
}

func (s *Human) sayHi() {
	fmt.Printf("i am %s, age %d", s.name, s.age)
}

func main() {
	stu := Student{Human{"zhangsan", 20}}
	stu.sayHi()
}
