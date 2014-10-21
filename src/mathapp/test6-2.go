package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "hello")
}

func main() {

	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal("listenAndServe: ", err)
	}
}
