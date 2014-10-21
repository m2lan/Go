package main

import "fmt"

func main() {
	if 10%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 9; num < 4 {
		fmt.Println("num<4")
	} else if num < 10 {
		fmt.Println("num<10")
	} else {
		fmt.Println(num)
	}
}
