package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"zhangsan", 20})
	fmt.Println(Person{name: "wangwu", age: 30})
	fmt.Println(Person{name: "lisi"})
	fmt.Println(&Person{name: "mmmm", age: 50})

	s := Person{name: "kkkk", age: 60}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 66
	fmt.Println(sp.age)
}
