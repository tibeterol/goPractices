package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type renk map[string]string

type biBakam struct{} // ici bos atruct tanimlanabiliyor istenirse ama tabi lazim bi kullanim degil

func main() {

	adana := biBakam{}
	adana.denensin()
	fmt.Println(sum(1.4, 6.8))
	/*colors := map[string]string{ // map tanimlama alternatif 3 ve ayni zamanda degerler direk ataniyor
		"red":    "#fff42",
		"yellow": "#fff423",
		"white":  "#546546",
	}*/

	//v.18 ile birlikte 2022 yılında cıkmıs bu versiyon go da genericlerde kullanilabiliyor
	// generic icin en guzel yol constraints.Ordered ile kullanmak ve bu bir import gerektiriyor golang.org/x/exp/constraints
	//fonk.larda,maplerde ve structlarda kullanilabiliyor

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for ", key, " is: ", value)
	}
}

func (r renk) printRenk() {
	for key, value := range r {
		fmt.Println("Hex code for ", key, " is: ", value)
	}
}

func sum[T constraints.Ordered](a T, b T) T {
	return a + b
}

func (biBakam) denensin() { // burada bi struct receiver olarak verildi ancak receiver olarak b biBakam seklinde instance verilmedi kullanilmayacagindan dolayi. Sayet b biBakam deseydim hata vermezdi ancak kullanilmayacak seyleri tanimlamak best practice degil
	println("deneme")
}
