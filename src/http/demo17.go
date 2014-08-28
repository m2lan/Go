package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandle{})
	mux.HandleFunc("/hello", sayHello)

	error := http.ListenAndServe(":8080", mux)

	if error != nil {
		log.Fatal(error)
	}
}

type myHandle struct{}

func (*myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: "+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say hello")
}
