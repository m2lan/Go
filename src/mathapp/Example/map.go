package main

import (
	"fmt"
)

func main() {

	/**
	 * map 两种用法
	 * 1. m:=make(map[string]int) 对map进行初始化
	 * 2. m:=map[string]int{"aa":1,"bb":2} 对map进行初始化并赋值
	 */

	m := make(map[string]int)
	m["java"] = 3
	m["php"] = 2
	fmt.Println(m)
	fmt.Println(len(m))

	v1 := m["java"]
	fmt.Println(v1)

	delete(m, "java")
	fmt.Println(m)

	/*
	 * @return 1:value 2:boolean(值存在为true，不存在为false)
	 */
	_, pes := m["php"]
	fmt.Println(pes)

	foot := map[string]int{"nodejs": 1, "go": 1}
	fmt.Println(foot)
}
