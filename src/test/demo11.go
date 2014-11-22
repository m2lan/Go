package main

import (
	"fmt"
)

type pperson struct {
	kk string
}

type person struct {
	pperson
	Name    string
	Age     int
	Address string
}

type person2 struct {
	pperson
	Name    string
	Age     int
	Contact struct {
		City, Phone string
	}
}

func main() {

	test2()

	p2 := person2{
		Name: "oop",
		Age:  33,
		pperson: pperson{
			kk: "kk",
		},
	}
	p2.Contact.City = "sz"
	p2.Contact.Phone = "13333333333"
	fmt.Println(p2)

	test1()
}

func test1() {
	test := person{
		Name:    "zs",
		Age:     10,
		Address: "sz",
		pperson: pperson{
			kk: "kkk",
		},
	}

	fmt.Println(test)
}

func test2() {
	test2 := struct {
		Name string
		Age  int
	}{
		Name: "ls",
		Age:  22,
	}
	fmt.Println(test2)
}
