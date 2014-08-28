package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandle{})
	mux.HandleFunc("/hello", sayHello)

	wd, error := os.Getwd()
	if error != nil {
		log.Fatal(error)
	}

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	error = http.ListenAndServe(":8080", mux)

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
