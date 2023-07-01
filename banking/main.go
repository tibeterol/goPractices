package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", greet)

	http.ListenAndServe("localhost:8080", nil)

}

func greet(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!")

}
