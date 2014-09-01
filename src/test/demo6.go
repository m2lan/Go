package main

import (
	"fmt"
)

func main() {

	/**
	 * slice append copy map
	 */

	println("------------------------slice-----------------------------")
	array := []string{"aa", "bb", "cc", "dd", "ee"}

	array1 := array[1 : 3]

	array2 := array[:]

	array3 := array[:3]

	array4 := array[:len(array)]

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)

	var array5 [2]int
	array5[0] = 22
	array5[1] = 33
	fmt.Println(array5)

	println("------------------------append-----------------------------")
	ads := []int{0, 0}
	ads1 := append(ads, 2)
	ads2 := append(ads1, 1, 2, 3, 4)
	ads3 := append(ads2, ads2...)
	fmt.Println(ads1)
	fmt.Println(ads2)
	fmt.Println(ads3)

	println("------------------------copy-----------------------------")
	var a = []int{1, 2, 3, 4, 5, 6, 7}
	var b = make([]int, 6)
	s := copy(b, a[:3])
	fmt.Println(s, b)
	s1 := copy(b, b[2:])
	fmt.Println(s1, b)

	println("------------------------map-----------------------------")
	mymap := map[string]int {
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	year := 0
	for _, days := range mymap {
		year += days
		fmt.Println(days)
	}
	fmt.Println("一年共有:", year, "天")

	// 向map中添加元素
	mymap["kkkk"] = 33
	println(mymap["kkkk"])

	// 删除map中某个元素
	delete(mymap, "kkkk")

	// 判断kkkk是否存在
	_, ok := mymap["kkkk"]

	if ok {
		println("有值")
	} else {
		println("无值")
	}

}