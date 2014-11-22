package main

import (
	"fmt"
)

type testFunc func(int) bool

func init() {
	fmt.Println("----init----")
}

func odd(a int) bool {
	if a%2 == 0 {
		return false
	}
	return true
}

func even(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, t testFunc) []int {
	var result []int
	for _, n := range slice {
		if t(n) {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice=", slice)
	fmt.Println("filter---odd---", filter(slice, odd))
	fmt.Println("filter---even---", filter(slice, even))
}
