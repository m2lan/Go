package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("this is 2")
	case 3:
		fmt.Println("this is 3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Thursday:
		fmt.Println(int64(time.Now().Weekday()))
	default:
		fmt.Println(time.Monday)
	}
}
