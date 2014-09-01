package main

import "fmt"

func main() {

	var ints [10]int

	ints[0] = 22
	ints[1] = 33

	for k, v := range ints {

		fmt.Printf("%d ----- %d\n", k, v)
	}

	array := []int{11, 22, 33}

	array1 := [3]int{11, 22, 33}

	array2 := [...]int{11, 22, 33}

	println(len(array), len(array1), len(array2))

	array3 := [3][2]int{{1, 2}, {2, 3}, {3, 4}}

	for k, v := range array3 {
		for kk, vv := range v {
			fmt.Printf("%d ---- %d ---- %d\n", k, kk, vv)
		}
	}
}