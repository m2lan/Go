package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

   /**
	* 练习一
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i > 5 {
			goto kk
		}
	}
	kk:println("haha")

	array := []string{"aa", "bb", "cc", "dd", "ee"}
	for _, v := range array {
		fmt.Println(v)
	}

   /**
	* 练习二
	*/
	for i := 1; i <= 100; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			println("FizzBuzz")
		case i % 3 == 0:
			println("Fizz")
		case i % 5 == 0:
			println("Buzz")
		default:
			println(i)
		}
	}

   /**
	* 练习三
	*/
	str := "A"
	for i := 1; i <= 100; i++ {
		fmt.Printf("%s\n", str)
		str += "A"
	}

	array1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v \n", array1)

   /**
	* 练习四
	*/
	ss := "sdfdsfssddsfds哈哈"
	println(utf8.RuneCount([]byte(ss)),",", len([]byte(ss)))

	r := []rune(ss)
	copy(r[4:4+3], []rune("abc"))
	fmt.Println(ss)
	fmt.Println(string(r))

	s1 := "yuiop"
	s2 := []rune(s1)

	for i, j := 0, len(s2) - 1; i < j; i, j = i + 1, j - 1 {
		s2[i], s2[j] = s2[j], s2[i]
	}
	fmt.Println(string(s2))

   /**
	* 练习五
	*/

}