package main

import (
	"fmt"
)

type MM interface {
	Name() string
	Test()
}

type Test struct {
	name string
}

func (t Test) Name() string {
	return t.name
}

func (t Test) Test() {
	fmt.Println("Test")
}

func main() {
	t := Test{"name"}
	t.Test()
	dis(t)
}

func dis(m MM) {
	if t, ok := m.(Test); ok {
		fmt.Println("okk", t.name)
	}
}
