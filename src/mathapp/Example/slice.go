package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "red"
	s[1] = "green"
	s[2] = "blue"
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println(len(s))

	s = append(s, "black")
	fmt.Println(s)
	s = append(s, "yello", "white")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := c[2:5]
	fmt.Println(l)

	l = c[:5]
	fmt.Println(l)

	l = c[2:]
	fmt.Println(l)

	t := []string{"aaa", "bbb", "ccc"}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
