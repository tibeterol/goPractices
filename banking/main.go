package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/greet", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World!")

	})

	http.ListenAndServe("localhost:8080", nil)

}
