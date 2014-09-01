package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------------------------------------------------")
	myFunc_range()
	fmt.Println("----------------------------------------------------")
	myFunc_range1()
	fmt.Println("----------------------------------------------------")
	println(unhex('b'))
	fmt.Println("----------------------------------------------------")
	myFunc_swith()
}

func myFunc_range() {
	list := []string{"a", "b", "c", "d", "e"};

	// range 是个迭代器,当被调用的时候,从它循环的内容中返回一个键值对
	for k, v := range list {
		fmt.Printf("%d---%s, ", k, v)
	}
	println()
}

func myFunc_range1() {
	for k, v := range "kms" {
		fmt.Printf("%d---%c, ", k, v)
	}
	println()
}

func unhex(c byte) byte {
	switch {
		case '0'<=c&&c<='9': return c - '0'
		case 'a'<=c&&c<='f': return c - 'a' + 10
		case 'A'<=c&&c<='F': return c - 'A' + 10
	}
	return 0
}

// 不使用fallthrough的时候，当i为0，将不会输出kkop
func myFunc_swith() {
	i := 0

	switch i {
	case 0: fallthrough
	case 1:
		println("kkop")
	}
}