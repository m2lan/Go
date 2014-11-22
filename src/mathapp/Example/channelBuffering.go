package main

import (
	"fmt"
)

func main() {

	// 创建channel缓冲区
	messages := make(chan string, 2)

	messages <- "hello"
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
