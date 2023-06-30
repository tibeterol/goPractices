package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	l := logWriter{}

	if err == nil {

		io.Copy(l, resp.Body)

	}

}

func (l logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("boyut = ", len(bs))
	return len(bs), nil
}
