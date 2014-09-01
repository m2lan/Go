package main

import (
	"fmt"
	"os"
)

func main() {
	var str string = "Hello world"

	fmt.Println(str)

	/**
     * 定义常量，使用iota生成枚举值
	 */
	/*
	const (
		a = iota
		b = iota
	)

	fmt.Printf("%d-----%d\n", a, b)
	*/

	fmt.Println("----------------------------------------------------")

	/*
	aa := 10

	fmt.Println(aa)
	*/

	fmt.Println("----------------------------------------------------")

	/**
     * 修改字符串中字符 
	 */
	/* 

	content := "huahuahua"

	// 将字符串content转换为rune数组
	c := []rune(content)

	// 修改数组第一个元素
	c[0] = 'X'

	// 创建新的字符串contents保存修改
	contents := string(c)
	fmt.Printf("%s\n", contents)

	*/

	fmt.Println("----------------------------------------------------")

	// 多行字符串
	// content := "my name is Jack," +
			   // " age is 22"
	content := `my name is Jack,
	             age is 22`
	fmt.Println(content)

	fmt.Println("----------------------------------------------------")

	/**
	 * 控制语句
	 */
	 if x := 10; x > 9 {
	 	fmt.Println(x)
	 } else {
	 	fmt.Println(9)
	 }

	 if true && true {
	 	fmt.Println("true")
	 }

	 if ! false {
	 	fmt.Println("true")
	 }

	 fmt.Println(os.O_RDONLY)

	 fmt.Println("----------------------------------------------------")

	 myFunc_goto()

	 fmt.Println("----------------------------------------------------")

	 myFunc_for()

	 fmt.Println("----------------------------------------------------")

	 myFunc_for1()

	 fmt.Println("----------------------------------------------------")

	 myFunc_for2()

	 fmt.Println("----------------------------------------------------")
}

func myFunc_goto() {
	i := 0
	Here:
	if i < 100 {
		fmt.Printf("%d,\n", i)
		i++
	} else {
		return
	}
	goto Here
}

func myFunc_for() {
	sum := 0

	for i := 0; i < 100; i++ {
		sum += i
	}

	println(sum)
}

func myFunc_for1() {
	for i := 0; i < 10; i++ {
		if i > 5 {
			continue
		}
		fmt.Printf("%d, ", i)
	}
	println()
}

func myFunc_for2() {
	J: for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			if j > 5 {
				break J
			}

			println(j)
		}
	}
}