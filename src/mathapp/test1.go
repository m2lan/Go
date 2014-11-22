package main

import (
	"fmt"
)

func main() {

	// var mymap map[string]int = make(map[string]int)
	// mymap["c#"] = 20
	// mymap["php"] = 30
	// mymap["java"] = 40

	mymap := map[string]int{"c#": 20, "php": 30, "java": 40}

	aaa, ok := mymap["c#"]

	if ok {
		fmt.Println(aaa)
	}

	fmt.Println(mymap)
	fmt.Println(max(10, 20))
	fmt.Println(sumAndProduct(10, 20))
	fmt.Println(moreArgs(10, 20, 30, 40))
	x := 10
	fmt.Println("x=", x)
	fmt.Println("add=", add(&x))
	fmt.Println("x=", x)

	// 后进先出
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, ",")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func moreArgs(a ...int) int {
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}

func add(a *int) int {
	*a = *a + 1
	return *a
}
