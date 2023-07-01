package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip"`
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

	// sayet json ya da xml oldugunu headerda belirtmezsek response plain text olarak gozukuyor

	if request.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customerList)
	} else { // json
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customerList)
	}

}
