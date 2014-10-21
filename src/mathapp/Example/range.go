package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)

	for i, j := range nums {
		if j == 2 {
			fmt.Println("index: ", i)
		}
	}

	mm := map[string]string{"Java": "Java WEB", "Javascript": "Javascript MVC"}
	for k, v := range mm {
		fmt.Println(k, "---", v)
	}

	for i, c := range "golang" {
		fmt.Printf("%d----%c\n", i, c)
	}
}
