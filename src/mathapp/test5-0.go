package main

import (
	"fmt"
	"strconv"
)

/* interface */
type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "---" + h.name + "---" + strconv.Itoa(h.age) + "----" + h.phone
}

func main() {
	hu := Human{"zhangsan", 22, "1333333333"}
	fmt.Println(hu)
}
