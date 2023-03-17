package dersler

import (
	"fmt"
)

var sayilar [10]int

func Eleman_ekle(sunu int, suraya int) {

	sayilar[suraya] = sunu

}

func Diziler() {

	fmt.Println(sayilar)
	var max int = sayilar[0]
	for i := 0; i < len(sayilar); i++ {
		if sayilar[i] > max {
			max = sayilar[i]
		}
	}
	fmt.Println("enbuyuk: ", max)

}

func Ikiboyutludizi() {
	sayac := 1
	var ikiboyutlu [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ikiboyutlu[i][j] = sayac
			sayac = sayac + 1
			fmt.Print(ikiboyutlu[i][j])

		}
		fmt.Println("")
	}
	fmt.Println(ikiboyutlu)
}
