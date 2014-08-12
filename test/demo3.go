package main

import "fmt"

func main() {
	
	if myFunc_switch('a') {
		println("aaaa")
	} else {
		println("bbbb")
	}

	println(compare([]byte{0x00}, []byte{0x00}))
}

func myFunc_for() {
	str := []string{"aa", "bb", "cc", "dd", "ee"};

	for k, v := range str {
		fmt.Printf("%d------%s\n", k, v);
	}
}

func myFunc_switch(c byte) bool {
	switch c {
		case 'a', 'b', 'c', 'd':
			return true
	}

	return false
}

func compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i] :
			return 1
		case a[i] < b[i] :
			return -1
		}
	}

	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}
