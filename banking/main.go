package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", greet)

	log.Fatal(http.ListenAndServe("localhost:80", nil))

}

func greet(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!")

}
