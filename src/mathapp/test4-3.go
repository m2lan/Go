package main

import (
	"fmt"
)

/* method继承 */
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("i am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	mark := Student{
		Human{
			"mark",
			25,
			"88888888",
		}, "TTTT"}
	mark2 := Employee{
		Human{
			"mark2",
			23,
			"66666666",
		}, "KKKK"}

	mark.SayHi()
	mark2.SayHi()
}
