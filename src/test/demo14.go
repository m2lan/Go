package main

import (
	"fmt"
)

func main() {
	test3()
}

func test1() {
	c := make(chan bool)
	go func() {
		fmt.Println("fuck !!! ")
		c <- true
	}()
	<-c
}

func test2() {
	c := make(chan bool)
	go func() {
		fmt.Println("fuck fuck!!!")
		c <- true
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func test3() {
	num := 0
	for i := 0; i < 1000000000; i++ {
		num += i
	}
	fmt.Println(num)
}
