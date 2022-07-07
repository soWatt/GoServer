package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//simple easy pz server
	//need to gather api ideas
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello World!\n")
		io.WriteString(w, "Testing\n")
	}
	http.HandleFunc("/hello", helloHandler)
	log.Println("listening for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
