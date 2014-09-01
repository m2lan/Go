package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/"] = hello
	mux["/sayhello"] = sayHello

	error := server.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if f, ok := mux[r.URL.String()]; ok {
		f(w, r)
		return
	}
	io.WriteString(w, "URL: "+r.URL.String())
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "sayhello")
}
