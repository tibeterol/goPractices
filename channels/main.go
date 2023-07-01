package main

import (
	"fmt"
	"net/http"
)

// go routinesler concurrencyi saglar. channellar go routineler arasi iletisimi saglar. channellar olmazsa calismiyor.
// mesela bu ornekte go sendRequest(url) denildiginde 5 tane child go routine olusturuyor ve main routine in baska isi olmadigindan childlari beklemeden programi sonlandiriyor
// bu sorunun olmamasi icin channel kullanimi zorunlu ve her channelin bi typei var

//sayet mainin sonuna fmt.Println(<-c) koysak bu sefer sadece ilkini basiyor 2 tane koyulsa ilk ikisini basiyor.
//Cunku listede 5 tane var for icinde 5 adet child go routinei olusturdu fmt.Println(<-c) dan dolayi veri bekliyor ve main go routine sleep moduna geciyo veri geldikten sonra baska satir olmadigindan programi sonlandiriyor. Bunun onune for ile gecildi

func main() {
	urlList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://go.dev",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, url := range urlList {
		go sendRequest(url, c)
	}

	/*	for { // surekli istek atiyor. Channela her bilgi geldiginde dongu tekrar ediyor. Bunun aynisi asagidaki for loopunda
		go sendRequest(<-c, c)
	}*/

	for l := range c {
		go sendRequest(l, c)
	}

	/*	for i := 0; i < len(urlList); i++ {
		fmt.Println(<-c)
	}*/

}

func sendRequest(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, " is not active")
		c <- url
	} else {
		fmt.Println(url, " is active")
		c <- url
	}
}
