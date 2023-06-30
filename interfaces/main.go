package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")

	if err == nil {

		io.Copy(os.Stdout, resp.Body)

	}

}
