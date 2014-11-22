package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "这是一个web测试页面")
	fmt.Fprintf(w, "\n")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	error := http.ListenAndServe("0.0.0.0:9090", nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}
}