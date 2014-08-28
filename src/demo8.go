package main

<<<<<<< HEAD
import (
	"github.com/hoisie/web"
)

func welcome(ctx *web.Context, val string) {
	for k, v := range ctx.Params {
		println(k, v)
	}
}

func main() {
	web.Get("/(.*)", welcome)
	web.Run("0.0.0.0:9999")
=======
import "fmt"

func main() {

LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break LABEL1
			}
		}
	}
	fmt.Println("OK")

	// 冒泡排序
	a := [...]int{2, 33, 11, 99, 55, 88}

	fmt.Println(a)

	for i, num := 0, len(a); i < num; i++ {
		for j := 0; j < i+1; j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}

	fmt.Println(a)

	a1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a1)

	s := a1[:]
	fmt.Println(s)

	s1 := make([]map[int]string, 10)

	for k := range s1 {
		s1[k] = make(map[int]string, 1)
		s1[k][1] = "OK"
	}

	fmt.Println(s1)

	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m1)
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
>>>>>>> 18255b2cbf6d68e0fe4d0fe56b2c062bb3bc5b30
}
