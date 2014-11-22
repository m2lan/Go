package main

import (
	"errors"
	"fmt"
)

const v = 100 + iota

const (
	x = 100 + iota
	y
	z
)

func main() {
	str := "hello"
	c := []byte(str)
	c[0] = 'm'
	str2 := string(c)
	fmt.Println(str2)

	str = "c" + str[1:]
	fmt.Println(str)

	if err := errors.New("hello world"); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d\n", v)

	fmt.Printf("%d---%d---%d\n", x, y, z)

	var mm [10]float32
	fmt.Println(mm)

	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(slice)

	// var numbers map[string]int = make(map[string]int)
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println(numbers)

	language := map[string]float32{"java": 22.5, "c#": 11, "php": 44}
	exitst, ok := language["java"]
	if ok {
		fmt.Println("ok", exitst)
	} else {
		fmt.Println("no ok")
	}

	// sum := 0
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)
}
