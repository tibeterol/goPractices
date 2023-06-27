package main

import "fmt"

// eger bir go dosyasi icinde fonk. disinda degisken tanimlanacaksa uzun sekilde yapilmasi ve bir deger atanacaksa bu sekilde atanmasi gerekmekte. Kisa yoldan deneme := 5 seklinde atanamiyor fonk. disinda ve ayni sekilde once var deneme int deyip alt satirda deneme = 6 yapsan gene olmuyor fonk. disinda yapildigi icin illa var deneme int = 5 dicen
// bunun sebebi kodun okunabilirligi. Fonk. disindaki degiskenlerin tipi net anlasilsin diye boyle tasarlanmis.

var deneme int = 5

func main() {
	//   deneme := 5
	fmt.Println(deneme)
}
