package main

import (
	"fmt"
)

type kkk interface {
	say()
}

type person struct {
}

func (p person) say() {
	fmt.Println("self")
}

func main() {
	p := person{}
	p.say()
}
