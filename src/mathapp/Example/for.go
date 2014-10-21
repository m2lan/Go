package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 3
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("loop")
		break
	}
}
