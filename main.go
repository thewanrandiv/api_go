package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/greet", Greet)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func Greet(w http.ResponseWriter, r *http.Request) {
	log.Panic(fmt.Fprintf(w, "hello world"))

}
