package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, err := http.Get("http://google.com")

	if err == nil {

		bs := make([]byte, 99999) // 99.999 yer ayirdik. Cunku read fonk. slicei resize etmiyor. eger 0 lengthli byte slice yollasaydik parametre olarak bu zaten dolu diyip verileri aktarmayacakti
		resp.Body.Read(bs)

		fmt.Println(string(bs)) // byte slice i dogrudan stringe cevirebiliyoruz

	}

}
