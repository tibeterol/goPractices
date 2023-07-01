package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customersList", getCustomers)

	log.Fatal(http.ListenAndServe("localhost:80", nil))

}

func greet(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!")

}

func getCustomers(w http.ResponseWriter, request *http.Request) {
	customerList := []customer{
		{Name: "Tibet", City: "İstanbul", ZipCode: "34718"},
		{Name: "Betül", City: "İzmir", ZipCode: "35718"},
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customerList)
}
